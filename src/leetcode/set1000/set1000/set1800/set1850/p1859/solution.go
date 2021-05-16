package p1859

const MOD = 1000000007

func sumOfFlooredPairs(nums []int) int {
	var upper int
	for _, num := range nums {
		if num > upper {
			upper = num
		}
	}
	cnt := make([]int, upper+1)
	for _, num := range nums {
		cnt[num]++
	}

	pre := make([]int64, upper+1)

	for i := 1; i <= upper; i++ {
		pre[i] = pre[i-1] + int64(cnt[i])
	}

	var ans int64

	for y := 1; y <= upper; y++ {
		if cnt[y] > 0 {
			for d := 1; d <= upper/y; d++ {
				tmp := int64(cnt[y]) * int64(d) * (pre[min((d+1)*y-1, upper)] - pre[d*y-1])
				ans += tmp
				ans %= MOD
			}
		}
	}
	return int(ans)
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
