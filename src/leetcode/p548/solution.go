package main

import (
	"fmt"
)

func main() {
	fmt.Println(splitArray([]int{1, 2, 1, 2, 1, 2, 1}))
	fmt.Println(splitArray([]int{1, 2, 1, 2, 1, 2, 1, 2}))

	fmt.Println(splitArray([]int{1, 2, -1, 1, 2, 5, 2, 5, 2}))

}

func splitArray(nums []int) bool {
	n := len(nums)
	sums := make([]int64, n+1)

	for i, num := range nums {
		sums[i+1] = sums[i] + int64(num)
	}

	for j := 3; j < n-3; j++ {
		firstHalf := make(map[int64]bool)

		for i := 1; i+1 < j; i++ {
			a := sums[i]
			b := sums[j] - sums[i+1]
			if a == b {
				firstHalf[a] = true
			}
		}

		for k := j + 2; k < n; k++ {
			a := sums[k] - sums[j+1]
			b := sums[n] - sums[k+1]
			//fmt.Printf("%d %d\n", a, b)
			if a == b && firstHalf[a] {
				return true
			}
		}

	}

	return false
}
