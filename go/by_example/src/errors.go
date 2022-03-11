package by_example

import (
  "fmt"
  "errors"
)

func throws(arg int) (int, error) {
  if arg == 42 {
    return -1, errors.New("can't work with 42")
  }

  return arg + 3, nil
}

type argError struct {
  arg int
  prob string
}

func (e *argError) Error() string {
  return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func throwsArg(arg int) (int, error) {
  if arg == 42 {
    return -1, &argError{arg, "can't work with it"}
  }
  return arg + 4, nil
}

func Errors() {
  for _, i := range []int{7, 42} {
    if r, e := throws(i); e != nil {
      fmt.Println("throws failed", e)
    } else {
      fmt.Println("throws passed", r)
    }
  }

  for _, i := range []int{7, 42} {
    if r, e := throwsArg(i); e != nil {
      fmt.Println("throwsArg failed", e)
    } else {
      fmt.Println("throwsArg worked", r)
    }
  }

  _, e := throwsArg(42)
  if ae, ok := e.(*argError); ok {
    fmt.Println(ae.arg)
    fmt.Println(ae.prob)
  }
}
