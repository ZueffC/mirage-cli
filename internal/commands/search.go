package commands

import (
	"github.com/urfave/cli/v2"
)

func Search() *cli.Command {
	var command cli.Command

	command.Name = "search"
	command.Aliases = []string{"sch", "s"}
	command.Description = "this command will search package by name on nodes from yor .config file"
	command.Action = searchPackageByName()

	return &command
}

func searchPackageByName() cli.ActionFunc { return nil }
