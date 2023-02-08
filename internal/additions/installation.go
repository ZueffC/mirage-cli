package additions

import (
	"mirage-cli/internal/parsers"
	log "mirage-cli/packages/logger"
	"os"
	"os/exec"
)

func Install(pkg parsers.PackageInfo) {
	if len(pkg.Dependencies) > 0 {
		InstallDependency(pkg.Dependencies)
	}

	(&log.Message{
		Type:    log.Good,
		Message: "Well, now dependency installation was done. ", //Its time to try running a project!
	}).Log()

	if len(pkg.AfterInstallScript) > 0 {
		cmd := exec.Command("/bin/sh", "-c", pkg.AfterInstallScript)
		cmd.Stdin, cmd.Stdout = os.Stdin, os.Stdout
		error := cmd.Run()

		if error != nil {
			(&log.Message{
				Type:    log.Error,
				Message: "There is an error in the final installation script.",
			}).Log()
		}

		return
	}

	(&log.Message{
		Type:    log.Info,
		Message: "Hmm, now it may have been installed. Let's try to run it!",
	}).Log()
}
