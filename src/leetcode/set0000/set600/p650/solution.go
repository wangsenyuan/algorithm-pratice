package main

import (
	"math"
	"fmt"
)

func minSteps(n int) int {

	f := make([][]int, n+1, n+1)

	for i := 0; i <= n; i++ {
		f[i] = make([]int, n+1)
		for j := 0; j <= n; j++ {
			f[i][j] = math.MaxInt32
		}
	}

	f[1][0] = 0
	f[1][1] = 1

	for i := 1; i <= n; i++ {
		for j := 1; j < i; j++ {
			f[i][i] = min(f[i][i], f[i][j]+1)
		}
		for j := 1; j <= i && i+j <= n; j++ {
			f[i+j][j] = min(f[i+j][j], f[i][j]+1)
		}
	}

	res := math.MaxInt32

	for j := 0; j <= n; j++ {
		res = min(res, f[n][j])
	}
	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func main() {
	fmt.Println(minSteps(1))
	fmt.Println(minSteps(3))

	fmt.Println(minSteps(10))

	//fmt.Println(minSteps(100))
	fmt.Println(minSteps(1000))

}
