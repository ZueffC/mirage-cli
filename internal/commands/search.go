package commands

import (
	"mirage-cli/internal/commands/actions"

	"github.com/urfave/cli/v2"
)

func Search() *cli.Command {
	return &cli.Command{
		Name:        "search",
		Aliases:     []string{"s", "-s", "search"},
		Description: "Search package by name on nodes from your nodes.toml file",
		Action:      actions.SearchAction,
	}
}
