package problem5

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
			tree.Update(j, j, coin)
		} else if op[0] == 2 {
			i := op[1] - 1
			coin := op[2]
			j := open[i]
			k := close[i]
			tree.Update(j, k-1, coin)
		} else {
			i := op[1] - 1
			j := open[i]
			k := close[i]
			res = append(res, tree.Get(j, k-1))
		}
	}

	return res
}

type MergeTree struct {
	arr  []int
	lazy []int
	size int
}

func NewMergeTree(n int) MergeTree {
	arr := make([]int, 4*n)
	lazy := make([]int, 4*n)
	return MergeTree{arr, lazy, n}
}

func (tree *MergeTree) Update(left, right int, value int) {
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
			lazy[i] += value
			arr[i] += lazy[i] * (end - start + 1)
			if 2*i+1 < len(lazy) {
				lazy[2*i+1] += value
			}
			if 2*i+2 < len(lazy) {
				lazy[2*i+2] += value
			}
			lazy[i] = 0
			return
		}
		mid := (start + end) / 2
		loop(2*i+1, start, mid)
		loop(2*i+2, mid+1, end)
		arr[i] = arr[2*i+1] + arr[2*i+2]
	}
	loop(0, 0, size-1)
}

func (tree *MergeTree) Get(left, right int) int {
	arr := tree.arr
	lazy := tree.lazy
	size := tree.size
	var loop func(i int, start int, end int) int
	loop = func(i int, start int, end int) int {
		if start > right || end < left {
			return 0
		}
		if left <= start && end <= right {
			if 2*i+1 < len(lazy) {
				lazy[2*i+1] += lazy[i]
			}
			if 2*i+2 < len(lazy) {
				lazy[2*i+2] += lazy[i]
			}
			lazy[i] = 0
			return arr[i]
		}
		mid := (start + end) / 2
		return loop(2*i+1, start, mid) + loop(2*i+2, mid+1, end)
	}

	return loop(0, 0, size-1)
}
