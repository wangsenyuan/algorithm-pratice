package main

import (
	"fmt"
	"math"
)

func main() {
	prices := []int{1, 2, 1, 3, 0, 4}
	fmt.Println(maxProfit2(prices))
}

func maxProfit2(prices []int) int {
	firstBuy, firstSell := math.MinInt32, 0
	secondBuy, secondSell := math.MinInt32, 0

	for _, price := range prices {
		if firstBuy < -price {
			firstBuy = -price
		}

		if firstSell < firstBuy+price {
			firstSell = firstBuy + price
		}

		if secondBuy < firstSell-price {
			secondBuy = firstSell - price
		}

		if secondSell < secondBuy+price {
			secondSell = secondBuy + price
		}
	}

	return secondSell
}

func maxProfit(prices []int) int {
	prices = compact1(prices)
	prices = compact2(prices)
	if len(prices) == 0 {
		return 0
	}
	n := len(prices)
	fx := make([][]int, n)

	for i := range fx {
		fx[i] = make([]int, n)
	}

	for k := 1; k < n; k++ {
		for i := 0; i+k < n; i++ {
			j := i + k
			fx[i][j] = max(0, prices[j]-prices[i])
			if k == 1 {
				continue
			}
			fx[i][j] = max(fx[i][j], fx[i+1][j])
			fx[i][j] = max(fx[i][j], fx[i][j-1])
		}
	}

	r := fx[0][n-1]
	for k := 1; k < n-1; k++ {
		if fx[0][k]+fx[k][n-1] > r {
			r = fx[0][k] + fx[k][n-1]
		}
	}

	return r
}

func max(a, b int) int {
	if a >= b {
		return a
	}

	return b
}

func compact1(prices []int) []int {
	com := make([]int, 0, len(prices))
	for i := 0; i < len(prices); {
		com = append(com, prices[i])
		j := i
		for j < len(prices)-1 && prices[j+1] >= prices[j] {
			j++
		}
		if j > i {
			com = append(com, prices[j])
			i = j + 1
		} else {
			i++
		}
	}

	return com
}

func compact2(prices []int) []int {
	com := make([]int, 0, len(prices))
	for i := 0; i < len(prices); {
		com = append(com, prices[i])
		j := i
		for j < len(prices)-1 && prices[j+1] <= prices[j] {
			j++
		}
		if j > i {
			com = append(com, prices[j])
			i = j + 1
		} else {
			i++
		}
	}

	return com
}
