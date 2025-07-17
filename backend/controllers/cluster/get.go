package cluster

import (
	"context"
	"krm-backend/config"
	"krm-backend/utils/logs"

	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Get(r *gin.Context) {
	logs.Debug(nil, "获取集群详情")
	clusterId := r.Query("id")
	returnData := config.ReturnData{}
	// 获取集群的配置/secret
	clusterSecret, err := config.InClusterClientSet.CoreV1().Secrets(config.MetadataNamespace).Get(context.TODO(), clusterId, metav1.GetOptions{})
	if err != nil {
		// 说明查询集群失败
		logs.Error(map[string]interface{}{"id": clusterId, "msg": err.Error()}, "查询集群详解失败")
		returnData.Status = 400
		returnData.Message = "获取集群详情失败: " + err.Error()
	} else {
		returnData.Status = 200
		returnData.Message = "查询成功"
		returnData.Data = make(map[string]interface{})
		clusterConfigMap := clusterSecret.Annotations
		clusterConfigMap["kubeconfig"] = string(clusterSecret.Data["kubeconfig"])
		returnData.Data["item"] = clusterConfigMap
	}
	r.JSON(200, returnData)
}
