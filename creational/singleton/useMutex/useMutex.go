package syncMutexSingleton

import (
	"math/rand"
	"sync"
	"time"
)

type singletonType struct {
	id int
}

var lock = &sync.Mutex{}
var singletonInstance *singletonType = nil

func getSingleton() *singletonType {
	lock.Lock()
	defer lock.Unlock()

	if singletonInstance == nil {
		// random seed by time
		t := time.Now().UnixNano()
		randInt := rand.New(rand.NewSource(t))

		singletonInstance = new(singletonType)
		singletonInstance.id = randInt.Intn(50)
	}
	return singletonInstance
}
