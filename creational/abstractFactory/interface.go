package main

type FactoryInterface interface {
	createShoes() shoesInterface
	createClothes() clothesInterface
	createBall() ballInterface
}

type shoesInterface interface {
	describe()
}

type clothesInterface interface {
	describe()
}

type ballInterface interface {
	describe()
}
