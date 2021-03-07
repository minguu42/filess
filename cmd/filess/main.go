package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	"github.com/minguu42/filess"
)

const version = "0.1.0"

var revision = ""

func main() {
	var IsVOptionActive bool
	flag.BoolVar(&IsVOptionActive, "v", false, "show version")
	flag.BoolVar(&IsVOptionActive, "version", false, "show version")
	flag.Parse()
	if IsVOptionActive {
		filess.ShowVersion(version, revision)
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
