package main

import "fmt"

func main() {
	M := [][]int{
		[]int{1, 1, 0},
		[]int{1, 1, 0},
		[]int{0, 0, 1},
	}

	fmt.Println(findCircleNum(M))
}

func findCircleNum(M [][]int) int {
	n := len(M)

	if n == 0 {
		return 0
	}
	checked := make([]bool, n)
	var dfs func(v int)

	dfs = func(v int) {
		checked[v] = true

		for w := 0; w < n; w++ {
			if !checked[w] && M[v][w] == 1 {
				dfs(w)
			}
		}
	}
	var ans = 0
	for v := 0; v < n; v++ {
		if !checked[v] {
			dfs(v)
			ans++
		}
	}

	return ans
}
