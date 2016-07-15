package main

import "fmt"

func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}

func trap(height []int) int {
	n := len(height)
	if n <= 2 {
		return 0
	}

	v := 0
	for i, j := 0, n-1; i < j; {
		if height[i] <= height[j] {
			k := i + 1
			for ; height[i] > height[k]; k++ {
				v += height[i] - height[k]
			}
			i = k
		} else {
			k := j - 1
			for ; height[j] > height[k]; k-- {
				v += height[j] - height[k]
			}
			j = k
		}
	}

	return v
}
