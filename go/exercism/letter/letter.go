package letter

import (
	"fmt"
	"sync"
)

type FreqMap map[rune]int

func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func worker(s string, c chan bool, sm *sync.Map) {
	freq := Frequency(s)

	for k, v := range freq {
		existing, exists := sm.Load(k)

		// Convert to int - not sure why this is needed
		okInt, _ := existing.(int)
		if exists {
			sm.Store(k, v+okInt)
		} else {
			sm.Store(k, v)
		}
	}

	c <- true
	// Removing the following fmt.Println makes this fail. Why?
	// fmt.Println("s", s[:1])
}

func ConcurrentFrequency(l []string) FreqMap {
	done := make(chan bool, len(l))
	var sm sync.Map
	// var wg sync.WaitGroup

	for _, passage := range l {
		// wg.Add(1)
		go worker(passage, done, &sm)
	}

	noted := <-done
	// wg.Wait()
	fmt.Println("done", noted)

	m := map[rune]int{}

	sm.Range(func(k interface{}, v interface{}) bool {
		key, _ := k.(rune)
		val, _ := v.(int)
		fmt.Println(key, ":", val)
		m[key] = val
		return true
	})

	return m
}
