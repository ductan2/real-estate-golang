package worker

import (
	"ecommerce/internal/services/project"
	"log"
	"time"
)

type CronjobWorker struct {
	projectService project.IProjectService
}

func NewCronjobWorker(projectService project.IProjectService) *CronjobWorker {
	return &CronjobWorker{
		projectService: projectService,
	}
}

func (w *CronjobWorker) Start() {
	// Run the job immediately when starting
	w.updateProjectStatus()

	// Schedule the job to run at midnight every day
	for {
		now := time.Now()
		next := time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, now.Location())
		duration := next.Sub(now)

		time.Sleep(duration)
		w.updateProjectStatus()
	}
}

func (w *CronjobWorker) updateProjectStatus() {
	log.Println("Starting midnight project status update...")

	// Update all expired projects
	updates := map[string]interface{}{
		"status": "expired",
	}

	err := w.projectService.UpdateProjectStatus(updates)
	if err != nil {
		log.Printf("Failed to update project statuses: %v", err)
		return
	}

	log.Println("Successfully completed midnight project status update")
}
