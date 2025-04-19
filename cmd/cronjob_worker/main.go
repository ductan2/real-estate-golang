package main

import (
	"ecommerce/internal/initialize"
	"ecommerce/internal/repositories"
	"ecommerce/internal/services/project"
	"ecommerce/internal/worker"
	"log"
)

func main() {
	initialize.InitEnv()
	initialize.LoadConfig()
	initialize.InitDB()

	// Create repositories and services
	repos := repositories.NewRepositories()
	projectService := project.NewProjectService(repos.Project)

	// Create and start cronjob worker
	cronjobWorker := worker.NewCronjobWorker(projectService)

	log.Println("Starting cronjob worker...")
	cronjobWorker.Start()
}
