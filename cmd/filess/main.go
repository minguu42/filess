package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Start of main")
	envPort, _ := strconv.Atoi(os.Getenv("PORT"))
	var port int
	flag.IntVar(&port, "port", envPort, "port to use")
	flag.IntVar(&port, "p", envPort, "port to use (short)")
	flag.Parse()

	fmt.Println(port)
	fmt.Println("End of main")
}
