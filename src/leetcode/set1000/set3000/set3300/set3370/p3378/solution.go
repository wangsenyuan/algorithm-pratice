package p3378

func countComponents(nums []int, threshold int) int {
	pos := make([]int, threshold+1)
	for i := 0; i <= threshold; i++ {
		pos[i] = -1
	}
	for i, num := range nums {
		if num <= threshold {
			pos[num] = i
		}
	}
	n := len(nums)

	find := func(g int) int {
		for x := g; x <= threshold; x += g {
			if pos[x] >= 0 {
				return x
			}
		}
		return -1
	}
	set := NewDSU(n)

	res := n
	for g := 1; g <= threshold; g++ {
		x := find(g)
		if x < 0 {
			continue
		}
		for y := x + g; y*x/g <= threshold; y += g {
			if pos[y] >= 0 {
				if set.Union(pos[x], pos[y]) {
					res--
				}
			}
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
