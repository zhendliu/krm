package cluster

import (
	"context"
	"krm-backend/config"
	"krm-backend/utils/logs"

	"github.com/gin-gonic/gin"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Delete(r *gin.Context) {
	logs.Debug(nil, "删除集群")
	// 1. 接收删除集群的ID
	clusterId := r.Query("id")
	// 2. 删除
	err := config.InClusterClientSet.CoreV1().Secrets(config.MetadataNamespace).Delete(
		context.TODO(), clusterId, v1.DeleteOptions{})
	returnData := config.ReturnData{}
	if err != nil {
		logs.Error(map[string]interface{}{"id": clusterId}, "删除失败")
		// 说明删除集群失败
		msg := "集群删除失败: " + err.Error()
		returnData.Status = 400
		returnData.Message = msg
	} else {
		logs.Warning(map[string]interface{}{"id": clusterId}, "删除成功")
		// 说明删除成功
		returnData.Status = 200
		returnData.Message = "删除成功"
		delete(config.ClusterKubeconfig, clusterId)
		// fmt.Println("当前集群配置:", config.ClusterKubeconfig)
	}
	r.JSON(200, returnData)
}
