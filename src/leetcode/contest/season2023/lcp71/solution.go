package lcp71

func reservoir(shape []string) int {
	n := len(shape)
	m := len(shape[0])
	up := make([][]int, n+1)
	dw := make([][]int, n+1)
	lf := make([][]int, n+1)
	rg := make([][]int, n+1)
	var id int
	for i := 0; i <= n; i++ {
		up[i] = make([]int, m+2)
		dw[i] = make([]int, m+2)
		lf[i] = make([]int, m+2)
		rg[i] = make([]int, m+2)
		if i < n {
			for j := 1; j <= m; j++ {
				id++
				up[i][j] = id
				id++
				dw[i][j] = id
				id++
				lf[i][j] = id
				id++
				rg[i][j] = id
			}
		}
	}
	set := NewUFSet(id + 1)

	ok := make([]bool, id+1)

	for i := n - 1; i >= 0; i-- {
		for j := 0; j <= m; j++ {
			set.Union(rg[i][j], lf[i][j+1])
		}
		for j := 1; j <= m; j++ {
			set.Union(dw[i][j], up[i+1][j])
			if shape[i][j-1] == '.' {
				set.Union(lf[i][j], rg[i][j])
				set.Union(up[i][j], dw[i][j])
				set.Union(lf[i][j], up[i][j])
			} else if shape[i][j-1] == 'l' {
				set.Union(up[i][j], rg[i][j])
				set.Union(lf[i][j], dw[i][j])
			} else {
				set.Union(up[i][j], lf[i][j])
				set.Union(rg[i][j], dw[i][j])
			}
		}

		for j := 1; j <= m; j++ {
			ok[lf[i][j]] = set.Find(lf[i][j]) != set.Find(0)
			ok[rg[i][j]] = set.Find(rg[i][j]) != set.Find(0)
			ok[up[i][j]] = set.Find(up[i][j]) != set.Find(0)
			ok[dw[i][j]] = set.Find(dw[i][j]) != set.Find(0)
		}
	}

	for j := 1; j <= m; j++ {
		set.Union(up[0][j], 0)
	}

	var ans int

	for i, b := range ok {
		if b && set.Find(i) == set.Find(0) {
			ans++
		}
	}
	return ans / 2
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

type UF struct {
	set  []int
	size []int
}

func NewUFSet(n int) *UF {
	set := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		set[i] = i
		size[i] = 1
	}
	return &UF{set, size}
}

func (this *UF) Find(x int) int {
	set := this.set
	var loop func(x int) int
	loop = func(x int) int {
		if x != set[x] {
			set[x] = loop(set[x])
		}
		return set[x]
	}
	return loop(x)
}

func (this *UF) Union(a, b int) {
	pa := this.Find(a)
	pb := this.Find(b)
	if pa == pb {
		return
	}
	size := this.size
	if size[pa] >= size[pb] {
		this.set[pb] = pa
		size[pa] += size[pb]
	} else {
		this.set[pa] = pb
		size[pb] += size[pa]
	}
}

func (this UF) Size(x int) int {
	return this.size[x]
}
