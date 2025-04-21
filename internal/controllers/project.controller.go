package controllers

import (
	"context"
	"ecommerce/global"
	"ecommerce/internal/filters"
	"ecommerce/internal/middlewares"
	"ecommerce/internal/model"
	projectService "ecommerce/internal/services/project"
	"ecommerce/internal/utils/convert"
	"ecommerce/internal/vo"
	"ecommerce/pkg/response"
	"encoding/json"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
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
		middlewares.HandleError(ctx, err, response.UnprocessableEntity)
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
		middlewares.HandleError(ctx, err, response.InternalServerError)
		return
	}

	response.SuccessResponse(ctx, response.Success, "Project created successfully")
}

func (c *ProjectController) GetById(ctx *gin.Context) {
	projectId := ctx.Param("id")
	project, err := c.projectService.GetProjectById(projectId)
	if err != nil {
		middlewares.HandleError(ctx, err, response.InternalServerError)
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
		userId, _ := ctx.Request.Context().Value(middlewares.UserUUIDKey).(string)

		// Use goroutine for Redis operations
		go func() {
			// pipeline use for execute multiple commands in a single round trip to the server
			pipeline := global.Redis.Pipeline()
			if userId != "" {
				userKey := "user_search:" + userId + "_project_name_suggest"
				pipeline.ZIncrBy(context.Background(), userKey, 1, name)
				pipeline.Expire(context.Background(), userKey, 24*time.Hour)
			}

			pipeline.ZIncrBy(context.Background(), "project_name_suggest", 1, name)
			pipeline.Expire(context.Background(), "project_name_suggest", 24*time.Hour)
			_, err := pipeline.Exec(context.Background())
			if err != nil {
				middlewares.HandleError(ctx, err, response.InternalServerError)
				return
			}
		}()
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
	if province := ctx.Query("province"); province != "" {
		filter.Province = &province
	}


	projects, total, err := c.projectService.GetAllProjects(page, limit, filter)
	if err != nil {
		middlewares.HandleError(ctx, err, response.InternalServerError)
		return
	}

	response.SuccessResponse(ctx, response.Success, gin.H{
		"projects": projects,
		"total":    total,
		"page":     page,
		"limit":    limit,
	})
}

func (c *ProjectController) Update(ctx *gin.Context) {
	projectId := ctx.Param("id")
	var req vo.ProjectUpdateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		middlewares.HandleError(ctx, err, response.UnprocessableEntity)
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
		middlewares.HandleError(ctx, err, response.InternalServerError)
		return
	}

	response.SuccessResponse(ctx, response.Success, "Project updated successfully")
}

func (c *ProjectController) Delete(ctx *gin.Context) {
	projectId := ctx.Param("id")
	if err := c.projectService.DeleteProject(projectId); err != nil {
		middlewares.HandleError(ctx, err, response.InternalServerError)
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
		middlewares.HandleError(ctx, err, response.InternalServerError)
		return
	}

	response.SuccessResponse(ctx, response.Success, gin.H{
		"projects": projects,
		"total":    total,
	})
}

func (c *ProjectController) Suggest(ctx *gin.Context) {
	keyword := ctx.Query("keyword")
	suggestions := make([]string, 0)
	seen := make(map[string]bool)

	userID, _ := ctx.Request.Context().Value(middlewares.UserUUIDKey).(string)

	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	var wg sync.WaitGroup
	results := make(chan []string, 2)

	getSearches := func(key string, count int64) {
		defer wg.Done()
		result, err := global.Redis.ZRevRangeByScore(ctxWithTimeout, key, &redis.ZRangeBy{
			Min:    "-inf",
			Max:    "+inf",
			Offset: 0,
			Count:  count,
		}).Result()
		if err == nil {
			results <- result // send result to channel
		} else {
			results <- []string{} // send empty result to channel
		}
	}

	// Start goroutines for both user and global searches
	if userID != "" {
		wg.Add(1) // Increment the WaitGroup counter to indicate a new goroutine is being added for user-specific searches
		go getSearches("user_search:"+userID+"_project_name_suggest", 10)
	}

	wg.Add(1)
	go getSearches("project_name_suggest", 10)

	// Wait for all goroutines to complete
	wg.Wait()
	close(results)

	// Process results
	for searches := range results {
		for _, search := range searches {
			if keyword == "" || convert.ContainsIgnoreCase(search, keyword) {
				if !seen[search] {
					suggestions = append(suggestions, search)
					seen[search] = true
				}
			}
		}
	}

	response.SuccessResponse(ctx, response.Success, suggestions)
}
