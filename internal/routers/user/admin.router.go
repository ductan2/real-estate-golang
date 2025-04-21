package user

import (
	"ecommerce/internal/middlewares"
	"ecommerce/internal/wire"

	"github.com/gin-gonic/gin"
)

type AdminRouter struct{}

func (ar *AdminRouter) InitAdminRouter(Router *gin.RouterGroup) {
	adminController, _ := wire.InitAdminRouterHandler()
	adminService, _ := wire.InitAdminService()
	
	adminRouter := Router.Group("admin")
	adminRouter.Use(middlewares.AuthenMiddleware())
	adminRouter.Use(middlewares.AdminMiddleware(adminService))
	{
		adminRouter.POST("apply", adminController.ApplyForAdmin)
		adminRouter.PUT("approve/seller", adminController.ApproveSellerRequest)
		adminRouter.DELETE("block/seller", adminController.BlockSeller)
	}

}