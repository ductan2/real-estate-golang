package repo

import (
	"ecommerce/internal/model"
	"errors"

	"gorm.io/gorm"
)

type IProjectRepository interface {
	CreateProject(payload *model.Project) error
	GetProjectById(projectId string) (*model.Project, error)
}

type projectRepository struct {
	db *gorm.DB
}

// CreateProject implements IProjectRepository.
func (s *projectRepository) CreateProject(payload *model.Project) error {
	return s.db.Create(&payload).Error
}

// GetSellerById implements ISellerRepository.
func (s *projectRepository) GetProjectById(projectId string) (*model.Project, error) {
	var project model.Project
	err := s.db.First(&project, "id = ?", projectId).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &project, nil
}

func NewProjectRepository(db *gorm.DB) IProjectRepository {
	return &projectRepository{db: db}
}
