package wire

import (
	"ecommerce/global"
	"ecommerce/internal/controllers"
	repo "ecommerce/internal/repositories/user"
	services "ecommerce/internal/services/user"
)


func InitUserRouterHanlder() (*controllers.UserController, error) {
	iUserRepository := repo.NewUserRepository(global.DB)
	iUserService := services.NewUserService(iUserRepository)
	userController := controllers.NewUserController(iUserService)
	return userController, nil
}

func InitUserService() (services.IUserService, error) {
	iUserRepository := repo.NewUserRepository(global.DB)
	iUserService := services.NewUserService(iUserRepository)
	return iUserService, nil
}

func InitSellerRouterHandler() (*controllers.SellerController, error) {
	iSellerRepository := repo.NewSellerRepository(global.DB)
	iSellerService := services.NewSellerService(iSellerRepository)
	sellerController := controllers.NewSellerController(iSellerService)
	return sellerController, nil
}

func InitAdminRouterHandler() (*controllers.AdminController, error) {
	iAdminRepository := repo.NewAdminRepository(global.DB)
	iAdminService := services.NewAdminService(iAdminRepository)
	iSellerRepository := repo.NewSellerRepository(global.DB)
	iSellerService := services.NewSellerService(iSellerRepository)
	adminController := controllers.NewAdminController(iAdminService, iSellerService)
	return adminController, nil
}

func InitAdminService() (services.IAdminService, error) {
	iAdminRepository := repo.NewAdminRepository(global.DB)
	iAdminService := services.NewAdminService(iAdminRepository)
	return iAdminService, nil
}
