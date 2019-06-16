package p777

func canTransform(start string, end string) bool {
	n := len(start)
	var a, b, c, d int
	var s, t string
	for i := 0; i < n; i++ {
		if start[i] == 'R' {
			s += start[i : i+1]
			a++
		}
		if start[i] == 'L' {
			s += start[i : i+1]
			b++
		}
		if end[i] == 'R' {
			t += end[i : i+1]
			c++
		}
		if end[i] == 'L' {
			t += end[i : i+1]
			d++
		}

		if c > a {
			return false
		}
		if b > d {
			return false
		}
	}

	return s == t
}
