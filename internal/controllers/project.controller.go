package controllers

import (
	"ecommerce/internal/middlewares"
	projectService "ecommerce/internal/services/project"
	userService "ecommerce/internal/services/user"
	"ecommerce/internal/vo"
	"ecommerce/pkg/enum"
	"ecommerce/pkg/response"

	"github.com/gin-gonic/gin"
)

type ProjectController struct {
	projectService projectService.IProjectService
	userService    userService.IUserService
}

func NewProjectController(projectService projectService.IProjectService, userService userService.IUserService) *ProjectController {
	return &ProjectController{
		projectService: projectService,
		userService:    userService,
	}
}

func (c *ProjectController) CreateProject(ctx *gin.Context) {
	userId := ctx.Request.Context().Value(middlewares.UserUUIDKey).(string)
	params := vo.ProjectCreateRequest{}
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
	// err = c.projectService.CreateProject(params)
	// if err != nil {
	// 	response.ErrorResponse(ctx, response.InternalServerError, err.Error())
	// 	return
	// }
	response.SuccessResponse(ctx, response.Success, "Create project successfully")
}
