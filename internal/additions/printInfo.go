package additions

import "github.com/fatih/color"

func PrintInfo(res *PackageData) {
	color.Cyan("[INFO] Found 1 package: %s", res.Name)
	color.Cyan("[INFO] It`s description:")
	color.Yellow(res.Description)
}
