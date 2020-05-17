package p1448

func largestNumber(cost []int, target int) string {
	pay := make([]int, 10)

	for i := 1; i < 10; i++ {
		pay[i] = cost[i-1]
	}
	// dp[x] = s
	// target 是 x 的时候，所能得到的最大的数字s
	dp := make([]string, target+1)
	dp[0] = ""
	for i := 1; i <= target; i++ {
		dp[i] = "-"
	}

	for x := 0; x < target; x++ {
		if dp[x] == "-" {
			continue
		}
		for i := 1; i < 10; i++ {
			if x+pay[i] <= target {
				dp[x+pay[i]] = maxNumber(dp[x+pay[i]], dp[x], i)
			}
		}
	}
	if dp[target] == "-" {
		return "0"
	}
	return dp[target]
}

func maxNumber(s string, x string, num int) string {
	y := insertNum(x, num)

	if s == "-" || len(s) < len(y) {
		return y
	}

	if len(s) == len(y) {
		var i int
		for i < len(s) && s[i] == y[i] {
			i++
		}
		if i == len(s) || s[i] < y[i] {
			return y
		}
	}
	return s
}

func insertNum(s string, num int) string {
	bs := []byte(s)
	buf := make([]byte, len(s)+1)

	var i int
	for i < len(bs) && int(bs[i]-'0') >= num {
		buf[i] = bs[i]
		i++
	}
	buf[i] = byte(num + '0')
	copy(buf[i+1:], bs[i:])
	return string(buf)
}
