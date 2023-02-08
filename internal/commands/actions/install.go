package actions

import (
	"fmt"
	"mirage-cli/internal/additions"
	"mirage-cli/internal/parsers"
	log "mirage-cli/packages/logger"
	"os"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/urfave/cli/v2"
)

/* for the future update
func removeFlags(slice []string) []string {
	for k, v := range slice {
		letters := strings.Split(v, "")

		if letters[0] == "-" {
			slice[k] = ""
		}
	}

	return slice
}
*/

func InstallAction(ctx *cli.Context) error {
	var agreement string
	name := ctx.Args().Get(0)

	//	args := ctx.Args().Slice()
	//	fmt.Println(removeFlags(args))
	nodes, err := parsers.ParseNodes()

	if len(strings.TrimSpace(name)) <= 1 {
		(log.Message{Type: log.Error, Message: "You`ve entered wrong package name"}).Log()
		return nil
	}

	isInstalled, nameOfDir := additions.IsInstalled(name)
	if isInstalled {
		path := parsers.ParseHomePath() + "/.mirage/" + nameOfDir
		r, _ := git.PlainOpen(path)
		w, _ := r.Worktree()
		w.Pull(&git.PullOptions{RemoteName: "origin"})

		(log.Message{
			Type:    log.Info,
			Message: "Package was successfully updated because it have been installed already",
		}).Log()

		return nil
	}

	if err != nil {
		(log.Message{Type: log.Error, Message: "Error parsing URL of nodes"}).Log()
	}

	for iter, url := range nodes {
		installationDirectory := parsers.ParseHomePath() + "/.mirage/" + name
		(log.Message{Type: log.Info, Message: "Searching for package " + name}).Log()
		pkg, error := additions.SearchByNameQuery(name, url)

		if error == nil && len(pkg.GitUrl) >= 19 {
			(log.Message{Type: log.Good, Message: "We`ve found a package by name, now we`re showing you package details; \n"}).Log()
			pkg.PrintPackageInfo()
			(log.Message{Type: log.Good, Message: "Are you sure it`s correct? (Y/N): ", NoBreak: true}).Log()

			if _, err := fmt.Scan(&agreement); err != nil {
				(log.Message{Type: log.Error, Message: "You have entered something strange, please try again or create issue"}).Log()
			} else if agreement = strings.ToLower(agreement); strings.Contains(agreement, "n") {
				(log.Message{Type: log.Info, Message: "Ok, we`ll interrupt this installation"}).Log()
				return nil
			}

			if agreement = strings.ToLower(agreement); strings.Contains(agreement, "y") {
				git.PlainClone(installationDirectory, false, &git.CloneOptions{
					URL:      pkg.GitUrl,
					Progress: os.Stdout,
				})

				additions.Install(parsers.ParsePackageInfo(installationDirectory))

				return nil
			}
		}

		if (iter + 1) == len(nodes) {
			(log.Message{
				Type:    log.Warning,
				Message: "We`ve found that we already have called all nodes from your config. ",
			}).Log()
		}

		if len(url) < 10 {
			(log.Message{
				Type:    log.Error,
				Message: "After calling all the APIs from your config, we didn't find anything. Maybe you have a typo?",
			}).Log()
		}
	}

	return nil
}
