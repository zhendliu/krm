package pod

import (
	"krm-backend/controllers"
	"krm-backend/utils/logs"

	"github.com/dotbalo/kubeutils/kubeutils"
	"github.com/gin-gonic/gin"
)

func Delete(r *gin.Context) {
	logs.Debug(nil, "删除Pod")
	var info controllers.Info
	kubeconfig := controllers.NewInfo(r, &info, "删除成功")
	// 改为接口形式的创建
	var kubeUtilser kubeutils.KubeUtilser
	instance := kubeutils.NewPod(kubeconfig, nil)
	// 把实例复制给kubeUtilser
	kubeUtilser = instance
	// 使用kubeUtilser创建资源
	info.Delete(r, kubeUtilser)
}

func DeleteList(r *gin.Context) {
	logs.Debug(nil, "删除Pod")
	var info controllers.Info
	kubeconfig := controllers.NewInfo(r, &info, "删除成功")
	// 改为接口形式的创建
	var kubeUtilser kubeutils.KubeUtilser
	instance := kubeutils.NewPod(kubeconfig, nil)
	// 把实例复制给kubeUtilser
	kubeUtilser = instance
	// 使用kubeUtilser创建资源
	info.DeleteList(r, kubeUtilser)
}
