package p3138

func minAnagramLength(s string) int {
	cnt := make([]int, 26)
	n := len(s)
	for i := 0; i < n; i++ {
		cnt[s[i]-'a']++
	}
	// 假设有l个个这样的字符串，
	// 那么有cnt[i] % l = 0
	var g int
	for i := 0; i < 26; i++ {
		if cnt[i] > 0 {
			g = gcd(g, cnt[i])
		}
	}
	g = gcd(g, n)

	return n / g
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
