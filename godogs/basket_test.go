package main

import "github.com/cucumber/godog"

type shopping struct{}

// Basket:
func (sh *shopping) iAddTheToTheBasket(arg1 string) error {
	return godog.ErrPending
}

func (sh *shopping) iShouldHaveProductsInTheBasket(arg1 int) error {
	return godog.ErrPending
}

func (sh *shopping) theOverallBasketPriceShouldBe(arg1 int) error {
	return godog.ErrPending
}

func (sh *shopping) thereIsAWhichCosts(arg1 string, arg2 int) error {
	return godog.ErrPending
}

func InitializeScenario(ctx *godog.ScenarioContext) {

	sh := &shopping{}

	// Basket:
	ctx.Step(`^I add the "([^"]*)" to the basket$`, sh.iAddTheToTheBasket)
	ctx.Step(`^I should have (\d+) products in the basket$`, sh.iShouldHaveProductsInTheBasket)
	ctx.Step(`^the overall basket price should be £(\d+)$`, sh.theOverallBasketPriceShouldBe)
	ctx.Step(`^there is a "([^"]*)", which costs £(\d+)$`, sh.thereIsAWhichCosts)
}
