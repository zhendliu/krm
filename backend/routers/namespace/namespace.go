package namespace

import (
	"krm-backend/controllers/namespace"

	"github.com/gin-gonic/gin"
)

func add(namespaceGroup *gin.RouterGroup) {
	namespaceGroup.POST("/create", namespace.Create)
}
func update(namespaceGroup *gin.RouterGroup) {
	namespaceGroup.POST("/update", namespace.Update)
}
func delete(namespaceGroup *gin.RouterGroup) {
	namespaceGroup.GET("/delete", namespace.Delete)
}
func get(namespaceGroup *gin.RouterGroup) {
	namespaceGroup.GET("/get", namespace.Get)
}
func list(namespaceGroup *gin.RouterGroup) {
	namespaceGroup.GET("/list", namespace.List)
}

func RegisterSubRouter(g *gin.RouterGroup) {
	namespaceGroup := g.Group("/namespace")
	add(namespaceGroup)
	update(namespaceGroup)
	delete(namespaceGroup)
	get(namespaceGroup)
	list(namespaceGroup)
}
