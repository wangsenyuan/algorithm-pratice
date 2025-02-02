package p3436

func maxDifference(s string) int {
	n := len(s)
	freq := make([]int, 26)
	res := -n
	for i := 0; i < n; i++ {
		x := int(s[i] - 'a')
		freq[x]++
	}

	for x := 0; x < 26; x++ {
		if freq[x] > 0 && freq[x]&1 == 1 {
			for y := 0; y < 26; y++ {
				if freq[y] > 0 && freq[y]&1 == 0 {
					res = max(res, freq[x]-freq[y])
				}
			}
		}

	}
	return res
}
