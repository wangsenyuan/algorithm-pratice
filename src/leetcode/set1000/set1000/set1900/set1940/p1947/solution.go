package p1947

func maxCompatibilitySum(students [][]int, mentors [][]int) int {
	m := len(students)
	n := len(students[0])
	S := make([]int, m)
	T := make([]int, m)

	for i := 0; i < m; i++ {
		S[i] = getAsBinary(students[i])
		T[i] = getAsBinary(mentors[i])
	}

	X := make([][]int, m)
	for i := 0; i < m; i++ {
		X[i] = make([]int, m)
		for j := 0; j < m; j++ {
			tmp := S[i] ^ T[j]
			X[i][j] = n
			for tmp > 0 {
				X[i][j] -= tmp & 1
				tmp >>= 1
			}
		}
	}

	// dp[i][state] = 匹配前i个学生，用state表示的mentors所能得到的最大值
	M := 1 << m
	dp := make([][]int, m+1)
	dp[0] = make([]int, M)
	for i := 1; i <= m; i++ {
		dp[i] = make([]int, M)
		cur := (1 << i) - 1
		last := cur << (m - i)

		for cur <= last {
			for j := 0; j < m; j++ {
				if cur&(1<<j) > 0 {
					// pair up i - 1 & j
					dp[i][cur] = max(dp[i][cur], X[i-1][j]+dp[i-1][cur^(1<<j)])
				}
			}
			cur = snoob(cur)
		}
	}
	return dp[m][M-1]
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func getAsBinary(arr []int) int {
	var res int
	n := len(arr)
	for i := 0; i < n; i++ {
		if arr[i] == 1 {
			res |= 1 << (n - 1 - i)
		}
	}
	return res
}

// this function returns next higher number with same number of set bits as x.
func snoob(x int) int {
	rightOne := x & -x
	nextHigherOneBit := x + rightOne
	rightOnesPattern := x ^ nextHigherOneBit
	rightOnesPattern = (rightOnesPattern) / rightOne
	rightOnesPattern >>= 2
	next := nextHigherOneBit | rightOnesPattern
	return next
}
