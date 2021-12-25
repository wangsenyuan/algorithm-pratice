package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)

	cities := make([]string, n)
	index := make(map[string]int)
	for i := 0; i < n; i++ {
		fmt.Scanf("%s", &cities[i])
		index[cities[i]] = i
	}

	roads := make([][]int, n)
	for i := 0; i < n; i++ {
		roads[i] = make([]int, n)
	}

	var m int
	fmt.Scanf("%d", &m)

	var a, b string
	var c int

	for i := 0; i < m; i++ {
		fmt.Scanf("%s %s %d", &a, &b, &c)
		ai, bi := index[a], index[b]
		roads[ai][bi] = c
	}

	var t int
	fmt.Scanf("%d", &t)
	//fmt.Println(t)
	for i := 0; i < t; i++ {
		pathLen := solve(roads, index)
		if pathLen >= 0 {
			fmt.Println(pathLen)
		} else {
			fmt.Println("ERROR")
		}
	}
}

func solve(roads [][]int, index map[string]int) int {
	var k int
	fmt.Scanf("%d", &k)
	path := make([]string, k)
	for i := 0; i < k; i++ {
		fmt.Scanf("%s", &path[i])
	}

	visited := make(map[string]bool)

	// check spell & duplicate
	for i := 0; i < k; i++ {
		city := path[i]
		if _, found := index[city]; !found {
			return -1
		}

		if visited[city] {
			return -1
		}

		visited[city] = true
	}

	length := 0
	for i := 0; i < k-1; i++ {
		a, b := path[i], path[i+1]
		ai, bi := index[a], index[b]
		if roads[ai][bi] == 0 {
			return -1
		}
		length += roads[ai][bi]
	}

	return length
}
