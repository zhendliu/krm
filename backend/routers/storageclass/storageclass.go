package storageclass

import (
	"krm-backend/controllers/storageclass"

	"github.com/gin-gonic/gin"
)

func list(storageclassGroup *gin.RouterGroup) {
	storageclassGroup.GET("/list", storageclass.List)
}

func RegisterSubRouter(g *gin.RouterGroup) {
	storageclassGroup := g.Group("/storageclass")
	list(storageclassGroup)
}
