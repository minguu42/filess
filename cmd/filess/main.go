package main

import (
	"flag"
	"log"
	"os"

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
	var target string
	flag.StringVar(&target, "t", "", "add the target path to config")
	flag.StringVar(&target, "target", "", "add the target path to config")
	var source string
	flag.StringVar(&source, "s", "", "add the source path to config")
	flag.StringVar(&source, "source", "", "add the source path to config")
	var inspection string
	flag.StringVar(&inspection, "i", "", "add the inspection path to config")
	flag.StringVar(&inspection, "inspection", "", "add the inspection path to config")
	flag.Parse()
	if IsVOptionActive {
		filess.ShowVersion(version, revision)
		return
	}
	if IsCOptionActive {
		userPath, err := os.UserHomeDir()
		if err != nil {
			log.Fatal(err)
		}
		filess.ShowConfig(filess.GetConfigFilePath(userPath))
		return
	}
	if target != "" {
		filess.AddToConfig(target, "targets")
		return
	}
	if source != "" {
		filess.AddToConfig(source, "sources")
		return
	}
	if inspection != "" {
		filess.AddToConfig(inspection, "inspections")
		return
	}

	firstArg := flag.Arg(0)
	if firstArg == "init" {
		filess.Init()
	} else {
		filess.Filess(firstArg)
	}
}
