package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func main() {
	var t int
	fmt.Scanf("%d", &t)
	for i := 1; i <= t; i++ {
		var n int
		fmt.Scanf("%d", &n)
		nums := make([]int, 0, 2*n)

		for j := 0; j < 2*n; j++ {
			var x int
			fmt.Scanf("%d", &x)
			nums = append(nums, x)
		}

		result := reverseSlice(play(reverseSlice(nums), n))
		fmt.Printf("Case #%d: %s\n", i, joinSliceAsString(result))
	}
}

func joinSliceAsString(xs []int) string {
	var buffer bytes.Buffer
	var sep = ""
	for _, x := range xs {
		buffer.WriteString(sep)
		buffer.WriteString(strconv.Itoa(x))
		sep = " "
	}

	return buffer.String()
}

func reverseSlice(xs []int) []int {
	n := len(xs)
	ys := make([]int, n, n)

	for i, x := range xs {
		ys[n-1-i] = x
	}

	return ys
}

func indexOf(slice []int, f func(int) bool) int {
	for i, x := range slice {
		if f(x) {
			return i
		}
	}
	return -1
}
func play(nums []int, n int) []int {
	var doIt func(nums, result []int) []int

	doIt = func(nums, result []int) []int {
		if len(nums) == 0 {
			return result
		}

		h := nums[0]
		tail := nums[1:]
		saleAt := indexOf(tail, func(x int) bool {
			return x == h*3/4
		})

		//fmt.Printf("%d at %s\n", saleAt, joinSliceAsString(tail))

		return doIt(append(tail[:saleAt], tail[saleAt+1:]...), append(result, h*3/4))
	}

	return doIt(nums, make([]int, 0, n))
}
