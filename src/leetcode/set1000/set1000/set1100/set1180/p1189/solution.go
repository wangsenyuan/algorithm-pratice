package p1189

func maxNumberOfBalloons(text string) int {
	// balloon
	cnt := make([]int, 26)
	for i := 0; i < len(text); i++ {
		cnt[text[i]-'a']++
	}

	// a b l o n
	need := cnt[0]
	need = min(need, cnt[1])
	need = min(need, cnt['l'-'a']/2)
	need = min(need, cnt['o'-'a']/2)
	need = min(need, cnt['n'-'a'])
	return need
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
