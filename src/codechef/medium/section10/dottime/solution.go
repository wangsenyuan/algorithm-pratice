package main

import (
	"bufio"
	"fmt"
	"os"
)

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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	for tc > 0 {
		tc--
		line := readNNums(reader, 3)
		n, m, q := line[0], line[1], line[2]
		A := readNNums(reader, n)
		ops := make([][]int, q)
		for i := 0; i < q; i++ {
			ops[i] = readNNums(reader, 2)
		}
		res := solve(n, m, q, A, ops)

		for i := 0; i < q; i++ {
			fmt.Println(res[i])
		}
	}
}

const MOD = 998244353

func solve(N, M, Q int, A []int, ops [][]int) []int {
	// when M = 2
	//G(1) = a1 * (a1 + a2 + a3)
	// G(2) = a2 * (a1 + a2 + a3)
	// G(3) = a3 * (a1 + a2 + a3)
	// P(2) = a2 * (a2 + a3 + a4)
	// P(3) = a3 * (a2 + a3 + a4)
	// P(4) = a4 * (a2 + a3 + a4)
	// S = (a1 + a2 + a3) ^^ 2 + (a2 + a3 + a4) ^^ 2
	var sum int64
	for i := 0; i < N-M; i++ {
		sum += int64(A[i])
	}

	window := make([]int64, M)

	for i := 0; i < M; i++ {
		sum = add(sum, int64(A[i+N-M]))
		window[i] = sum
		sum = add(sum, int64(MOD-A[i]))
	}

	tree := NewSegTree(window)

	res := make([]int, Q)

	for i := 0; i < Q; i++ {
		p, v := ops[i][0], ops[i][1]
		p--
		delta := v - A[p]
		A[p] += delta
		if A[p] >= MOD {
			A[p] -= MOD
		}
		if A[p] < 0 {
			A[p] += MOD
		}

		left := max(0, p-(N-M))
		right := min(p, M-1)

		tree.Update(left, right, delta)

		res[i] = int(tree.Result())
	}

	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

type SegTree struct {
	h           int
	t, t2, lazy []int64
}

func NewSegTree(arr []int64) SegTree {
	var h int = 1
	for h < len(arr) {
		h <<= 1
	}
	t := make([]int64, h<<1)
	t2 := make([]int64, h<<1)
	lazy := make([]int64, h<<1)

	for i := 0; i < len(arr); i++ {
		t[i+h] = int64(arr[i])
		t2[i+h] = mul(int64(arr[i]), int64(arr[i]))
	}

	for i := h - 1; i > 0; i-- {
		t[i] = add(t[i<<1], t[i<<1|1])
		t2[i] = add(t2[i<<1], t2[i<<1|1])
	}

	return SegTree{h, t, t2, lazy}
}

func (tree *SegTree) push(i int, ll, rr int) {
	if tree.lazy[i] != 0 {
		sum := tree.t[i]
		sum2 := tree.t2[i]
		delta := tree.lazy[i]
		sz := rr - ll + 1
		tree.t[i] = add(sum, mul(delta, int64(sz)))
		tree.t2[i] = add(sum2, add(mul(int64(sz), mul(delta, delta)), mul(2, mul(delta, int64(sum)))))

		if i < tree.h {
			tree.lazy[i<<1] = add(tree.lazy[i<<1], tree.lazy[i])
			tree.lazy[i<<1|1] = add(tree.lazy[i<<1|1], tree.lazy[i])
		}
		tree.lazy[i] = 0
	}
}

func (tree *SegTree) Update(l int, r int, x int) {
	X := int64(x)

	var loop func(i int, l int, r int, ll int, rr int)

	loop = func(i int, l int, r int, ll int, rr int) {
		if l == ll && r == rr {
			tree.lazy[i] = add(tree.lazy[i], X)
			tree.push(i, ll, rr)
			return
		}
		tree.push(i, ll, rr)

		mid := (ll + rr) / 2

		if r <= mid {
			loop(l, r, i<<1, ll, mid)
			tree.push(i<<1|1, mid+1, rr)
		} else if l > mid {
			tree.push(i<<1, ll, mid)
			loop(l, r, i<<1|1, mid+1, rr)
		} else {
			loop(i<<1, l, mid, ll, mid)
			loop(i<<1|1, mid+1, r, mid+1, rr)
		}

		tree.t[i] = add(tree.t[i<<1], tree.t[i<<1|1])
		tree.t2[i] = add(tree.t2[i<<1], tree.t2[i<<1|1])
	}

	loop(1, l, r, 0, tree.h-1)
}

func (tree SegTree) Result() int64 {
	return tree.t2[1]
}

func mul(a, b int64) int64 {
	res := int64(a) * int64(b)
	res %= MOD
	return res
}

func add(a, b int64) int64 {
	a += b
	if a >= MOD {
		a -= MOD
	}
	if a < 0 {
		a += MOD
	}
	return a
}
