package command

import (
	"flag"
	"fmt"

	"github.com/norhe/transit-benchmark/queue"

	"github.com/mitchellh/cli"
)

// SeedCommand : x
type SeedCommand struct {
	NumRecords    int
	QueueAddr     string
	MaxRecordSize int
	Ui            cli.Ui
}

// Run : x
func (c *SeedCommand) Run(args []string) int {
	cmdFlags := flag.NewFlagSet("seed", flag.ContinueOnError)
	cmdFlags.Usage = func() { c.Ui.Output(c.Help()) }

	cmdFlags.IntVar(&c.NumRecords, "num-records", 500, "The number of records to seed for benchmarking")
	cmdFlags.IntVar(&c.MaxRecordSize, "max-record-size", 1024, "The max record size to generate")
	cmdFlags.StringVar(&c.QueueAddr, "queue-addr", "amqp://guest:guest@localhost:5672/", "The rabbitmq addr")
	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}

	c.Ui.Output(fmt.Sprintf("Would seed %d records with max size %d to queue at %s", c.NumRecords, c.MaxRecordSize, c.QueueAddr))

	queue.SeedQueueRandom(c.QueueAddr, c.NumRecords, c.MaxRecordSize)

	return 0
}

// Help : x
func (c *SeedCommand) Help() string {
	return "Generate N number of random messagesd with lengths between 0 and Y"
}

// Synopsis : x
func (c *SeedCommand) Synopsis() string {
	return "Seed messages into the queue"
}
