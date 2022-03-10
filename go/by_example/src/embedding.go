package by_example

import "fmt"

type base struct {
  num int
}

func (b base) describe() string {
  return fmt.Sprintf("base with num: %v", b.num)
}

type Container struct {
  base
  str string
}

func Embedding() {
  co := Container{
    base: base{
      num: 1,
    },
    str: "some name",
  }

  fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)
  fmt.Println("also num:", co.base.num)
  fmt.Println("describe:", co.describe())

  type describer interface {
    describe() string
  }

  var d describer = co
  fmt.Println("describer:", d.describe())
}
