package p1049

import "sort"

func lastStoneWeightII(stones []int) int {
	sort.Ints(stones)
	dp := make([]bool, 1001)
	fp := make([]bool, 1001)

	dp[0] = true

	for i := 0; i < len(stones); i++ {
		for j := 0; j <= 1000; j++ {
			fp[j] = false
		}

		for j := 0; j <= 1000; j++ {
			if dp[j] {
				fp[abs(j-stones[i])] = true

				if j+stones[i] <= 1000 {
					fp[j+stones[i]] = true
				}
			}
		}

		dp, fp = fp, dp
	}

	for i := 0; i <= 1000; i++ {
		if dp[i] {
			return i
		}
	}

	return -1
}


func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}