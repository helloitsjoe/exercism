package main

import (
	"flag"
	"fmt"
)

func main() {
	name := flag.String("name", "", "A name")

	flag.Parse()

	fmt.Printf("Hello, %s!", name)
}
