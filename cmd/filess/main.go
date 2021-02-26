package main

import (
	"flag"
	"fmt"

	"github.com/minguu42/filess"
)

var version = "0.1.0"
var revision = ""

func main() {
	var showVersion bool
	flag.BoolVar(&showVersion, "v", false, "show version")
	flag.BoolVar(&showVersion, "version", false, "show version")
	flag.Parse()
	if showVersion {
		fmt.Printf("version: %v-%v\n", version, revision)
		return
	}

	arg := flag.Arg(0)
	if arg == "" {
		fmt.Println("JSONファイルが指定されていません。")
	} else {
		config := filess.LoadJSON(arg)
		filess.Organize(config)
	}
}
