package consecstrs

func LongestConsec(strarr []string, k int) string {
	// your code
	n := len(strarr)

	if n == 0 || k > n || k <= 0 {
		return ""
	}

	var best string

	i := 0
	for i < k {
		best += strarr[i]
		i++
	}

	tmp := best
	for i < n {
		tmp = tmp[len(strarr[i-k]):] + strarr[i]
		if len(tmp) > len(best) {
			best = tmp
		}

		i++
	}

	return best
}
