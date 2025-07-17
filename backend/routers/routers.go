// 路由层 管理程序的路由信息
package routers

import (
	"krm-backend/routers/auth"
	"krm-backend/routers/cluster"
	"krm-backend/routers/configmap"
	"krm-backend/routers/cronjob"
	"krm-backend/routers/daemonset"
	"krm-backend/routers/deployment"
	"krm-backend/routers/ingress"
	"krm-backend/routers/namespace"
	"krm-backend/routers/node"
	"krm-backend/routers/pod"
	"krm-backend/routers/pv"
	"krm-backend/routers/pvc"
	"krm-backend/routers/replicaset"
	"krm-backend/routers/secret"
	"krm-backend/routers/service"
	"krm-backend/routers/statefulset"
	"krm-backend/routers/storageclass"

	"github.com/gin-gonic/gin"
)

// 注册路由的方法
func RegisterRouters(r *gin.Engine) {
	// 登录的路由配置
	// 1. 登录：login
	// 2. 退出：logout
	// 3. 	/api/auth/login
	// 		/api/auth/logout
	apiGroup := r.Group("/api")
	auth.RegisterSubRouter(apiGroup)
	cluster.RegisterSubRouter(apiGroup)
	namespace.RegisterSubRouter(apiGroup)
	pod.RegisterSubRouter(apiGroup)
	deployment.RegisterSubRouter(apiGroup)
	statefulset.RegisterSubRouter(apiGroup)
	daemonset.RegisterSubRouter(apiGroup)
	cronjob.RegisterSubRouter(apiGroup)
	replicaset.RegisterSubRouter(apiGroup)
	node.RegisterSubRouter(apiGroup)
	service.RegisterSubRouter(apiGroup)
	ingress.RegisterSubRouter(apiGroup)
	configmap.RegisterSubRouter(apiGroup)
	secret.RegisterSubRouter(apiGroup)
	pv.RegisterSubRouter(apiGroup)
	pvc.RegisterSubRouter(apiGroup)
	storageclass.RegisterSubRouter(apiGroup)

	// Prometheus istio

}
