package command

import (
	"github.com/mitchellh/cli"
)

// SeedCommand : x
type SeedCommand struct {
	Ui cli.Ui
}

// Run : x
func (c *SeedCommand) Run(_ []string) int {
	c.Ui.Output("Would seed here")
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
