package by_example

import "fmt"

type Person struct {
	name string
	age  int
}

func newPerson(name string) *Person {
	p := Person{name: name}
	p.age = 42
	return &p
}

func Structs() {
	fmt.Println(Person{"Bob", 20})
	fmt.Println(Person{name: "Alice", age: 30})
	fmt.Println(Person{name: "Fred"})
	fmt.Println(&Person{name: "Ann", age: 40})
	fmt.Println(newPerson("John"))

	s := Person{name: "Sean", age: 32}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 52
	fmt.Println(sp)
}
