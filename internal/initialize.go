package internal

import (
	"mirage-cli/internal/commands"
	"strconv"
	"time"

	"github.com/urfave/cli/v2"
)

var App cli.App

func Initialize() {
	yearNow := strconv.Itoa(time.Now().Year())

	App.Name = "mirage"
	App.Usage = "blazingly fast package manager"
	App.UsageText = "mirage [command] [name of package] [flags]"
	App.Copyright = "(C) " + yearNow + " 0xFFCat. By WTFPL License"

	App.Authors = []*cli.Author{
		{
			Name:  "Зуев Даниил",
			Email: "zueffc@gmail.com",
		},
	}

	App.Commands = []*cli.Command{
		commands.Search(), //Search command
	}

}
