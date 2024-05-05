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
	best := n

	for i := 1; i <= g/i; i++ {
		if g%i == 0 {
			if n%i == 0 {
				best = min(best, n/i)
			}
			j := g / i
			if n%j == 0 {
				best = min(best, n/j)
			}
		}
	}

	return best
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
