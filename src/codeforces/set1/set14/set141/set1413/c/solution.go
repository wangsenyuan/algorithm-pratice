package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	A := readNNums(reader, 6)
	n := readNum(reader)
	B := readNNums(reader, n)
	res := solve(A, B)
	fmt.Println(res)
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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

const inf = 1 << 30

func solve(A []int, B []int) int {
	sort.Ints(A)
	m := len(A)
	n := len(B)

	vals := make([]Pair, m*n)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			vals[i*n+j] = Pair{B[j] - A[i], j}
		}
	}

	sort.Slice(vals, func(i, j int) bool {
		return vals[i].first < vals[j].first
	})

	cnt := make(map[int]int)

	ans := inf

	for l, r := 0, 0; r < len(vals); {
		for r < len(vals) && len(cnt) < n {
			cnt[vals[r].second]++
			r++
		}
		if len(cnt) < n {
			break
		}
		for len(cnt) == n {
			ans = min(ans, vals[r-1].first-vals[l].first)
			if cnt[vals[l].second] == 1 {
				delete(cnt, vals[l].second)
			} else {
				cnt[vals[l].second]--
			}
			l++
		}
	}
	return ans
}

func solve1(A []int, B []int) int {
	// len(A) = 6
	// len(B) = n < 1e5
	//sort.Ints(B)
	// 当B[0]在某个上面被弹时
	// 当B[i]在某个线上做为最大值时，
	// 其他的
	//m := len(A)
	sort.Ints(A)
	n := len(B)

	pos := make([]int, n)
	vals := make([]Pair, n)
	for i := 0; i < n; i++ {
		vals[i] = Pair{B[i] - A[0], i}
	}

	mx := NewSegTree(vals, maxPair)
	mn := NewSegTree(vals, minPair)

	ans := inf
	for {
		a := mx.Get(0, n, Pair{0, -1})
		b := mn.Get(0, n, Pair{inf, -1})
		ans = min(ans, a.first-b.first)
		id := a.second
		pos[id]++
		if pos[id] == 6 {
			break
		}

		mx.Update(a.second, Pair{B[id] - A[pos[id]], a.second})
		mn.Update(a.second, Pair{B[id] - A[pos[id]], a.second})
	}

	return ans
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
