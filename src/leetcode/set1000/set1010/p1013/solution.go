package p1013

func numPairsDivisibleBy60(time []int) int {
	cnt := make([]int, 60)
	for i := 0; i < len(time); i++ {
		cnt[time[i]%60]++
	}

	res := cnt[30] * (cnt[30] - 1) / 2
	res += cnt[0] * (cnt[0] - 1) / 2

	for i := 1; i < 60; i++ {
		j := 60 - i
		if i < j {
			res += cnt[i] * cnt[j]
		}
	}

	return res
}
