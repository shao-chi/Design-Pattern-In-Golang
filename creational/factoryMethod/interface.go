package main

// product interface
type productInterface interface {
	operate()
}

// factory interface
type factoryInterface interface {
	createProduct() productInterface
}
