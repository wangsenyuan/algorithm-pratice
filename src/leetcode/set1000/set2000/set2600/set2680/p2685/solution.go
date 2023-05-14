package p2685

func countCompleteComponents(n int, edges [][]int) int {
	dsu := NewDSU(n)
	deg := make([]int, n)

	for _, e := range edges {
		u, v := e[0], e[1]
		deg[u]++
		deg[v]++
		dsu.Union(u, v)
	}

	comp := make([]int, n)

	for i := 0; i < n; i++ {
		comp[dsu.Find(i)]++
	}

	good := make([]bool, n)
	for i := 0; i < n; i++ {
		good[i] = true
	}

	for i := 0; i < n; i++ {
		r := dsu.Find(i)
		m := comp[r]
		if deg[i] != m-1 {
			good[r] = false
		}
	}
	var res int

	for i := 0; i < n; i++ {
		r := dsu.Find(i)
		if good[r] {
			res++
			good[r] = false
		}
	}

	return res
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
