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
		V := readNNums(reader, n)
		Q := make([][]int, m)
		for i := 0; i < m; i++ {
			Q[i] = readNNums(reader, 4)
		}

		res := solve(V, Q)

		for _, ans := range res {
			buf.WriteString(fmt.Sprintf("%d\n", ans))
		}
	}

	fmt.Print(buf.String())
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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

	for len(bs) == 0 || bs[0] == '\n' || bs[0] == '\r' {
		bs, _ = reader.ReadBytes('\n')
	}

	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func solve(V []int, Q [][]int) []int64 {
	n := len(V)
	nodes := make([]Node, 4*n)

	la := make([]int, 4*n)
	ld := make([]int, 4*n)

	var build func(i int, l int, r int)
	build = func(i int, l int, r int) {
		la[i] = -1
		ld[i] = -1
		if l == r {
			nodes[i] = NewNode(V[l])
			return
		}
		mid := (l + r) / 2
		build(2*i, l, mid)
		build(2*i+1, mid+1, r)
		nodes[i] = Merge(nodes[2*i], nodes[2*i+1])
	}

	build(1, 0, n-1)

	push := func(i int, l int, r int) {
		if la[i] != -1 {
			nodes[i] = NewNode2(la[i], ld[i], r-l+1)
			if l < r {
				mid := (l + r) / 2
				la[2*i] = la[i]
				la[2*i+1] = la[i] + (mid-l+1)*ld[i]
				ld[2*i] = ld[i]
				ld[2*i+1] = ld[i]
			}
			la[i] = -1
			ld[i] = -1
		}
	}

	var update func(L int, R int, x int, y int, i int, l int, r int)

	update = func(L int, R int, x int, y int, i int, l int, r int) {
		push(i, l, r)
		if R < l || r < L {
			return
		}
		if L <= l && r <= R {
			nodes[i] = NewNode2(x+y*(l-L), y, r-l+1)
			if l < r {
				mid := (l + r) / 2
				push(2*i, l, mid)
				push(2*i+1, mid+1, r)
				la[2*i] = x + y*(l-L)
				la[2*i+1] = x + y*(mid+1-L)
				ld[2*i] = y
				ld[2*i+1] = y
			}
			return
		}
		mid := (l + r) / 2
		update(L, R, x, y, 2*i, l, mid)
		update(L, R, x, y, 2*i+1, mid+1, r)
		nodes[i] = Merge(nodes[2*i], nodes[2*i+1])
	}
	ans := make([]int64, 0, len(Q)+1)
	ans = append(ans, nodes[1].ans)

	for _, cur := range Q {
		L, R, x, y := cur[0], cur[1], cur[2], cur[3]
		update(L-1, R-1, x, y, 1, 0, n-1)
		ans = append(ans, nodes[1].ans)
	}

	return ans
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

func max(arr ...int64) int64 {
	res := arr[0]
	for i := 1; i < len(arr); i++ {
		if res < arr[i] {
			res = arr[i]
		}
	}
	return res
}

type Pair struct {
	first  int
	second int
}

func (this Pair) AddSecond(v int) Pair {
	return Pair{this.first, this.second + v}
}

// All the distinct prefix and suffix gcd's of the sub-array represented by the node in the segment tree
// stored as (gcd, size) pairs where size is the number of prefixes/suffixes with the same gcd
// The size of these is always $O(\log C)$
type Node struct {
	pref []Pair
	suf  []Pair
	ans  int64 // Answer for the sub-array represented by the node in the segment tree
}

func NewNode(x int) Node {
	pref := []Pair{{x, 1}}
	suf := []Pair{{x, 1}}
	ans := int64(x)
	return Node{pref, suf, ans}
}

func NewNode2(a int, d int, sz int) Node {
	if sz == 1 {
		return NewNode(a)
	}
	pref := []Pair{{a, 1}, {gcd(a, d), sz - 1}}
	suf := []Pair{{a + (sz-1)*d, 1}, {gcd(a, d), sz - 1}}
	ans := max(int64(a), int64(a)+int64(sz-1)*int64(d), int64(gcd(a, d))*int64(sz))

	return Node{pref, suf, ans}
}

func Merge(a, b Node) Node {
	if len(a.pref) == 0 {
		return b
	}
	if len(b.pref) == 0 {
		return a
	}
	pref := mergeGcd(a.pref, b.pref)
	suf := mergeGcd(b.suf, a.suf)
	ans := max(a.ans, b.ans)
	var ptr, l, r int
	for _, p := range b.pref {
		r += p.second
		for ptr < len(a.suf) && a.suf[ptr].first%p.first == 0 {
			l += a.suf[ptr].second
			ptr++
		}
		ans = max(ans, int64(l+r)*int64(p.first))
	}
	ptr, l, r = 0, 0, 0
	for _, p := range a.suf {
		l += p.second
		for ptr < len(b.pref) && b.pref[ptr].first%p.first == 0 {
			r += b.pref[ptr].second
			ptr++
		}
		ans = max(ans, int64(l+r)*int64(p.first))
	}
	return Node{pref, suf, ans}
}

func mergeGcd(c []Pair, b []Pair) []Pair {
	a := make([]Pair, len(c), len(c)+len(b))
	copy(a, c)
	for _, p := range b {
		x := a[len(a)-1]
		g := gcd(x.first, p.first)
		if g == x.first {
			a[len(a)-1] = x.AddSecond(p.second)
		} else {
			a = append(a, Pair{g, p.second})
		}
	}
	return a
}
