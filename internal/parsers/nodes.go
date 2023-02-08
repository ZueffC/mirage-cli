package parsers

import (
	log "mirage-cli/packages/logger"

	"github.com/BurntSushi/toml"
)

func NormalizeURL(url string) string {
	if string(url[len(url)-1]) == "/" {
		url = string(url[:len(url)-1])
		return NormalizeURL(url)
	}

	return url
}

func ParseNodes() ([]string, error) {
	var nodes []string
	var nodesList struct {
		Nodes []string
	}

	_, err := toml.DecodeFile(_pathToConfig, &nodesList)

	if err != nil {
		(&log.Message{
			Type:    log.Error,
			Message: "While checking available nodes... Please, check your configuration file at: " + _pathToConfig,
		}).Log()

		return nil, err
	}

	for _, v := range nodesList.Nodes {
		nodes = append(nodes, NormalizeURL(v))
	}

	return nodes, nil
}
