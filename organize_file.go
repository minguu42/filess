package filess

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
)

func Organize(config Config) {
	// ターゲットディレクトリとソースディレクトリの取得
	targets := config.Targets
	sources := config.Sources

	// ソースディレクトリを一つずつ調べる
	for _, target := range targets {
		for _, source := range sources {
			organizeFile(target, source)
		}
	}
}

// 一つのソースディレクトリから１つのターゲットディレクトリにファイルを整理する。
func organizeFile(target, source string) {
	prefix := filepath.Base(target) + "_"
	files, err := ioutil.ReadDir(source)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if strings.HasPrefix(file.Name(), prefix) {
			nextPath := filepath.Join(target, file.Name())
			// TODO: ソースからターゲットにファイルを移動させる
			fmt.Println(nextPath)
		}
	}
}
