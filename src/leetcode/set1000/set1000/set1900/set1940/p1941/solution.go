package p1941

func areOccurrencesEqual(s string) bool {
	n := len(s)
	cnt := make([]int, 26)
	var x int
	for i := 0; i < n; i++ {
		cnt[int(s[i]-'a')]++
		if cnt[int(s[i]-'a')] == 1 {
			x++
		}
	}

	if n%x != 0 {
		return false
	}

	n /= x

	for i := 0; i < 26; i++ {
		if cnt[i] != 0 && cnt[i] != n {
			return false
		}
	}
	return true
}
