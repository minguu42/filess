package filess

import (
	"reflect"
	"testing"
)

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
