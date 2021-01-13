package main

func NewBasket() *Basket {
	return &Basket{
		products: make(map[string]float64),
	}
}

type Basket struct {
	products map[string]float64
}
