package syncMutexSingleton

import (
	"fmt"
	"testing"
)

func TestMutexSingleton(t *testing.T) {
	singleton := getSingleton()

	for i := 0; i < 5; i++ {
		s := getSingleton()

		fmt.Printf("(s.id: %d, singleton.id: %d)", s.id, singleton.id)
		if s == singleton {
			t.Log("Singleton using sync.Mutex PASS !!")
		} else {
			t.Error("Singleton using sync.Mutex FAIL !!")
		}
	}
}
