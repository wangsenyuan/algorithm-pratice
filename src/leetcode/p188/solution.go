package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{1, 2, 4}
	//fmt.Println(len(nums))
	fmt.Println(maxProfit(2, nums))
	//fmt.Println(math.MaxInt32)
	//fmt.Println(1000000000)
}

func maxProfit(k int, prices []int) int {
	if k == 0 {
		return 0
	}

	if k > len(prices) {
		k = len(prices)
	}

	profits := make([]int, 2*k)

	//profits[0] = math.MinInt32
	for i := 0; i < k; i++ {
		profits[2*i] = math.MinInt32
	}

	for _, price := range prices {
		if -price > profits[0] {
			profits[0] = -price
		}

		if profits[0]+price > profits[1] {
			profits[1] = profits[0] + price
		}

		for i := 1; i < k; i++ {
			if profits[2*i-1]-price > profits[2*i] {
				profits[2*i] = profits[2*i-1] - price
			}

			if profits[2*i]+price > profits[2*i+1] {
				profits[2*i+1] = profits[2*i] + price
			}
		}
	}

	return profits[2*k-1]
}
