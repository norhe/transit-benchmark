package queue

import (
	"log"
	"math/rand"

	"github.com/norhe/transit-benchmark/utils"
	"github.com/norhe/transit-benchmark/workunit"
	"github.com/streadway/amqp"
)

// SeedQueueRandom : seeds random messages to be transitted.  Should work with OperationTypes Encrypt, SignData, HashData
func SeedQueueRandom(queueAddr string, opType workunit.OperationType, numRecords, maxRecordSize int) {
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

	// Seed the queue
	for n := 0; n <= numRecords; n++ {
		pload := []byte(utils.RandSeq(rand.Intn(maxRecordSize)))
		wu := workunit.WorkUnit{
			Operation:   opType,
			Payload:     pload,
			PayloadSize: len(pload),
		}

		err := ch.Publish(
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				ContentType: "text/json",
				Body:        workunit.ToJSON(wu),
			})
		utils.FailOnError(err, "Failed to publish a message")
	}

	log.Printf("Seeded the queue with %d messages with max length %d", numRecords, maxRecordSize)
}
