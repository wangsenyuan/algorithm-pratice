package p3365

func isPossibleToRearrange(s string, t string, k int) bool {
	n := len(s)
	if n%k != 0 {
		return false
	}
	m := n / k
	freq := make(map[string]int)

	for i := 0; i < n; i += m {
		freq[s[i:i+m]]++
		freq[t[i:i+m]]--
	}

	for _, v := range freq {
		if v != 0 {
			return false
		}
	}
	return true
}
