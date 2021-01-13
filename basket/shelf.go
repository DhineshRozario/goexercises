package main

// NewShelf instantiates a new Shelf object
func NewShelf() *Shelf {
	return &Shelf{
		products: make(map[string]float64),
	}
}

// Shelf stores a list of products which are available for purchase
type Shelf struct {
	products map[string]float64
}