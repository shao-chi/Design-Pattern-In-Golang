package main

import "fmt"

type jobStrategy interface {
	do()
	introduce()
	getTitle() string
}

type job struct {
	strategy jobStrategy
	title    string
}

func (j *job) introduce() {
	fmt.Printf("I'm a %s.\n", j.title)
}

func (j *job) getTitle() string {
	return j.title
}

type softwareEngineer struct {
	job
}

func (job *softwareEngineer) do() {
	fmt.Println("I'm coding...")
}

type chef struct {
	job
}

func (job *chef) do() {
	fmt.Println("I'm cooking...")
}

type student struct {
	job
}

func (job *student) do() {
	fmt.Println("I'm doing homework...")
}
