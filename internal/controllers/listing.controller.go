package controllers

import (
	"ecommerce/internal/middlewares"
	userService "ecommerce/internal/services/user"
	listingService "ecommerce/internal/services/listing"
	"ecommerce/internal/vo"
	"ecommerce/pkg/enum"
	"ecommerce/pkg/response"

	"github.com/gin-gonic/gin"
)

type ListingController struct {
	listingService listingService.IListingService
	userService    userService.IUserService
}

func NewListingController(listingService listingService.IListingService, userService userService.IUserService) *ListingController {
	return &ListingController{
		listingService: listingService,
		userService:    userService,
	}
}

func (c *ListingController) CreateListing(ctx *gin.Context) {
	userId := ctx.Request.Context().Value(middlewares.UserUUIDKey).(string)
	params := vo.ListingCreateRequest{}
	err := ctx.ShouldBindBodyWithJSON(&params)
	if err != nil {
		response.ErrorResponse(ctx, response.UnprocessableEntity, err.Error())
		return
	}
	user := c.userService.GetUserInfo(userId)
	if user == nil {
		response.ErrorResponse(ctx, response.NotFound, "User not found")
		return
	}
	if user.Role != enum.UserRole.Seller {
		response.ErrorResponse(ctx, response.Forbidden, "You are not a seller")
		return
	}
	// TODO: Create listing
	// err = c.listingService.CreateListing(params)
	// if err != nil {
	// 	response.ErrorResponse(ctx, response.InternalServerError, err.Error())
	// 	return
	// }
	response.SuccessResponse(ctx, response.Success, "Create listing successfully")
}

func (c *ListingController) GetListingById(ctx *gin.Context) {
	listingId := ctx.Param("id")
	listing, err := c.listingService.GetListingById(listingId)
	if err != nil {
		response.ErrorResponse(ctx, response.InternalServerError, err.Error())
		return
	}
	if listing == nil {
		response.ErrorResponse(ctx, response.NotFound, "Listing not found")
		return
	}
	response.SuccessResponse(ctx, response.Success, listing)
}

func (c *ListingController) GetListingsByUserId(ctx *gin.Context) {
	userId := ctx.Param("userId")
	listings, err := c.listingService.GetListingsByUserId(userId)
	if err != nil {
		response.ErrorResponse(ctx, response.InternalServerError, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.Success, listings)
}
