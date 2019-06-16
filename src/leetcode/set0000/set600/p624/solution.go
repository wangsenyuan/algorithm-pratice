package main

import "fmt"

func main() {
	arrays := [][]int{
		{1, 5},
		{3, 4},
	}

	fmt.Println(maxDistance(arrays))
}

func maxDistance(arrays [][]int) int {
	a, b := findMinMax(arrays, -1)

	if a != b {
		return arrays[b][len(arrays[b])-1] - arrays[a][0]
	}

	c, d := findMinMax(arrays, a)

	x := arrays[b][len(arrays[b])-1] - arrays[c][0]
	y := arrays[d][len(arrays[d])-1] - arrays[a][0]

	return max(x, y)
}

func findMinMax(arrays [][]int, skip int) (int, int) {
	a, b := -1, -1

	for i := 0; i < len(arrays); i++ {
		if i == skip {
			continue
		}
		if a == - 1 || arrays[i][0] < arrays[a][0] {
			a = i
		}

		if b == -1 || arrays[i][len(arrays[i])-1] > arrays[b][len(arrays[b])-1] {
			b = i
		}
	}

	return a, b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
