package commands

import (
	"mirage-cli/internal/commands/actions"

	"github.com/urfave/cli/v2"
)

func Run() *cli.Command {
	return &cli.Command{
		Name:        "run",
		Aliases:     []string{"run", "-r", "r"},
		Description: "This command will run package you want",
		Action:      actions.RunAction,
	}
}
