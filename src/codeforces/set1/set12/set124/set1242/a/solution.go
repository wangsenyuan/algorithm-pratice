package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	res := solve(n)
	fmt.Println(res)
}

func solve(n int) int {
	if n == 1 {
		return 1
	}

	var arr []int
	for x := 2; x <= n/x && len(arr) < 2; x++ {
		if n%x == 0 {
			arr = append(arr, x)
			for n%x == 0 {
				n /= x
			}
		}
	}

	if n > 1 {
		arr = append(arr, n)
	}
	if len(arr) > 1 {
		return 1
	}

	return arr[0]
}
