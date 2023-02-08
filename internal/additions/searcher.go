package additions

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	log "mirage-cli/packages/logger"

	"github.com/fatih/color"
	"github.com/go-resty/resty/v2"
	"github.com/jedib0t/go-pretty/table"
)

type PackageData struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	GitUrl      string `json:"git_url"`
}

type Inf struct {
	AuthorID    uint   `json:"author_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	GitURL      string `json:"git_url"`
}

func (pkg PackageData) PrintPackageInfo() {
	t := table.NewWriter()
	t.SetOutputMirror(color.Output)

	t.AppendHeader(table.Row{"Name", "Description", "Url"})
	t.AppendRow(table.Row{pkg.Name, pkg.Description, pkg.GitUrl})

	t.Render()
	println()
}

func SearchByNameQuery(name, url string) (*PackageData, error) {
	client := resty.New()
	var result PackageData

	client.SetTLSClientConfig(&tls.Config{
		InsecureSkipVerify: false,
	})

	if IsCorrectUrl(url) {
		resp, err := client.R().
			SetBody(map[string]interface{}{"type": "current", "name": name}).
			Post(url + "/packages/get")

		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(resp.Body(), &result)
		if err != nil {
			return nil, err
		}

		return &result, nil
	} else {
		(&log.Message{
			Type:    log.Error,
			Message: "There is incorrect node url",
		}).Log()

		return nil, fmt.Errorf("incorrect node url")
	}
}
