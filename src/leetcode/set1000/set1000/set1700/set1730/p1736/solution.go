package p1736

func maximumTime(time string) string {

	buf := []byte(time)

	hour := time[:2]

	if hour == "??" {
		buf[0] = '2'
		buf[1] = '3'
	} else if hour[0] == '?' {
		if time[1] > '3' {
			buf[0] = '1'
		} else {
			buf[0] = '2'
		}
	} else if hour[1] == '?' {
		if time[0] == '0' || time[0] == '1' {
			buf[1] = '9'
		} else {
			buf[1] = '3'
		}
	}

	min := time[3:]

	if min[0] == '?' {
		buf[3] = '5'
	}
	if min[1] == '?' {
		buf[4] = '9'
	}

	return string(buf)
}
