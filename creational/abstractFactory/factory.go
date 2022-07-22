package main

type basketballFactory struct{}
type volleyballFactory struct{}

func (factory basketballFactory) createShoes() shoesInterface {
	return basketballShoes{}
}

func (factory basketballFactory) createClothes() clothesInterface {
	return basketballClothes{}
}

func (factory basketballFactory) createBall() ballInterface {
	return basketball{}
}

func (factory volleyballFactory) createShoes() shoesInterface {
	return volleyballShoes{}
}

func (factory volleyballFactory) createClothes() clothesInterface {
	return volleyballClothes{}
}

func (factory volleyballFactory) createBall() ballInterface {
	return volleyball{}
}
