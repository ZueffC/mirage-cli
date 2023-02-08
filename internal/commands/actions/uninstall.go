package actions

import (
	"mirage-cli/internal/additions"
	"mirage-cli/internal/parsers"
	log "mirage-cli/packages/logger"
	"os"

	"github.com/urfave/cli/v2"
)

func UninstallAction(ctx *cli.Context) error {
	name := ctx.Args().Get(0)

	(log.Message{Type: log.Info, Message: "Ok, we`ll remove this package from your device"}).Log()

	if flag, dirName := additions.IsInstalled(name); flag == true {
		path := parsers.ParseHomePath() + "/.mirage/" + dirName
		if err := os.RemoveAll(path); err != nil {
			(log.Message{
				Type:    log.Info,
				Message: "There was an error removing directory '" + path + "' Please try again",
			}).Log()
		}

		(log.Message{Type: log.Good, Message: "Ok, we`ve remove this package"}).Log()
	}

	return nil
}
