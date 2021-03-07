package filess

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"os"
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
	configDirPath := getConfigDirPath()
	perm := os.ModeDir + (fs.ModePerm - 0022)
	if err := os.MkdirAll(configDirPath, perm); err != nil {
		log.Fatal(err)
	}

	initialConfig := Config{
		[]string{""},
		[]string{""},
		[]string{""},
	}
	contents, err := json.MarshalIndent(initialConfig, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	configFilePath := getConfigFilePath()
	file, err := os.Create(configFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	if _, err := file.Write(contents); err != nil {
		log.Fatal(err)
	}
	log.Printf("Initialize %s\n", configFilePath)
}

// Filess は、filessコマンドの実装である。
// jsonPathが指定されてる時はそのJSONを、されていない時は設定の規定のJSONを使用する。
func Filess(jsonPath string) {
	if jsonPath == "" {
		configFilePath := getConfigFilePath()
		if existsFile(configFilePath) {
			targets, sources, inspections := loadJSON(configFilePath)
			move(targets, sources)
			log.Println(inspections)
		} else {
			log.Println("設定ファイルが存在しません")
		}
	} else {
		targets, sources, inspections := loadJSON(jsonPath)
		move(targets, sources)
		log.Println(inspections)
	}
}
