package p3403

func answerString(word string, k int) string {
	if k == 1 {
		return word
	}
	// 除了最大的其他，都是1个长度的
	// 选中的长度应该是 n - (k - 1)
	n := len(word)
	// pref[i] = i前面的最大的字符，不包括i
	pref := make([]byte, n)
	pref[0] = 'a'
	for i := 1; i < n; i++ {
		pref[i] = max(word[i-1], pref[i-1])
	}
	suf := make([]byte, n)
	suf[n-1] = 'a'
	for i := n - 2; i >= 0; i-- {
		suf[i] = max(word[i+1], suf[i+1])
	}
	m := n - (k - 1)
	// 如果i是当前的选中的开始位置
	var ans string
	for i := 0; i < n; i++ {
		if pref[i] > word[i] || suf[min(i+m, n)-1] > word[i] {
			// 它不是最大的
			continue
		}
		if len(ans) == 0 || ans < word[i:min(i+m, n)] {
			ans = word[i:min(i+m, n)]
		}
	}

	return ans
}
