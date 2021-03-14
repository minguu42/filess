package filess

import (
	"reflect"
	"testing"
)

func TestGetConfigDirPath(t *testing.T) {
	if "/Users/hoge/.config/filess" != getConfigDirPath("/Users/hoge") {
		t.Fatal("Error")
	}
}

func TestGetConfigFilePath(t *testing.T) {
	if "/Users/hoge/.config/filess/config.json" != GetConfigFilePath("/Users/hoge") {
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
