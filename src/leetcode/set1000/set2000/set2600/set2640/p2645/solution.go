package p2645

func addMinimum(word string) int {
	n := len(word)
	var res int
	for i := 0; i < n; i++ {
		if word[i] == 'a' {
			if i+2 < n && word[i+1] == 'b' && word[i+2] == 'c' {
				i += 2
				continue
			}
			if i+1 == n || word[i+1] == 'a' {
				res += 2
				continue
			}
			// insert one ab or ac
			res++
			i++
			continue
		} else if word[i] == 'b' {
			// must insert a before
			res++
			if i+1 < n && word[i+1] == 'c' {
				i++
				continue
			}
			res++
		} else {
			res += 2
		}
	}

	return res
}
