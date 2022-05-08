package p2267

func hasValidPath(grid [][]byte) bool {
	m := len(grid)
	n := len(grid[0])
	if grid[0][0] == ')' || grid[m-1][n-1] == '(' {
		return false
	}

	if (m+n-1)%2 == 1 {
		return false
	}

	H := (m + n + 1) / 2

	dp := make([][][]bool, m)
	for i := 0; i < m; i++ {
		dp[i] = make([][]bool, n)
		for j := 0; j < n; j++ {
			dp[i][j] = make([]bool, H)
		}
	}

	dp[0][0][1] = true

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			}
			for k := 0; k < H; k++ {
				if grid[i][j] == '(' {
					if k > 0 && i-1 >= 0 && dp[i-1][j][k-1] {
						dp[i][j][k] = true
					}
					if k > 0 && j-1 >= 0 && dp[i][j-1][k-1] {
						dp[i][j][k] = true
					}
				} else {
					if k+1 < H && i-1 >= 0 && dp[i-1][j][k+1] {
						dp[i][j][k] = true
					}
					if k+1 < H && j-1 >= 0 && dp[i][j-1][k+1] {
						dp[i][j][k] = true
					}
				}
			}
		}
	}

	return dp[m-1][n-1][0]
}

func hasValidPath2(grid [][]byte) bool {
	m := len(grid)
	n := len(grid[0])
	if grid[0][0] == ')' || grid[m-1][n-1] == '(' {
		return false
	}

	dp := make([][]TwoUints, m)
	fp := make([][]TwoUints, m)
	vis := make([][]bool, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]TwoUints, n)
		fp[i] = make([]TwoUints, n)
		vis[i] = make([]bool, n)
	}
	dp[0][0] = NewTwoUints(0, 2)
	fp[m-1][n-1] = NewTwoUints(0, 2)

	vis[0][0] = true

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i+j > (m+n)/2 {
				continue
			}
			if i+1 < m {
				if grid[i+1][j] == '(' {
					dp[i+1][j] = dp[i+1][j].Union(dp[i][j].LeftShift())
				} else {
					dp[i+1][j] = dp[i+1][j].Union(dp[i][j].RightShift())
				}
				vis[i+1][j] = true
			}
			if j+1 < n {
				if grid[i][j+1] == '(' {
					dp[i][j+1] = dp[i][j+1].Union(dp[i][j].LeftShift())
				} else {
					dp[i][j+1] = dp[i][j+1].Union(dp[i][j].RightShift())
				}
				vis[i][j+1] = true
			}
		}
	}

	check := func(a, b TwoUints) bool {
		c := a.Intersect(b)
		return c[0] > 0 || c[1] > 0
	}

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i+j < (m+n)/2 {
				continue
			}
			if i-1 >= 0 {
				if vis[i-1][j] {
					if check(fp[i][j], dp[i-1][j]) {
						return true
					}
				} else {
					if grid[i-1][j] == ')' {
						fp[i-1][j] = fp[i-1][j].Union(fp[i][j].LeftShift())
					} else {
						fp[i-1][j] = fp[i-1][j].Union(fp[i][j].RightShift())
					}
				}
			}
			if j-1 >= 0 {
				if vis[i][j-1] {
					if check(fp[i][j], dp[i][j-1]) {
						return true
					}
				} else {
					if grid[i][j-1] == ')' {
						fp[i][j-1] = fp[i][j-1].Union(fp[i][j].LeftShift())
					} else {
						fp[i][j-1] = fp[i][j-1].Union(fp[i][j].RightShift())
					}
				}
			}
		}
	}

	return false
}

type TwoUints [2]uint64

func NewTwoUints(a, b uint64) TwoUints {
	return TwoUints([2]uint64{a, b})
}

func (this TwoUints) LeftShift() TwoUints {
	a, b := this[0], this[1]
	a <<= 1
	a |= ((b >> 63) & 1)
	b <<= 1
	return NewTwoUints(a, b)
}

func (this TwoUints) RightShift() TwoUints {
	a, b := this[0], this[1]
	b >>= 1
	b &= ((1 << 63) - 1)
	b |= ((a & 1) << 63)
	a >>= 1
	a &= ((1 << 63) - 1)
	return NewTwoUints(a, b)
}

func (this TwoUints) Union(that TwoUints) TwoUints {
	a, b := this[0], this[1]
	a |= that[0]
	b |= that[1]
	return NewTwoUints(a, b)
}

func (this TwoUints) Intersect(that TwoUints) TwoUints {
	a, b := this[0], this[1]
	a &= that[0]
	b &= that[1]
	return NewTwoUints(a, b)
}

func hasValidPath1(grid [][]byte) bool {
	m := len(grid)
	n := len(grid[0])

	dp := make([][]map[int]bool, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]map[int]bool, n)
	}

	if grid[0][0] == ')' {
		return false
	}

	dp[0][0] = map[int]bool{1: true}

	increase := func(cur map[int]bool) map[int]bool {
		res := make(map[int]bool)
		for k := range cur {
			res[k+1] = true
		}
		return res
	}

	decrease := func(cur map[int]bool) map[int]bool {
		res := make(map[int]bool)
		for k := range cur {
			if k-1 >= 0 {
				res[k-1] = true
			}
		}
		return res
	}

	union := func(a, b map[int]bool) map[int]bool {
		if a == nil {
			a = make(map[int]bool)
		}
		for k := range b {
			a[k] = true
		}
		return a
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			cur := dp[i][j]
			if cur == nil || len(cur) == 0 {
				continue
			}
			if i+1 < m {
				// go down
				if grid[i+1][j] == '(' {
					// will increase
					dp[i+1][j] = union(dp[i+1][j], increase(dp[i][j]))
				} else {
					// need to decrease
					dp[i+1][j] = union(dp[i+1][j], decrease(dp[i][j]))
				}
			}
			if j+1 < n {
				// go right
				if grid[i][j+1] == '(' {
					dp[i][j+1] = union(dp[i][j+1], increase(dp[i][j]))
				} else {

					dp[i][j+1] = union(dp[i][j+1], decrease(dp[i][j]))
				}
			}
		}
	}
	if dp[m-1][n-1] == nil {
		return false
	}

	return dp[m-1][n-1][0]
}
