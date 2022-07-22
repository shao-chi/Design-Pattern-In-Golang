package main

func main() {
	v_factory := volleyballFactory{}
	v_factory.createBall().describe()
	v_factory.createShoes().describe()
	v_factory.createClothes().describe()

	b_factory := basketballFactory{}
	b_factory.createBall().describe()
	b_factory.createShoes().describe()
	b_factory.createClothes().describe()
}
