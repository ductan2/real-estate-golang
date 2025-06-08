package main

import (
	"log"

	"ecommerce/global"
	"ecommerce/internal/initialize"
	repoNotif "ecommerce/internal/repositories/notification"
	repoUser "ecommerce/internal/repositories/user"
	serviceNotif "ecommerce/internal/services/notification"
	"ecommerce/internal/worker"
)

func main() {
	initialize.InitEnv()
	initialize.LoadConfig()
	initialize.InitLogger()
	initialize.InitDB()
	initialize.InitRabbitMQ()

	notifRepo := repoNotif.NewNotificationRepository(global.DB)
	adminRepo := repoUser.NewAdminRepository(global.DB)
	notifService := serviceNotif.NewNotificationService(notifRepo, adminRepo)

	queueName := global.Config.RabbitMQ.Queues["notification"]
	w := worker.NewNotificationWorker(global.RabbitMQ, notifService, queueName)
	if err := w.Start(); err != nil {
		log.Fatalf("failed to start notification worker: %v", err)
	}

	select {}
}
