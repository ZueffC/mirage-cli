package additions

import (
	"encoding/json"

	"github.com/go-resty/resty/v2"
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

func SearchByNameQuery(name, url string) *PackageData {
	client := resty.New()
	var result PackageData

	resp, err := client.R().
		SetBody(map[string]interface{}{"type": "current", "name": name}).
		Post(url + "/packages/get")

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		panic(err)
	}

	return &result
}
