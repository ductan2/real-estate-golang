package project

import (
	"ecommerce/internal/middlewares"
	"ecommerce/internal/wire"

	"github.com/gin-gonic/gin"
)

type ProjectRouter struct {}


func (r *ProjectRouter) InitProjectRouter(Router *gin.RouterGroup) {
	projectController, _ := wire.InitProjectRouterHanlder()
	projectRouter := Router.Group("/project")
	projectRouter.Use(middlewares.OptionalAuthMiddleware())
	{
		projectRouter.POST("/", projectController.Create)
		projectRouter.GET("/", projectController.GetAll)
		projectRouter.GET("/suggest", projectController.Suggest)
		projectRouter.GET("/:id", projectController.GetById)
		projectRouter.PUT("/:id", projectController.Update)
		projectRouter.DELETE("/:id", projectController.Delete)
		projectRouter.GET("/investor/:investorId", projectController.GetByInvestor)
	}
}
