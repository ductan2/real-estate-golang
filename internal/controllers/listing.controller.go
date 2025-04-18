package controllers

import (
	"ecommerce/internal/filters"
	"ecommerce/internal/middlewares"
	"ecommerce/internal/model"
	listingService "ecommerce/internal/services/listing"
	userService "ecommerce/internal/services/user"
	"ecommerce/internal/vo"
	"ecommerce/pkg/enum"
	"ecommerce/pkg/response"
	"strconv"
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

	err = c.listingService.CreateListing(&model.Listing{
		Title:           params.Title,
		Description:     params.Description,
		Price:           params.Price,
		Images:          params.Images,
		Unit:            params.Unit,
		PropertyType:    params.PropertyType,
		Area:            params.Area,
		Bedroom:         params.Bedroom,
		Bathroom:        params.Bathroom,
		Floor:           params.Floor,
		Direction:       params.Direction,
		LegalStatus:     params.LegalStatus,
		IsForRent:       params.IsForRent,
		VideoURL:        &params.VideoURL,
		DurationListing: params.DurationListing,
		IsPublished:     params.IsPublished,
		StartDate:       params.StartDate,
		EndDate:         params.EndDate,
	})
	if err != nil {
		response.ErrorResponse(ctx, response.InternalServerError, err.Error())
		return
	}
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

func (c *ListingController) GetListingsBySellerId(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	sellerId := ctx.Param("sellerId")
	listings, total, err := c.listingService.GetListingsBySellerId(sellerId, page, limit)
	if err != nil {
		response.ErrorResponse(ctx, response.InternalServerError, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.Success, gin.H{
		"listings": listings,
		"total":    total,
		"page":     page,
		"limit":    limit,
	})
}

func (c *ListingController) GetAllListings(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	filters := &filters.ListingFilter{}
	
	listings, total, err := c.listingService.GetAllListings(page, limit, filters)
	if err != nil {
		response.ErrorResponse(ctx, response.InternalServerError, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.Success, gin.H{
		"listings": listings,
		"total":    total,
		"page":     page,
		"limit":    limit,
	})
}

func (c *ListingController) UpdateListing(ctx *gin.Context) {
	listingId := ctx.Param("id")
	listing, err := c.listingService.GetListingById(listingId)
	if err != nil {
		response.ErrorResponse(ctx, response.InternalServerError, err.Error())
		return
	}
	params := vo.ListingUpdateRequest{}
	err = ctx.ShouldBindBodyWithJSON(&params)
	if err != nil {
		response.ErrorResponse(ctx, response.UnprocessableEntity, err.Error())
		return
	}
	if listing == nil {
		response.ErrorResponse(ctx, response.NotFound, "Listing not found")
		return
	}
	err = c.listingService.UpdateListing(listingId, &model.Listing{
		Title:           params.Title,
		Description:     params.Description,
		Price:           params.Price,
		Images:          params.Images,
		Unit:            params.Unit,
		PropertyType:    params.PropertyType,
		Area:            params.Area,
		Bedroom:         params.Bedroom,
		Bathroom:        params.Bathroom,
		Floor:           params.Floor,
		Direction:       params.Direction,
		LegalStatus:     params.LegalStatus,
		IsForRent:       params.IsForRent,
		VideoURL:        &params.VideoURL,
		DurationListing: params.DurationListing,
		IsPublished:     params.IsPublished,
		StartDate:       params.StartDate,
		EndDate:         params.EndDate,
	})
	if err != nil {
		response.ErrorResponse(ctx, response.InternalServerError, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.Success, "Update listing successfully")
}	

func (c *ListingController) DeleteListing(ctx *gin.Context) {
	listingId := ctx.Param("id")
	err := c.listingService.DeleteListing(listingId)
	if err != nil {
		response.ErrorResponse(ctx, response.InternalServerError, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.Success, "Delete listing successfully")
}
