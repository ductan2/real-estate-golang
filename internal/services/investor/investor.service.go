package investor

import (
	"ecommerce/internal/filters"
	"ecommerce/internal/model"
	repo "ecommerce/internal/repositories/investor"
	"ecommerce/internal/vo"

	"github.com/google/uuid"
)

type IInvestorService interface {
	Create(investor vo.InvestorCreateRequest, userId uuid.UUID) error
	GetById(id string) (*model.Investor, error)
	GetAll(page int, limit int, filter *filters.InvestorFilter) ([]model.Investor, int64, error)
	Update(id string, investor *model.Investor) error
	Delete(id string) error
	GetInvestorByUserId(userId string) ([]model.Investor, error)
}

type investorService struct {
	investorRepo repo.IInvestorRepository
}

func NewInvestorService(investorRepo repo.IInvestorRepository) IInvestorService {
	return &investorService{
		investorRepo: investorRepo,
	}
}

func (s *investorService) GetInvestorByUserId(userId string) ([]model.Investor, error) {
	return s.investorRepo.GetInvestorByUserId(userId)
}

func (s *investorService) Create(investor vo.InvestorCreateRequest, userId uuid.UUID) error {
	return s.investorRepo.Create(&model.Investor{
		Name:        investor.Name,
		Address:     investor.Address,
		Email:       investor.Email,
		Phone:       investor.Phone,
		Website:     investor.Website,
		Description: investor.Description,
		Logo:        investor.Logo,
		Background:  investor.Background,
		Type:        investor.Type,
		UserId:      userId,
	})
}

func (s *investorService) GetById(id string) (*model.Investor, error) {
	return s.investorRepo.GetById(id)
}

func (s *investorService) GetAll(page int, limit int, filter *filters.InvestorFilter) ([]model.Investor, int64, error) {
	return s.investorRepo.GetAll(page, limit, filter)
}

func (s *investorService) Update(id string, investor *model.Investor) error {
	return s.investorRepo.Update(id, investor)
}

func (s *investorService) Delete(id string) error {
	return s.investorRepo.Delete(id)
}
