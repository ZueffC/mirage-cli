package additions

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/fatih/color"
)

func InstallDependency(deps []string) {
	out, err := exec.Command("cat", "/proc/version").Output()

	if err != nil {
		log.Fatal(err)
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
		color.Green("[INFO] This package has no dependencies (unlike you)")
		return
	}

	for _, v := range deps {
		cmd := exec.Command("/bin/sh", "-c", "sudo apt install -y "+v)
		cmd.Stderr, cmd.Stdin, cmd.Stdout = os.Stderr, os.Stdin, os.Stdout
		cmd.Run()
	}
}

func onArchBased(deps []string) {
	if len(deps) <= 0 {
		color.Green("[INFO] This package has no dependencies (unlike you)")
		return
	}

	for _, v := range deps {
		cmd := exec.Command("/bin/sh", "-c", "sudo pacman -S "+v)
		cmd.Stderr, cmd.Stdin, cmd.Stdout = os.Stderr, os.Stdin, os.Stdout
		cmd.Run()
	}
}
