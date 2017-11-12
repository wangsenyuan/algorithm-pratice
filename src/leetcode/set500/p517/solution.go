package main

import "fmt"

func main() {
	machines := []int{4, 0, 0, 4}
	fmt.Println(findMinMoves(machines))
}

func findMinMoves(machines []int) int {
	n := len(machines)
	sum := make([]int, n + 1)

	for i, num := range machines {
		sum[i + 1] = sum[i] + num
	}

	if sum[n]%n != 0 {
		return -1
	}

	avg := sum[n] / n
	res := 0
	for i := 0; i < n; i++ {
		l := i*avg - sum[i]
		r := (n-i-1)*avg - (sum[n] - sum[i + 1])

		if l > 0 && r > 0 {
			if l+r > res {
				res = l + r
			}
		} else {
			if l > res {
				res = l
			}
			if r > res {
				res = r
			}
		}
	}

	return res
}
