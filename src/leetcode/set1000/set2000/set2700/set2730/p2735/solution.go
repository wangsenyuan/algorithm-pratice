package p2735

func smallestString(s string) string {
	n := len(s)
	buf := []byte(s)

	var i int
	for i < n && buf[i] == 'a' {
		i++
	}
	if i == n {
		buf[n-1] = 'z'
		return string(buf)
	}
	// 开始的位置已经定下了
	for i < n {
		x := buf[i]
		if x == 'a' {
			break
		}
		buf[i] = byte(x - 1)
		i++
	}
	return string(buf)
}
