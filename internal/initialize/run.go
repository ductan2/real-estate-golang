package initialize

import (
	"github.com/gin-gonic/gin"
)

func Run() *gin.Engine {
	// Load configuration
	// initialize database
	// initialize redis
	LoadConfig()
	InitDB()
	InitRedis()
	r:=InitRouter()
	// initialize kafka
	
	return r
}