package p2607

import "sort"

func makeSubKSumEqual(arr []int, k int) int64 {
	n := len(arr)

	// i % k 相同位的要一致
	// [1, 2, 3, 4, 5], and k = 3
	// 1 & 4, 2 & 5, 3 & 1, 4 & 2, 所有的要一样
	// [1, 2, 3, 4, 5, 6], and k = 4
	// 1 & 5, 2 & 6, 3 & 1, 4 & 2, 5 & 3
	var res int64
	vis := make([]bool, n)
	buf := make([]int, n)
	for i := 0; i < k; i++ {
		if vis[i] {
			continue
		}
		var p int
		for j := i; !vis[j]; j = (j + k) % n {
			vis[j] = true
			buf[p] = arr[j]
			p++
		}
		x := findMedian(buf[:p])
		for j := 0; j < p; j++ {
			res += int64(abs(x - buf[j]))
		}
	}

	return res
}

func findMedian(arr []int) int {
	sort.Ints(arr)
	n := len(arr)
	return arr[n/2]
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
