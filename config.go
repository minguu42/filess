package filess

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// Config はターゲットディレクトリのパスのスライスとソースディレクトリのパスのスライスを管理します。
type Config struct {
	Targets []string `json:"targets"`
	Sources []string `json:"sources"`
}

func existsFile(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func loadJSON(path string) Config {
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var config Config
	if err := json.Unmarshal(raw, &config); err != nil {
		log.Fatal(err)
	}

	return config
}
