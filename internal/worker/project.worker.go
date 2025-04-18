package worker

import (
	"ecommerce/internal/services/project"
	"ecommerce/internal/services/queue"
	"log"
	"time"
)

type ProjectWorker struct {
	queue          queue.IQueueService
	projectService project.IProjectService
}

func NewProjectWorker(queue queue.IQueueService, projectService project.IProjectService) *ProjectWorker {
	return &ProjectWorker{
		queue:          queue,
		projectService: projectService,
	}
}

func (w *ProjectWorker) Start() {
	tasks, err := w.queue.ConsumeProjectTasks()
	if err != nil {
		log.Fatalf("Failed to consume tasks: %v", err)
	}

	for task := range tasks {
		// Check if it's time to execute the task
		if time.Now().Before(task.ExecuteAt) {
			// Re-queue the task if it's not time yet
			time.Sleep(time.Until(task.ExecuteAt))
		}

		// Process the task
		switch task.Action {
		case "expire":
			updates := map[string]interface{}{
				"status": "expired",
			}
			if err := w.projectService.UpdateProject(task.ProjectID, updates); err != nil {
				log.Printf("Failed to expire project %s: %v", task.ProjectID, err)
				continue
			}
			log.Printf("Successfully expired project %s", task.ProjectID)
		}
	}
}
