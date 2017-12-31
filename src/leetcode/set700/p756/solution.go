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
