package p1883

func maxValue(n string, x int) string {
	res := make([]byte, len(n)+1)
	var i int
	if n[0] == '-' {
		i = 1
		for i < len(n) && int(n[i]-'0') <= x {
			i++
		}
	} else {
		for i < len(n) && int(n[i]-'0') >= x {
			i++
		}
	}
	copy(res, n[:i])
	res[i] = byte(x + '0')

	if i < len(n) {
		copy(res[i+1:], n[i:])
	}

	return string(res)
}
