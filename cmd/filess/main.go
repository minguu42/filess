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
	var IsCOptionActive bool
	flag.BoolVar(&IsCOptionActive, "c", false, "show config")
	flag.BoolVar(&IsCOptionActive, "config", false, "show config")
	flag.Parse()
	if IsVOptionActive {
		filess.ShowVersion(version, revision)
		return
	}
	if IsCOptionActive {
		filess.ShowConfig()
		return
	}

	firstArg := flag.Arg(0)
	if firstArg == "init" {
		filess.Init()
	} else {
		filess.Filess(firstArg)
	}
}
