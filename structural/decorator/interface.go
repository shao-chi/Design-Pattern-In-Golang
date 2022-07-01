package main

import "fmt"

type componentInterface interface {
	getCost()
	getDescription()
}

type drink struct {
	description string
	cost        int
}

func (drink *drink) getCost() int {
	return drink.cost
}

func (drink *drink) getDescription() {
	fmt.Printf("%s: $%d\n", drink.description, drink.getCost())
}

// type addition struct {
// 	drink drink
// }
