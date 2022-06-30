package main

import "fmt"

func main() {
	h1 := human{name: "Tom"}
	h1.introduce()

	h1.setJob(&softwareEngineer{job: job{title: "junior backend engineer"}})
	h1.introduce()
	h1.work()

	h1.setJob(&softwareEngineer{job: job{title: "senior backend engineer"}})

	fmt.Println("------------------------------------------------")

	h2 := human{
		name: "Simon",
		job:  &chef{job: job{title: "chef"}},
	}
	h2.introduce()
	h2.work()

	h2.setJob(&softwareEngineer{job: job{title: "frontend engineer"}})
	h2.work()

	fmt.Println("------------------------------------------------")

	h3 := human{
		name: "Greenie",
		job:  &student{job: job{title: "high school student"}},
	}
	h3.introduce()
	h3.work()

	h3.setJob(&student{job: job{title: "university student"}})
}
