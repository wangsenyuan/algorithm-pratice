package p756

func pyramidTransition(bottom string, allowed []string) bool {
	allow := make([][][]bool, 26)
	for x := 0; x < 26; x++ {
		allow[x] = make([][]bool, 26)
		for y := 0; y < 26; y++ {
			allow[x][y] = make([]bool, 26)
		}
	}

	for i := 0; i < len(allowed); i++ {
		s := allowed[i]
		a, b, c := int(s[0]-'A'), int(s[1]-'A'), int(s[2]-'A')
		allow[a][b][c] = true
	}

	var generate func(bottom string) []string

	generate = func(bottom string) []string {
		res := make([]string, 0, 10)

		var f func(i int, s string)

		f = func(i int, s string) {
			if i == len(bottom)-1 {
				res = append(res, s)
				return
			}
			a, b := int(bottom[i]-'A'), int(bottom[i+1]-'A')
			for x := 0; x < 26; x++ {
				if allow[a][b][x] {
					f(i+1, s+string(byte(x+'A')))
				}
			}
		}

		f(0, "")

		return res
	}

	// cache := make(map[string]int)

	var process func(bottom string) bool

	process = func(bottom string) bool {
		// fmt.Printf("[debug] process %s\n", bottom)
		if len(bottom) == 1 {
			return true
		}

		var ans bool
		strs := generate(bottom)
		if len(strs) > 0 {
			for _, str := range strs {
				if process(str) {
					ans = true
					break
				}
			}
		}
		return ans
	}

	return process(bottom)
}

func pyramidTransition1(bottom string, allowed []string) bool {
	allow := make([][][]bool, 26)
	for x := 0; x < 26; x++ {
		allow[x] = make([][]bool, 26)
		for y := 0; y < 26; y++ {
			allow[x][y] = make([]bool, 26)
		}
	}

	for i := 0; i < len(allowed); i++ {
		s := allowed[i]
		a, b, c := int(s[0]-'A'), int(s[1]-'A'), int(s[2]-'A')
		allow[a][b][c] = true
	}
	n := len(bottom)
	f := make([][]bool, n)
	g := make([][]bool, n)
	for i := 0; i < n; i++ {
		f[i] = make([]bool, 26)
		f[i][int(bottom[i]-'A')] = true
		g[i] = make([]bool, 26)
	}

	for n > 1 {
		for i := 0; i < n-1; i++ {
			a, b := f[i], f[i+1]
			for x := 0; x < 26; x++ {
				g[i][x] = false
				for j := 0; j < 26 && !g[i][x]; j++ {
					if a[j] == false {
						continue
					}
					for k := 0; k < 26 && !g[i][x]; k++ {
						if b[k] == false {
							continue
						}
						g[i][x] = allow[j][k][x]
					}
				}
			}
		}
		f, g = g, f
		n--
	}
	for i := 0; i < 26; i++ {
		if f[0][i] {
			return true
		}
	}
	return false
}
