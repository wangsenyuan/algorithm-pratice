package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	A := readNNums(reader, n)
	B := readNNums(reader, n)
	fmt.Println(solve(A, B))
}

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')

	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(reader *bufio.Reader) (a int) {
	bs, _ := reader.ReadBytes('\n')
	readInt(bs, 0, &a)
	return
}

func readTwoNums(reader *bufio.Reader) (a int, b int) {
	res := readNNums(reader, 2)
	a, b = res[0], res[1]
	return
}

func readThreeNums(reader *bufio.Reader) (a int, b int, c int) {
	res := readNNums(reader, 3)
	a, b, c = res[0], res[1], res[2]
	return
}

func readNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, n)
	x := 0
	bs, _ := reader.ReadBytes('\n')
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

const inf = 1 << 29

func solve(A []int, B []int) int {
	// i < j < k
	// A[i] < A[j] < A[k]
	// min (B[i] + B[j] + B[k])
	n := len(A)
	//L[i] 表示 min(B[j], where j < i && A[j] < A[i])
	type Pair struct {
		first  int
		second int
	}
	nums := make([]Pair, n)

	for i := 0; i < n; i++ {
		nums[i] = Pair{A[i], i}
	}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i].first < nums[j].first
	})

	L := make([]int, n)

	t := NewSegTree(n)

	for i := 0; i < n; {
		j := i
		for i < n && nums[i].first == nums[j].first {
			L[nums[i].second] = t.Get(0, nums[i].second)
			i++
		}

		for j < i {
			t.Set(nums[j].second, B[nums[j].second])
			j++
		}
	}

	R := make([]int, n)

	t.Reset()

	for i := n - 1; i >= 0; {
		j := i
		for i >= 0 && nums[i].first == nums[j].first {
			R[nums[i].second] = t.Get(nums[i].second+1, n)
			i--
		}
		for j > i {
			t.Set(nums[j].second, B[nums[j].second])
			j--
		}
	}

	best := inf
	for i := 0; i < n; i++ {
		best = min(best, L[i]+B[i]+R[i])
	}
	if best >= inf {
		return -1
	}
	return best
}

type SegTree struct {
	arr []int
	n   int
}

func NewSegTree(n int) *SegTree {
	arr := make([]int, 2*n)
	for i := 0; i < len(arr); i++ {
		arr[i] = inf
	}
	return &SegTree{arr, n}
}

func (t *SegTree) Set(p int, v int) {
	p += t.n
	t.arr[p] = min(t.arr[p], v)

	for p > 1 {
		t.arr[p>>1] = min(t.arr[p], t.arr[p^1])
		p >>= 1
	}
}

func (t *SegTree) Get(l int, r int) int {
	res := inf
	l += t.n
	r += t.n

	for l < r {
		if l&1 == 1 {
			res = min(res, t.arr[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = min(res, t.arr[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}

func (t *SegTree) Reset() {
	for i := 0; i < len(t.arr); i++ {
		t.arr[i] = inf
	}
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
