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
	InitLogger()
	cors := InitCors()
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
	r.Use(cors)
	InitCloudinary()
	InitDB()
	InitRedis()
	InitRabbitMQ()
	InitRouter(r)
	return r
}
