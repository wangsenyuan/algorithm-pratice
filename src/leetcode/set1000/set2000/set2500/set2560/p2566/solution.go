package p2566

func handleQuery(nums1 []int, nums2 []int, queries [][]int) []int64 {
	n := len(nums1)

	var res []int64

	t := NewSegTree(nums1)

	var sum int64

	for _, num := range nums2 {
		sum += int64(num)
	}

	for _, cur := range queries {
		if cur[0] == 1 {
			// flip
			l, r := cur[1], cur[2]
			t.Flip(1, 0, n-1, l, r)
		} else if cur[0] == 2 {
			// add
			sum += int64(t.arr[1]) * int64(cur[1])
		} else {
			res = append(res, sum)
		}
	}

	return res
}

type SegTree struct {
	arr  []int
	lazy []int
}

func NewSegTree(num []int) *SegTree {
	t := new(SegTree)
	n := len(num)
	t.arr = make([]int, n*4)
	t.lazy = make([]int, n*4)

	var loop func(i int, l int, r int)
	loop = func(i int, l int, r int) {
		if l == r {
			t.arr[i] = num[l]
			return
		}
		mid := (l + r) / 2
		loop(2*i, l, mid)
		loop(2*i+1, mid+1, r)
		t.arr[i] = t.arr[2*i] + t.arr[2*i+1]
	}

	loop(1, 0, n-1)

	return t
}

func (t *SegTree) push(i int, l, r int) {
	if t.lazy[i] == 1 && l < r {
		mid := (l + r) / 2
		t.Flip(2*i, l, mid, l, mid)
		t.Flip(2*i+1, mid+1, r, mid+1, r)
		t.lazy[i] ^= 1
	}
}

func (t *SegTree) Flip(i int, l int, r int, L int, R int) {
	if r < L || R < l {
		return
	}
	if L <= l && r <= R {
		t.arr[i] = r - l + 1 - t.arr[i]
		t.lazy[i] ^= 1
		return
	}
	t.push(i, l, r)
	mid := (l + r) / 2
	t.Flip(2*i, l, mid, L, R)
	t.Flip(2*i+1, mid+1, r, L, R)
	t.arr[i] = t.arr[2*i] + t.arr[2*i+1]
}
