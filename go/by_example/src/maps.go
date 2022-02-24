package by_example

import (
	"fmt"
)

func Maps() {
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)
	fmt.Println("map length:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	_, exists := m["k2"]
	fmt.Println("exists:", exists)

	n := map[string]string{"foo": "bar", "bar": "baz"}
	fmt.Println("n:", n)
}
