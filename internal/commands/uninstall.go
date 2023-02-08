package commands

import (
	"mirage-cli/internal/commands/actions"

	"github.com/urfave/cli/v2"
)

func Uninstall() *cli.Command {
	return &cli.Command{
		Name:        "uninstall",
		Aliases:     []string{"uninstall", "-u", "u"},
		Description: "This command will be remove package from your device",
		Action:      actions.UninstallAction,
	}
}
