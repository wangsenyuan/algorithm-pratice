package main

import (
	"fmt"
)

func DblLinear(n int) int {
	q2 := make([]int, 0, n)
	q3 := make([]int, 0, n)
	j, k := 0, 0
	res := 1
	for i := 0; i < n; i++ {
		q2 = append(q2, 2*res+1)
		q3 = append(q3, 3*res+1)
		a, b := q2[j], q3[k]

		res = min(a, b)
		if res == a {
			j++
		}

		if res == b {
			k++
		}
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	for i := 0; i <= 10; i++ {
		fmt.Println(DblLinear(i))
	}
	fmt.Println(DblLinear(50))
	fmt.Println(DblLinear(1000))

}
