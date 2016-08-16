package main

import "fmt"

func main() {
	//fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	//fmt.Println(maxProfit([]int{7, 6, 5, 4, 3, 2, 1}))
	//fmt.Println(maxProfit([]int{1, 2, 3, 4, 5, 6}))
	//fmt.Println(maxProfit([]int{1, 1, 1, 1, 1}))
	//fmt.Println(maxProfit([]int{2, 1, 2, 0, 1}))
	//fmt.Println(maxProfit([]int{3, 2, 6, 5, 0, 3}))
	fmt.Println(maxProfit([]int{2, 1, 2, 1, 0, 1, 2}))
}

func maxProfit(prices []int) int {
	st := make([]int, len(prices))
	i := 0
	profit := 0

	for _, price := range prices {
		for i > 0 && price < st[i-1] {
			profit = max(st[i-1]-st[0], profit)
			i--
		}
		st[i] = price
		i++
	}

	if i > 0 {
		profit = max(st[i-1]-st[0], profit)
	}

	return profit
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
