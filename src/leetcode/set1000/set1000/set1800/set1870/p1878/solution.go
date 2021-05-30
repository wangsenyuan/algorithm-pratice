package p1878

func getBiggestThree(grid [][]int) []int {
	m := len(grid)
	n := len(grid[0])

	s1 := make([][]int, m)
	s2 := make([][]int, m)

	for i := 0; i < m; i++ {
		s1[i] = make([]int, n)
		s2[i] = make([]int, n)
		for j := 0; j < n; j++ {
			s1[i][j] = grid[i][j]
			if i > 0 && j > 0 {
				s1[i][j] += s1[i-1][j-1]
			}
			s2[i][j] = grid[i][j]
			if i > 0 && j+1 < n {
				s2[i][j] += s2[i-1][j+1]
			}
		}
	}

	res := make([]int, 3)
	update := func(value int) {
		if value > res[0] {
			res[2] = res[1]
			res[1] = res[0]
			res[0] = value
		} else if value != res[0] && value > res[1] {
			res[2] = res[1]
			res[1] = value
		} else if value != res[0] && value != res[1] && value > res[2] {
			res[2] = value
		}
	}

	getSum2 := func(i, j, u, v int) int {
		ans := s2[i][j]
		if u > 0 && v+1 < n {
			ans -= s2[u-1][v+1]
		}
		return ans
	}

	getSum1 := func(i, j, u, v int) int {
		ans := s1[i][j]
		if u > 0 && v > 0 {
			ans -= s1[u-1][v-1]
		}
		return ans
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			update(grid[i][j])
			for k := i + 2; k < m; k += 2 {
				b := k
				x := (b + i) / 2
				l := j - (k-i)/2
				r := j + (k-i)/2
				if l < 0 || r >= n {
					break
				}
				tmp1 := getSum2(x, l, i, j)
				tmp2 := getSum1(x, r, i, j)
				tmp3 := getSum2(b, j, x, r)
				tmp4 := getSum1(b, j, x, l)
				v := tmp1 + tmp2 + tmp3 + tmp4 - grid[i][j] - grid[x][r] - grid[b][j] - grid[x][l]
				update(v)
			}
		}
	}
	for i := 0; i < 3; i++ {
		if res[i] == 0 {
			return res[:i]
		}
	}
	return res
}

func getBiggestThree1(grid [][]int) []int {
	m := len(grid)
	n := len(grid[0])

	lp := make([][]int, m+n-1)
	rp := make([][]int, m+n-1)

	for i := 0; i < m+n-1; i++ {
		lp[i] = make([]int, m+n-1)
		rp[i] = make([]int, m+n-1)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// for rp, k is i + j
			k := i + j
			rp[k][i] = grid[i][j]
			if i > 0 {
				rp[k][i] += rp[k][i-1]
			}
			// top-right one is 0
			// bottom-left one is m + n - 2
			k = n - 1 - j + i
			lp[k][i] = grid[i][j]
			if i > 0 {
				lp[k][i] += lp[k][i-1]
			}
		}
	}
	res := make([]int, 3)
	update := func(value int) {
		if value == res[0] || value == res[1] || value == res[2] {
			return
		}
		if value > res[0] {
			res[2] = res[1]
			res[1] = res[0]
			res[0] = value
		} else if value > res[1] {
			res[2] = res[1]
			res[1] = value
		} else if value > res[2] {
			res[2] = value
		}
	}

	get := func(arr []int, i, j int) int {
		ans := arr[i]
		if j > 0 {
			ans -= arr[j-1]
		}
		return ans
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			update(grid[i][j])
			var x = 1
			for {
				t := i - x
				r := j + x
				b := i + x
				l := j - x
				if t < 0 || r >= n || b >= m || l < 0 {
					break
				}
				tmp1 := get(rp[l+i], i, t)
				tmp2 := get(lp[n-1-r+i], i, t)
				tmp3 := get(rp[r+i], b, i)
				tmp4 := get(lp[n-1-l+i], b, i)
				update(tmp1 + tmp2 + tmp3 + tmp4 - grid[t][j] - grid[i][r] - grid[b][j] - grid[i][l])
				x++
			}
		}
	}
	for i := 0; i < 3; i++ {
		if res[i] == 0 {
			return res[:i]
		}
	}
	return res
}
