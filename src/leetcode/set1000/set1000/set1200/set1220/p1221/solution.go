package p5222

func balancedStringSplit(s string) int {
	var res int
	var cnt int
	for i := 0; i < len(s); i++ {
		if s[i] == 'L' {
			cnt++
		} else {
			cnt--
		}
		if cnt == 0 {
			res++
		}
	}
	return res
}
