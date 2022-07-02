package main

func main() {
	bot := chatbot{}
	user_1 := user{name: "John"}
	user_2 := user{name: "Tom"}
	user_3 := user{name: "Jay"}

	bot.addUser(user_1)
	bot.addUser(user_2)
	bot.notifyEveryone("Hello everyone !!")

	user_4 := user{name: "Andy"}
	bot.addUser(user_3)
	bot.addUser(user_4)
	bot.notifyEveryone("We have new members !!")

	bot.removeUser(user_3)
	bot.notifyEveryone("Go Go Go !!")

	bot.removeUser(user_2)
	bot.notifyEveryone("QQ")
}
