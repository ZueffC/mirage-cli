package additions

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func InstallDependency(deps []string) {
	out, err := exec.Command("cat", "/proc/version").Output()

	if err != nil {
		log.Fatal(err)
	}

	nameOfDistr := strings.ReplaceAll(strings.ToLower(fmt.Sprintf("%s", out)), " ", "")
	ctn := strings.Contains

	if ctn(nameOfDistr, "mint") || ctn(nameOfDistr, "debian") || ctn(nameOfDistr, "ubuntu") {
		fmt.Println("debian based")
		onDebianBased(deps)
	} else if ctn(nameOfDistr, "arch") {
		println("Arch based")
	}
}

func onDebianBased(deps []string) {
	for _, v := range deps {
		exec.Command("/bin/sh", "sudo apt install", v).CombinedOutput()
	}
}
func onArchBased(deps []string) {}
