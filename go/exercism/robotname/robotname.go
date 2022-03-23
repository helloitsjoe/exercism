// Package robotname gets and resets a robot's name
package robotname

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

var caps = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
var names = map[string]bool{}

const TRIES = 10

type Robot struct {
	name string
}

var r = rand.Intn

func generateName() string {
	first := string(caps[rand.Intn(len(caps))])
	second := string(caps[rand.Intn(len(caps))])

	return fmt.Sprintf("%s%s%d%d%d", first, second, r(10), r(10), r(10))
}

// Name returns the Robot's name.
func (r *Robot) Name() (string, error) {
	// Robot has a unique name, 2 uppercase letters and 3 numbers
	if r.name == "" {
		name, err := r.Reset()
		if err != nil {
			return "", err
		}
		r.name = name
	}
	return r.name, nil
}

var tries = TRIES

// Reset generates a new random name and sets it on the robot.
func (r *Robot) Reset() (string, error) {
	rand.Seed(time.Now().UnixNano())
	r.name = generateName()
	if names[r.name] == true {
		tries--
		if tries == 0 {
			return "", errors.New("names have run out!")
		}
		r.Reset()
	}
	tries = TRIES
	names[r.name] = true
	return r.name, nil
}
