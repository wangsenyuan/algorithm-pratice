package p2286

func rearrangeCharacters(s string, target string) int {
	cnt1 := countLetters(s)
	cnt2 := countLetters(target)
	res := len(s)
	for i := 0; i < 26; i++ {
		if cnt2[i] > 0 {
			res = min(res, cnt1[i]/cnt2[i])
		}
	}
	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func countLetters(s string) []int {
	cnt := make([]int, 26)
	for i := 0; i < len(s); i++ {
		cnt[int(s[i]-'a')]++
	}
	return cnt
}
