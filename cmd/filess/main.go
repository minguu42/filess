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
		filess.LoadJson(arg)
	}
	fmt.Println("END: main")
}
