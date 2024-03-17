package p3085

func countSubstrings(s string, c byte) int64 {
	var cnt int64
	for i := 0; i < len(s); i++ {
		if s[i] == c {
			cnt++
		}
	}
	return cnt * (cnt + 1) / 2
}
