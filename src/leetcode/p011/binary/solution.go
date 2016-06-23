package main

import "fmt"

func main() {
	fmt.Println(maxArea([]int{1, 1}))
	fmt.Println(maxArea([]int{1, 2, 4, 3}))
}

func maxArea(height []int) int {
	len := len(height)
	low := 0
	high := len - 1
	maxArea := 0
	for low < high {
		maxArea = max(maxArea,
			(high-low)*min(height[low], height[high]))
		if height[low] < height[high] {
			low++
		} else {
			high--
		}
	}
	return maxArea
}

func max(a, b int) int {
	if a >= b {
		return a
	}

	return b
}
func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
