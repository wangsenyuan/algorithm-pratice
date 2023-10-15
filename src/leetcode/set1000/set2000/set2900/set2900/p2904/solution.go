package p2904

func shortestBeautifulSubstring(s string, k int) string {
	n := len(s)

	pref := make([]int, n+1)

	ans := ""

	check := func(cur string) bool {
		if len(ans) == 0 {
			return true
		}
		return len(ans) > len(cur) || len(ans) == len(cur) && ans > cur
	}

	var l int
	for i := 0; i < n; i++ {
		pref[i+1] = pref[i]
		if s[i] == '1' {
			pref[i+1]++
		}

		for pref[i+1]-pref[l] >= k {
			if pref[i+1]-pref[l] == k && s[l] == '1' {
				if check(s[l : i+1]) {
					ans = s[l : i+1]
				}
			}
			l++
		}
	}

	return ans
}
