package p1307

func isSolvable(words []string, result string) bool {
	n := len(result)
	// n <= 7 => result => 10000000

	for i := 0; i < len(words); i++ {
		if len(words[i]) > n {
			return false
		}
		words[i] = reverse(words[i])
	}

	result = reverse(result)

	assign := make([]int, 26)

	for i := 0; i < 26; i++ {
		assign[i] = -1
	}

	canNotBeZeros := make(map[int]bool)

	for i := 0; i < len(words); i++ {
		if len(words[i]) == 1 {
			continue
		}
		x := int(words[i][len(words[i])-1] - 'A')
		canNotBeZeros[x] = true
	}

	if n > 1 {
		canNotBeZeros[int(result[n-1]-'A')] = true
	}

	var solve func(i, j int, carry int, used int) bool

	solve = func(i, j int, carry int, used int) bool {
		if i == n {
			return carry == 0
		}

		if j == len(words) {
			sum := carry

			for k := 0; k < len(words); k++ {
				if i >= len(words[k]) {
					continue
				}

				x := int(words[k][i] - 'A')

				sum += assign[x]
			}

			r := sum % 10

			z := int(result[i] - 'A')

			if assign[z] >= 0 {
				return r == assign[z] && solve(i+1, 0, sum/10, used)
			}

			if used&(1<<uint(r)) != 0 {
				return false
			}
			assign[z] = r
			res := solve(i+1, 0, sum/10, used|(1<<uint(r)))
			assign[z] = -1
			return res
		}

		if i >= len(words[j]) {
			return solve(i, j+1, carry, used)
		}

		x := int(words[j][i] - 'A')
		if assign[x] >= 0 {
			return solve(i, j+1, carry, used)
		}
		// try all possibilities
		var num int
		if canNotBeZeros[x] {
			num++
		}
		for num < 10 {
			if used&(1<<uint(num)) == 0 {
				assign[x] = num
				if solve(i, j+1, carry, used|(1<<uint(num))) {
					return true
				}
				assign[x] = -1
			}

			num++
		}
		return false
	}

	return solve(0, 0, 0, 0)
}

func reverse(word string) string {
	bs := []byte(word)

	for i, j := 0, len(word)-1; i < j; i, j = i+1, j-1 {
		bs[i], bs[j] = bs[j], bs[i]
	}

	return string(bs)
}
