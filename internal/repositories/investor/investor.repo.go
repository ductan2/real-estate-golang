package investor

import (
	"ecommerce/internal/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IInvestorRepository interface {
	Create(investor *model.Investor) error
	GetById(id string) (*model.Investor, error)
	GetAll() ([]model.Investor, error)
	Update(id string, investor *model.Investor) error
	Delete(id string) error
	GetInvestorByUserId(userId string) ([]model.Investor, error)
}

type investorRepository struct {
	db *gorm.DB
}


func NewInvestorRepository(db *gorm.DB) IInvestorRepository {
	return &investorRepository{db: db}
}
func (r *investorRepository) GetInvestorByUserId(userId string) ([]model.Investor, error) {
	var investors []model.Investor
	err := r.db.Where("user_id = ?", userId).Find(&investors).Error
	if err != nil {
		return nil, err
	}
	return investors, nil
}

func (r *investorRepository) Create(investor *model.Investor) error {
	return r.db.Create(investor).Error
}

func (r *investorRepository) GetById(id string) (*model.Investor, error) {
	var investor model.Investor
	err := r.db.Where("id = ?", id).First(&investor).Error
	if err != nil {
		return nil, err
	}
	return &investor, nil
}

func (r *investorRepository) GetAll() ([]model.Investor, error) {
	var investors []model.Investor
	err := r.db.Find(&investors).Error
	if err != nil {
		return nil, err
	}
	return investors, nil
}

func (r *investorRepository) Update(id string, investor *model.Investor) error {
	investor.ID = uuid.MustParse(id)
	return r.db.Save(investor).Error
}

func (r *investorRepository) Delete(id string) error {
	return r.db.Delete(&model.Investor{}, "id = ?", id).Error
}
