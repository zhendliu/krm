package pv

import (
	"krm-backend/controllers"
	"krm-backend/utils/logs"

	"github.com/dotbalo/kubeutils/kubeutils"
	"github.com/gin-gonic/gin"
	corev1 "k8s.io/api/core/v1"
)

func Create(r *gin.Context) {
	logs.Debug(nil, "创建PersistentVolume")
	// 需要看一下是否要自动创建headless service
	var pv corev1.PersistentVolume
	var info controllers.Info
	info.Item = &pv
	kubeconfig := controllers.NewInfo(r, &info, "创建成功")
	// 改为接口形式的创建
	var kubeUtilser kubeutils.KubeUtilser
	instance := kubeutils.NewPersistentVolume(kubeconfig, &pv)
	// 把实例复制给kubeUtilser
	kubeUtilser = instance
	// 使用kubeUtilser创建资源
	info.Create(r, kubeUtilser)
}
