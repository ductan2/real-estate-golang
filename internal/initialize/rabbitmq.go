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
	queueService, err := queue.NewQueueService(rabbitmqUrl, global.Config.RabbitMQ.QueueName,"seller_approval_exchange")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}

	global.RabbitMQ = queueService
}
