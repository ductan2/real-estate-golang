package wire

import (
	"ecommerce/global"
	"ecommerce/internal/controllers"
	repo "ecommerce/internal/repositories/investor"
	services "ecommerce/internal/services/investor"
)


func InitInvestorRouterHanlder() (*controllers.InvestorController, error) {
	iUserService, err := InitUserService()
	if err != nil {
		return nil, err
	}
	iInvestorRepository := repo.NewInvestorRepository(global.DB)
	iInvestorService := services.NewInvestorService(iInvestorRepository)
	investorController := controllers.NewInvestorController(iInvestorService, iUserService)
	return investorController, nil
}