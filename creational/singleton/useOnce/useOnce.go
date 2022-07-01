package syncOnceSingleton

import (
	"math/rand"
	"sync"
	"time"
)

type singletonType struct {
	id int
}

var once sync.Once
var singletonInstance *singletonType

func getSingleton() *singletonType {
	// just do once
	once.Do(func() {
		// random seed by time
		t := time.Now().UnixNano()
		randInt := rand.New(rand.NewSource(t))

		singletonInstance = &singletonType{id: randInt.Intn(50)}
	})
	return singletonInstance
}
