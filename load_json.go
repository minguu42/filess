package filess

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Config struct {
	Targets []string `json:"targets"`
	Sources []string `json:"sources"`
}

func LoadJson(path string) {
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var config Config
	json.Unmarshal(raw, &config)

	fmt.Println(config)
}
