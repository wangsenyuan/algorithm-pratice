package main

import "fmt"

func main() {
	var t int
	fmt.Scanf("%d", &t)

	for i := 0; i < t; i++ {
		var n, m int
		fmt.Scanf("%d %d", &n, &m)
		edges := make([][]int, m)

		for i := 0; i < m; i++ {
			var a, b int
			fmt.Scanf("%d %d", &a, &b)
			edges[i] = []int{a, b}
		}
		res, found := solve(n, m, edges)
		if !found {
			fmt.Println("IMPOSSIBLE")
		} else {
			fmt.Println(res)
		}
	}
}

func solve(n int, m int, edges [][]int) (int, bool) {

	arr := make([]int, n+1)

	ans := 0
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		if arr[a] == 0 {
			arr[a] = 1
		}

		if arr[b] == 0 {
			arr[b] = arr[a] + 1
		}

		if arr[a] > arr[b] {
			return -1, false
		}

		if arr[b] > ans {
			ans = arr[b]
		}
	}

	return ans, true
}
