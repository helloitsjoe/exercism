package by_example

import "fmt"

func adder() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func Closures() {
	getNext := adder()

	fmt.Println("Next:")
	fmt.Println(getNext())
	fmt.Println(getNext())
	fmt.Println(getNext())

	anotherGetNext := adder()

	fmt.Println("Another:")
	fmt.Println(anotherGetNext())
}
