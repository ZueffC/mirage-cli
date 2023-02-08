package additions

import (
	"bytes"
	"errors"
	"fmt"
	"mirage-cli/internal/parsers"
	log "mirage-cli/packages/logger"
	"net/url"
	"os"
	"strings"

	"github.com/BurntSushi/toml"
)

func IsInstalled(name string) (bool, string) {
	arrayOfDirs := parsers.GetArrayOfDirs()

	for _, dirName := range arrayOfDirs {
		if strings.Contains(dirName, name) {
			return true, dirName
		}
	}

	return false, ""
}

func IsCorrectUrl(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != "" && u.Scheme != "http"
}

func InputUrl() string {
	var input string
	fmt.Scanln(&input)

	if len(input) <= 5 && IsCorrectUrl(input) {
		(&log.Message{Type: log.Error, Message: "Your URL is invalid, please try again: "}).Log()
		return InputUrl()
	}

	return input
}

func CheckExistConf() {
	homePath, _ := os.UserHomeDir()
	pathToConfigFolder := homePath + "/.config/mirage/"
	cnfPath := pathToConfigFolder + "nodes.toml"

	nodesArray, _ := parsers.ParseNodes()

	if _, err := os.Stat(cnfPath); errors.Is(err, os.ErrNotExist) || len(nodesArray) <= 0 {
		(&log.Message{Type: log.Error, Message: "No one config file exists, creating one..."}).Log()
		(&log.Message{
			Type:    log.Error,
			Message: "Please write node url (e.g. https://zueffc.ml): ",
			NoBreak: true,
		}).Log()

		url := InputUrl()

		buf := new(bytes.Buffer)
		err = toml.NewEncoder(buf).Encode(map[string]interface{}{
			"Nodes": []string{parsers.NormalizeURL(url)},
		})

		if err != nil {
			(&log.Message{Type: log.Error, Message: "Error while adding url, please retry or create issue..."}).Log()
			return
		}

		os.Mkdir(pathToConfigFolder, os.ModePerm)
		os.WriteFile(pathToConfigFolder+"nodes.toml", buf.Bytes(), 0755)
	}
}
