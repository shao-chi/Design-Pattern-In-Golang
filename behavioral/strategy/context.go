package main

import "fmt"

type humanInterface interface {
	work()
	setJob(job jobStrategy)
	introduce()
}

type human struct {
	name string
	job  jobStrategy
}

func (h *human) work() {
	h.job.do()
}

func (h *human) setJob(job jobStrategy) {
	h.job = job
	fmt.Printf("I become a %s.\n", h.job.getTitle())
}

func (h *human) introduce() {
	fmt.Printf("I'm %s.\n", h.name)

	if h.job == nil {
		fmt.Println("I don't have job.")
	} else {
		h.job.introduce()
	}
}
