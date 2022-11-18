package internal

import (
	"fmt"
	"mirage-cli/internal/additions"
	"os"
	"os/exec"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
)

var App cli.App

type Nodes struct {
	Nodes []string
}

func printInfo(res *additions.PackageData) {
	color.Cyan("[INFO] Found 1 package: %s", res.Name)
	color.Cyan("[INFO] It`s description:")
	color.Yellow(res.Description)
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

				if len(res.Name) > 0 {
					printInfo(res)
				} else {
					color.Red("No one package was found")
				}

				return nil
			},
		},

		&cli.Command{
			Name:        "install",
			Aliases:     []string{"ins", "i"},
			Description: "this command will install package on your machine",
			Action: func(ctx *cli.Context) error {
				var agreement string
				name := ctx.Args().Get(0)
				res := additions.SearchByNameQuery(name, nodes[0][1:len(nodes[0])-1])

				if len(res.Description) > 0 {
					printInfo(res)
					fmt.Print("Do y wanna install it? [Y/N]: ")
					fmt.Scan(&agreement)
					agreement = strings.ToLower(agreement)

					if agreement == "yes" || agreement == "y" {
						(exec.Command("mkdir ~/mirage-packages")).Run()
						(exec.Command("cd ~/mirage-packages")).Run()
						(exec.Command("git clone " + res.GitUrl + " .")).Run()
					} else {
						color.HiRed("Installation was stopped")
					}
				} else {
					color.Red("No one package found")
				}

				return nil
			},
		},
	}

}
