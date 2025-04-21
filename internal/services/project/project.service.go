package project

import (
	"ecommerce/internal/filters"
	"ecommerce/internal/model"
	"ecommerce/internal/repositories/project"
	"errors"
)

type ProjectFilter struct {
	Name       *string
	Status     *string
	IsPublish  *bool
	InvestorID *string
}

type IProjectService interface {
	CreateProject(payload *model.Project) error
	GetProjectById(projectId string) (*model.Project, error)
	GetAllProjects(page int, limit int, filter *filters.ProjectFilter) ([]*model.Project, int64, error)
	UpdateProject(projectId string, updates map[string]interface{}) error
	DeleteProject(projectId string) error
	GetProjectsByInvestor(investorId string, page, limit int) ([]*model.Project, int64, error)
	UpdateProjectStatus(updates map[string]interface{}) error
}

type projectService struct {
	projectRepo project.IProjectRepository
}

// CreateProject implements IProjectService.
func (s *projectService) CreateProject(payload *model.Project) error {
	return s.projectRepo.CreateProject(payload)
}

func (s *projectService) GetProjectById(projectId string) (*model.Project, error) {
	project, err := s.projectRepo.GetProjectById(projectId)
	if err != nil {
		return nil, err
	}
	if project == nil {
		return nil, errors.New("project not found")
	}
	return project, nil
}

func (s *projectService) GetAllProjects(page int, limit int, filter *filters.ProjectFilter) ([]*model.Project, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	return s.projectRepo.GetAllProjects(page, limit, filter)
}

func (s *projectService) UpdateProject(projectId string, updates map[string]interface{}) error {
	if name, ok := updates["name"].(string); ok && name == "" {
		return errors.New("project name cannot be empty")
	}
	if areaLand, ok := updates["area_land"].(float64); ok && areaLand <= 0 {
		return errors.New("land area must be greater than 0")
	}
	return s.projectRepo.UpdateProject(projectId, updates)
}

func (s *projectService) DeleteProject(projectId string) error {
	return s.projectRepo.DeleteProject(projectId)
}

func (s *projectService) GetProjectsByInvestor(investorId string, page, limit int) ([]*model.Project, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	return s.projectRepo.GetProjectsByInvestor(investorId, page, limit)
}

func (s *projectService) UpdateProjectStatus(updates map[string]interface{}) error {
	return s.projectRepo.UpdateProjectStatus(updates)
}

func NewProjectService(projectRepo project.IProjectRepository) IProjectService {
	return &projectService{projectRepo: projectRepo}
}
