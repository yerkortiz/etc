package portfolio

import (
	"fmt"
	"log/slog"
	"math"
	"time"
)

const (
	layoutISO     = "2006-01-02" // ISO format for date strings
	dayInHours    = 24           // Hours in a day
	avgDaysInYear = 365.2422     // Average days in a year
)

var (
	// error when a in a range of dates the endDate is less than or equal to startDate
	invalidDateRangeErr = fmt.Errorf("invalid date range")
	// error when portfolio value is zero
	portfolioZeroValueErr = fmt.Errorf("portfolio has zero value")
)

// Portfolio contains a collection of Stocks
type Portfolio struct {
	stocks []Stock
}

// Stock contains prices mapped by date, where date is a string and the price an integer.
type Stock struct {
	datePriceMap map[string]int
}

// Profit returns the sum of the price difference between two dates for a Stock's collection.
func (p *Portfolio) Profit(startDate time.Time, endDate time.Time) (int, error) {
	if endDate.Before(startDate) || endDate.Sub(startDate).Hours() == 0 {
		return 0, invalidDateRangeErr
	}

	profit := 0
	for _, stock := range p.stocks {
		startPrice, ok := stock.Price(startDate)
		// if there is no price at given date, the function will continue its execution,
		// considerating the profit of the current stock as zero.
		if !ok {
			slog.Warn("there is no price available for the given date", "date", startDate.Format(layoutISO))
			continue
		}

		endPrice, ok := stock.Price(endDate)
		if !ok {
			slog.Warn("there is no price available for the given date", "date", endDate.Format(layoutISO))
			continue
		}

		profit += (endPrice - startPrice)
	}
	return profit, nil
}

// AnnualizedReturn calculates the annualized return between two dates.
func (p *Portfolio) AnnualizedReturn(startDate, endDate time.Time) (float64, error) {
	if endDate.Before(startDate) || endDate.Sub(startDate).Hours() == 0 {
		return 0, invalidDateRangeErr
	}
	startValue := 0
	endValue := 0
	for _, stock := range p.stocks {
		startPrice, ok := stock.Price(startDate)
		if !ok {
			slog.Warn("there is no price available for the given date", "date", startDate.Format(layoutISO))
			continue
		}

		endPrice, ok := stock.Price(endDate)
		if !ok {
			slog.Warn("there is no price available for the given date", "date", endDate.Format(layoutISO))
			continue
		}

		startValue += startPrice
		endValue += endPrice
	}

	if startValue == 0 || endValue == 0 {
		return 0.0, portfolioZeroValueErr
	}
	totalValue := float64(endValue-startValue) / float64(startValue)

	// there is already a validation when the expression endDate.Sub(startDate).Hours() is equal to zero
	// at the start of the function, so there is no need to validate if years is equal to zero
	years := hoursToYears(endDate.Sub(startDate).Hours())

	return math.Pow(1+totalValue, 1/years) - 1, nil
}

// Price returns the price of the stock for a given date, and a bool to handle cases
// when there is no existing prices for a given date
func (s *Stock) Price(date time.Time) (int, bool) {
	price, exists := s.datePriceMap[date.Format(layoutISO)]
	return price, exists
}

// Instance a new portfolio with an array of stocks
func NewPortfolio(stocks []Stock) Portfolio {
	return Portfolio{
		stocks: stocks,
	}
}

// Instance a zero initialized Stock struct
func NewStock() Stock {
	return Stock{
		datePriceMap: map[string]int{},
	}
}

// AddPrice adds a new price to the Stock for a given date,
// if there is already a price for a given date, it will overwrite it.
func (s *Stock) AddPrice(date time.Time, price int) {
	s.datePriceMap[date.Format(layoutISO)] = price
}

// X hours are equal to X/24 days
// Y days are equal to Y/365.2422 years
// there is no risk of division by zero because
// dayInHours and avgDaysInYear are constants
// distinct to zero
func hoursToYears(hours float64) float64 {
	return hours / dayInHours / avgDaysInYear
}
