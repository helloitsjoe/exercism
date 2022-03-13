package by_example

import (
 "fmt"
 "time"
)

func f(arg string) {
  for i, _ := range make([]int, 3) {
    fmt.Println(arg, ":", i)
  }
}


func Goroutines() {
  f("direct")

  go f("goroutine")

  go func(msg string) {
    fmt.Println(msg)
  }("going")

  time.Sleep(time.Second)
  fmt.Println("done")
}
