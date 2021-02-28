package main

import (
	"flag"
	"fmt"
	"github.com/minguu42/filess"
	"log"
	"os"
	"path/filepath"
)

const version = "0.1.0"
var revision = ""

func main() {
	var showVersion bool
	flag.BoolVar(&showVersion, "v", false, "show version")
	flag.BoolVar(&showVersion, "version", false, "show version")
	flag.Parse()
	if showVersion {
		fmt.Println("version:", version)
		fmt.Println("revision:", revision)
		return
	}

	arg := flag.Arg(0)
	if arg == "init" {
		filess.Init()
	} else if arg == "" {
		userPath, _ := os.UserHomeDir()
		configFilePath := filepath.Join(userPath, ".config", "filess", "config.json")
		if filess.ExistsFile(configFilePath) {
			filess.Organize(configFilePath)
		} else {
			log.Printf("JSONパスが指定されていません。")
		}
	} else {
		filess.Organize(arg)
	}
}
