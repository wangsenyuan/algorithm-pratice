package p2062

func countVowelSubstrings(word string) int {
	n := len(word)
	L := make([]int, n)

	for i := 0; i < n; i++ {
		j := index(word[i])
		if j >= 0 {
			// a vowel
			if i > 0 && L[i-1] > 0 {
				L[i] = L[i-1] + 1
			} else {
				L[i] = 1
			}
		}
	}
	cnt := make(map[int]int)
	var res int
	var l int
	for r := 0; r < n; r++ {
		if L[r] == 0 {
			cnt = make(map[int]int)
			l = r + 1
			continue
		}
		// a vowel
		cnt[index(word[r])]++
		for len(cnt) == 5 {
			j := index(word[l])
			if cnt[j] == 1 {
				res += L[l]
				break
			}
			cnt[j]--
			l++
		}
	}

	return res
}

func index(x byte) int {
	if x == 'a' {
		return 0
	}
	if x == 'o' {
		return 1
	}
	if x == 'e' {
		return 2
	}
	if x == 'i' {
		return 3
	}
	if x == 'u' {
		return 4
	}
	return -1
}
