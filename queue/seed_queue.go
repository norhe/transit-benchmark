package queue

import (
	"log"
	"math/rand"

	"github.com/norhe/transit-benchmark/utils"
	"github.com/streadway/amqp"
)

// SeedQueueRandom : seeds random messages to be transitted
func SeedQueueRandom(queueAddr string, numRecords, maxRecordSize int) {
	conn, err := amqp.Dial(queueAddr)
	utils.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	utils.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"benchmark", // name
		false,       // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	utils.FailOnError(err, "Failed to declare a queue")

	// Seed the queue
	for n := 0; n <= numRecords; n++ {
		err = ch.Publish(
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(utils.RandSeq(rand.Intn(maxRecordSize))),
			})
		utils.FailOnError(err, "Failed to publish a message")
	}

	log.Printf("Seeded the queue with %d messages with max length %d", numRecords, maxRecordSize)

	/*msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	utils.FailOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever*/
}
