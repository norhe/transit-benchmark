package queue

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"

	"github.com/norhe/transit-benchmark/utils"
	"github.com/streadway/amqp"
)

// pass in NUM_RECORDS and MAX_RECORD_SIZE as env vars
func getEnv() (int, int) {
	num_records, err := strconv.Atoi(os.Getenv("NUM_RECORDS"))
	utils.FailOnError(err, "Couldn't retrieve NUMBER_RECORDS")

	max_size, err := strconv.Atoi(os.Getenv("MAX_RECORD_SIZE"))
	utils.FailOnError(err, "Couldn't retrieve MAX_RECORD_SIZE")
	return num_records, max_size
}

// pass in NUM_RECORDS and MAX_RECORD_SIZE as env vars
func SeedQueueRandom() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
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

	num_records, max_size := getEnv()

	// Seed the queue
	for n := 0; n <= num_records; n++ {
		err = ch.Publish(
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(utils.RandSeq(rand.Intn(max_size))),
			})
		utils.FailOnError(err, "Failed to publish a message")
		fmt.Println(n)
	}

	log.Printf("Seeded the queue with %d messages with max length %d", num_records, max_size)

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
