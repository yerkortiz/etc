package main

import (
	"fmt"
	"math/rand"
	"time"
	"yerkortiz/fintual/portfolio/portfolio"
)

// // create stock with random price
func createStock(startDate time.Time, endDate time.Time) portfolio.Stock {
	stock := portfolio.NewStock()
	stock.AddPrice(startDate, rand.Int()%(1<<16))
	stock.AddPrice(endDate, rand.Int()%(1<<16))
	return stock
}

// create stock with random price
// uncomment to pass the stock prices manually
// func createStock(startDate time.Time, endDate time.Time) portfolio.Stock {
// 	stock := portfolio.NewStock()
// 	startPrice := 0
// 	endPrice := 0
// 	fmt.Print("Enter the startPrice of the stock: ")
// 	if _, err := fmt.Scanf("%d\n", &startPrice); err != nil {
// 		fmt.Printf("Error reading startPrice of the stock: %v\n", err)
// 		return portfolio.Stock{}
// 	}
// 	fmt.Print("Enter the endPrice of the stock: ")
// 	if _, err := fmt.Scanf("%d\n", &endPrice); err != nil {
// 		fmt.Printf("Error reading endPrice of the stock: %v\n", err)
// 		return portfolio.Stock{}
// 	}
// 	stock.AddPrice(startDate, startPrice)
// 	stock.AddPrice(endDate, endPrice)
// 	return stock
// }

func main() {
	// instance input variables
	rand.New(rand.NewSource(40))
	nStocks := 0
	startDateStr := ""
	endDateStr := ""
	fmt.Printf("welcome to portfolio.go: you can try different combinations of dates and stock prices")
	fmt.Print("Enter number of stocks: ")
	if _, err := fmt.Scanf("%d\n", &nStocks); err != nil {
		fmt.Printf("Error reading number of stocks: %v\n", err)
		return
	}

	fmt.Print("Enter start date (YYYY-MM-DD): ")
	if _, err := fmt.Scanf("%s\n", &startDateStr); err != nil {
		fmt.Printf("Error reading start date: %v\n", err)
		return
	}

	fmt.Print("Enter end date (YYYY-MM-DD): ")
	if _, err := fmt.Scanf("%s\n", &endDateStr); err != nil {
		fmt.Printf("Error reading end date: %v\n", err)
		return
	}

	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		fmt.Printf("Error parsing start date: %v\n", err)
		return
	}

	endDate, err := time.Parse("2006-01-02", endDateStr)
	if err != nil {
		fmt.Printf("Error parsing end date: %v\n", err)
		return
	}

	// create stocks
	stocks := make([]portfolio.Stock, 0, nStocks)
	for i := 0; i < nStocks; i++ {
		stock := createStock(startDate, endDate)
		stocks = append(stocks, stock)
	}

	// create portfolio
	portfolio := portfolio.NewPortfolio(stocks)

	// calculate profit
	profit, err := portfolio.Profit(startDate, endDate)
	if err != nil {
		fmt.Printf("Error calculating profit: %v\n", err)
		return
	}
	fmt.Printf("Profit: %d\n", profit)

	// calculate annualized return
	annualReturn, err := portfolio.AnnualizedReturn(startDate, endDate)
	if err != nil {
		fmt.Printf("Error calculating annualized return: %v\n", err)
		return
	}
	fmt.Printf("Annualized Return: %.2f\n", annualReturn)
}
