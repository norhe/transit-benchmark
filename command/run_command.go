package command

import (
	"github.com/mitchellh/cli"
)

// RunCommand : x
type RunCommand struct {
	Ui cli.Ui
}

// Run : x
func (c *RunCommand) Run(_ []string) int {
	c.Ui.Output("Would Run here")
	return 0
}

// Help : x
func (c *RunCommand) Help() string {
	return "Generate N number of random messagesd with lengths between 0 and Y"
}

// Synopsis : x
func (c *RunCommand) Synopsis() string {
	return "Run messages into the queue"
}
