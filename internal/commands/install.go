package commands

import (
	"mirage-cli/internal/commands/actions"

	"github.com/urfave/cli/v2"
)

func Install() *cli.Command {
	return &cli.Command{
		Name:        "install",
		Aliases:     []string{"i", "install", "-i"},
		Description: "for installing you should specify package name",
		Action:      actions.InstallAction,
	}
}
