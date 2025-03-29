package project

import (
	"ecommerce/internal/filters"
	"ecommerce/internal/model"
	"errors"

	"gorm.io/gorm"
)

type IProjectRepository interface {
	CreateProject(payload *model.Project) error
	GetProjectById(projectId string) (*model.Project, error)
	GetAllProjects(page int, limit int, filter *filters.ProjectFilter) ([]*model.Project, int64, error)
	UpdateProject(projectId string, updates map[string]interface{}) error
	DeleteProject(projectId string) error
	GetProjectsByInvestor(investorId string, page, limit int) ([]*model.Project, int64, error)
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
	err := s.db.Preload("Investor").First(&project, "id = ?", projectId).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &project, nil
}

func (s *projectRepository) GetAllProjects(page int, limit int, filter *filters.ProjectFilter) ([]*model.Project, int64, error) {
	var projects []*model.Project
	var total int64

	offset := (page - 1) * limit
	query := s.db.Model(&model.Project{})

	// Apply filters if they are provided
	if filter != nil {
		if filter.Name != nil && *filter.Name != "" {
			query = query.Where("name ILIKE ?", "%"+*filter.Name+"%")
		}
		if filter.Status != nil && *filter.Status != "" {
			query = query.Where("status = ?", *filter.Status)
		}
		if filter.IsPublish != nil {
			query = query.Where("is_publish = ?", *filter.IsPublish)
		}
		if filter.InvestorID != nil && *filter.InvestorID != "" {
			query = query.Where("investor_id = ?", *filter.InvestorID)
		}
	}

	// Get total count
	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// Get paginated results
	err = query.Preload("Investor").
		Offset(offset).
		Limit(limit).
		Order("created_at DESC").
		Find(&projects).Error
	if err != nil {
		return nil, 0, err
	}

	return projects, total, nil
}

func (s *projectRepository) UpdateProject(projectId string, updates map[string]interface{}) error {
	result := s.db.Model(&model.Project{}).Where("id = ?", projectId).Updates(updates)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("project not found")
	}
	return nil
}

func (s *projectRepository) DeleteProject(projectId string) error {
	result := s.db.Delete(&model.Project{}, "id = ?", projectId)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("project not found")
	}
	return nil
}

func (s *projectRepository) GetProjectsByInvestor(investorId string, page, limit int) ([]*model.Project, int64, error) {
	var projects []*model.Project
	var total int64

	offset := (page - 1) * limit
	err := s.db.Model(&model.Project{}).Where("investor_id = ?", investorId).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = s.db.Preload("Investor").
		Where("investor_id = ?", investorId).
		Offset(offset).
		Limit(limit).
		Order("created_at DESC").
		Find(&projects).Error
	if err != nil {
		return nil, 0, err
	}

	return projects, total, nil
}

func NewProjectRepository(db *gorm.DB) IProjectRepository {
	return &projectRepository{db: db}
}
