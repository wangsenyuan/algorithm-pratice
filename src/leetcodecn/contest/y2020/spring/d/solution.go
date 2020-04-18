package d

const INF = 1 << 22

func minJump(jump []int) int {
	n := len(jump)

	arr := make([]int, 2*n+3)

	for i := 0; i < len(arr); i++ {
		arr[i] = INF
	}

	update := func(pos int, v int) {
		pos += n

		arr[pos] = v
		for pos > 1 {
			arr[pos>>1] = min(arr[pos>>1], arr[pos])
			pos >>= 1
		}
	}

	get := func(left, right int) int {
		res := INF

		left += n
		right += n

		for left < right {
			if left&1 == 1 {
				res = min(res, arr[left])
				left++
			}

			if right&1 == 1 {
				right--
				res = min(res, arr[right])
			}
			left >>= 1
			right >>= 1
		}

		return res
	}

	dp := make([]int, n)

	for i := 0; i < n; i++ {
		dp[i] = INF
	}
	dp[0] = 0

	res := INF

	for i := 0; i < n; i++ {
		x := jump[i]
		y := get(i+1, n)

		if i+x < n {
			dp[i+x] = min(dp[i+x], dp[i]+1)
			if y+2 < dp[i+x] {
				dp[i+x] = y + 2
			}
			update(i+x, dp[i+x])
		} else {
			res = min(res, dp[i]+1)
			res = min(res, y+2)
		}
	}

	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
