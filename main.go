package main

import (
	"log"
	"mirage-cli/internal"
	"os"
)

func main() {
	internal.Initialize()

	if err := internal.App.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
