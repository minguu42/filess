package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/minguu42/filess"
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
		// TODO: initコマンドの実装処理
		log.Println("now developing...")
	} else {
		filess.Organize(arg)
	}
}
