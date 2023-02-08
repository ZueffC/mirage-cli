package commands

import (
	"mirage-cli/internal/commands/actions"

	"github.com/urfave/cli/v2"
)

func List() *cli.Command {
	return &cli.Command{
		Name:        "list",
		Aliases:     []string{"list", "-l", "l"},
		Description: "This command will show all installed packages",
		Action:      actions.ListAction,
	}
}
