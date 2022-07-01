package main

func main() {
	americano := makeAmericano()
	americano.getDescription()
	americano = addSugar(americano)
	americano.getDescription()
	americano = addIce(americano)
	americano.getDescription()

	espresso := makeExpresso()
	espresso.getDescription()
	espresso = addSugar(espresso)
	espresso.getDescription()
	espresso = addSugar(espresso)
	espresso.getDescription()

	black_tea := makeBlackTea()
	black_tea.getDescription()
	black_tea = addMilk(black_tea)
	black_tea.getDescription()
}
