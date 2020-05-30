package main

import "fmt"

func main() {
	fmt.Println(largestRectangleArea([]int{2, 1, 5, 6, 2, 3}))
}

func largestRectangleArea(heights []int) int {
	n := len(heights)
	left := make([]int, n)
	right := make([]int, n)
	for i := 0; i < n; i++ {
		right[i] = n
		left[i] = -1
	}

	stack := make([]int, n)
	var p int

	for i := 0; i < n; i++ {
		for p > 0 && heights[stack[p-1]] >= heights[i] {
			right[stack[p-1]] = i
			p--
		}
		if p > 0 {
			left[i] = stack[p-1]
		}
		stack[p] = i
		p++
	}
	var res int
	for i := 0; i < n; i++ {
		res = max(res, max(1, (right[i]-left[i]-1))*heights[i])
	}
	return res
}

func largestRectangleArea1(heights []int) int {
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
