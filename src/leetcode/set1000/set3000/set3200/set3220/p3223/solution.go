package p3223

func minimumLength(s string) int {
	freq := make([]int, 26)
	for i := 0; i < len(s); i++ {
		freq[int(s[i]-'a')]++
	}
	var res int
	for i := 0; i < 26; i++ {
		if freq[i] < 3 {
			res += freq[i]
		} else if freq[i]%2 == 1 {
			res++
		} else {
			res += 2
		}
	}

	return res
}
