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
	return search(haystack, needle)
}

func search(str, pat string) int {
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
