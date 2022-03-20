package by_example

import (
  "fmt"
  "sync"
  "time"
)

func waitGroupsWorker(id int) {
  fmt.Printf("Worker %d starting\n", id)
  time.Sleep(time.Second)
  fmt.Printf("Worker %d done\n", id)
}

func WaitGroups() {
  var wg sync.WaitGroup

  for i := 1; i <= 5; i++ {
    wg.Add(1)

    // Capture value in closure
    i := i

    go func() {
      defer wg.Done()
      waitGroupsWorker(i)
    }()
  }

  wg.Wait()
}

