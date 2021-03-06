package filess

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Organize はConfigを受け取り、ファイルを整理します。
func Organize(jsonPath string) {
	// ターゲットディレクトリとソースディレクトリの取得
	config := LoadJSON(jsonPath)
	targets := config.Targets
	sources := config.Sources

	// ソースディレクトリを一つずつ調べる
	for _, target := range targets {
		for _, source := range sources {
			if target == "" || source == "" {
				continue
			}
			organizeFile(target, source)
		}
	}
}

// organizeFile は、一つのソースから１つのターゲットにファイルを整理する。
func organizeFile(target, source string) {
	prefix := filepath.Base(target) + "_"
	files, err := ioutil.ReadDir(source)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if strings.HasPrefix(file.Name(), prefix) {
			currentPath := filepath.Join(source, file.Name())
			nextPath := filepath.Join(target, file.Name())
			if err := os.Rename(currentPath, nextPath); err != nil {
				log.Fatal(err)
			} else {
				fmt.Printf("Move to %v from %v\n", currentPath, nextPath)
			}
		}
	}
}
