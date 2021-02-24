package filess

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

func OrganizeFile(config Config) {
	// ターゲットディレクトリとソースディレクトリの取得
	targets := config.Targets
	sources := config.Sources

	// 分類するためのプリフィックスの取得


	// ソースディレクトリを一つずつ調べる
	for _, source := range sources {
		files, err := ioutil.ReadDir(source)
		if err != nil {
			log.Fatal(err)
		}
		for _, file := range files {
			fmt.Println(filepath.Join(source, file.Name()))
		}
	}
	fmt.Println(targets)
}
