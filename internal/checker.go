package internal

import (
	"bytes"
	"errors"
	"fmt"
	"mirage-cli/packages/informer"
	"net/url"
	"os"

	"github.com/BurntSushi/toml"
)

func IsUrl(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}

func inpUrl() string {
	var input string
	fmt.Scanln(&input)

	if len(input) <= 5 && IsUrl(input) {
		informer.Inform("error", "Your URL is invalid, please try again: ")
		return inpUrl()
	}

	return input
}

func checkExistConf() {
	homePath, _ := os.UserHomeDir()
	pathToConfigFolder := homePath + "/.config/mirage/"

	if _, err := os.Stat(pathToConfigFolder + "nodes.toml"); errors.Is(err, os.ErrNotExist) {
		informer.Inform("error", "No one config file exists, creating one...")
		informer.Inform("error", "Please input one node url (e.g. http://zueffc.ml:1984): ")
		url := inpUrl()

		buf := new(bytes.Buffer)
		err := toml.NewEncoder(buf).Encode(map[string]interface{}{
			"Nodes": []string{url},
		})

		if err != nil {
			informer.Inform("error", "Error while addind url, please retry or create issue...")
			return
		}

		os.Mkdir(pathToConfigFolder, os.ModePerm)
		os.WriteFile(pathToConfigFolder+"nodes.toml", buf.Bytes(), 0755)
	}
}
