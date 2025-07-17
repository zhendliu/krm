// 项目的总入口
package main

import (
	"krm-backend/config"
	_ "krm-backend/controllers/initcontroller"
	"krm-backend/middlerwares"
	"krm-backend/routers"
	"krm-backend/utils/logs"

	"github.com/gin-gonic/gin"
)

// import (
// 	"krm-backend/config"
// 	"krm-backend/middlerwares"
// 	"krm-backend/routers"
// 	"krm-backend/utils/logs"

// 	"github.com/gin-gonic/gin"
// )

func main() {
	// 1. 加载程序的配置
	// 2. 配置gin
	r := gin.Default()
	r.Use(middlerwares.LoggerMiddleware())
	r.Use(middlerwares.CORSMiddleware())
	r.Use(middlerwares.JWTAuth)
	logs.Info(nil, "程序启动成功")

	routers.RegisterRouters(r)
	r.Run(config.Port)
}
