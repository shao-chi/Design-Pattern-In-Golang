package main

func addSugar(drink drink) drink {
	drink.cost += 5
	drink.description += " + sugar"
	return drink
}

func addIce(drink drink) drink {
	drink.cost += 1
	drink.description += " + ice"
	return drink
}

func addMilk(drink drink) drink {
	drink.cost += 10
	drink.description += " + milk"
	return drink
}
