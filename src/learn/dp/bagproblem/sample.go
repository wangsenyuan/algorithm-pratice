package main

import "fmt"

func main() {
	ws := []int{5, 10, 20}
	vs := []int{50, 60, 140}
	v := dp(ws, vs, 30)
	fmt.Printf("%d\n", v)
}

func dp(ws []int, vs []int, weight int) int {
	n := len(ws)

	fx := make([]int, weight+1)
	//f[i][w] = max(f[i][w], f[i - 1][w - wi] + vi| j from 1 until i)

	for i := 1; i <= n; i++ {
		for w := weight; w >= ws[i-1]; w-- {
			fx[w] = max(fx[w], fx[w-ws[i-1]]+vs[i-1])
		}
	}

	return fx[weight]
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func backTrack(ws []int, vs []int, W int) int {
  var knapsack func(i int, profit int, weight int) int
  maxPrifit := 0

  knapsack = func(i int, profit int, weight int) int {

  }
}

}
