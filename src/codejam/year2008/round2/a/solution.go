package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("./A-large-practice.in")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	var tc int
	fmt.Fscanf(f, "%d", &tc)

	for i := 1; i <= tc; i++ {
		var M, V int
		fmt.Fscanf(f, "%d %d", &M, &V)
		internal := make([][]int, M/2)
		for j := 0; j < M/2; j++ {
			internal[j] = make([]int, 2)
			fmt.Fscanf(f, "%d %d", &internal[j][0], &internal[j][1])
		}
		leaf := make([]int, (M+1)/2)
		for j := 0; j < (M+1)/2; j++ {
			fmt.Fscanf(f, "%d", &leaf[j])
		}
		res := solve(M, V, internal, leaf)
		if res < 0 {
			fmt.Printf("Case #%d: IMPOSSIBLE\n", i)
		} else {
			fmt.Printf("Case #%d: %d\n", i, res)
		}
	}
}

const INF = 1000000

func solve(M, V int, internal [][]int, leaf []int) int {

	gates := [](func([]int, []int) []int){orGate, andGate}

	var dfs func(i int) []int

	dfs = func(i int) []int {
		if 2*i > M {
			//leaf
			j := i - M/2 - 1
			res := []int{INF, INF}
			res[leaf[j]] = 0
			return res
		}
		left := dfs(2 * i)
		right := dfs(2*i + 1)
		j, can := internal[i-1][0], internal[i-1][1]
		res := gates[j](left, right)
		if can == 1 {
			res2 := gates[1-j](left, right)
			res[0] = min(res[0], res2[0]+1)
			res[1] = min(res[1], res2[1]+1)
		}
		return res
	}

	res := dfs(1)
	if res[V] >= INF {
		return -1
	}
	return res[V]
}

func orGate(a, b []int) []int {
	c := make([]int, 2)
	c[0] = a[0] + b[0]
	c[1] = a[0] + b[1]
	c[1] = min(c[1], a[1]+b[1])
	c[1] = min(c[1], a[1]+b[0])

	return c
}

func andGate(a, b []int) []int {
	c := make([]int, 2)
	c[0] = a[0] + b[0]
	c[0] = min(c[0], a[0]+b[1])
	c[0] = min(c[0], a[1]+b[0])
	c[1] = a[1] + b[1]
	return c
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
