package main

import "fmt"

type subject interface {
	addUser(user)
	removeUser(user)
	notifyEveryone(string)
}

type chatbot struct {
	userList []user
}

func (bot *chatbot) addUser(u user) {
	bot.userList = append(bot.userList, u)
	fmt.Printf("Add user(%s)\n", u.getName())
}

func (bot *chatbot) removeUser(u user) {
	numberUser := len(bot.userList)
	for i, userInList := range bot.userList {
		if u.getName() == userInList.getName() {
			tmp := bot.userList[numberUser-1]
			bot.userList[numberUser-1] = bot.userList[i]
			bot.userList[i] = tmp
		}
	}
	bot.userList = bot.userList[:numberUser-1]
	fmt.Printf("Remove user(%s)\n", u.getName())
}

func (bot *chatbot) notifyEveryone(message string) {
	fmt.Printf("Sending message: %s\n", message)
	for _, user := range bot.userList {
		user.notify(message)
	}
}
