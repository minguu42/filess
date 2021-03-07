package filess

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func move(targets, sources []string) {


	for _, target := range targets {
		for _, source := range sources {
			if target == "" || source == "" {
				continue
			}
			moveFile(target, source)
		}
	}
}

func moveFile(target, source string) {
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
