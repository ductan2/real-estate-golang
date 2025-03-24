package services

import (
	"ecommerce/internal/model"
	repo "ecommerce/internal/repositories/project"

)

type IProjectService interface {
	CreateProject(payload *model.Project) error
}

type projectService struct {
	projectRepo repo.IProjectRepository
}

// CreateProject implements IProjectService.
func (s *projectService) CreateProject(payload *model.Project) error {
	err := s.projectRepo.CreateProject(payload)
	if err != nil {
		return err
	}
	return nil
}

func NewProjectService(projectRepo repo.IProjectRepository) IProjectService {
	return &projectService{projectRepo: projectRepo}
}
