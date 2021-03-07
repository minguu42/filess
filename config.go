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

// ExistsFile はファイルの存在を確認する関数である。
func ExistsFile(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

// LoadJSON はJSONファイルへのパスを受け取り、読み込んでConfigを返す。
func LoadJSON(path string) Config {
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
