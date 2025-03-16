package controllers

import (
	"ecommerce/internal/middlewares"
	services "ecommerce/internal/services/user"
	"ecommerce/pkg/response"

	"github.com/gin-gonic/gin"
)

type SellerController struct {
	sellerService services.ISellerService
}

func NewSellerController(sellerService services.ISellerService) *SellerController {
	return &SellerController{sellerService: sellerService}
}

func (c *SellerController) ApplyForSeller(ctx *gin.Context) {
	userId := ctx.Request.Context().Value(middlewares.UserUUIDKey).(string)
	err := c.sellerService.ApplyForSeller(userId)
	if err != nil {
		response.ErrorResponse(ctx, response.InternalServerError, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.Success, "Apply for seller successfully")
	
}
