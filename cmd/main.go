package main

import (
	"mirage-cli/internal"
	log "mirage-cli/packages/logger"
	"os"
)

func main() {
	internal.Initialize()

	if err := internal.App.Run(os.Args); err != nil {
		(&log.Message{Type: log.Error, Message: "There is some kind of error"}).Log()
	}
}
