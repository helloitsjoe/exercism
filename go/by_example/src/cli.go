package by_example

import (
	"fmt"
	"os"
)

func Cli() {
	// name := flag.String("name", "nobody", "A name")

	// flag.Parse()

	args := os.Args

	name := args[1:]

	fmt.Printf("Hello, %s!\n", name)

	// f, err := os.ReadFile(name)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("Hello %s", f)
}
