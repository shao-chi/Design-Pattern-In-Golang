package main

import "fmt"

type basketballShoes struct{}
type basketballClothes struct{}
type basketball struct{}

type volleyballShoes struct{}
type volleyballClothes struct{}
type volleyball struct{}

func (shoes basketballShoes) describe() {
	fmt.Println("They're the shoes for playing basketball.")
}

func (clothes basketballClothes) describe() {
	fmt.Println("They're the clothes for playing basketball")
}

func (ball basketball) describe() {
	fmt.Println("It's a basketball.")
}

func (shoes volleyballShoes) describe() {
	fmt.Println("They're the shoes for playing volley ball.")
}

func (clothe volleyballClothes) describe() {
	fmt.Println("They're the clothes for playing volley ball.")
}

func (ball volleyball) describe() {
	fmt.Println("It's a volley ball")
}
