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

	n := readNum(reader)
	a := readNNums(reader, n)
	res := solve(a)

	fmt.Println(res)
}

const inf = 1 << 60

func solve(a []int) int {
	n := len(a)

	arr := make([]Pair, n)
	for i := 0; i < n; i++ {
		arr[i] = Pair{a[i], i}
	}

	tr := NewSegTree(arr, minPair)

	var dfs func(l int, r int, x int) int

	dfs = func(l int, r int, x int) int {
		if l > r {
			return 0
		}
		if l == r {
			return 1
		}
		ln := r - l + 1
		p := tr.Get(l, r+1, Pair{inf, -1})

		res := p.first - x

		for i := l; i <= r; {
			if a[i] == p.first {
				i++
				continue
			}
			j := i
			for i <= r && a[i] > p.first {
				i++
			}
			res += dfs(j, i-1, p.first)
		}

		return min(res, ln)
	}

	return dfs(0, n-1, 0)
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

type Pair struct {
	first  int
	second int
}

func maxPair(a, b Pair) Pair {
	if a.first >= b.first {
		return a
	}
	return b
}

func minPair(a, b Pair) Pair {
	if a.first <= b.first {
		return a
	}
	return b
}

type SegTree struct {
	arr []Pair
	f   func(Pair, Pair) Pair
	n   int
}

func NewSegTree(arr []Pair, f func(Pair, Pair) Pair) *SegTree {
	n := len(arr)
	set := make([]Pair, 2*n)
	copy(set[n:], arr)
	for i := n - 1; i > 0; i-- {
		set[i] = f(set[i*2], set[i*2+1])
	}
	return &SegTree{set, f, n}
}

func (t *SegTree) Update(p int, v Pair) {
	p += t.n
	t.arr[p] = v
	for p > 1 {
		t.arr[p>>1] = t.f(t.arr[p], t.arr[p^1])
		p >>= 1
	}
}

func (t *SegTree) Get(l int, r int, iv Pair) Pair {
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
