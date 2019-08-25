package p1003

func isValid(S string) bool {
	if len(S) == 0 {
		return false
	}

	n := len(S)

	//cnt[a] >= cnt[b] >= cnt[c]
	cnt := make([]int, 3)
	for i := 0; i < n; i++ {
		cnt[int(S[i]-'a')]++
		if cnt[0] < cnt[1] {
			return false
		}
		if cnt[0] < cnt[2] {
			return false
		}
		if cnt[1] < cnt[2] {
			return false
		}
	}
	return cnt[0] == cnt[1] && cnt[1] == cnt[2]
}
