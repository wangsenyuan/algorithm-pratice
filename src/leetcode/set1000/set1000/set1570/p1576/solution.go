package p1576

func modifyString(s string) string {
	n := len(s)
	buf := []byte(s)

	var i int
	for i < n && buf[i] == '?' {
		i++
	}
	if i > 0 {
		var x int
		if i < n && buf[i] == byte('a'+x) {
			x++
		}
		for j := i - 1; j >= 0; j-- {
			buf[j] = byte('a' + x)
			x = (x + 1) % 26
		}
		if i == n {
			return string(buf)
		}
	}
	for i < n {
		// buf[i] is always not ?
		if i+1 < n && buf[i+1] == '?' {
			j := i + 1
			for j < n && buf[j] == '?' {
				j++
			}
			var x int
			if buf[i] == byte(x+'a') {
				x++
			}
			if j == i+2 {
				// only one ?
				if i+2 < n && buf[i+2] == byte(x+'a') {
					x++
				}
				if buf[i] == byte(x+'a') {
					x++
				}
				buf[i+1] = byte('a' + x)
				i = j - 1
			} else {
				// multiple ?
				for k := i + 1; k < j-1; k++ {
					buf[k] = byte('a' + x)
					x = (x + 1) % 26
				}
				i = j - 3
			}
		}
		i++
	}

	return string(buf)
}
