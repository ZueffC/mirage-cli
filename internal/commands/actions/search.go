package actions

import (
	"fmt"
	"mirage-cli/internal/additions"
	"mirage-cli/internal/parsers"
	log "mirage-cli/packages/logger"
	"strings"

	"github.com/urfave/cli/v2"
)

func SearchAction(ctx *cli.Context) error {
	var agreement string
	name := ctx.Args().Get(0)

	if len(strings.TrimSpace(name)) <= 1 {
		(log.Message{Type: log.Error, Message: "No one package was found"}).Log()
		return nil
	}

	nodes, err := parsers.ParseNodes()

	for iter, url := range nodes {
		(log.Message{Type: log.Info, Message: "Searching for package " + name}).Log()
		(log.Message{Type: log.Info, Message: "Now we`re calling for " + url + " node;"}).Log()

		pkg, error := additions.SearchByNameQuery(name, url)

		if error == nil && len(pkg.GitUrl) >= 19 {
			(log.Message{Type: log.Good, Message: "We`ve found a package by name, now we`re showing you package details; \n"}).Log()
			pkg.PrintPackageInfo()
			(log.Message{Type: log.Good, Message: "Are you sure it`s correct? (Y/N): ", NoBreak: true}).Log()

			if _, err := fmt.Scan(&agreement); err != nil {
				(log.Message{Type: log.Error, Message: "You have entered something strange, please try again or create issue"}).Log()
			}

			if agreement = strings.ToLower(agreement); strings.Contains(agreement, "y") {
				(log.Message{Type: log.Good, Message: "We`ll take stock this choice"}).Log()
				(log.Message{Type: log.Good, Message: "Well, we found the right package. It's time to end your search!"}).Log()
				return nil
			}

			(log.Message{Type: log.Info, Message: "Ok, we`ll search more..."}).Log()
		}

		if error != nil || len(pkg.GitUrl) <= 19 {
			(log.Message{
				Type:    log.Warning,
				Message: fmt.Sprintf("Well, problems in the %d node. Maybe it doesn't work, maybe the package just doesn't exist.", iter),
			}).Log()
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

	if err != nil {
		(log.Message{Type: log.Error, Message: "Some problem caused by parsing your config. Pls check it in default folder"}).Log()
	}

	return nil
}
