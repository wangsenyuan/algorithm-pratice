package main

import "fmt"

func NbMonths(startPriceOld, startPriceNew, savingperMonth int, percentLossByMonth float64) [2]int {
	var process func(priceOld float64, priceNew float64, percentLoss float64, months int) (int, float64)

	process = func(priceOld float64, priceNew float64, percentLoss float64, months int) (int, float64) {
		if priceOld+float64(months)*float64(savingperMonth) >= priceNew {
			return months, priceOld + float64(months)*float64(savingperMonth) - priceNew
		}

		if months%2 == 1 {
			percentLoss += 0.5
		}

		return process(priceOld*(1.0-percentLoss*0.01), priceNew*(1.0-percentLoss*0.01), percentLoss, months+1)
	}

	months, money := process(float64(startPriceOld), float64(startPriceNew), percentLossByMonth, 0)
	//fmt.Printf("%d %.2f\n", months, money)
	return [...]int{months, int(money + 0.5)}
}

func main() {
	fmt.Println(NbMonths(2000, 8000, 1000, 1.5))
	fmt.Println(NbMonths(8000, 12000, 500, 1.0))
	fmt.Println(NbMonths(18000, 32000, 1500, 1.25))
	fmt.Println(NbMonths(7500, 32000, 300, 1.55))
}
