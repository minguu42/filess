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

// ShowConfig はfiless -cの実装である。
// 設定を表示する。
func ShowConfig(configFilePath string) {
	targets, sources, inspections := loadConfig(configFilePath)
	fmt.Println("[targets]")
	for _, target := range targets {
		if target == "" {
			continue
		}
		fmt.Println(target)
	}
	fmt.Println()
	fmt.Println("[sources]")
	for _, source := range sources {
		if source == "" {
			continue
		}
		fmt.Println(source)
	}
	fmt.Println()
	fmt.Println("[inspections]")
	for _, inspection := range inspections {
		if inspection == "" {
			continue
		}
		fmt.Println(inspection)
	}
	fmt.Println()
}

// AddToConfig はfiless -s | -t | -iの実装である。
func AddToConfig(path, class string) {
	if !existsDir(path) {
		return
	}
	configFilePath := GetConfigFilePath()
	targets, sources, inspections := loadConfig(configFilePath)
	switch class {
	case "targets":
		targets = appendPath(targets, path)
		targets = removeDuplicate(targets)
	case "sources":
		sources = appendPath(sources, path)
		sources = removeDuplicate(sources)
	case "inspections":
		inspections = appendPath(inspections, path)
		inspections = removeDuplicate(inspections)
	}
	config := Config{
		targets,
		sources,
		inspections,
	}
	contents, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.Create(configFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	if _, err := file.Write(contents); err != nil {
		log.Fatal(err)
	}
	log.Printf("Add %v to %v\n", path, class)
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

	configFilePath := GetConfigFilePath()
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
		configFilePath := GetConfigFilePath()
		if existsFile(configFilePath) {
			targets, sources, inspections := loadConfig(configFilePath)
			move(targets, sources)
			inspect(inspections)
		} else {
			log.Println("設定ファイルが存在しません")
		}
	} else {
		targets, sources, inspections := loadConfig(jsonPath)
		move(targets, sources)
		inspect(inspections)
	}
}
