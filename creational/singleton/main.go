package main

import "fmt"

func main() {
	counter := getSingletonCounter()
	fmt.Printf("counter id: %d, count: %d\n", counter.id, counter.count)

	counter.countOne()
	fmt.Printf("counter id: %d, count: %d\n", counter.id, counter.count)

	counter = getSingletonCounter()
	fmt.Printf("counter id: %d, count: %d\n", counter.id, counter.count)
	counter.countOne()
	counter.countOne()
	counter.countOne()
	counter.countOne()

	counter = getSingletonCounter()
	fmt.Printf("counter id: %d, count: %d\n", counter.id, counter.count)
	counter.countOne()
	counter.countOne()
	counter.countOne()

	counter = getSingletonCounter()
	fmt.Printf("counter id: %d, count: %d\n", counter.id, counter.count)
}
