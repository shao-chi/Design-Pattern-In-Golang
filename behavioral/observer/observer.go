package main

import "fmt"

type observer interface {
	notify(string)
	getName() string
}

type user struct {
	name string
}

func (user *user) getName() string {
	return user.name
}

func (user *user) notify(message string) {
	fmt.Printf("User(%s) got message: %s\n", user.name, message)
}
