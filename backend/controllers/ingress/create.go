package ingress

import (
	"krm-backend/controllers"
	"krm-backend/utils/logs"

	"github.com/dotbalo/kubeutils/kubeutils"
	"github.com/gin-gonic/gin"

	// 1.19+ v1  1.19 extension/v1beta1
	networkingv1 "k8s.io/api/networking/v1"
)

func Create(r *gin.Context) {
	logs.Debug(nil, "创建Ingress")
	var ingress networkingv1.Ingress
	var info controllers.Info
	info.Item = &ingress
	kubeconfig := controllers.NewInfo(r, &info, "创建成功")
	// 改为接口形式的创建
	var kubeUtilser kubeutils.KubeUtilser
	instance := kubeutils.NewIngress(kubeconfig, &ingress)
	// 把实例复制给kubeUtilser
	kubeUtilser = instance
	// 使用kubeUtilser创建资源
	info.Create(r, kubeUtilser)
}
