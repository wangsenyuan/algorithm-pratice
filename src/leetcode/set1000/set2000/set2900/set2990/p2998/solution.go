package p2998

func numberOfPowerfulInt(start int64, finish int64, limit int, s string) int64 {
	suf := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		suf[i] = int(s[i] - '0')
	}

	res := solve(finish, suf, limit)
	res -= solve(start-1, suf, limit)
	return res
}

func solve(num int64, suf []int, limit int) int64 {
	var arr []int
	for num > 0 {
		arr = append(arr, int(num%10))
		num /= 10
	}
	reverse(arr)
	if len(arr) < len(suf) {
		return 0
	}

	dp := make([]int64, 2)
	dp[1] = 1

	m := len(arr)
	n := len(suf)

	for i := 0; i < m; i++ {
		fp := make([]int64, 2)
		for j := 0; j < 2; j++ {
			if i < m-n {
				for u := 0; u <= limit; u++ {
					nj := 0
					if j == 1 {
						if u > arr[i] {
							break
						} else if u == arr[i] {
							nj = 1
						}
					}
					fp[nj] += dp[j]
				}
			} else {
				u := suf[n-(m-i)]
				nj := 0
				if j == 1 {
					if u > arr[i] {
						continue
					} else if u == arr[i] {
						nj = 1
					}
				}
				fp[nj] += dp[j]
			}
		}
		copy(dp, fp)
	}

	return dp[0] + dp[1]
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
