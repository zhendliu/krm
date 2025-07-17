package namespace

import (
	"context"
	"krm-backend/config"
	"krm-backend/controllers"
	"krm-backend/utils/logs"

	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Delete(r *gin.Context) {
	logs.Debug(nil, "删除Namespace")
	clientset, basicInfo, err := controllers.BasicInit(r, nil)
	returnData := config.ReturnData{}
	if err != nil {
		returnData.Message = err.Error()
		returnData.Status = 400
		r.JSON(200, returnData)
		return
	}
	if basicInfo.Name == "kube-system" {
		returnData.Message = "禁止删除kube-system"
		returnData.Status = 400
		r.JSON(200, returnData)
		return
	}
	err = clientset.CoreV1().Namespaces().Delete(context.TODO(), basicInfo.Name, metav1.DeleteOptions{})
	if err != nil {
		msg := "删除Namespace失败: " + err.Error()
		returnData.Message = msg
		returnData.Status = 400
	} else {
		returnData.Message = "删除成功"
		returnData.Status = 200
	}
	r.JSON(200, returnData)
}
