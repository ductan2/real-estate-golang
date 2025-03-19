package initialize

import (
	"ecommerce/global"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitCors() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{global.Config.Server.FrontendUrl},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Authorization", "Content-Type", "X-Requested-With"},
		AllowCredentials: true,
		ExposeHeaders:    []string{"Set-Cookie"},
	})
}
