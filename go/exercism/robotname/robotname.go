// Package robotname gets and resets a robot's name
package robotname

import (
	"errors"
	"fmt"
	"math/rand"
)

type Robot struct {
	name string
}

var names = generateNames()

func getName() string {
	// Previous (slow) solution
	// i := rand.Intn(len(names))
	// name := names[i]
	// names = append(names[:i], names[i+1:]...)
	// return name

	name := names[0]
	names = names[1:]
	return name
}

// Create a slice of robot names going sequentially.
// When Name is called, remove that name from the list.
// When the list is empty, return an error
var generateNames = func() []string {
	// Robot has a unique name, 2 uppercase letters and 3 numbers
	var names = make([]string, 0)
	var caps = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

	for _, cap1 := range caps {
		for _, cap2 := range caps {
			for i := 0; i < 1000; i++ {
				num := fmt.Sprintf("%03d", i)
				name := string(cap1) + string(cap2) + num
				names = append(names, name)
			}
		}
	}

	rand.Shuffle(26*26*1000, func(i, j int) {
		names[i], names[j] = names[j], names[i]
	})

	return names
}

// Name returns the Robot's name.
func (r *Robot) Name() (string, error) {
	if r.name == "" {
		r.Reset()
		if r.name == "" {
			return "", errors.New("names have run out!")
		}
	}
	return r.name, nil
}

// Reset generates a new random name and sets it on the robot.
func (r *Robot) Reset() {
	if len(names) == 0 {
		r.name = ""
	} else {
		r.name = getName()
	}
}
