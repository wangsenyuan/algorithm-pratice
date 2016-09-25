package main

import "fmt"

func main() {
	heightMap := [][]int{
		[]int{1, 4, 3, 1, 3, 2},
		[]int{3, 2, 1, 3, 2, 4},
		[]int{2, 3, 3, 2, 3, 1},
	}

	fmt.Println(trapRainWater(heightMap))
}

func trapRainWater(heightMap [][]int) int {
	m := len(heightMap)
	if m <= 2 {
		return 0
	}
	n := len(heightMap[0])
	if n <= 2 {
		return 0
	}
	capacity := make([][]int, m)
	for i := 0; i < m; i++ {
		capacity[i] = make([]int, n)
	}

	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			a := max(capacity[i - 1][j], heightMap[i - 1][j])
			b := max(capacity[i][j - 1], heightMap[i][j - 1])
			capacity[i][j] = min(a, b)
		}
	}

	for i := m - 2; i >= 1; i-- {
		for j := m - 2; j >= 1; j-- {
			a := max(capacity[i + 1][j], heightMap[i + 1][j])
			b := max(capacity[i][j + 1], heightMap[i][j + 1])
			c := min(a, b)
			if c < capacity[i][j] {
				capacity[i][j] = c
			}
		}
	}

	v := 0

	for a := 1; a < m-1; a++ {
		cp := capacity[a]
		height := heightMap[a]
		for i, j := 0, n-1; i < j; {
			if height[i] <= height[j] {
				k := i + 1
				for ; height[i] > height[k]; k++ {
					u := min(height[i], cp[k]) - height[k]
					if u > 0 {
						v += u
					}
				}
				i = k
			} else {
				k := j - 1
				for ; height[j] > height[k]; k-- {
					u := min(height[j], cp[k]) - height[k]
					if u > 0 {
						v += u
					}
				}
				j = k
			}
		}
	}
	return v
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
