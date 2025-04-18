package investor

import (
	"ecommerce/internal/filters"
	"ecommerce/internal/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IInvestorRepository interface {
	Create(investor *model.Investor) error
	GetById(id string) (*model.Investor, error)
	GetAll(page int, limit int, filter *filters.InvestorFilter) ([]model.Investor, int64, error)
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

func (r *investorRepository) GetAll(page int, limit int, filter *filters.InvestorFilter) ([]model.Investor, int64, error) {
	var investors []model.Investor
	var total int64
	if filter.Email != nil {
		r.db = r.db.Where("email = ?", *filter.Email)
	}
	if filter.Phone != nil {
		r.db = r.db.Where("phone = ?", *filter.Phone)
	}
	if filter.Address != nil {
		r.db = r.db.Where("address LIKE ?", "%"+*filter.Address+"%")
	}
	if filter.Website != nil {
		r.db = r.db.Where("website LIKE ?", "%"+*filter.Website+"%")
	}
	err := r.db.Offset((page - 1) * limit).Limit(limit).Find(&investors).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	return investors, total, nil
}

func (r *investorRepository) Update(id string, investor *model.Investor) error {
	investor.ID = uuid.MustParse(id)
	return r.db.Save(investor).Error
}

func (r *investorRepository) Delete(id string) error {
	return r.db.Delete(&model.Investor{}, "id = ?", id).Error
}
