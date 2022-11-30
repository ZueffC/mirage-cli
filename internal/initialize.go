package internal

import (
	"fmt"
	"mirage-cli/internal/additions"
	"os"
	"os/exec"
	"os/user"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
)

var App cli.App

func parseNodes(path string) []string {
	type nodesStruct struct {
		Nodes []string
	}

	var config nodesStruct

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

	return nodes
}

func parseDepencies(path string) []string {
	type dependenciesStruct struct {
		Dependencies []string
	}

	var dependencies dependenciesStruct

	_, err := toml.DecodeFile(path, &dependencies)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	typ, val := reflect.TypeOf(dependencies), reflect.ValueOf(dependencies)
	deps := make([]string, typ.NumField())

	for i := 0; i < typ.NumField(); i++ {
		deps[i] = fmt.Sprintf("%s", val.Field(i).Interface())
	}

	return deps
}

func Initialize() {
	yearNow := strconv.Itoa(time.Now().Year())

	App.Name = "mirage"
	App.Usage = "lightweight package manager"
	App.UsageText = "mirage [command] [name of package] [flags]"
	App.Copyright = "Copyright (c) " + yearNow + " ZueffC. Licensed under WTFPL License."

	App.Authors = []*cli.Author{
		{
			Name:  "Зуев Даниил",
			Email: "zueffc@gmail.com",
		},
	}

	App.Commands = []*cli.Command{
		&cli.Command{
			Name:        "search",
			Aliases:     []string{"s"},
			Description: "Search package by name on nodes from your .config file",
			Action: func(ctx *cli.Context) error {
				name := ctx.Args().Get(0)

				nodes := parseNodes("nodes.toml")
				res := additions.SearchByNameQuery(name, nodes[0][1:len(nodes[0])-1])

				if len(res.Name) > 0 {
					additions.PrintInfo(res)
				} else {
					color.Red("No package was found")
				}

				return nil
			},
		},

		&cli.Command{
			Name:        "install",
			Aliases:     []string{"i"},
			Description: "Install package on your machine",
			Action: func(ctx *cli.Context) error {
				var agreement string
				name := ctx.Args().Get(0)

				nodes := parseNodes("nodes.toml")
				res := additions.SearchByNameQuery(name, nodes[0][1:len(nodes[0])-1])

				if len(res.Description) > 0 {
					additions.PrintInfo(res)

					fmt.Print("Proceed with installation? [Y/n]: ")
					fmt.Scan(&agreement)

					agreement = strings.ToLower(agreement)

					if agreement == "yes" || agreement == "y" {
						user, _ := user.Current()
						homedir := strings.ToLower(user.HomeDir)
						err := os.MkdirAll(homedir+"/.mirage", os.ModePerm)

						cmd := exec.Command("git", "clone", res.GitUrl)
						cmd.Dir = homedir + "/.mirage"
						cmd.Run()

						pathToPackage := homedir + "/.mirage/" + name + "/Mirage.toml"
						deps := parseDepencies(pathToPackage)

						additions.InstallDependency(deps)

						if err != nil {
							panic(err)
						}
					} else {
						color.HiRed("Installation was stopped")
					}
				} else {
					color.Red("No package was found")
				}

				return nil
			},
		},
	}

}
