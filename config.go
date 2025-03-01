package filess

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// Config はターゲットディレクトリのパスのスライスとソースディレクトリのパスのスライスを管理します。
type Config struct {
	Targets     []string `json:"targets"`
	Sources     []string `json:"sources"`
	Inspections []string `json:"inspections"`
}

func getConfigDirPath() string {
	userHomePath, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	return filepath.Join(userHomePath, ".config", "filess")
}

func GetConfigFilePath() string {
	userHomePath, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	return filepath.Join(userHomePath, ".config", "filess", "config.json")
}

func loadConfig(path string) ([]string, []string, []string) {
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	var config Config
	if err := json.Unmarshal(raw, &config); err != nil {
		log.Fatal(err)
	}

	return config.Targets, config.Sources, config.Inspections
}
