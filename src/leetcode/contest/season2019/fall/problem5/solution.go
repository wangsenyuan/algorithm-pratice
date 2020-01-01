package problem5

const MOD = 1000000007

func bonus1(n int, leadership [][]int, operations [][]int) []int {
	open := make([]int, n)
	close := make([]int, n)
	// tour := make([]int, n)

	outs := make([][]int, n)

	for i := 0; i < n; i++ {
		outs[i] = make([]int, 0, 2)
	}

	for _, op := range leadership {
		a, b := op[0]-1, op[1]-1
		outs[a] = append(outs[a], b)
	}

	var dfs func(u int, pos *int)

	dfs = func(u int, pos *int) {
		open[u] = *pos

		*pos++

		for _, v := range outs[u] {
			dfs(v, pos)
		}

		close[u] = *pos - 1
	}

	dfs(0, new(int))

	bit1 := CreateBitTree(n)
	bit2 := CreateBitTree(n)

	updateRange := func(v int, l int, r int) {
		bit1.update(l, v)
		bit1.update(r+1, -v)

		bit2.update(l, v*(l-1))
		bit2.update(r+1, -v*r)
	}

	sum := func(x int) int64 {
		a := bit1.getSum(x) * int64(x)
		a %= MOD
		b := bit2.getSum(x) % MOD
		a -= b
		if a < 0 {
			a += MOD
		}
		return a
	}

	rangeSum := func(l, r int) int {
		res := sum(r) - sum(l-1)
		if res < 0 {
			res += MOD
		}
		return int(res % MOD)
	}

	var res []int

	for _, op := range operations {
		if op[0] == 1 {
			updateRange(op[2], open[op[1]-1], open[op[1]-1])
		} else if op[0] == 2 {
			updateRange(op[2], open[op[1]-1], close[op[1]-1])
		} else {
			res = append(res, rangeSum(open[op[1]-1], close[op[1]-1]))
		}
	}

	return res
}

type BitTree struct {
	arr []int64
	n   int
}

func CreateBitTree(n int) BitTree {
	arr := make([]int64, n+1)
	return BitTree{arr, n}
}

func (bit *BitTree) update(p int, v int) {
	V := int64(v)
	p++
	n := bit.n
	for p <= n {
		bit.arr[p] += V
		p += p & -p
	}
}

func (bit BitTree) getSum(p int) int64 {
	var res int64
	p++
	for p > 0 {
		res += bit.arr[p]
		p -= p & -p
	}
	return res
}

func bonus(n int, leadership [][]int, operations [][]int) []int {
	outs := make([][]int, n)

	for i := 0; i < n; i++ {
		outs[i] = make([]int, 0, 3)
	}

	for _, leader := range leadership {
		a := leader[0] - 1
		b := leader[1] - 1
		outs[a] = append(outs[a], b)
	}

	open := make([]int, n)
	close := make([]int, n)

	var dfs func(u int, time *int)

	dfs = func(u int, time *int) {
		open[u] = *time
		*time++

		for _, v := range outs[u] {
			dfs(v, time)
		}
		close[u] = *time
	}

	dfs(0, new(int))

	tree := NewMergeTree(n)

	var res []int

	for _, op := range operations {
		if op[0] == 1 {
			i := op[1] - 1
			coin := op[2]
			j := open[i]
			tree.Update(j, j, int64(coin))
		} else if op[0] == 2 {
			i := op[1] - 1
			coin := op[2]
			j := open[i]
			k := close[i]
			tree.Update(j, k-1, int64(coin))
		} else {
			i := op[1] - 1
			j := open[i]
			k := close[i]
			res = append(res, int(tree.Get(j, k-1)))
		}
	}

	return res
}

func modAdd(a *int64, b int64) {
	*a += b
	if *a >= MOD {
		*a -= MOD
	}
}

type MergeTree struct {
	arr  []int64
	lazy []int64
	size int
}

func NewMergeTree(n int) MergeTree {
	arr := make([]int64, 4*n)
	lazy := make([]int64, 4*n)
	return MergeTree{arr, lazy, n}
}

func (tree *MergeTree) Update(left, right int, value int64) {
	arr := tree.arr
	lazy := tree.lazy
	size := tree.size
	var loop func(i int, start int, end int)

	loop = func(i int, start int, end int) {
		if start > right || end < left {
			return
		}
		if left <= start && end <= right {
			// total in the range
			modAdd(&arr[i], int64(end-start+1)*value%MOD)
			modAdd(&lazy[i], value)
			return
		}
		mid := (start + end) / 2
		loop(2*i+1, start, mid)
		loop(2*i+2, mid+1, end)
		arr[i] = 0
		modAdd(&arr[i], arr[2*i+1])
		modAdd(&arr[i], arr[2*i+2])
		modAdd(&arr[i], int64(end-start+1)*lazy[i]%MOD)
	}
	loop(0, 0, size-1)
}

func (tree *MergeTree) Get(left, right int) int64 {
	arr := tree.arr
	lazy := tree.lazy
	size := tree.size
	var loop func(i int, start int, end int, v int64) int64
	loop = func(i int, start int, end int, v int64) int64 {
		if start > right || end < left {
			return 0
		}
		var res int64
		if left <= start && end <= right {
			modAdd(&res, arr[i])
			modAdd(&res, int64(end-start+1)*v%MOD)
			return res
		}
		v += lazy[i]
		if v >= MOD {
			v -= MOD
		}
		mid := (start + end) / 2
		a := loop(2*i+1, start, mid, v)
		b := loop(2*i+2, mid+1, end, v)
		modAdd(&res, a)
		modAdd(&res, b)
		return res
	}

	return loop(0, 0, size-1, 0)
}
