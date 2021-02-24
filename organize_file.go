package filess

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
)

func OrganizeFile(config Config) {
	// ターゲットディレクトリとソースディレクトリの取得
	targets := config.Targets
	sources := config.Sources

	// ソースディレクトリを一つずつ調べる
	for _, source := range sources {
		// ディレクトリの中身のコレクションを取得する。
		files, err := ioutil.ReadDir(source)
		if err != nil {
			log.Fatal(err)
		}
		// ターゲットディレクトリを一つずつ見ていく。
		for _, target := range targets {
			prefix := filepath.Base(target) + "_"
			// ファイルを一つずつ検査する。
			for _, file := range files {
				if strings.HasPrefix(file.Name(), prefix) {
					nextPath := filepath.Join(target, file.Name())
					// TODO: ここでファイルの移動を実装する。ファイルを移動させたらログを出す。
					fmt.Println(nextPath)
				} else {
				}
			}
		}
	}
}
