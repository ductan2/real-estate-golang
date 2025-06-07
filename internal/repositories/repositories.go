package repositories

import (
	"ecommerce/global"
	investorRepo "ecommerce/internal/repositories/investor"
	listingRepo "ecommerce/internal/repositories/listing"
	projectRepo "ecommerce/internal/repositories/project"
	userRepo "ecommerce/internal/repositories/user"
)

type Repositories struct {
	User     userRepo.IUserRepository
	Admin    userRepo.IAdminRepository
	Seller   userRepo.ISellerRepository
	Project  projectRepo.IProjectRepository
	Listing  listingRepo.IListingRepository
	Investor investorRepo.IInvestorRepository
	UserSession userRepo.IUserSessionRepository
}

func NewRepositories() *Repositories {
	return &Repositories{
		User:     userRepo.NewUserRepository(global.DB),
		Admin:    userRepo.NewAdminRepository(global.DB),
		Seller:   userRepo.NewSellerRepository(global.DB),
		Project:  projectRepo.NewProjectRepository(global.DB),
		Listing:  listingRepo.NewListingRepository(global.DB),
		Investor: investorRepo.NewInvestorRepository(global.DB),
		UserSession: userRepo.NewUserSessionRepository(global.DB),
	}
}
