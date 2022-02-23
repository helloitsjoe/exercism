package by_example

import "fmt"

// Slices are like Vecs in Rust?
func Slices() {
	s := make([]string, 3)
	fmt.Println("empty string slice:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	// This gives index out of range error
	// s[3] = "d"

	// Different printing fns
	fmt.Println("added some letters:", s)
	fmt.Printf("s is %v elements long\n", len(s))
	newString := fmt.Sprintf("%s is a string %d characters long", "foo", len("foo"))
	fmt.Println(newString)

	// Appending to the slice
	s = append(s, "d")
	s2 := append(s, "e", "f", "g")
	fmt.Println("append doesn't mutate:", s, s2)

	// Copy
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copy:", c)

	// Slice of a slice
	l := s2[3:5]
	fmt.Println("sliced slice:", l)

	l = s2[:4]
	fmt.Println("up to index 4:", l)

	l = s2[2:]
	fmt.Println("from index 2:", l)

	// Similar to array
	// a := [3]string{"x", "y", "z"}
	t := []string{"x", "y", "z"}
	fmt.Println("initialized:", t)

	// Also similar to array except slice lengths can vary
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = innerLen
		}
	}
	fmt.Println("Two D array,", twoD)
}
