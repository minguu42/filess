package main

import (
	"flag"
	"fmt"
	"github.com/minguu42/filess"
)

func main() {
	fmt.Println("START: main")
	flag.Parse()
	arg := flag.Arg(0)

	if arg == "" {
		fmt.Println("JSONファイルが指定されていません。")
	} else {
		config := filess.LoadJson(arg)
		filess.Organize(config)
	}
	fmt.Println("END: main")
}
