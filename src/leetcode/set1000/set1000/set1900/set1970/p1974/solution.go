package p1974

func minTimeToType(word string) int {
	var ans int
	var cur byte = 'a'
	for i := 0; i < len(word); i++ {
		// move it
		a, b := 26, 26
		if word[i] <= cur {
			a = int(cur - word[i])
			b = int('z'-cur) + int(word[i]-'a') + 1
		} else {
			a = int(word[i] - cur)
			b = int('z'-word[i]) + int(cur-'a') + 1
		}
		ans += min(a, b)
		// type it
		ans++
		cur = word[i]
	}
	return ans
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
