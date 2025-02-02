package p3440

func findValidPair(s string) string {
	cnt := make([]int, 10)
	for _, x := range []byte(s) {
		cnt[int(x-'0')]++
	}

	for i := 0; i+1 < len(s); i++ {
		if s[i] != s[i+1] {
			x := int(s[i] - '0')
			y := int(s[i+1] - '0')
			if cnt[x] == x && cnt[y] == y {
				return s[i : i+2]
			}
		}
	}
	return ""
}
