package commands

import (
	"mirage-cli/internal/commands/actions"

	"github.com/urfave/cli/v2"
)

func Run() *cli.Command {
	return &cli.Command{
		Name:        "search",
		Aliases:     []string{"r", "-r", "run"},
		Description: "Can run installed application",
		Action:      actions.RunAction,
	}
}
