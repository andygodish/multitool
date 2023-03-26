package main

import (
	"fmt"
	"flag"
)

func main() {
	helloPtr := flag.Bool("hello", false, "print hello world")
	flag.Parse()

	if *helloPtr {
		fmt.Println("Hello World")
	} else {
		fmt.Println("Unknown Command")
	}
}

func createFile() error {
	
}