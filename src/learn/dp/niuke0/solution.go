package niuke0

func numberOf1BeforeN(n int) int {
	var nums []int

	num := n

	for num > 0 {
		nums = append(nums, num%10)
		num /= 10
	}
	m := len(nums)

	P := make([]int, 10)
	P[0] = 1
	for i := 1; i < 10; i++ {
		P[i] = P[i-1] * 10
	}

	dp := make([][][]int, 2)
	for i := 0; i < 2; i++ {
		dp[i] = make([][]int, 10)
		for j := 0; j < 10; j++ {
			dp[i][j] = make([]int, 2)
		}
	}

	dp[0][0][1] = 1

	var res int

	for i := 0; i < m; i++ {
		a := i & 1
		b := 1 - a
		for e := 0; e < 2; e++ {
			for d := 0; d < 10; d++ {
				dp[b][d][e] = 0
			}
		}

		for e := 0; e < 2; e++ {
			for d := 0; d < 10; d++ {
				for dd := 0; dd < 10; dd++ {
					if e == 1 && dd > nums[m-1-i] {
						break
					}
					var f int
					if e == 1 && dd == nums[m-1-i] {
						f = 1
					}
					dp[b][dd][f] += dp[a][d][e]
					if dd == 1 {
						if f == 1 {
							res += dp[a][d][e] * (n%P[m-1-i] + 1)
						} else {
							res += dp[a][d][e] * P[m-1-i]
						}
					}
				}
			}
		}
	}

	return res
}
