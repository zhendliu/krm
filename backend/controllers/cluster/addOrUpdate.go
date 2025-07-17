// 用来更新或者创建集群
package cluster

import (
	"context"
	"krm-backend/config"
	"krm-backend/utils"
	"krm-backend/utils/logs"
	"net/http"

	"github.com/gin-gonic/gin"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// method: update，  create
func addOrUpdate(r *gin.Context, method string) {
	var arg string
	if method == "Create" {
		arg = "添加"
	} else {
		arg = "更新"
	}
	// 声明一个集群配置
	clusterConfig := ClusterConfig{}
	returnData := config.NewReturnData()
	if err := r.ShouldBindJSON(&clusterConfig); err != nil {
		msg := arg + "集群的配置信息不完整: " + err.Error()
		returnData.Status = 400
		returnData.Message = msg
		r.JSON(200, returnData)
		return
	}
	// 判断集群是否正常
	clusterStatus, err := clusterConfig.getClusterStatus()
	if err != nil {
		msg := "无法获取集群信息: " + err.Error()
		returnData.Status = 400
		returnData.Message = msg
		r.JSON(http.StatusOK, returnData)
		logs.Error(map[string]interface{}{"error": err.Error()}, arg+"集群失败,无法获取集群信息")
		return
	}
	logs.Info(map[string]interface{}{"集群名称": clusterConfig.DisplayName, "集群ID": clusterConfig.Id}, "开始"+arg+"集群")
	// 创建一个集群配置的secret
	var clusterConfigSecret corev1.Secret
	clusterConfigSecret.Name = clusterConfig.Id
	clusterConfigSecret.Labels = make(map[string]string)
	clusterConfigSecret.Labels[config.ClusterConfigSecretLabelKey] = config.ClusterConfigSecretLabelValue
	// 添加注释，保存集群的配置信息
	clusterConfigSecret.Annotations = make(map[string]string)
	// 把集群的状态结构体转成map
	m := utils.Struct2Map(clusterStatus)
	clusterConfigSecret.Annotations = m
	// 保存kubeconfig
	clusterConfigSecret.StringData = make(map[string]string)
	clusterConfigSecret.StringData["kubeconfig"] = clusterConfig.Kubeconfig
	// 创建secret
	if method == "Create" {
		_, err = config.InClusterClientSet.CoreV1().Secrets(config.MetadataNamespace).Create(context.TODO(), &clusterConfigSecret, metav1.CreateOptions{})
	} else {
		_, err = config.InClusterClientSet.CoreV1().Secrets(config.MetadataNamespace).Update(context.TODO(), &clusterConfigSecret, metav1.UpdateOptions{})
	}
	if err != nil {
		// 说明创建失败
		logs.Error(map[string]interface{}{"集群ID": clusterConfig.Id, "集群名字:": clusterConfig.DisplayName, "msg": err.Error()}, "集群"+arg+"失败")
		msg := arg + "集群失败: " + err.Error()
		returnData.Message = msg
		returnData.Status = 400
		r.JSON(200, returnData)
		return
	}
	config.ClusterKubeconfig[clusterConfig.Id] = clusterConfig.Kubeconfig
	// fmt.Println("当前集群配置:", config.ClusterKubeconfig)
	logs.Info(map[string]interface{}{"集群ID": clusterConfig.Id, "集群名字:": clusterConfig.DisplayName}, "集群"+arg+"成功")
	returnData.Status = 200
	returnData.Message = arg + "成功"
	r.JSON(200, returnData)
}
