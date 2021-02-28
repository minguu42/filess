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

// config はターゲットディレクトリのパスのスライスとソースディレクトリのパスのスライスを管理します。
type config struct {
	Targets []string `json:"targets"`
	Sources []string `json:"sources"`
}

func Init() {
	userPath, _ := os.UserHomeDir()
	configDirPath := filepath.Join(userPath, ".config", "filess")
	perm := os.ModeDir + (fs.ModePerm - 0022)
	// ディレクトリの作成
	if err := os.MkdirAll(configDirPath, perm); err != nil {
		log.Fatal(err)
	}

	// 書き込み内容の作成
	initialization := config{
		[]string{""},
		[]string{""},
	}
	b, err := json.MarshalIndent(initialization, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	// config.jsonファイルの作成
	configFilePath := filepath.Join(configDirPath, "config.json")
	file, err := os.Create(configFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	if _, err := file.Write(b); err != nil {
		log.Fatal(err)
	}
	log.Printf("Create %s\n", configFilePath)
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
