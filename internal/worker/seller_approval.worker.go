package worker

import (
	"ecommerce/global"
	"ecommerce/internal/services/queue"
	mail "ecommerce/internal/utils/email"
	"encoding/json"
	"log"
)

type SellerApprovalWorker struct {
	queue queue.IQueueService
}

func NewSellerApprovalWorker(queue queue.IQueueService) *SellerApprovalWorker {
	return &SellerApprovalWorker{
		queue: queue,
	}
}


func (w *SellerApprovalWorker) Start() {
	// Create queue and bind to exchange
	queueName := "project_tasks"
	exchange := "seller_approval_exchange"
	routingKey := "user.seller.approval"

	err := w.queue.CreateQueueAndBind(exchange, queueName, routingKey)
	if err != nil {
		log.Fatalf("Failed to create queue and bind: %v", err)
	}

	msgs, err := w.queue.ConsumeMessages(queueName)
	if err != nil {
		log.Fatalf("Failed to consume messages: %v", err)
	}

	log.Printf("Started consuming messages from queue: %s", queueName)

	for msg := range msgs {
		var task queue.ProjectTask
		if err := json.Unmarshal(msg.Body, &task); err != nil {
			log.Printf("Failed to unmarshal message: %v", err)
			continue
		}

		// Get seller's email from database
		var seller struct {
			Email string
		}
		if err := global.DB.Table("users").Select("email").Where("id = ?", task.ProjectID).First(&seller).Error; err != nil {
			log.Printf("Failed to get seller email: %v", err)
			continue
		}

		// Send email notification
		subject := "Seller Approval Notification"
		body := "Congratulations! Your seller account has been approved. You can now start selling on our platform."

		err = mail.SendEmail(
			[]string{seller.Email},
			global.Config.SMTP.Username,
			subject,
			body,
		)
		if err != nil {
			log.Printf("Failed to send email: %v", err)
			continue
		}

		log.Printf("Successfully sent approval email to seller %s at %s", task.ProjectID, seller.Email)
	}
}
