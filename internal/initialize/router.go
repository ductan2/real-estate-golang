package initialize

import (
	"ecommerce/internal/middlewares"
	"ecommerce/internal/routers"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	// Add logger middleware
	r.Use(middlewares.Logger())

	mainRouter := routers.RouterGroupApp
	MainGroup := r.Group("/api/v1")
	{
		MainGroup.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}
	{
		mainRouter.User.InitUserRouter(MainGroup)
		mainRouter.Seller.InitSellerRouter(MainGroup)
		mainRouter.Admin.InitAdminRouter(MainGroup)
		mainRouter.Investor.InitInvestorRouter(MainGroup)
		mainRouter.Project.InitProjectRouter(MainGroup)
	}
}
