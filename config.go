package filess

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// Config はターゲットディレクトリのパスのスライスとソースディレクトリのパスのスライスを管理します。
type Config struct {
	Targets     []string `json:"targets"`
	Sources     []string `json:"sources"`
	Inspections []string `json:"inspections"`
}

func loadJSON(path string) ([]string, []string) {
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	var config Config
	if err := json.Unmarshal(raw, &config); err != nil {
		log.Fatal(err)
	}

	return config.Targets, config.Sources
}
