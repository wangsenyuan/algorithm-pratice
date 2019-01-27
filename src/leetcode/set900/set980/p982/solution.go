package p982

func countTriplets(A []int) int {
	N := 1 << 16
	dp := make([][]int, 3)
	for i := 0; i < 3; i++ {
		dp[i] = make([]int, N)
	}

	for i := 0; i < len(A); i++ {
		dp[0][A[i]]++
	}

	for i := 1; i < 3; i++ {
		for _, a := range A {
			for b := 0; b < N; b++ {
				dp[i][a&b] += dp[i-1][b]
			}
		}
	}
	return dp[2][0]
}

func countTriplets1(A []int) int {
	n := len(A)

	var res int

	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			for k := j; k < n; k++ {
				if A[i]&A[j]&A[k] == 0 {
					if i == j && i == k {
						res++
					} else if i == j || j == k {
						res += 3
					} else {
						res += 6
					}
				}
			}
		}
	}

	return res
}
