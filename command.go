package filess

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

// ShowVersion はfiless -vの実装である。
// バージョンを表示する。
func ShowVersion(version, revision string) {
	fmt.Println("version:", version)
	fmt.Println("revision:", revision)
}

// Init はfiless initコマンドの実装である。
// 設定ファイルを作成する。既に存在する場合は、上書きする。
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
	defer file.Close()
	if _, err := file.Write(b); err != nil {
		log.Fatal(err)
	}
	log.Printf("Initialize %s\n", configFilePath)
}

// Filess は、filessコマンドの実装である。
// jsonPathが指定されてる時はそのJSONを、されていない時は設定の規定のJSONを使用する。
func Filess(jsonPath string) {
	if jsonPath == "" {
		configPath := getConfigPath()
		if existsFile(configPath) {
			move(configPath)
		} else {
			log.Println("設定ファイルが存在しません")
		}
	} else {
		move(jsonPath)
	}
}
