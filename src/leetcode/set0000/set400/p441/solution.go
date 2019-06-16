package main

import "fmt"

func main() {
	fmt.Println(arrangeCoins(1))
	fmt.Println(arrangeCoins(2))
	fmt.Println(arrangeCoins(5))
	fmt.Println(arrangeCoins(8))
	fmt.Println(arrangeCoins(9))
	fmt.Println(arrangeCoins(10))
}

func arrangeCoins(n int) int {

	i, j := 1, n

	for i <= j {
		m := i + (j-i)/2
		if (1+m)*m > 2*n {
			j = m - 1
		} else {
			i = m + 1
		}
	}

	return j
}
