package worker

import (
	"encoding/json"
	"log"

	"ecommerce/internal/services/notification"
	"ecommerce/internal/services/queue"
)

type NotificationPayload struct {
	Message string `json:"message"`
}

type NotificationWorker struct {
	queue     queue.IQueueService
	service   notification.INotificationService
	queueName string
}

func NewNotificationWorker(q queue.IQueueService, s notification.INotificationService, queueName string) *NotificationWorker {
	return &NotificationWorker{queue: q, service: s, queueName: queueName}
}

func (w *NotificationWorker) Start() error {
	if err := w.queue.CreateQueueAndBind("seller_approval_exchange", w.queueName, "seller.approved"); err != nil {
		return err
	}
	msgs, err := w.queue.ConsumeMessages(w.queueName)
	if err != nil {
		return err
	}
	go func() {
		for m := range msgs {
			var payload NotificationPayload
			if err := json.Unmarshal(m.Body, &payload); err != nil {
				continue
			}
			if err := w.service.NotifyAdmins(payload.Message); err != nil {
				log.Printf("failed to create notification: %v", err)
			}
		}
	}()
	return nil
}
