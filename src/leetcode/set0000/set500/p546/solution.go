package main

import "fmt"

func main() {
	fmt.Println(removeBoxes([]int{1, 3, 2, 2, 2, 3, 4, 3, 1}))
	fmt.Println(removeBoxes([]int{3, 2, 2, 2, 3}))
}

func removeBoxes(boxes []int) int {

	n := len(boxes)

	dp := make([][][]int, n)
	for i := range dp {
		dp[i] = make([][]int, n)
		for j := range dp[i] {
			dp[i][j] = make([]int, n)
		}
	}

	//dp[i][j][k] means from box i to j, having k boxes, the maximum points

	var process func(i, j, k int) int

	process = func(i, j, k int) int {
		if i > j {
			return 0
		}
		if i == j {
			return k * k
		}

		if dp[i][j][k] > 0 {
			return dp[i][j][k]
		}

		ans := process(i+1, j, 1) + k*k
		for m := i + 1; m <= j; m++ {
			if boxes[m] == boxes[i] {
				subAns := process(i + 1, m-1, 1) + process(m, j, k+1)
				if subAns > ans {
					ans = subAns
				}
			}
		}
		dp[i][j][k] = ans
		return ans
	}

	return process(0, n-1, 1)
}
