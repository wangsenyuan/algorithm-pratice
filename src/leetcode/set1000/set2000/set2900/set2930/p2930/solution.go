package p2930

func nC2(n int) int64 {
	if n < 0 {
		return 0
	}
	return int64(n) * int64(n-1) / 2
}

func distributeCandies(n int, limit int) int64 {
	return nC2(n+2) - 3*nC2(n-(limit+1)+2) + 3*nC2(n-2*(limit+1)+2) - nC2(n-3*(limit+1)+2)
}
