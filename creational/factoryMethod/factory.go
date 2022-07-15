package main

import "fmt"

type hawaiianPizzaFactory struct{}
type margheritaPizzaFactory struct{}

func (factory hawaiianPizzaFactory) createProduct() productInterface {
	fmt.Println("Created a Hawaiian Pizza.")
	return hawaiianPizza{}
}

func (factory margheritaPizzaFactory) createProduct() productInterface {
	fmt.Println("Created a Margherita Pizza.")
	return margheritaPizza{}
}
