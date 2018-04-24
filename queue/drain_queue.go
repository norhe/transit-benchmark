package queue

import (
	"log"

	"github.com/norhe/transit-benchmark/utils"
	"github.com/streadway/amqp"
)

// DrainQueueTransit : When a test is run it will drain a queue to find messages to send
func DrainQueueTransit(queueAddr, vaultAddr, vaultToken, transitKeyName string) {
	conn, err := amqp.Dial(queueAddr)
	utils.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	utils.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		BQueue.Name,             // name
		BQueue.Durable,          // durable
		BQueue.DeleteWhenUnused, // delete when unused
		BQueue.Exclusive,        // exclusive
		BQueue.NoWait,           // no-wait
		BQueue.Args,             // arguments
	)
	utils.FailOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	utils.FailOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for msg := range msgs {
			exec_work.ExecuteWorkUnit(vaultAddr, vaultToken, keyName, msg.Body)
			//log.Printf("Received a message: %s", msg.Body)
			msg.Ack(false)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
