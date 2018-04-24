package command

import (
	"flag"

	"github.com/mitchellh/cli"
	"github.com/norhe/transit-benchmark/persist"
)

// ResultsCommand : Properties for the results command
type ResultsCommand struct {
	UI         cli.Ui
	QueueAddr  string
	DbAddr     string
	DbUsername string
	DbPassword string
	DbName     string
}

// Run : I assume this will write to a mariadb/mysql compatible system
func (c *ResultsCommand) Run(args []string) int {
	cmdFlags := flag.NewFlagSet("results", flag.ContinueOnError)
	cmdFlags.Usage = func() { c.UI.Output(c.Help()) }

	cmdFlags.StringVar(&c.QueueAddr, "queue-addr", "amqp://guest:guest@localhost:5672/", "The rabbitmq addr")
	cmdFlags.StringVar(&c.DbAddr, "db-addr", "tcp(127.0.0.1:3306)", "The databse server address")
	cmdFlags.StringVar(&c.DbUsername, "db-username", "root", "The username to use when authenticating to the db")
	cmdFlags.StringVar(&c.DbPassword, "db-password", "root", "The transit key to use.  This key must already exist on the Vault server")
	cmdFlags.StringVar(&c.DbName, "db-name", "benchmark", "The database to write results too")

	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}

	dCfg := persist.Config{
		Addr:     c.DbAddr,
		Username: c.DbUsername,
		Password: c.DbPassword,
		DbName:   c.DbName,
	}

	persist.CreateTables(dCfg)
	/*c.Ui.Output(fmt.Sprintf("Would connect to queue at %s and send messages to %s Vault server for key %s", c.QueueAddr, c.VaultAddr, c.TransitKeyName))

	vCfg := vault.Config{
		Address:        c.VaultAddr,
		Token:          c.VaultToken,
		TransitKeyName: c.TransitKeyName,
	}

	//queue.SeedQueueRandom(c.QueueAddr, c.NumRecords, c.MaxRecordSize)
	queue.DrainQueueTransit(c.QueueAddr, vCfg)*/

	return 0
}

// Help : x
func (c *ResultsCommand) Help() string {
	return "Execute tests "
}

// Synopsis : x
func (c *ResultsCommand) Synopsis() string {
	return "Execute tests by draining messages from the Queue"
}
