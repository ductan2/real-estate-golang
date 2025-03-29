package wire

import (
	"ecommerce/internal/controllers"
	repo "ecommerce/internal/repositories"
	services "ecommerce/internal/services/project"
)


func InitProjectRouterHanlder() (*controllers.ProjectController, error) {
	iProjectRepository := repo.NewRepositories().Project
	iProjectService := services.NewProjectService(iProjectRepository)
	projectController := controllers.NewProjectController(iProjectService)
	return projectController, nil
}