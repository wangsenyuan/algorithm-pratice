package p028

func strStr(haystack string, needle string) int {
	if len(haystack) == 0 {
		if len(needle) == 0 {
			return 0
		}
		return -1
	}
	if len(needle) == 0 {
		return 0
	}
	return searchByKMP(haystack, needle)
}

func searchByKMP(str, pat string) int {
	m := len(pat)
	lps := make([]int, m)
	for l, i := 0, 1; i < m; {
		if pat[l] == pat[i] {
			l++
			lps[i] = l
			i++
		} else {
			if l > 0 {
				l = lps[l-1]
			} else {
				lps[i] = 0
				i++
			}
		}
	}

	n := len(str)
	for i, j := 0, 0; i < n; {
		if str[i] == pat[j] {
			i++
			j++
		}
		if j == m {
			return i - m
		} else if i < n && str[i] != pat[j] {
			if j > 0 {
				j = lps[j-1]
			} else {
				i++
			}
		}
	}
	return -1
}

func searchByZ(str, pat string) int {
	concat := pat + "$" + str
	n := len(concat)
	z := make([]int, n)
	var l, r int

	for i := 1; i < n; i++ {
		if i > r {
			l = i
			r = i
			for r < n && concat[r-l] == concat[r] {
				r++
			}
			z[i] = r - l
			r--
		} else {
			k := i - l
			if z[k] < r-i+1 {
				z[i] = z[k]
			} else {
				l = i
				for r < n && concat[r-l] == concat[r] {
					r++
				}
				z[i] = r - l
				r--
			}
		}
	}
	for i := 0; i < n; i++ {
		if z[i] == len(pat) {
			return i - len(pat) - 1
		}
	}
	return -1
}
