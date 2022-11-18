package internal

import (
	"fmt"
	"mirage-cli/internal/additions"
	"os"
	"reflect"
	"strconv"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/urfave/cli/v2"
)

var App cli.App

type Nodes struct {
	Nodes []string
}

func Initialize() {
	var config Nodes

	yearNow := strconv.Itoa(time.Now().Year())
	_, err := toml.DecodeFile("nodes.toml", &config)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	typ, val := reflect.TypeOf(config), reflect.ValueOf(config)
	nodes := make([]string, typ.NumField())

	for i := 0; i < typ.NumField(); i++ {
		nodes[i] = fmt.Sprintf("%v", val.Field(i).Interface())
	}

	App.Name = "mirage"
	App.Usage = "blazingly fast package manager"
	App.UsageText = "mirage [command] [name of package] [flags]"
	App.Copyright = "(C) " + yearNow + " ZueffC. By WTFPL License"

	App.Authors = []*cli.Author{
		{
			Name:  "Зуев Даниил",
			Email: "zueffc@gmail.com",
		},
	}

	App.Commands = []*cli.Command{
		&cli.Command{
			Name:        "search",
			Aliases:     []string{"sch", "s"},
			Description: "this command will search package by name on nodes from yor .config file",
			Action: func(ctx *cli.Context) error {
				name := ctx.Args().Get(0)
				res := additions.SearchByNameQuery(name, nodes[0][1:len(nodes[0])-1])
				println(res.GitUrl)

				return nil
			},
		},
	}

}
