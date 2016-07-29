package main

import "fmt"

func main() {
	fmt.Println(largestRectangleArea([]int{2, 1, 5, 6, 2, 3}))
}

func largestRectangleArea(heights []int) int {
	n := len(heights)

	stack := make([]int, n)

	p := 0
	area := 0
	for i, h := range heights {
		k := i
		for p > 0 && heights[stack[p-1]] > h {
			j := stack[p-1]
			area = max(area, heights[j]*(i-j))
			heights[j] = h
			k = j
			p--
		}
		stack[p] = k
		p++
		area = max(area, h)
	}

	for p > 0 {
		i := stack[p-1]
		area = max(area, (n-i)*heights[i])
		p--
	}

	return area
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
