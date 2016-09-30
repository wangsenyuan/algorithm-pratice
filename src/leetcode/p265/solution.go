package main

import "fmt"

func main() {
	costs := [][]int{[]int{1, 5, 3}, []int{2, 9, 4}}
	fmt.Println(minCostII(costs))
}

func minCostII(costs [][]int) int {
	m := len(costs)
	if m == 0 {
		return 0
	}

	n := len(costs[0])
	if n == 0 {
		return 0
	}

	selectedIndex := -1
	firstMin, secondMin := 0, 0

	for i := 0; i < m; i++ {
		k := selectedIndex
		a, b := -1, -1
		for j := 0; j < n; j++ {
			cost := costs[i][j]
			if selectedIndex == -1 || selectedIndex != j {
				cost += firstMin
			} else {
				cost += secondMin
			}
			if a == -1 || cost < a {
				a, b = cost, a
				k = j
			} else if b == -1 || cost < b {
				b = cost
			}
		}
		selectedIndex, firstMin, secondMin = k, a, b
	}

	return firstMin
}
