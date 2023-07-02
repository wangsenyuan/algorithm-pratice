package p2755

const inf = 1 << 30

func continuousSubarrays(nums []int) int64 {
	n := len(nums)

	// abs(nums[i] - nums[j]) <= 2

	var res int64

	// 假设以以当前的r为结束位置
	// 那么需要找到一个位置l
	// 要满足max(l...r) - min(l...r) > 2的
	// 如果对于r-1已经有一个l，对于r，这个l只会往后走
	a := NewSegTree(nums, max)
	b := NewSegTree(nums, min)

	for l, r := 0, 0; r < n; r++ {
		for l < r && a.Get(l, r+1, 0)-b.Get(l, r+1, inf) > 2 {
			l++
		}
		res += int64(r - l + 1)
	}

	return res
}

type SegTree struct {
	arr []int
	n   int
	f   func(int, int) int
}

func NewSegTree(arr []int, f func(int, int) int) *SegTree {
	n := len(arr)
	rr := make([]int, 2*n)
	copy(rr[n:], arr)
	for i := n - 1; i > 0; i-- {
		rr[i] = f(rr[i*2], rr[i*2+1])
	}
	return &SegTree{rr, n, f}
}

func (t *SegTree) Get(l int, r int, iv int) int {
	res := iv
	l += t.n
	r += t.n

	for l < r {
		if l&1 == 1 {
			res = t.f(res, t.arr[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = t.f(res, t.arr[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	return a + b - max(a, b)
}
