package controllers

import (
	"ecommerce/internal/filters"
	"ecommerce/internal/model"
	projectService "ecommerce/internal/services/project"
	"ecommerce/internal/vo"
	"ecommerce/pkg/response"
	"encoding/json"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ProjectController struct {
	projectService projectService.IProjectService
}

func NewProjectController(projectService projectService.IProjectService) *ProjectController {
	return &ProjectController{
		projectService: projectService,
	}
}

func (c *ProjectController) Create(ctx *gin.Context) {
	var req vo.ProjectCreateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(ctx, response.UnprocessableEntity, err.Error())
		return
	}
	imagesJson, _ := json.Marshal(req.Images)

	project := &model.Project{
		Name:            req.Name,
		Description:     req.Description,
		LongDescription: req.LongDescription,
		Status:          req.Status,
		AreaLand:        req.AreaLand,
		ProjectType:     req.ProjectType,
		Images:          imagesJson,
		Address:         req.Address,
		LegalStatus:     req.LegalStatus,
		StartDate:       req.StartDate,
		EndDate:         req.EndDate,
		Apartment:       req.Apartment,
		InvestorID:      uuid.MustParse(req.InvestorID),
	}

	if err := c.projectService.CreateProject(project); err != nil {
		response.ErrorResponse(ctx, response.InternalServerError, err.Error())
		return
	}

	response.SuccessResponse(ctx, response.Success, "Project created successfully")
}

func (c *ProjectController) GetById(ctx *gin.Context) {
	projectId := ctx.Param("id")
	project, err := c.projectService.GetProjectById(projectId)
	if err != nil {
		response.ErrorResponse(ctx, response.InternalServerError, err.Error())
		return
	}

	response.SuccessResponse(ctx, response.Success, project)
}

func (c *ProjectController) GetAll(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))

	// Create filter struct
	filter := &filters.ProjectFilter{}

	// Get optional filters from query parameters
	if name := ctx.Query("name"); name != "" {
		filter.Name = &name
	}
	if status := ctx.Query("status"); status != "" {
		filter.Status = &status
	}
	if isPublish := ctx.Query("is_publish"); isPublish != "" {
		isPublishBool, _ := strconv.ParseBool(isPublish)
		filter.IsPublish = &isPublishBool
	}
	if investorID := ctx.Query("investor_id"); investorID != "" {
		filter.InvestorID = &investorID
	}

	projects, total, err := c.projectService.GetAllProjects(page, limit, filter)
	if err != nil {
		response.ErrorResponse(ctx, response.InternalServerError, err.Error())
		return
	}

	response.SuccessResponse(ctx, response.Success, gin.H{
		"projects": projects,
		"total":    total,
	})
}

func (c *ProjectController) Update(ctx *gin.Context) {
	projectId := ctx.Param("id")
	var req vo.ProjectUpdateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(ctx, response.UnprocessableEntity, err.Error())
		return
	}

	updates := map[string]interface{}{
		"name":             req.Name,
		"description":      req.Description,
		"long_description": req.LongDescription,
		"status":           req.Status,
		"area_land":        req.AreaLand,
		"area_build":       req.AreaBuild,
		"project_type":     req.ProjectType,
		"images":           req.Images,
		"address":          req.Address,
		"apartment":        req.Apartment,
	}

	if err := c.projectService.UpdateProject(projectId, updates); err != nil {
		response.ErrorResponse(ctx, response.InternalServerError, err.Error())
		return
	}

	response.SuccessResponse(ctx, response.Success, "Project updated successfully")
}

func (c *ProjectController) Delete(ctx *gin.Context) {
	projectId := ctx.Param("id")
	if err := c.projectService.DeleteProject(projectId); err != nil {
		response.ErrorResponse(ctx, response.InternalServerError, err.Error())
		return
	}

	response.SuccessResponse(ctx, response.Success, "Project deleted successfully")
}

func (c *ProjectController) GetByInvestor(ctx *gin.Context) {
	investorId := ctx.Param("investorId")
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))

	projects, total, err := c.projectService.GetProjectsByInvestor(investorId, page, limit)
	if err != nil {
		response.ErrorResponse(ctx, response.InternalServerError, err.Error())
		return
	}

	response.SuccessResponse(ctx, response.Success, gin.H{
		"projects": projects,
		"total":    total,
	})
}
