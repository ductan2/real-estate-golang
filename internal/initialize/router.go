package initialize

import (
	"ecommerce/global"
	"ecommerce/internal/routers"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	var r *gin.Engine
	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}
	userRouter := routers.RouterGroupApp
	MainGroup := r.Group("/api/v1")
	{
		MainGroup.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}
	{
		userRouter.User.InitUserRouter(MainGroup)
		userRouter.Seller.InitSellerRouter(MainGroup)
		userRouter.Admin.InitAdminRouter(MainGroup)
	}
	

	return r
}
