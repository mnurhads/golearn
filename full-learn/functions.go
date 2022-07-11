package main

import (
	"fmt"
	"math"
	"errors"
)
func avg(x float64, y float64) float64 {
	return (x + y) / 2
}

func getStockPriceChange(prevPrice, currentPrice float64) (float64, float64) {
	change := currentPrice - prevPrice
	percentChange := (change / prevPrice) * 100
	return change, percentChange
}

func getStockPriceChangeWithError(prevPrices, currentPrices float64) (float64, float64, error) {
	if prevPrices == 0 {
		err := errors.New("Previous price cannot be zero")
		return 0, 0, err
	}
	changes := currentPrices - prevPrices
	percentChanges := (changes / prevPrices) * 100
	return changes, percentChanges, nil
}

func main() {
	x := 5.75
	y := 6.25

	result := avg(x, y)

	fmt.Printf("Average of %.2f and %.2f = %.2f\n", x, y, result)

	// new
	prevStockPrice := 0.0
	currentStockPrice := 100000.0

	change, percentChange := getStockPriceChange(prevStockPrice, currentStockPrice)

	if change < 0 {
		fmt.Printf("The Stock Price decreased by $%.2f which is %.2f%% of the prev price\n", math.Abs(change), math.Abs(percentChange))
	} else {
		fmt.Printf("The Stock Price increased by $%.2f which is %.2f%% of the prev price\n", change, percentChange)
	}

	// new
	prevStockPrices := 0.0
	currentStockPrices := 100000.0

	changes, percentChanges, err := getStockPriceChangeWithError(prevStockPrices, currentStockPrices)

	if err != nil {
		fmt.Println("Sorry! There was an error: ", err)
	} else {
		if change < 0 {
			fmt.Printf("The Stock Price decreased by $%.2f which is %.2f%% of the prev price\n", math.Abs(changes), math.Abs(percentChanges))
		} else {
			fmt.Printf("The Stock Price increased by $%.2f which is %.2f%% of the prev price\n", changes, percentChanges)
		}
	}
}