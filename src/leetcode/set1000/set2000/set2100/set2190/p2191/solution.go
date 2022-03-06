package p2191

func minMovesToMakePalindrome(s string) int {
	// n <= 2000
	n := len(s)
	cnt := make([]int, 26)
	for i := 0; i < n; i++ {
		cnt[int(s[i]-'a')]++
	}
	used := make([]bool, n)

	check := func(x int, l, r *int) int {
		// find x between l & r
		var a, b int
		for i := *l; ; i++ {
			if !used[i] {
				a++
				if int(s[i]-'a') == x {
					a--
					break
				}
			}
		}
		for i := *r; ; i-- {
			if !used[i] {
				b++
				if int(s[i]-'a') == x {
					b--
					break
				}
			}
		}
		return a + b
	}

	mark := func(x int, l, r *int) {
		for i := *l; ; i++ {
			if !used[i] && int(s[i]-'a') == x {
				used[i] = true
				break
			}
		}
		for i := *r; ; i-- {
			if !used[i] && int(s[i]-'a') == x {
				used[i] = true
				break
			}
		}
		for *l < *r && used[*l] {
			*l++
		}
		for *r > *l && used[*r] {
			*r--
		}
	}

	var res int
	left, right := 0, n-1
	for left < right {
		tmp := n
		it := -1
		for k := 0; k < 26; k++ {
			if cnt[k] >= 2 {
				x := check(k, &left, &right)
				if x < tmp {
					tmp = x
					it = k
				}
			}
		}
		cnt[it] -= 2
		mark(it, &left, &right)
		res += tmp
	}

	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
