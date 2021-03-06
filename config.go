package filess

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// Config はターゲットディレクトリのパスのスライスとソースディレクトリのパスのスライスを管理します。
type Config struct {
	Targets []string `json:"targets"`
	Sources []string `json:"sources"`
}

// Init はfiless initコマンドの実装です。
// 設定ファイルを作成、上書きします。
func Init() {
	userPath, _ := os.UserHomeDir()
	configDirPath := filepath.Join(userPath, ".Config", "filess")
	perm := os.ModeDir + (fs.ModePerm - 0022)
	// ディレクトリの作成
	if err := os.MkdirAll(configDirPath, perm); err != nil {
		log.Fatal(err)
	}

	// 書き込み内容の作成
	initialization := Config{
		[]string{""},
		[]string{""},
	}
	b, err := json.MarshalIndent(initialization, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	// Config.jsonファイルの作成
	configFilePath := filepath.Join(configDirPath, "Config.json")
	file, err := os.Create(configFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	if _, err := file.Write(b); err != nil {
		log.Fatal(err)
	}
	log.Printf("Create %s\n", configFilePath)
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
