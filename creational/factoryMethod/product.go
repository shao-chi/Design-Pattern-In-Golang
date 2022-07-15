package main

import "fmt"

type hawaiianPizza struct{}
type margheritaPizza struct{}

func (pizza hawaiianPizza) operate() {
	fmt.Println("I'm a Hamaiian Pizza ~")
}

func (pizza margheritaPizza) operate() {
	fmt.Println("I'm a Margherita Pizza.")
}
