package filess

import (
	"log"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestGetConfigDirPath(t *testing.T) {
	userHomePath, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	expectedConfigDirPath := filepath.Join(userHomePath, ".config", "filess")
	if getConfigDirPath() != expectedConfigDirPath {
		t.Fatal("Error")
	}
}

func TestGetConfigFilePath(t *testing.T) {
	userHomePath, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	expectedConfigFilePath := filepath.Join(userHomePath, ".config", "filess", "config.json")
	if GetConfigFilePath() != expectedConfigFilePath {
		t.Fatal("Error")
	}
}

func TestLoadConfig(t *testing.T) {
	targets, sources, inspections := loadConfig("testdata/config.json")

	exceptedTargets := []string{"/example/hoge/tmp"}
	if !reflect.DeepEqual(targets, exceptedTargets) {
		t.Fatal("Error")
	}

	exceptedSources := []string{"/example/tmp"}
	if !reflect.DeepEqual(sources, exceptedSources) {
		t.Fatal("Error2")
	}

	exceptedInspections := []string{"/example/hoge/Downloads"}
	if !reflect.DeepEqual(inspections, exceptedInspections) {
		t.Fatal("Error3")
	}
}
