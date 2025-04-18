package main

import (
	"ecommerce/global"
	"ecommerce/internal/initialize"
	"ecommerce/internal/worker"
	"log"
)

func main() {
	initialize.InitEnv()
	initialize.LoadConfig()

	initialize.InitRabbitMQ()

	sellerApprovalWorker := worker.NewSellerApprovalWorker(global.RabbitMQ)

	log.Println("Starting seller approval worker...")
	sellerApprovalWorker.Start()
}
