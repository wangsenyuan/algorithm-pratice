package p076

func minWindow(s string, t string) string {
	cnt := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		cnt[t[i]]++
	}

	nu := len(t)
	var ans string
	for i, j := 0, 0; i < len(s); i++ {
		if cnt[s[i]] > 0 {
			nu--
		}
		cnt[s[i]]--
		for nu == 0 {
			if len(ans) == 0 || len(ans) > i-j+1 {
				ans = s[j : i+1]
			}
			if cnt[s[j]] == 0 {
				nu++
			}
			cnt[s[j]]++
			j++
		}
	}
	return ans
}

func minWindow1(s string, t string) string {
	cnt := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		cnt[t[i]]++
	}
	var ans string
	tmp1 := copyMap(cnt)
	tmp2 := make(map[byte]int)
	for i, j := 0, 0; i < len(s); i++ {
		for j <= i && cnt[s[j]] == 0 {
			j++
		}
		tmp2[s[i]]++
		if tmp1[s[i]] > 0 {
			if tmp1[s[i]] == 1 {
				delete(tmp1, s[i])
			} else {
				tmp1[s[i]]--
			}
			if len(tmp1) == 0 {
				//a candidate
				for j <= i && tmp2[s[j]] >= cnt[s[j]] && len(tmp1) == 0 {
					if tmp2[s[j]] == cnt[s[j]] && cnt[s[j]] > 0 {
						tmp1[s[j]]++
						if len(ans) == 0 || len(ans) > i-j+1 {
							ans = s[j : i+1]
						}
					}
					tmp2[s[j]]--
					j++
				}
			}
		}
	}
	return ans
}

func copyMap(src map[byte]int) map[byte]int {
	dst := make(map[byte]int)
	for k, v := range src {
		dst[k] = v
	}
	return dst
}
