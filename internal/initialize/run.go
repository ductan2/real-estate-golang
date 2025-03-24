package initialize

import (
	"ecommerce/global"

	"github.com/gin-gonic/gin"
)

func Run() *gin.Engine {
	// Load configuration
	// initialize database
	// initialize redis
	InitEnv()
	var r *gin.Engine
	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}
	LoadConfig()
	InitCloudinary()
	InitDB()
	InitRedis()
	cors := InitCors()
	r.Use(cors)
	InitRouter(r)
	// initialize kafka

	return r
}
