package p3186

func countCompleteDayPairs(hours []int) int64 {
	freq := make([]int, 24)

	var res int
	for _, h := range hours {
		h %= 24
		if h == 0 {
			res += freq[0]
		} else {
			res += freq[24-h]
		}
		freq[h]++
	}
	return int64(res)
}
