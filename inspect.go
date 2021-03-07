package filess

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func inspect(inspections []string) {
	for _, inspection := range inspections {
		if inspection == "" {
			continue
		}

		files, err := ioutil.ReadDir(inspection)
		if err != nil {
			log.Fatal(err)
		}

		for _, file := range files {
			if file.IsDir() {
				continue
			}
			shouldDelete := confirmIfRemoveFile(file.Name())
			if shouldDelete {
				path := filepath.Join(inspection, file.Name())
				if err := os.Remove(path); err != nil {
					log.Fatal(err)
				} else {
					fmt.Printf("Delete %v\n", path)
				}
			} else {
				continue
			}
		}
	}
}

func confirmIfRemoveFile(fileName string) bool {
	for {
		var answer string
		fmt.Printf("%s を削除しますか？（Y/n）: ", fileName)
		if _, err := fmt.Scan(&answer); err != nil {
			log.Fatal(err)
		}
		switch answer {
		case "Y":
			return true
		case "n":
			return false
		default:
			fmt.Println(`"Y" もしくは "n" を入力してください`)
		}
	}
}
