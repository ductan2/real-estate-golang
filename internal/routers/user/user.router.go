package user

import (
	"ecommerce/internal/middlewares"
	"ecommerce/internal/wire"

	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (us *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userController, _ := wire.InitUserRouterHanlder()
	userRouter := Router.Group("user")
	{
		userRouter.POST("register", userController.Register)
		userRouter.POST("verify-otp", userController.VerifyOtp)
		userRouter.POST("login", userController.Login)
		userRouter.POST("resend-otp", userController.ResendOtp)
	}
	privateRouter := Router.Group("user")
	privateRouter.Use(middlewares.AuthenMiddleware())
	{
		privateRouter.DELETE("logout", userController.Logout)
		privateRouter.GET("profile", userController.GetUserInfo)
		privateRouter.PATCH("profile", userController.UpdateUserInfo)
		privateRouter.PATCH("upload-avatar", userController.UploadAvatar)
	}
}