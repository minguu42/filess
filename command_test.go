package filess

func ExampleShowVersion() {
	ShowVersion("0.1.0", "6c14e05")
	// Output:
	// version: 0.1.0
	// revision: 6c14e05
}

func ExampleShowConfig() {
	ShowConfig("testdata/config.json")
	// Output:
	// [targets]
	// /example/hoge/tmp
	//
	// [sources]
	// /example/tmp
	//
	// [inspections]
	// /example/hoge/Downloads
}