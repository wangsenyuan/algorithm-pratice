package p3184

func countOfPeaks(nums []int, queries [][]int) []int {
	n := len(nums)

	tr := NewSegTree(n)

	val := make([]int, n)
	for i := 1; i+1 < n; i++ {
		if nums[i] > nums[i-1] && nums[i] > nums[i+1] {
			val[i] = 1
			tr.Set(i, 1)
		}
	}

	find := func(l int, r int) int {
		if r-l+1 <= 2 {
			return 0
		}
		return tr.Get(l+1, r)
	}

	checkAndUpdate := func(i int) {
		if i < 0 || i == n {
			return
		}
		if i > 0 && nums[i] > nums[i-1] && i+1 < n && nums[i] > nums[i+1] {
			val[i] = 1
			tr.Set(i, 1)
		}
	}

	update := func(i int, v int) {
		if i > 0 && val[i-1] == 1 {
			val[i-1] = 0
			tr.Set(i-1, 0)
		}
		if val[i] == 1 {
			val[i] = 0
			tr.Set(i, 0)
		}
		if i+1 < n && val[i+1] == 1 {
			val[i] = 0
			tr.Set(i+1, 0)
		}

		nums[i] = v
		checkAndUpdate(i - 1)
		checkAndUpdate(i)
		checkAndUpdate(i + 1)
	}

	var res []int

	for _, qr := range queries {
		if qr[0] == 1 {
			l, r := qr[1], qr[2]
			res = append(res, find(l, r))
		} else {
			i, v := qr[1], qr[2]
			update(i, v)
		}
	}

	return res
}

type SegTree struct {
	arr []int
	n   int
}

func NewSegTree(n int) *SegTree {
	arr := make([]int, 2*n)
	return &SegTree{arr, n}
}

func (tr *SegTree) Set(i int, v int) {
	i += tr.n
	tr.arr[i] = v
	for i > 1 {
		tr.arr[i>>1] = tr.arr[i] + tr.arr[i^1]
		i >>= 1
	}
}
func (tr *SegTree) Get(l int, r int) int {
	var res int
	l += tr.n
	r += tr.n
	for l < r {
		if l&1 == 1 {
			res += tr.arr[l]
			l++
		}
		if r&1 == 1 {
			r--
			res += tr.arr[r]
		}
		l >>= 1
		r >>= 1
	}
	return res
}
