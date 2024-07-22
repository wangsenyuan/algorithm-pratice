package p3228

func maxOperations(s string) int {
	var res int
	var cnt int
	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			cnt++
		} else if i > 0 && s[i-1] == '1' {
			res += cnt
		}
	}
	return res
}
