package parsers

import (
	"fmt"
	"io/ioutil"
	log "mirage-cli/packages/logger"

	"github.com/BurntSushi/toml"
)

type PackageInfo struct {
	Name               string
	Description        string
	Authors            []string
	Dependencies       []string
	AfterInstallScript string
	RunApplication     []string
}

func ParsePackageInfo(path string) PackageInfo {
	var pkgInfo PackageInfo

	_, err := toml.DecodeFile(path+"/Mirage.toml", &pkgInfo)
	if err != nil {
		(&log.Message{
			Type:    log.Error,
			Message: "There are some problem caused parsing Mirage.toml file.",
		}).Log()

		fmt.Println(err)
	}

	return pkgInfo
}

func GetArrayOfDirs() []string {
	var dirs []string
	DirsStruct, err := ioutil.ReadDir(_homePath + "/.mirage")

	if err != nil {
		(log.Message{
			Type:    log.Error,
			Message: "Error reading directories from default mirage path.",
		}).Log()
	}

	for _, DirStruct := range DirsStruct {
		dirs = append(dirs, DirStruct.Name())
	}

	return dirs
}
