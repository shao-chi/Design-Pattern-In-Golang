package main

import (
	"math/rand"
	"sync"
	"time"
)

type webCounter struct {
	id    int
	count int
}

var once sync.Once
var globalWebCounter *webCounter

func getSingletonCounter() *webCounter {
	// just do once
	once.Do(func() {
		// random seed by time
		t := time.Now().UnixNano()
		randInt := rand.New(rand.NewSource(t))

		globalWebCounter = &webCounter{}
		globalWebCounter.id = randInt.Intn(50)
		globalWebCounter.count = 1
	})
	return globalWebCounter
}

func (counter *webCounter) countOne() {
	counter.count += 1
}
