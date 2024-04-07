package p3106

func getSmallestString(s string, k int) string {
	buf := []byte(s)
	n := len(s)
	var cur int
	for i := 0; i < n; i++ {
		for j := 0; j < 26; j++ {
			x := byte('a' + j)
			if cur+dist(x, s[i]) <= k {
				buf[i] = x
				cur += dist(x, s[i])
				break
			}
		}
	}

	return string(buf)
}

func dist(a, b byte) int {
	if a > b {
		a, b = b, a
	}
	return min(int(b-a), int(a-'a')-int(b-'a')+26)
}
