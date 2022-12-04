package internal

import (
	"fmt"
	"mirage-cli/internal/additions"
	"os"
	"os/exec"
	"os/user"
	"strconv"
	"strings"
	"time"

	"github.com/urfave/cli/v2"
)

var App cli.App

var userInfo, _ = user.Current()
var pathToMirageFolder = strings.ToLower(userInfo.HomeDir) + "/.mirage"
var informer = additions.Informer

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

				nodes := additions.ParseNodes("nodes.toml")
				res := additions.SearchByNameQuery(name, nodes[0][1:len(nodes[0])-1])

				if len(res.Name) > 0 {
					additions.PrintInfo(res)
				} else {
					informer("error", "No package was found")
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

				nodes := additions.ParseNodes("nodes.toml")
				res := additions.SearchByNameQuery(name, nodes[0][1:len(nodes[0])-1])

				if len(res.Description) > 0 {
					additions.PrintInfo(res)

					fmt.Print("Proceed with installation? [Y/n]: ")
					fmt.Scan(&agreement)

					agreement = strings.ToLower(agreement)

					if agreement == "yes" || agreement == "y" {
						err := os.MkdirAll(pathToMirageFolder, os.ModePerm)

						cmd := exec.Command("git", "clone", res.GitUrl)
						cmd.Dir = pathToMirageFolder
						cmd.Run()

						pathToPackage := pathToMirageFolder + "/" + name + "/Mirage.toml"
						deps := additions.ParseDependencies(pathToPackage)

						additions.InstallDependency(deps)

						if err != nil {
							panic(err)
						}
					} else {
						informer("error", "Installation was stopped")
					}
				} else {
					informer("error", "No package was found")
				}

				return nil
			},
		},

		&cli.Command{
			Name:        "run",
			Aliases:     []string{"r", "start"},
			Description: "Starts app by entered name",
			Action: func(ctx *cli.Context) error {
				name := ctx.Args().Get(0)
				packagePath := pathToMirageFolder + "/" + name

				if _, err := os.Stat(packagePath); os.IsNotExist(err) {
					informer("error", "You entered incorrect name of package. Please retry.")
					return nil
				}

				informer("info", "Package found, starting app...")
				runScript := (additions.ParseRunScript(packagePath + "/Mirage.toml"))[0]

				cmd := exec.Command("/bin/sh", "-c", runScript)
				cmd.Dir = packagePath
				cmd.Stderr, cmd.Stdin, cmd.Stdout = os.Stderr, os.Stdin, os.Stdout
				cmd.Run()

				return nil
			},
		},
	}
}
