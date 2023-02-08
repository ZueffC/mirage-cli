package actions

import (
	"mirage-cli/internal/additions"
	"mirage-cli/internal/parsers"
	log "mirage-cli/packages/logger"
	"os"
	"os/exec"

	"github.com/urfave/cli/v2"
)

func RunAction(ctx *cli.Context) error {
	name := ctx.Args().Get(0)

	if flag, dirName := additions.IsInstalled(name); flag == true {
		path := parsers.ParseHomePath() + "/.mirage/" + dirName
		pkgInfo := parsers.ParsePackageInfo(path)

		for _, v := range pkgInfo.RunApplication {
			cmd := exec.Command("/bin/sh", "-c", v)
			cmd.Stderr, cmd.Stdin, cmd.Stdout = os.Stderr, os.Stdin, os.Stdout
			err := cmd.Run()

			if err != nil {
				(&log.Message{
					Type:    log.Error,
					Message: "There is an error in run script.",
				}).Log()
			}
		}

	}

	return nil
}
