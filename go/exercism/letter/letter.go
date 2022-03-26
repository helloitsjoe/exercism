package letter

import "sync"

type FreqMap map[rune]int

func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func worker(s string, sm *sync.Map, wg *sync.WaitGroup) {
	freq := Frequency(s)

	for k, v := range freq {
		existing, exists := sm.Load(k)
		// convert interface{} type to int
		okInt, _ := existing.(int)
		if exists {
			sm.Store(k, v+okInt)
		} else {
			sm.Store(k, v)
		}
	}

	wg.Done()
}

func ConcurrentFrequency(l []string) FreqMap {
	// We want these to be pointers so we use new() Could also create them as
	// structs and pass them with &, but this makes it clearer that they should
	// be used as pointers.
	sm := new(sync.Map)
	wg := new(sync.WaitGroup)

	for _, passage := range l {
		wg.Add(1)
		go worker(passage, sm, wg)
	}

	wg.Wait()

	m := map[rune]int{}

	sm.Range(func(k interface{}, v interface{}) bool {
		key, _ := k.(rune)
		val, _ := v.(int)
		m[key] = val
		return true
	})

	return m
}
