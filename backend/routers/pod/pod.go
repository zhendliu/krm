package pod

import (
	"krm-backend/controllers/pod"

	"github.com/gin-gonic/gin"
)

func add(podGroup *gin.RouterGroup) {
	podGroup.POST("/create", pod.Create)
}
func update(podGroup *gin.RouterGroup) {
	podGroup.POST("/update", pod.Update)
}
func delete(podGroup *gin.RouterGroup) {
	podGroup.POST("/delete", pod.Delete)
}
func deleteList(podGroup *gin.RouterGroup) {
	podGroup.POST("/deleteList", pod.DeleteList)
}
func get(podGroup *gin.RouterGroup) {
	podGroup.GET("/get", pod.Get)
}
func list(podGroup *gin.RouterGroup) {
	podGroup.GET("/list", pod.List)
}

func RegisterSubRouter(g *gin.RouterGroup) {
	podGroup := g.Group("/pod")
	add(podGroup)
	update(podGroup)
	delete(podGroup)
	get(podGroup)
	list(podGroup)
	deleteList(podGroup)
}
