package p2156

func subStrHash(s string, power int, modulo int, k int, hashValue int) string {
	s = reverse(s)
	hv := int64(hashValue)
	var pk int64 = 1
	var cur int64
	P := int64(power)
	M := int64(modulo)
	for i := 0; i < k; i++ {
		cur = (cur*P%M + int64(s[i]-'a'+1)) % M
		if i > 0 {
			pk = pk * P % M
		}
	}

	var ans int

	for i := k; i < len(s); i++ {
		cur += M - (int64(s[i-k]-'a'+1)*pk)%M
		cur = (cur*P%M + int64(s[i]-'a'+1)) % M
		cur %= M
		if cur == hv {
			ans = i - k + 1
		}
	}
	x := s[ans : ans+k]
	return reverse(x)
}

func reverse(s string) string {
	buf := []byte(s)
	for i, j := 0, len(buf)-1; i < j; i, j = i+1, j-1 {
		buf[i], buf[j] = buf[j], buf[i]
	}
	return string(buf)
}
