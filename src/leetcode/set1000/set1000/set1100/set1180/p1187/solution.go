package p1187

import "sort"

const INF = 1 << 30

func makeArrayIncreasing2(arr1 []int, arr2 []int) int {
	m := len(arr1)
	arr3 := make([]int, len(arr2))
	copy(arr3, arr2)
	// arr3[len(arr3)-1] = INF
	sort.Ints(arr3)
	n := 0

	// remove duplicates from arr2
	for i := 1; i <= len(arr3); i++ {
		if i == len(arr3) || (arr3[i] > arr3[i-1]) {
			arr3[n] = arr3[i-1]
			n++
		}
	}
	// dp[i][j] = (arr1[i] ends at arr2[j], operations)
	// fp[j] = min result before j
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = INF
		}
	}
	for j := 0; j < n; j++ {
		if arr1[0] > arr3[j] {
			dp[0][j] = 1
		} else {
			dp[0][j] = 0
		}
	}

	for i := 0; i < m-1; i++ {
		for j := 0; j < n-1; j++ {
			if dp[i][j] < INF {
				if arr1[i+1] > arr3[j] && arr1[i+1] <= arr3[j+1] {
					dp[i+1][j+1] = dp[i][j]
				} else {
					dp[i+1][j+1] = dp[i][j] + 1
				}
			}
		}

		for j := 1; j < n; j++ {
			dp[i+1][j] = min(dp[i+1][j], dp[i+1][j-1])
		}
	}

	res := dp[m-1][n-1]

	for i := m - 2; i >= 0 && arr1[i+1] > arr1[i]; i-- {
		for j := n - 1; j >= 0; j-- {
			if arr1[i+1] > arr3[j] {
				res = min(res, dp[i][j])
			}
		}
	}

	if res >= INF {
		return -1
	}
	return res
}

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
