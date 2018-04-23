package queue

import (
	"github.com/norhe/transit-benchmark/utils"
	"github.com/streadway/amqp"
)

type benchmarkQueue struct {
	Name             string
	Durable          bool
	DeleteWhenUnused bool
	Exclusive        bool
	NoWait           bool
	Args             amqp.Table
}

type resultsQueue struct {
	Name             string
	Durable          bool
	DeleteWhenUnused bool
	Exclusive        bool
	NoWait           bool
	Args             amqp.Table
}

// BQueue : Queue properties for the test queue
var BQueue benchmarkQueue

// RQueue : Queue properties for the results queue
var RQueue resultsQueue

func init() {
	// benchmark queue properties
	BQueue.Name = "benchmark"
	BQueue.Durable = true
	BQueue.DeleteWhenUnused = false
	BQueue.Exclusive = false
	BQueue.NoWait = false
	BQueue.Args = nil

	RQueue.Name = "results"
	RQueue.Durable = true
	RQueue.DeleteWhenUnused = false
	RQueue.Exclusive = false
	RQueue.NoWait = false
	RQueue.Args = nil
}

func connectToQueue(queueAddr string) *amqp.Channel {
	conn, err := amqp.Dial(queueAddr)
	utils.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	utils.FailOnError(err, "Failed to open a channel")

	return ch
}

func declareQueue(ch *amqp.Channel) amqp.Queue {
	q, err := ch.QueueDeclare(
		BQueue.Name,             // name
		BQueue.Durable,          // durable
		BQueue.DeleteWhenUnused, // delete when unused
		BQueue.Exclusive,        // exclusive
		BQueue.NoWait,           // no-wait
		BQueue.Args,             // arguments
	)
	utils.FailOnError(err, "Failed to declare a queue")

	return q
}
