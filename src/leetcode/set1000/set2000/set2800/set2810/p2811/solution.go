package p2811

var dd = []int{-1, 0, 1, 0, -1}

func maximumSafenessFactor(grid [][]int) int {
	n := len(grid)

	que := make([]int, n*n)
	var front, end int
	pushBack := func(i, j int) {
		que[end] = i*n + j
		end++
	}
	popFront := func() (int, int) {
		cur := que[front]
		front++
		i, j := cur/n, cur%n
		return i, j
	}
	dist := make([][]int, n)
	for i := 0; i < n; i++ {
		dist[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dist[i][j] = -1
			if grid[i][j] == 1 {
				dist[i][j] = 0
				pushBack(i, j)
			}
		}
	}

	for front < end {
		u, v := popFront()
		for i := 0; i < 4; i++ {
			x, y := u+dd[i], v+dd[i+1]
			if x >= 0 && x < n && y >= 0 && y < n && dist[x][y] < 0 {
				dist[x][y] = dist[u][v] + 1
				pushBack(x, y)
			}
		}
	}

	dsu := NewDSU(n * n)

	check := func(expect int) bool {
		if dist[0][0] < expect || dist[n-1][n-1] < expect {
			return false
		}
		dsu.Reset()
		// only when dist[i][j] >= expect, union it
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if dist[i][j] < expect {
					continue
				}
				for k := 0; k < 4; k++ {
					u, v := i+dd[k], j+dd[k+1]
					if u >= 0 && u < n && v >= 0 && v < n && dist[u][v] >= expect {
						dsu.Union(i*n+j, u*n+v)
					}
				}
			}
		}
		return dsu.Find(0) == dsu.Find(n*n-1)
	}

	l, r := 0, n*2

	for l < r {
		mid := (l + r) / 2
		if check(mid) {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return r - 1
}

type DSU struct {
	arr []int
	cnt []int
}

func NewDSU(n int) *DSU {
	arr := make([]int, n)
	cnt := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
		cnt[i] = 1
	}
	return &DSU{arr, cnt}
}

func (this *DSU) Find(x int) int {
	if this.arr[x] != x {
		this.arr[x] = this.Find(this.arr[x])
	}
	return this.arr[x]
}

func (this *DSU) Union(x int, y int) bool {
	px := this.Find(x)
	py := this.Find(y)

	if px == py {
		return false
	}
	if this.cnt[px] < this.cnt[py] {
		px, py = py, px
	}
	this.cnt[px] += this.cnt[py]
	this.arr[py] = px
	return true
}

func (this *DSU) Reset() {
	for i := 0; i < len(this.arr); i++ {
		this.arr[i] = i
		this.cnt[i] = 1
	}
}
