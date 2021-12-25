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
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(scanner *bufio.Scanner) (a int) {
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &a)
	return
}

func readTwoNums(scanner *bufio.Scanner) (a int, b int) {
	res := readNNums(scanner, 2)
	a, b = res[0], res[1]
	return
}

func readNNums(scanner *bufio.Scanner, n int) []int {
	res := make([]int, n)
	x := 0
	scanner.Scan()
	for i := 0; i < n; i++ {
		for x < len(scanner.Bytes()) && scanner.Bytes()[x] == ' ' {
			x++
		}
		x = readInt(scanner.Bytes(), x, &res[i])
	}
	return res
}

func fillNNums(scanner *bufio.Scanner, n int, res []int) {
	x := 0
	scanner.Scan()
	for i := 0; i < n; i++ {
		for x < len(scanner.Bytes()) && scanner.Bytes()[x] == ' ' {
			x++
		}
		x = readInt(scanner.Bytes(), x, &res[i])
	}
}

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		line := readNNums(scanner, 5)
		N, Q, K, L, R := line[0], line[1], line[2], line[3], line[4]
		C := make([]int, N)
		P := make([]int, N)
		for i := 0; i < N; i++ {
			C[i], P[i] = readTwoNums(scanner)
		}
		res := solve(N, Q, K, L, R, C, P)

		for i := 0; i < N; i++ {
			if i < N-1 {
				fmt.Printf("%d ", res[i])
			} else {
				fmt.Printf("%d\n", res[i])
			}
		}
	}
}

const INF = 1 << 30

func solve(N, Q, K, L, R int, C []int, P []int) []int {
	tree := CreateSegTree(1 + Q)

	for i := 0; i < N; i++ {
		c, p := C[i], P[i]
		st, en := 0, -1

		if K < L {
			if c <= L {
				// no contribution
			} else if c <= R+1 {
				st = 1
				en = c - L
			} else {
				st = c - R
				en = c - L
			}
		} else if K > R {
			if c >= R {

			} else if c >= L-1 {
				st = 1
				en = min(Q, R-c)
			} else {
				st = L - c
				en = min(Q, R-c)
			}
		} else {
			if c < L-1 {
				st = L - c
			} else if c > R+1 {
				st = c - R
			} else {
				st = 1
			}
			en = Q
		}

		st = max(1, st)
		en = min(Q, en)
		if st <= en {
			tree.Update(st, en, p)
		}
	}
	ans := make([]int, Q)

	for i := 1; i <= Q; i++ {
		r := tree.GetMin(i)
		if r == INF {
			ans[i-1] = -1
		} else {
			ans[i-1] = r
		}
	}

	return ans
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

type SegTree struct {
	sum  []int
	lazy []int
	n    int
}

// const INF = 1 << 30

func CreateSegTree(n int) SegTree {
	m := 1
	for m <= n {
		m <<= 1
	}
	lazy := make([]int, m<<1)
	sum := make([]int, m)

	for i := 0; i < m; i++ {
		sum[i] = INF
	}

	for i := 0; i < len(lazy); i++ {
		lazy[i] = INF
	}

	return SegTree{sum, lazy, m}
}

func (tree *SegTree) push(i int) {
	n := tree.n
	if i >= n {
		tree.sum[i-n] = min(tree.sum[i-n], tree.lazy[i])
	} else {
		tree.lazy[i<<1] = min(tree.lazy[i<<1], tree.lazy[i])
		tree.lazy[i<<1|1] = min(tree.lazy[i<<1|1], tree.lazy[i])
	}
	tree.lazy[i] = INF
}

func (tree *SegTree) Update(l, r int, x int) {

	var loop func(l, r, ll, rr int, i int)

	loop = func(l, r, ll, rr int, i int) {
		tree.push(i)
		if l == ll && r == rr {
			tree.lazy[i] = x
			return
		}
		mid := (ll + rr) / 2

		if r <= mid {
			loop(l, r, ll, mid, i<<1)
		} else if l > mid {
			loop(l, r, mid+1, rr, (i<<1)|1)
		} else {
			loop(l, mid, ll, mid, i<<1)
			loop(mid+1, r, mid+1, rr, (i<<1)|1)
		}
	}

	loop(l, r, 0, tree.n-1, 1)
}

func (tree *SegTree) GetMin(pos int) int {
	var loop func(ll int, rr int, i int) int

	loop = func(ll int, rr int, i int) int {
		tree.push(i)
		if i >= tree.n {
			return tree.sum[i-tree.n]
		}
		mid := (ll + rr) / 2
		if pos <= mid {
			return loop(ll, mid, i<<1)
		}
		return loop(mid+1, rr, (i<<1)|1)
	}

	return loop(0, tree.n-1, 1)
}
