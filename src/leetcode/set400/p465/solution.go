package main

import (
	"fmt"
	"math"
)

func main() {
	//transactions := [][]int{[]int{0, 1, 10}, []int{2, 0, 5}}
	transactions := [][]int{[]int{0, 1, 10}, []int{1, 0, 1}, []int{1, 2, 5}, []int{2, 0, 5}}
	fmt.Println(minTransfers(transactions))
}

func minTransfers(transactions [][]int) int {
	net := make(map[int]int)
	for _, tx := range transactions {
		x, y, z := tx[0], tx[1], tx[2]
		net[x] -= z
		net[y] += z
	}

	arr := make([]int, 0, len(net))

	for _, v := range net {
		if v != 0 {
			arr = append(arr, v)
		}
	}

	if len(arr) < 2 {
		return 0
	}

	if len(arr) == 2 {
		return 1
	}

	var transfer func(start int, cur int)
	ans := math.MaxInt32

	transfer = func(start int, cur int) {
		if cur >= ans {
			return
		}

		if ans == (len(arr)+1)/2 {
			return
		}
		max, at := math.MinInt32, -1

		for i := start; i < len(arr); i++ {
			if abs(arr[i]) > max {
				max = abs(arr[i])
				at = i
			}
		}

		if max == 0 || start == len(arr) {
			if cur < ans {
				ans = cur
			}
			return
		}

		arr[start], arr[at] = arr[at], arr[start]

		for i := start + 1; i < len(arr); i++ {
			if diffSign(arr[start], arr[i]) {
				tmp := arr[i]
				arr[i] += arr[start]
				transfer(start+1, cur+1)
				arr[i] = tmp
			}
		}
		arr[start], arr[at] = arr[at], arr[start]
	}

	transfer(0, 0)

	return ans
}

func diffSign(a, b int) bool {
	if a > 0 && b < 0 {
		return true
	}

	if a < 0 && b > 0 {
		return true
	}
	return false
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
