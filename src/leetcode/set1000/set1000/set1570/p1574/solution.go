package p1574

import "sort"

func findLengthOfShortestSubarray(arr []int) int {
	n := len(arr)
	j := n - 1
	for j > 0 && arr[j-1] <= arr[j] {
		j--
	}
	if j == 0 {
		return 0
	}
	// j is the starting point of non-decreasing suffix
	var best = j
	for i := 0; i < j; i++ {
		// if we delete after i, the answer is
		x := arr[i]
		k := sort.Search(n-j, func(k int) bool {
			return arr[j+k] >= x
		})
		k += j
		if k-i-1 < best {
			best = k - i - 1
		}
		if arr[i+1] < arr[i] {
			break
		}
	}
	return best
}
