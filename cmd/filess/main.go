package main

import (
	"flag"
	"fmt"

	"github.com/minguu42/filess"
)

var version = "0.1.0"

func main() {
	var showVersion bool
	flag.BoolVar(&showVersion, "v", false, "show version")
	flag.BoolVar(&showVersion, "version", false, "show version")
	flag.Parse()
	if showVersion {
		fmt.Println("version:", version)
		return
	}

	arg := flag.Arg(0)
	if arg == "" {
		fmt.Println("JSONファイルが指定されていません。")
	} else {
		config := filess.LoadJson(arg)
		filess.Organize(config)
	}
}
