package controllers

import (
	"ecommerce/global"
	"ecommerce/internal/middlewares"
	services "ecommerce/internal/services/user"
	"ecommerce/internal/storage/cloudinary"
	"ecommerce/internal/utils/auth"
	"ecommerce/internal/vo"
	"ecommerce/pkg/response"
	"errors"

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
	user, err := c.userService.GetUserByEmail(params.Email)
	if err != nil {
		response.ErrorResponse(ctx, response.InternalServerError, err.Error())
		return
	}
	if user != nil {
		response.ErrorResponse(ctx, response.BadRequest, "User already exists")
		return
	}
	ip := ctx.ClientIP()
	err = c.userService.Register(params.Username, params.Email, params.Password, ip, params.Purpose)
	if err != nil {
		response.ErrorResponse(ctx, response.InternalServerError, err.Error())
		return
	}
	go func(email string) {
		if err := c.userService.SendOtp(email); err != nil {
			global.Logger.Errorf("Send OTP error: %v", err)
		}
	}(params.Email)

	response.SuccessResponse(ctx, response.Success, "User registered successfully")
}

func (c *UserController) VerifyOtp(ctx *gin.Context) {
	params := vo.UserVerifyOtpRequest{}
	err := ctx.ShouldBindBodyWithJSON(&params)
	if err != nil {
		middlewares.HandleError(ctx, err, response.BadRequest)
		return
	}
	err = c.userService.VerifyOtp(params.Email, params.Otp)
	if err != nil {
		middlewares.HandleError(ctx, err, response.BadRequest)
		return
	}

	response.SuccessResponse(ctx, response.Success, "Verify OTP successfully")
}

func (c *UserController) ResendOtp(ctx *gin.Context) {
	params := vo.UserResendOtpRequest{}
	err := ctx.ShouldBindBodyWithJSON(&params)
	if err != nil {
		middlewares.HandleError(ctx, err, response.UnprocessableEntity)
		return
	}
	go func(email string) {
		if err := c.userService.SendOtp(email); err != nil {
			global.Logger.Errorf("Send OTP error: %v", err)
		}
	}(params.Email)
	response.SuccessResponse(ctx, response.Success, "Resend OTP successfully")
}

func (c *UserController) Login(ctx *gin.Context) {
	params := vo.UserLoginRequest{}
	err := ctx.ShouldBindBodyWithJSON(&params)
	if err != nil {
		middlewares.HandleError(ctx, err, response.UnprocessableEntity)
		return
	}
	userSession := vo.UserSession{
 		IpAddress: ctx.ClientIP(),
		Location:  auth.GetLocationFromIP(ctx.ClientIP()),
		Device:    auth.GetUserAgentDetails(ctx.Request.Header.Get("User-Agent")).Platform(),
 		UserAgent: ctx.Request.Header.Get("User-Agent"),
 	}
	token, err := c.userService.Login(params.Email, params.Password, userSession)
	if err != nil {
		middlewares.HandleError(ctx, err, response.BadRequest)
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
		middlewares.HandleError(ctx, err, response.InternalServerError)
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
		middlewares.HandleError(ctx, err, response.UnprocessableEntity)
		return
	}

	err := c.userService.UpdateUserInfo(userId, params)
	if err != nil {
		middlewares.HandleError(ctx, err, response.InternalServerError)
		return
	}

	response.SuccessResponse(ctx, response.Success, "User info updated successfully")
}

// UploadAvatar handles the avatar upload for a user
func (c *UserController) UploadAvatar(ctx *gin.Context) {
	// Get user ID from context
	userId := ctx.Request.Context().Value(middlewares.UserUUIDKey).(string)

	// Get the file from the request
	file, err := ctx.FormFile("avatar")
	if err != nil {
		middlewares.HandleError(ctx, err, response.BadRequest)
		return
	}

	// Get Cloudinary service
	imageService := cloudinary.GetImageService()
	if imageService == nil {
		middlewares.HandleError(ctx, errors.New("image service not initialized"), response.InternalServerError)
		return
	}

	// Upload image to Cloudinary
	imageUrl, err := imageService.UploadImage(file, "avatars")
	if err != nil {
		middlewares.HandleError(ctx, err, response.InternalServerError)
		return
	}

	// Update user avatar in database
	err = c.userService.UpdateUserAvatar(userId, imageUrl)
	if err != nil {
		// If database update fails, try to remove the uploaded image
		_ = imageService.RemoveImage(imageUrl)
		middlewares.HandleError(ctx, err, response.InternalServerError)
		return
	}

	response.SuccessResponse(ctx, response.Success, gin.H{
		"avatar_url": imageUrl,
	})
}
