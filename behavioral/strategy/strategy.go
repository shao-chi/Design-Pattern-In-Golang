package main

import "fmt"

type jobStrategy interface {
	do()
	introduce()
	getTitle() string
}

// Parent Struct
type job struct {
	title string
}

func (j *job) introduce() {
	fmt.Printf("I'm a %s.\n", j.title)
}

func (j *job) getTitle() string {
	return j.title
}

// Child Struct
type softwareEngineer struct {
	job
}

func (job *softwareEngineer) do() {
	fmt.Println("I'm coding...")
}

// Child Struct
type chef struct {
	job
}

func (job *chef) do() {
	fmt.Println("I'm cooking...")
}

// Child Struct
type student struct {
	job
}

func (job *student) do() {
	fmt.Println("I'm doing homework...")
}
