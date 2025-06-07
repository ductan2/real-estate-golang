package controllers

import (
	"ecommerce/internal/middlewares"
	services "ecommerce/internal/services/user"
	"ecommerce/internal/vo"
	"ecommerce/pkg/response"

	"github.com/gin-gonic/gin"
)

type AdminController struct {
	adminService services.IAdminService
	sellerService services.ISellerService
}

func NewAdminController(adminService services.IAdminService, sellerService services.ISellerService) *AdminController {
	return &AdminController{adminService: adminService, sellerService: sellerService}
}

func (c *AdminController) ApplyForAdmin(ctx *gin.Context) {
	userId := ctx.Request.Context().Value(middlewares.UserUUIDKey).(string)
	err := c.adminService.ApplyForAdmin(userId)
	if err != nil {
		response.ErrorResponse(ctx, response.InternalServerError, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.Success, "Apply for admin successfully")
}

func (c *AdminController) ApproveSellerRequest(ctx *gin.Context) {
	userId := ctx.Request.Context().Value(middlewares.UserUUIDKey).(string)
	var params vo.SellerApproveRequest
	err := ctx.ShouldBindJSON(&params)
	if err != nil {
		response.ErrorResponse(ctx, response.UnprocessableEntity, err.Error())
		return
	}
	sellerExist := c.sellerService.GetSeller(params.SellerID)
	if sellerExist == nil || sellerExist.IsVerified {
		response.ErrorResponse(ctx, response.NotFound, "Seller not found")
		return
	}
	err = c.adminService.ApproveSellerRequest(userId, params.SellerID, sellerExist.User.Email, params.Approved)
	if err != nil {
		response.ErrorResponse(ctx, response.InternalServerError, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.Success, "Approve seller request successfully")
}

func (c *AdminController) BlockSeller(ctx *gin.Context) {
	userId := ctx.Request.Context().Value(middlewares.UserUUIDKey).(string)
	var params vo.SellerBlockRequest
	err := ctx.ShouldBindJSON(&params)
	if err != nil {
		response.ErrorResponse(ctx, response.UnprocessableEntity, err.Error())
		return
	}
	sellerExist := c.sellerService.GetSeller(params.SellerID)
	if sellerExist == nil || !sellerExist.IsVerified {
		response.ErrorResponse(ctx, response.NotFound, "Seller not found")
		return
	}
	err = c.adminService.BlockSeller(userId, params.SellerID, params.Reason)
	if err != nil {
		response.ErrorResponse(ctx, response.InternalServerError, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.Success, "Block seller request successfully")
}
