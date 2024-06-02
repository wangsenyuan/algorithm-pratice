package p3171

func minimumChairs(s string) int {
	var res int
	var active int
	for i := 0; i < len(s); i++ {
		if s[i] == 'E' {
			active++
		} else {
			active--
		}
		res = max(res, active)
	}
	return res
}
