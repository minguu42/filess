package filess

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// config はターゲットディレクトリのパスのスライスとソースディレクトリのパスのスライスを管理します。
type config struct {
	Targets []string `json:"targets"`
	Sources []string `json:"sources"`
}

// LoadJSON はJSONファイルへのパスを受け取り、読み込んでConfigを返します。
func LoadJSON(path string) config {
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var config config
	if err := json.Unmarshal(raw, &config); err != nil {
		log.Fatal(err)
	}

	return config
}
