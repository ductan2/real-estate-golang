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
	r.Use(cors)
	LoadConfig()
	InitCloudinary()
	InitDB()
	InitRedis()
	// InitELK()
	InitRabbitMQ()
	InitRouter(r)
	// initialize kafka
	return r
}
