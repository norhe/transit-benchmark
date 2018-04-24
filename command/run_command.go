package command

import (
	"flag"
	"fmt"

	"github.com/mitchellh/cli"
	"github.com/norhe/transit-benchmark/queue"
)

// RunCommand : x
type RunCommand struct {
	Ui             cli.Ui
	QueueAddr      string
	ShouldBatch    bool
	VaultAddr      string
	VaultToken     string
	TransitKeyName string
}

// Run : x
func (c *RunCommand) Run(args []string) int {
	cmdFlags := flag.NewFlagSet("run", flag.ContinueOnError)
	cmdFlags.Usage = func() { c.Ui.Output(c.Help()) }

	cmdFlags.StringVar(&c.QueueAddr, "queue-addr", "amqp://guest:guest@localhost:5672/", "The rabbitmq addr")
	cmdFlags.StringVar(&c.VaultAddr, "vault-addr", "http://localhost:8200", "The Vault server address")
	cmdFlags.StringVar(&c.VaultToken, "vault-token", "", "The token to use when authenticating to Vault")
	cmdFlags.StringVar(&c.TransitKeyName, "transit-key", "benchmark", "The transit key to use.  This key must already exist on the Vault server")
	cmdFlags.BoolVar(&c.ShouldBatch, "batch", false, "Should transit messages be batched")

	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}

	c.Ui.Output(fmt.Sprintf("Would connect to queue at %s and send messages to %s Vault server for key %s", c.QueueAddr, c.VaultAddr, c.TransitKeyName))

	//queue.SeedQueueRandom(c.QueueAddr, c.NumRecords, c.MaxRecordSize)
	queue.DrainQueue(c.QueueAddr, c.VaultAddr, c.VaultToken, c.TransitKeyName)

	return 0
}

// Help : x
func (c *RunCommand) Help() string {
	return "Execute tests "
}

// Synopsis : x
func (c *RunCommand) Synopsis() string {
	return "Execute tests by draining messages from the Queue"
}
