package p3007

func beautifulIndices(s string, a string, b string, k int) []int {
	la := kmp(a)
	lb := kmp(b)

	var arr []int

	for i, j := 0, 0; i < len(s); {
		if s[i] == b[j] {
			i++
			j++
		}
		if j == len(b) {
			arr = append(arr, i-j)
			j = lb[j-1]
		} else if i < len(s) && s[i] != b[j] {
			if j > 0 {
				j = lb[j-1]
			} else {
				i++
			}
		}
	}
	l, r := 0, 0
	check := func(i int) bool {
		for r < len(arr) && arr[r] < i {
			r++
		}
		for l < r && arr[l] < i && i-arr[l] > k {
			l++
		}
		if r < len(arr) && abs(i-arr[r]) <= k {
			return true
		}
		return l < r
	}

	var res []int
	for i, j := 0, 0; i < len(s); {
		if s[i] == a[j] {
			i++
			j++
		}
		if j == len(a) {
			// a good place
			if check(i - j) {
				res = append(res, i-j)
			}
			j = la[j-1]
		} else if i < len(s) && s[i] != a[j] {
			if j > 0 {
				j = la[j-1]
			} else {
				i++
			}
		}
	}

	return res
}

func kmp(pattern string) []int {
	n := len(pattern)

	lps := make([]int, n)

	for i := 1; i < n; i++ {
		j := lps[i-1]
		for j > 0 && pattern[i] != pattern[j] {
			j = lps[j-1]
		}
		lps[i] = j
		if pattern[i] == pattern[j] {
			lps[i]++
		}
	}
	return lps
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
