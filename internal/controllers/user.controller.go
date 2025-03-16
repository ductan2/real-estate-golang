package controllers

import (
	"ecommerce/internal/middlewares"
	services "ecommerce/internal/services/user"
	"ecommerce/internal/vo"
	"ecommerce/pkg/response"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService services.IUserService
}

func NewUserController(userService services.IUserService) *UserController {
	return &UserController{userService: userService}
}

func (c *UserController) Register(ctx *gin.Context) {
	params := vo.UserRegisterRequest{}
	err := ctx.ShouldBindBodyWithJSON(&params)
	if err != nil {
		response.ErrorResponse(ctx, response.UnprocessableEntity, err.Error())
		return
	}

	ip := ctx.ClientIP()
	err = c.userService.Register(params.Email, params.Password, ip, params.Purpose)
	if err != nil {
		response.ErrorResponse(ctx, response.InternalServerError, err.Error())
		return
	}
	err = c.userService.SendOtp(params.Email)
	if err != nil {
		response.ErrorResponse(ctx, response.InternalServerError, err.Error())
		return
	}

	response.SuccessResponse(ctx, response.Success, "User registered successfully")
}

func (c *UserController) VerifyOtp(ctx *gin.Context) {
	params := vo.UserVerifyOtpRequest{}
	err := ctx.ShouldBindBodyWithJSON(&params)
	if err != nil {
		response.ErrorResponse(ctx, response.UnprocessableEntity, err.Error())
		return
	}
	err = c.userService.VerifyOtp(params.Email, params.Otp)
	if err != nil {
		response.ErrorResponse(ctx, response.BadRequest, err.Error())
		return
	}

	response.SuccessResponse(ctx, response.Success, "Verify OTP successfully")

}

func (c *UserController) Login(ctx *gin.Context) {
	params := vo.UserLoginRequest{}
	err := ctx.ShouldBindBodyWithJSON(&params)
	if err != nil {
		response.ErrorResponse(ctx, response.UnprocessableEntity, err.Error())
		return
	}
	ip := ctx.ClientIP()
	token, err := c.userService.Login(params.Email, params.Password, ip)
	if err != nil {
		response.ErrorResponse(ctx, response.BadRequest, err.Error())
		return
	}
	// asign token to cookie
	ctx.SetCookie("access_token", token, 60*60*24, "/", "", false, true)

	response.SuccessResponse(ctx, response.Success, gin.H{"token": token})
}

func (c *UserController) Logout(ctx *gin.Context) {
	userId := ctx.Request.Context().Value(middlewares.UserUUIDKey).(string)
	err := c.userService.Logout(userId)
	if err != nil {
		response.ErrorResponse(ctx, response.InternalServerError, err.Error())
		return
	}
	ctx.SetCookie("access_token", "", -1, "/", "", false, true)
	response.SuccessResponse(ctx, response.Success, "Logout successfully")
}

func (c *UserController) GetUserInfo(ctx *gin.Context) {
	userId := ctx.Request.Context().Value(middlewares.UserUUIDKey).(string)
	user := c.userService.GetUserInfo(userId)
	response.SuccessResponse(ctx, response.Success, user)
}

func (c *UserController) UpdateUserInfo(ctx *gin.Context) {
	userId := ctx.Request.Context().Value(middlewares.UserUUIDKey).(string)

	var params vo.UserUpdateInfoRequest
	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(ctx, response.UnprocessableEntity, err.Error())
		return
	}

	err := c.userService.UpdateUserInfo(userId, params)
	if err != nil {
		response.ErrorResponse(ctx, response.InternalServerError, err.Error())
		return
	}

	response.SuccessResponse(ctx, response.Success, "User info updated successfully")
}
