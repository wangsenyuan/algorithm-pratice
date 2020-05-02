package countdigitone

func countDigitOne(n int) int {

	digits := make([]int, 11)

	P := make([]int, 11)
	P[0] = 1

	var h int

	for num := n; num > 0; num /= 10 {
		digits[h] = num % 10
		h++
		P[h] = 10 * P[h-1]
	}

	for i, j := 0, h-1; i < j; i, j = i+1, j-1 {
		digits[i], digits[j] = digits[j], digits[i]
	}

	dp := make([][]int, h+1)

	for i := 0; i <= h; i++ {
		dp[i] = make([]int, 2)
	}

	var res int

	dp[0][1] = 1

	for i := 0; i < h; i++ {
		for tight := 0; tight <= 1; tight++ {
			if dp[i][tight] == 0 {
				continue
			}

			for x := 0; x <= 9; x++ {
				if tight == 1 && x > digits[i] {
					break
				}
				var newTight int
				if tight == 1 && x == digits[i] {
					newTight = 1
				}
				dp[i+1][newTight] += dp[i][tight]

				if x == 1 {
					if newTight == 1 {
						res += n%(P[h-(i+1)]) + 1
					} else {
						res += dp[i][tight] * (P[h-(i+1)])
					}
				}
			}
		}
	}

	return res
}
