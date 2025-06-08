package initialize

import (
	"ecommerce/global"
	"ecommerce/internal/services/queue"
	"fmt"
	"log"
)

func InitRabbitMQ() {
	rabbitmqUrl := fmt.Sprintf(
		"amqp://%s:%s@%s:%d/",
		global.Config.RabbitMQ.Username,
		global.Config.RabbitMQ.Password,
		global.Config.RabbitMQ.Host,
		global.Config.RabbitMQ.Port,
	)
	fmt.Println(rabbitmqUrl)
	queueService, err := queue.NewQueueService(rabbitmqUrl)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}

	global.RabbitMQ = queueService

	// ensure default notification queue and exchange exist
	if q, ok := global.Config.RabbitMQ.Queues["notification"]; ok {
		if err := queueService.CreateQueueAndBind("seller_approval_exchange", q, "seller.approved"); err != nil {
			log.Fatalf("Failed to declare queue: %v", err)
		}
	}
}
