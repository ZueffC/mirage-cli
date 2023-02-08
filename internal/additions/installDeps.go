package additions

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	log "mirage-cli/packages/logger"
)

func InstallDependency(deps []string) {
	out, err := exec.Command("cat", "/proc/version").Output()

	if err != nil {
		(&log.Message{
			Type:    log.Error,
			Message: "There was a problem to understood what disto you use. ",
		}).Log()
	}

	nameOfDistr := strings.ReplaceAll(strings.ToLower(fmt.Sprintf("%s", out)), " ", "")
	ctn := strings.Contains

	if ctn(nameOfDistr, "mint") || ctn(nameOfDistr, "debian") || ctn(nameOfDistr, "ubuntu") {
		fmt.Println("Debian-based")
		onDebianBased(deps)
	} else if ctn(nameOfDistr, "arch") {
		onArchBased(deps)
	}
}

func onDebianBased(deps []string) {
	if len(deps) <= 0 {
		(&log.Message{
			Type:    log.Info,
			Message: "This package has no dependencies (unlike you)",
		}).Log()
		return
	}

	var doneChannel = make(chan bool)

	go func() {
		for _, v := range deps {
			cmd := exec.Command("/bin/sh", "-c", "sudo apt install -y "+strings.ToLower(v))
			cmd.Stdin, cmd.Stdout = os.Stdin, os.Stdout
			if error := cmd.Run(); error != nil {
				(&log.Message{
					Type:    log.Error,
					Message: "There was a problem installing some dependencies for this package. Try installing them manually: " + v,
				}).Log()

				doneChannel <- true
				return
			}
		}
		doneChannel <- true
	}()
	<-doneChannel
}

func onArchBased(deps []string) {
	if len(deps) <= 0 {
		(&log.Message{
			Type:    log.Info,
			Message: "This package has no dependencies (unlike you)",
		}).Log()

		return
	}

	for _, v := range deps {
		cmd := exec.Command("/bin/sh", "-c", "sudo pacman -S "+v)
		cmd.Stdin, cmd.Stdout = os.Stdin, os.Stdout
		if error := cmd.Run(); error != nil {
			(&log.Message{
				Type:    log.Error,
				Message: "There was a problem installing some dependencies for this package. Try installing them manually: " + v,
			}).Log()

			return
		}
	}
}
