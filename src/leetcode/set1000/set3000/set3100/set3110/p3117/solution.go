package p3117

func findLatestTime(s string) string {
	buf := []byte(s)
	if buf[0] == '?' {
		if buf[1] == '?' || buf[1] < '2' {
			buf[0] = '1'
		} else {
			buf[0] = '0'
		}
	}
	if buf[1] == '?' {
		if buf[0] == '1' {
			buf[1] = '1'
		} else {
			buf[1] = '9'
		}
	}
	if buf[3] == '?' {
		buf[3] = '5'
	}
	if buf[4] == '?' {
		buf[4] = '9'
	}

	return string(buf)
}
