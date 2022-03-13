package by_example

import (
  "fmt"
  "time"
)

func Channels() {
  messages := make(chan string)

  go func() { messages <- "ping" }()

  msg := <-messages
  fmt.Println(msg)

  bufferedMessages := make(chan string, 2)

  bufferedMessages <- "buffered"
  bufferedMessages <- "channel"

  fmt.Println(<-bufferedMessages)
  fmt.Println(<-bufferedMessages)

  worker := func (done chan bool) {
    fmt.Print("working...")
    time.Sleep(time.Second)
    fmt.Println("done")

    done <- true
  }

  done := make(chan bool, 1)
  go worker(done)

  <-done
}
