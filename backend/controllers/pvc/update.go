package pvc

import (
	"krm-backend/controllers"
	"krm-backend/utils/logs"

	"github.com/dotbalo/kubeutils/kubeutils"
	"github.com/gin-gonic/gin"
	corev1 "k8s.io/api/core/v1"
)

func Update(r *gin.Context) {
	logs.Debug(nil, "更新PersistentVolumeClaim")
	var pvc corev1.PersistentVolumeClaim
	var info controllers.Info
	info.Item = &pvc
	kubeconfig := controllers.NewInfo(r, &info, "更新成功")
	// 改为接口形式的创建
	var kubeUtilser kubeutils.KubeUtilser
	instance := kubeutils.NewPersistentVolumeClaim(kubeconfig, &pvc)
	// 把实例复制给kubeUtilser
	kubeUtilser = instance
	// 使用kubeUtilser创建资源
	info.Update(r, kubeUtilser)
}
