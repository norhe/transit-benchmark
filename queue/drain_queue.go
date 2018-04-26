package queue

import (
	"log"

	"github.com/norhe/transit-benchmark/exec_work"
	"github.com/norhe/transit-benchmark/stats"
	"github.com/norhe/transit-benchmark/utils"
	"github.com/norhe/transit-benchmark/vault"
	"github.com/streadway/amqp"
)

// DrainQueueTransit : When a test is run it will drain a queue to find messages to send
//func DrainQueueTransit(queueAddr, vaultAddr, vaultToken, transitKeyName string) {
func DrainQueueTransit(queueAddr string, vCfg vault.Config) {
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

	count := 0

	go func() {
		for msg := range msgs {
			execwork.ExecuteWorkUnit(vCfg, msg.Body)
			count++

			if count%100 == 0 {
				log.Printf("Performed %d operations:", count)
				for k, v := range stats.OpStatsMap {
					if v.Count > 0 {
						log.Printf("Calculated %d operations of type %v.  Average duration of operation: %s, max duration: %s, least duration: %s", v.Count, k, v.AverageDuration.String(), v.MaxDuration.String(), v.LeastDuration.String())
					}
				}
			}

			msg.Ack(false)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
