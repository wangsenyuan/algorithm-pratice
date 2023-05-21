package p2699

func minLength(s string) int {
	n := len(s)
	buf := make([]byte, n)
	var p int

	for i := 0; i < n; i++ {
		if s[i] == 'B' {
			if p > 0 && buf[p-1] == 'A' {
				p--
				continue
			}
		}
		if s[i] == 'D' {
			if p > 0 && buf[p-1] == 'C' {
				p--
				continue
			}
		}
		buf[p] = s[i]
		p++
	}
	return p
}
