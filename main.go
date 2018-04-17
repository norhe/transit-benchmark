package main

import (
	"os"

	"github.com/mitchellh/cli"
	"github.com/norhe/transit-benchmark/command"
	"github.com/norhe/transit-benchmark/utils"
)

/* App can be run in a variety of ways.  One can seed random
*  strings into the queue.  One can also drain the queue issuing
*  transit calls based upon the queue contents.  This way it is possible
*  for a user to tune the test data to their liking.  To do a generic test
*  seed the data with the first run, and then run again in test mode.
 */
func main() {
	ui := &cli.BasicUi{
		Reader:      os.Stdin,
		Writer:      os.Stdout,
		ErrorWriter: os.Stderr,
	}

	c := cli.NewCLI("app", "0.0.1")

	c.Args = os.Args[1:]

	c.Commands = map[string]cli.CommandFactory{
		"seed": func() (cli.Command, error) {
			return &command.SeedCommand{
				Ui: &cli.ColoredUi{
					Ui:          ui,
					OutputColor: cli.UiColorBlue,
				},
			}, nil
		},
		"run": func() (cli.Command, error) {
			return &command.RunCommand{
				Ui: &cli.ColoredUi{
					Ui:          ui,
					OutputColor: cli.UiColorGreen,
				},
			}, nil
		},
	}
	/*if os.Getenv("SEED") == "true" {
		queue.SeedQueueRandom()
	} else if os.Getenv("RUN_TEST") == "true" {
		log.Printf("Executing test...")
	} else {
		log.Println("Please pass in SEED=true or RUN_TEST=true as environment variables.")
	}*/

	exitStatus, err := c.Run()
	utils.FailOnError(err, "Failed to run command")
	os.Exit(exitStatus)
}

func benchmark() {

}
