package command

import (
	"flag"
	"fmt"

	"github.com/mitchellh/cli"
	"github.com/norhe/transit-benchmark/queue"
	"github.com/norhe/transit-benchmark/vault"
)

// RunCommand : x
type RunCommand struct {
	UI             cli.Ui
	QueueAddr      string
	ShouldBatch    bool
	VaultAddr      string
	VaultToken     string
	TransitKeyName string
}

// Run : x
func (c *RunCommand) Run(args []string) int {
	cmdFlags := flag.NewFlagSet("run", flag.ContinueOnError)
	cmdFlags.Usage = func() { c.UI.Output(c.Help()) }

	cmdFlags.StringVar(&c.QueueAddr, "queue-addr", "amqp://guest:guest@localhost:5672/", "The rabbitmq addr")
	cmdFlags.StringVar(&c.VaultAddr, "vault-addr", "http://localhost:8200", "The Vault server address")
	cmdFlags.StringVar(&c.VaultToken, "vault-token", "root", "The token to use when authenticating to Vault")
	cmdFlags.StringVar(&c.TransitKeyName, "transit-key", "benchmark", "The transit key to use.  This key must already exist on the Vault server")
	cmdFlags.BoolVar(&c.ShouldBatch, "batch", false, "Should transit messages be batched")

	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}

	c.UI.Output(fmt.Sprintf("Connecting to queue at %s and sending messages to Vault server %s for key %s", c.QueueAddr, c.VaultAddr, c.TransitKeyName))

	vCfg := vault.Config{
		Address:        c.VaultAddr,
		Token:          c.VaultToken,
		TransitKeyName: c.TransitKeyName,
	}

	//queue.SeedQueueRandom(c.QueueAddr, c.NumRecords, c.MaxRecordSize)
	queue.DrainQueueTransit(c.QueueAddr, vCfg)

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
