package p2405

func partitionString(s string) int {
	cnt := make(map[byte]int)

	res := 1

	for i := 0; i < len(s); i++ {
		if cnt[s[i]] > 0 {
			res++
			cnt = make(map[byte]int)
		}
		cnt[s[i]]++
	}

	return res
}
