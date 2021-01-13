package main

import (
	"fmt"

	"github.com/cucumber/godog"
)

type shopping struct{}

type shoppingBasket struct {
	basket *Basket
	shelf  *Shelf
}

// Basket:
func (sb *shoppingBasket) iShouldHaveProductsInTheBasket(productCount int) error {
	if sb.basket.GetBasketSize() != productCount {
		return fmt.Errorf(
			"expected %d products but there are %d",
			sb.basket.GetBasketSize(),
			productCount,
		)
	}

	return nil
}

func (sb *shoppingBasket) theOverallBasketPriceShouldBe(basketTotal float64) error {
	if sb.basket.GetBasketTotal() != basketTotal {
		return fmt.Errorf(
			"expected basket total to be %.2f but it is %.2f",
			basketTotal,
			sb.basket.GetBasketTotal(),
		)
	}

	return nil
}

func (sb *shoppingBasket) addToBasket(productName string) (err error) {
	sb.basket.AddItem(productName, sb.shelf.GetProductPrice(productName))
	return
}

func (b *Basket) AddItem(productName string, price float64) {
	b.products[productName] = price
}

func (s *Shelf) GetProductPrice(productName string) float64 {
	return s.products[productName]
}

func (s *Shelf) AddProduct(productName string, price float64) {
	s.products[productName] = price
}

func (sb *shoppingBasket) addProduct(productName string, price float64) (err error) {
	sb.shelf.AddProduct(productName, price)
	return
}

func (b *Basket) GetBasketSize() int {
	return len(b.products)
}

func (b *Basket) GetBasketTotal() float64 {
	basketTotal := 0.00
	shippingPrice := 0.00

	for _, value := range b.products {
		basketTotal += value
	}

	basketTotal = basketTotal * 1.2

	if basketTotal <= 10 {
		shippingPrice = 3
	} else {
		shippingPrice = 2
	}

	return basketTotal + shippingPrice
}

func InitializeScenario(ctx *godog.ScenarioContext) {

	sb := &shoppingBasket{}

	ctx.BeforeScenario(func(sc *godog.Scenario) {
		sb.shelf = NewShelf()
		sb.basket = NewBasket()
	})

	// Basket:
	ctx.Step(`^I add the "([^"]*)" to the basket$`, sb.addToBasket)
	ctx.Step(`^I should have (\d+) products in the basket$`, sb.iShouldHaveProductsInTheBasket)
	ctx.Step(`^the overall basket price should be £(\d+)$`, sb.theOverallBasketPriceShouldBe)
	ctx.Step(`^there is a "([^"]*)", which costs £(\d+)$`, sb.addProduct)
}
