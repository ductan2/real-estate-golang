package investor

import (
	"ecommerce/internal/middlewares"
	"ecommerce/internal/wire"

	"github.com/gin-gonic/gin"
)

type InvestorRouter struct{}

func (ir *InvestorRouter) InitInvestorRouter(Router *gin.RouterGroup) {
	investorController, _ := wire.InitInvestorRouterHanlder()
	investorRouter := Router.Group("investor")
	investorRouter.Use(middlewares.AuthenMiddleware())
	{
		investorRouter.POST("", investorController.Create)
		investorRouter.GET("", investorController.GetAll)
		investorRouter.GET(":id", investorController.GetById)
		investorRouter.PUT(":id", investorController.Update)
		investorRouter.DELETE(":id", investorController.Delete)
	}
}
