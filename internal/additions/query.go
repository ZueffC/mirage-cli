package additions

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"

	"github.com/BurntSushi/toml"
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

func TOMLParser(config interface{}, filename string) []string {
	_, err := toml.DecodeFile(filename, &config)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	typ, val := reflect.TypeOf(config), reflect.ValueOf(&config)
	nodes := make([]string, typ.NumField())

	for i := 0; i < typ.NumField(); i++ {
		nodes[i] = fmt.Sprintf("%v", val.Field(i).Interface())
	}

	return nodes
}
