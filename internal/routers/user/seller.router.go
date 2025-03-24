package user

import (
	"ecommerce/internal/middlewares"
	"ecommerce/internal/wire"

	"github.com/gin-gonic/gin"
)

type SellerRouter struct{}

func (sr *SellerRouter) InitSellerRouter(Router *gin.RouterGroup) {
	sellerController, _ := wire.InitSellerRouterHandler()
	sellerRouter := Router.Group("seller")
	sellerRouter.Use(middlewares.AuthenMiddleware())
	{
		sellerRouter.POST("apply", sellerController.ApplyForSeller)
	}
}