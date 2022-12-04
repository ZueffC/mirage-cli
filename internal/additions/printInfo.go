package additions

import (
	"github.com/fatih/color"
)

func PrintInfo(res *PackageData) {
	color.Cyan("[INFO] Found 1 package: %s", res.Name)
	color.Cyan("[INFO] Its description:")
	color.Yellow(res.Description)
}

func Informer(typeOf string, message string) {
	if typeOf == "error" {
		color.Red("[ERROR] " + message)
	} else if typeOf == "info" {
		color.Blue("[INFO] " + message)
	}
}
