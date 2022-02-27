package by_example

import "fmt"

func add(a int, b int) int {
	return a + b
}

func addAnother(a, b, c int) int {
	return a + b + c
}

// Multiple return values
func swap(a, b int) (int, int) {
	return b, a
}

// Variadic
func addAll(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func Functions() {
	res := add(1, 2)
	fmt.Println("add:", res)

	res = addAnother(1, 2, 3)
	fmt.Println("addAnother:", res)

	secondArg, firstArg := swap(1, 2)
	fmt.Println("swapped:", secondArg, firstArg)

	variadic := addAll(1, 2, 3, 4)
	fmt.Println("variadic:", variadic)

	nums := []int{2, 3, 4, 5}
	spread := addAll(nums...)
	fmt.Println("spread:", spread)
}
