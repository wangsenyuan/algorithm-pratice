package p1187

import "sort"

const INF = 1 << 20

func makeArrayIncreasing(arr1 []int, arr2 []int) int {
	sort.Ints(arr2)

	n := 0

	// remove duplicates from arr2
	for i := 1; i <= len(arr2); i++ {
		if i == len(arr2) || (arr2[i] > arr2[i-1]) {
			arr2[n] = arr2[i-1]
			n++
		}
	}

	dp := make(map[int]int)
	dp[-1] = 0

	for _, num := range arr1 {
		tmp := make(map[int]int)
		for k, v := range dp {
			if num > k {
				if x, found := tmp[num]; found {
					tmp[num] = min(x, v)
				} else {
					tmp[num] = v
				}
			}
			j := sort.Search(n, func(j int) bool {
				return arr2[j] > k
			})

			if j < n {
				if x, found := tmp[arr2[j]]; found {
					tmp[arr2[j]] = min(x, v+1)
				} else {
					tmp[arr2[j]] = v + 1
				}
			}
		}
		dp = tmp
	}

	res := -1
	for _, v := range dp {
		if res < 0 || res > v {
			res = v
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
