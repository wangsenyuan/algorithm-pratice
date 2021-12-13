package p2053

func kthDistinct(arr []string, k int) string {
	cnt := make(map[string]int)

	for _, str := range arr {
		cnt[str]++
	}

	for _, str := range arr {
		if cnt[str] == 1 {
			if k == 1 {
				return str
			}
			k--
		}
	}

	return ""
}
