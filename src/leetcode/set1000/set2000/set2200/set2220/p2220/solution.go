package p2220

func minBitFlips(start int, goal int) int {
	num := start ^ goal
	var cnt int
	for num > 0 {
		cnt++
		num -= num & -num
	}

	return cnt
}
