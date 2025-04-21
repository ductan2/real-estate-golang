package controllers

import (
	"ecommerce/internal/filters"
	"ecommerce/internal/middlewares"
	"ecommerce/internal/model"
	investorService "ecommerce/internal/services/investor"
	"ecommerce/internal/vo"
	"ecommerce/pkg/response"
	"strconv"

	userService "ecommerce/internal/services/user"

	"github.com/gin-gonic/gin"
)

type InvestorController struct {
	investorService investorService.IInvestorService
	userService     userService.IUserService
}

func NewInvestorController(investorService investorService.IInvestorService, userService userService.IUserService) *InvestorController {
	return &InvestorController{
		investorService: investorService,
		userService:     userService,
	}
}

func (c *InvestorController) GetMe(ctx *gin.Context) {
	userId, exists := ctx.Request.Context().Value(middlewares.UserUUIDKey).(string)
	if !exists || userId == "" {
		response.ErrorResponse(ctx, response.Unauthorized, "User not authenticated")
		return
	}
	investors, err := c.investorService.GetInvestorByUserId(userId)
	if err != nil {
		response.ErrorResponse(ctx, response.InternalServerError, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.Success, investors)
}

// Create creates a new investor
func (c *InvestorController) Create(ctx *gin.Context) {
	var req vo.InvestorCreateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(ctx, response.UnprocessableEntity, err.Error())
		return
	}

	userId, exists := ctx.Request.Context().Value(middlewares.UserUUIDKey).(string)
	if !exists || userId == "" {
		response.ErrorResponse(ctx, response.Unauthorized, "User not authenticated")
		return
	}
	user := c.userService.GetUserInfo(userId)
	println("Test2", user)
	if user == nil {
		response.ErrorResponse(ctx, response.NotFound, "User not found")
		return
	}

	if err := c.investorService.Create(req, user.ID); err != nil {
		response.ErrorResponse(ctx, response.InternalServerError, err.Error())
		return
	}

	response.SuccessResponse(ctx, response.Success, "Investor created successfully")
}

// GetById gets an investor by ID
func (c *InvestorController) GetById(ctx *gin.Context) {
	id := ctx.Param("id")
	investor, err := c.investorService.GetById(id)
	if err != nil {
		response.ErrorResponse(ctx, response.NotFound, "Investor not found")
		return
	}

	response.SuccessResponse(ctx, response.Success, investor)
}

// GetAll gets all investors
func (c *InvestorController) GetAll(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))

	filter := &filters.InvestorFilter{}

	investors, total, err := c.investorService.GetAll(page, limit, filter)
	if err != nil {
		response.ErrorResponse(ctx, response.InternalServerError, err.Error())
		return
	}

	response.SuccessResponse(ctx, response.Success, gin.H{
		"investors": investors,
		"total":     total,
		"page":      page,
		"limit":     limit,
	})
}

// Update updates an investor
func (c *InvestorController) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var req vo.InvestorUpdateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(ctx, response.UnprocessableEntity, err.Error())
		return
	}

	investor := &model.Investor{
		Name:        req.Name,
		Address:     req.Address,
		Email:       req.Email,
		Phone:       req.Phone,
		Website:     req.Website,
		Description: req.Description,
		Logo:        req.Logo,
		Background:  req.Background,
		Type:        req.Type,
	}

	if err := c.investorService.Update(id, investor); err != nil {
		response.ErrorResponse(ctx, response.InternalServerError, err.Error())
		return
	}

	response.SuccessResponse(ctx, response.Success, "Investor updated successfully")
}

// Delete deletes an investor
func (c *InvestorController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := c.investorService.Delete(id); err != nil {
		response.ErrorResponse(ctx, response.InternalServerError, err.Error())
		return
	}

	response.SuccessResponse(ctx, response.Success, "Investor deleted successfully")
}
