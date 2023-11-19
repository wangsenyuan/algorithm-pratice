package p2939

func minimumSteps(s string) int64 {
	var res int
	var cnt int

	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			res += cnt
			continue
		}
		cnt++
	}
	return int64(res)
}
