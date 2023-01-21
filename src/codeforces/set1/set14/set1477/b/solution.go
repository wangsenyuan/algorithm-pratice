package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n, m := readTwoNums(reader)
		s, _ := reader.ReadString('\n')
		f, _ := reader.ReadString('\n')
		L, R := make([]int, m), make([]int, m)
		for i := 0; i < m; i++ {
			L[i], R[i] = readTwoNums(reader)
		}
		res := solve(n, s, f, L, R)
		if res {
			buf.WriteString("YES\n")
		} else {
			buf.WriteString("NO\n")
		}
	}
	fmt.Print(buf.String())
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

func solve(n int, s, f string, L, R []int) bool {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = int(f[i] - '0')
	}
	tree := NewSegTree(n, arr)

	for i := len(L) - 1; i >= 0; i-- {
		l, r := L[i], R[i]
		cnt := tree.Query(l-1, r-1)
		ll := r - l + 1
		if cnt*2 == ll {
			// half 0/1
			return false
		}
		if cnt*2 < ll {
			// before should be all zeros in range [l, r], some changed to 1
			tree.Update(l-1, r-1, 0)
		} else {
			tree.Update(l-1, r-1, 1)
		}
	}
	for i := 0; i < n; i++ {
		x := tree.Query(i, i)
		if x != int(s[i]-'0') {
			return false
		}
	}
	return true
}

type SegTree struct {
	lazy []int
	sum  []int
	n    int
}

func NewSegTree(n int, arr []int) *SegTree {
	m := 1
	for m < n {
		m <<= 1
	}
	lazy := make([]int, m<<1)
	sum := make([]int, m<<1)

	var build func(i int, l, r int)
	build = func(i int, l, r int) {
		if l == r {
			sum[i] = arr[l]
			return
		}
		mid := (l + r) >> 1
		build(i<<1, l, mid)
		build((i<<1)|1, mid+1, r)
		sum[i] = sum[i<<1] + sum[(i<<1)|1]
	}

	build(1, 0, n-1)

	return &SegTree{lazy, sum, n}
}

func (tree *SegTree) push(i int, l, r int) {
	if tree.lazy[i] != 0 {
		if tree.lazy[i] == 1 {
			// all change to 1
			tree.sum[i] = r - l + 1
		} else {
			// all change to 0
			tree.sum[i] = 0
		}
		if l < r {
			tree.lazy[i<<1] = tree.lazy[i]
			tree.lazy[(i<<1)|1] = tree.lazy[i]
		}
		tree.lazy[i] = 0
	}

}

// set all values between l, r to v
func (tree *SegTree) Update(l, r int, v int) {
	var loop func(i int, ll, rr int)
	loop = func(i int, ll, rr int) {
		tree.push(i, ll, rr)
		if r < ll || rr < l {
			return
		}
		if l <= ll && rr <= r {
			if v == 1 {
				tree.sum[i] = rr - ll + 1
				if ll < rr {
					tree.lazy[i<<1] = 1
					tree.lazy[(i<<1)|1] = 1
				}
			} else {
				tree.sum[i] = 0
				if ll < rr {
					tree.lazy[i<<1] = -1
					tree.lazy[(i<<1)|1] = -1
				}
			}
			return
		}
		mid := (ll + rr) >> 1
		loop(i<<1, ll, mid)
		loop((i<<1)|1, mid+1, rr)
		tree.sum[i] = tree.sum[i<<1] + tree.sum[(i<<1)|1]
	}

	loop(1, 0, tree.n-1)
}

func (tree *SegTree) Query(l, r int) int {
	var loop func(i int, ll, rr int) int

	loop = func(i int, ll, rr int) int {
		tree.push(i, ll, rr)

		if rr < l || r < ll {
			return 0
		}
		if l <= ll && rr <= r {
			return tree.sum[i]
		}
		mid := (ll + rr) >> 1
		a := loop(i<<1, ll, mid)
		b := loop((i<<1)|1, mid+1, rr)
		return a + b
	}
	return loop(1, 0, tree.n-1)
}
