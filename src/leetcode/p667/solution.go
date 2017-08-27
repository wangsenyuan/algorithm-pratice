package main

import "fmt"

func constructArray(n int, k int) []int {
	arr := make([]int, n)

	m := k

	for i := m; i < n; i++ {
		arr[i] = i + 1
	}

	x := k
	for i := m; i > 0; i-- {
		a := arr[i]
		b := a + x
		c := a - x
		if c > 0 {
			arr[i-1] = c
		} else {
			arr[i-1] = b
		}
		x--
	}
	return arr
}

func main() {
	fmt.Println(constructArray(6, 5))
	fmt.Println(constructArray(6, 3))
	fmt.Println(constructArray(9999, 9998))
}
