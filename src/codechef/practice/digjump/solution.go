package main

import (
	"fmt"
	"strings"
)

func main() {

	var digits string
	fmt.Scanf("%s", &digits)
	res := findBestPath(strings.TrimSpace(digits))
	fmt.Println(res)
}
func findBestPath(digits string) int {
	if len(digits) <= 1 {
		return 0
	}

	n := len(digits)

	if digits[0] == digits[n - 1] {
		return 1
	}

	dist := make([][]int, 10)
	best := make([][]int, 10)

	for i := 0; i < 10; i++ {
		dist[i] = []int{n + 1, n + 1, n + 1, n + 1, n + 1, n + 1, n + 1, n + 1, n + 1, n + 1}
		best[i] = []int{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1}
	}

	first := []int{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1}
	last := []int{n, n, n, n, n, n, n, n, n, n}

	for i := 0; i < n; i++ {
		cur := int(digits[i] - '0')
		if first[cur] == -1 {
			first[cur] = i
		}
		last[cur] = i
		for x := 0; x < 10; x++ {
			if last[x] == n {
				continue
			}
			if i-last[x] <= dist[cur][x] {
				dist[cur][x] = i - last[x]
				dist[x][cur] = i - last[x]
				best[cur][x] = i
				best[x][cur] = last[x]
			}
		}
	}

	visited := make([]bool, 10)
	cost := make([]int, 10)
	jumpTo := make([]int, 10)

	for i := 0; i < 10; i++ {
		if first[i] >= 0 {
			cost[i] = first[i] + 1
		} else {
			cost[i] = n + 1
		}
	}

	v := int(digits[0] - '0')
	visited[v] = true

	for v != int(digits[n - 1] - '0') {
		for x := 0; x < 10; x++ {
			if visited[x] {
				continue
			}
			if cost[x] >= cost[v]+dist[v][x]+1 {
				cost[x] = cost[v] + dist[v][x] + 1
				jumpTo[x] = v
			}
		}

		w, minCost := -1, n + 1
		for x := 0; x < 10; x++ {
			if visited[x] {
				continue
			}
			if cost[x] < minCost {
				minCost = cost[x]
				w = x
			}
		}
		visited[w] = true
		v = w
	}

	if first[v] == n-1 || (jumpTo[v] >= 0 && best[v][jumpTo[v]] == n-1) {
		return cost[v] - 1
	}
	return cost[v]
}
