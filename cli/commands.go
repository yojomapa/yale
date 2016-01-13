package cli

import "github.com/codegangsta/cli"

var commands = []cli.Command{
	{
		Name:    "deploy",
		Aliases: []string{"d"},
		Usage:   "despliega un servicio",
		Flags:   deployFlags(),
		Before:  deployBefore,
		Action:  deployCmd,
	},
}
