package cmd

import (
	"github.com/urfave/cli"
)

var Store = cli.Command{
	Name:        "store",
	Usage:       "Start store server",
	Description: "start store server",
	Action:      runStore,
	Flags: []cli.Flag{
		stringFlag("port, p", "5000", "Temporary port number to prevent conflict"),
		stringFlag("config, c", "", "Custom configuration file path"),
	},
}

func runStore(c *cli.Context) error {
	return nil
}
