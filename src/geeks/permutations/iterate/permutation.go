package iterate

//Permutations is used to get all Permutations of given str
func Permutations(str string) []string {
	n := len(str)
	if n == 0 {
		return make([]string, 0, 0)
	}

	p := fact(len(str))

	ps := make([]string, 0, p)
	//ps = append(ps, str)
	m := 0
	for i := 0; i < p; {
		s := copyAsSlice(str)

		s[0], s[m] = s[m], s[0]

		k := 0
		for k != p / n {
			for j := 1; j < n - 1; j++ {
				s[j], s[j + 1] = s[j + 1], s[j]
				i++
				k++
				ps = append(ps, string(s))
			}
		}
		m++
		if m == n {
			break
		}
	}

	return ps
}

func fact(n int) int {
	if n <= 1 {
		return 1
	}
	return n * fact(n - 1)
}

func copyAsSlice(str string) []rune {
	rs := make([]rune, len(str), len(str))
	copy(rs, []rune(str))
	return rs
}
