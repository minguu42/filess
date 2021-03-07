package main

import (
	"flag"
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

	firstArg := flag.Arg(0)
	if firstArg == "init" {
		filess.Init()
	} else {
		filess.Filess(firstArg)
	}
}
