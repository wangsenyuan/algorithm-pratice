package main

import "fmt"

func countNumbersWithUniqueDigits(n int) int {
	if n >= 11 || n == 0 {
		return 0
	}

	fx := make([]int, n+3, n+3)
	fx[0] = 0
	fx[1] = 10
	fx[2] = 91

	for i := 3; i <= n; i++ {
		fx[i] = (10-(i-1))*(fx[i-1]-fx[i-2]) + fx[i-1]
		//fmt.Println("fx[%d] = %d", i)
	}

	return fx[n]
}

func main() {
	fmt.Println(countNumbersWithUniqueDigits(0))
	fmt.Println(countNumbersWithUniqueDigits(1))
	fmt.Println(countNumbersWithUniqueDigits(4))
}
