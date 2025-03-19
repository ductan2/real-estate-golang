package initialize

import (
	"ecommerce/internal/middleware"
	"ecommerce/internal/routers"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	// Add logger middleware
	r.Use(middleware.Logger())

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
	}
}
