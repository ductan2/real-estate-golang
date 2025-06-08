package initialize

import (
	"ecommerce/internal/middlewares"
	"ecommerce/internal/routers"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.elastic.co/apm"
	"go.elastic.co/apm/module/apmgin"
)

func InitRouter(r *gin.Engine) {
	// Configure APM middleware with error tracking
	apmMiddleware := apmgin.Middleware(r, apmgin.WithTracer(apm.DefaultTracer))

	// Add middlewares
	r.Use(middlewares.Logger())
	r.Use(apmMiddleware)
	r.Use(middlewares.Prometheus())

	// expose metrics endpoint
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

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
