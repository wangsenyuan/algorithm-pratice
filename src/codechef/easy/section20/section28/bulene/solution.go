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
		S := readNNums(reader, n)
		B := make([][]int, m)
		P := readNNums(reader, m)
		C := readNNums(reader, m)
		for i := 0; i < m; i++ {
			B[i] = make([]int, 2)
			B[i][0] = P[i]
			B[i][1] = C[i]
		}
		res := solve(S, B)
		for i := 0; i < n; i++ {
			buf.WriteString(fmt.Sprintf("%d ", res[i]))
		}
		buf.WriteByte('\n')
	}
	fmt.Print(buf.String())
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			return s[:i]
		}
	}
	return s
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

const INF = 1 << 30

func solve(S []int, B [][]int) []int {
	n := len(S)
	m := len(B)

	fen := NewBIT(m)

	for i := 0; i < m; i++ {
		fen.Add(i, B[i][0])
	}

	cap := make([]int, m)

	for i := 0; i < m; i++ {
		cap[i] = B[i][1]
	}

	tree := NewSegTree(cap)

	totPow := fen.Get(m - 1)

	ans := make([]int, n)

	for i := 0; i < n; i++ {
		ans[i] = S[i]

		var r int
		if totPow <= int64(S[i]) {
			r = m - 1
			ans[i] -= int(totPow)
		} else {
			r = fen.Search(int64(S[i]))
			ans[i] = 0
		}

		tree.Update2(0, r)

		z := tree.FindMin(0, r)

		for z.first == 0 {
			totPow -= int64(B[z.second][0])
			fen.Add(z.second, -B[z.second][0])
			tree.Update1(z.second)
			z = tree.FindMin(0, r)
		}
	}

	return ans
}

type BIT []int64

func NewBIT(n int) BIT {
	arr := make([]int64, n+1)
	return BIT(arr)
}

func (this BIT) Add(p int, v int) {
	for p < len(this) {
		this[p] += int64(v)
		p |= (p + 1)
	}
}

func (this BIT) Get(p int) int64 {
	var res int64
	for p >= 0 {
		res += this[p]
		p &= (p + 1)
		p--
	}
	return res
}

func (this BIT) Search(sum int64) int {
	var ret = -1
	n := len(this)
	for i := 20; i >= 0; i-- {
		if ret+(1<<uint(i)) >= n {
			continue
		}
		if this[ret+(1<<uint(i))] >= sum {
			continue
		}
		ret += 1 << uint(i)
		sum -= this[ret]
	}
	return ret + 1
}

type Pair struct {
	first  int
	second int
}

type SegTree struct {
	v    []int
	lazy []int
	id   []int
	sz   int
}

func NewSegTree(arr []int) *SegTree {
	n := len(arr)

	v := make([]int, 4*n)
	for i := 0; i < len(v); i++ {
		v[i] = INF
	}
	lazy := make([]int, 4*n)

	id := make([]int, 4*n)

	var build func(i int, l int, r int)
	build = func(i int, l int, r int) {
		if l == r {
			v[i] = arr[l]
			id[i] = l
			return
		}
		m := (l + r) / 2
		build(2*i+1, l, m)
		build(2*i+2, m+1, r)

		v[i] = min(v[2*i+1], v[2*i+2])
		if v[i] == v[2*i+1] {
			id[i] = id[2*i+1]
		} else {
			id[i] = id[2*i+2]
		}
	}
	build(0, 0, n-1)
	return &SegTree{v, lazy, id, n}
}

func (this *SegTree) relax(i int) {
	if this.lazy[i] != 0 {
		if 2*i+1 < len(this.v) {
			this.v[2*i+1] += this.lazy[i]
			this.lazy[2*i+1] += this.lazy[i]
		}
		if 2*i+2 < len(this.v) {
			this.v[2*i+2] += this.lazy[i]
			this.lazy[2*i+2] += this.lazy[i]
		}
	}

	this.lazy[i] = 0
}

func (this *SegTree) Update1(pos int) {

	v := this.v
	id := this.id

	var loop func(i int, l int, r int)

	loop = func(i int, l int, r int) {
		this.relax(i)
		if l == r {
			v[i] = INF
			return
		}
		m := (l + r) / 2
		if pos <= m {
			loop(2*i+1, l, m)
		} else {
			loop(2*i+2, m+1, r)
		}
		v[i] = min(v[2*i+1], v[2*i+2])
		if v[i] == v[2*i+1] {
			id[i] = id[2*i+1]
		} else {
			id[i] = id[2*i+2]
		}
	}

	loop(0, 0, this.sz-1)
}

func (this *SegTree) Update2(L int, R int) {

	v := this.v
	id := this.id
	lazy := this.lazy

	var loop func(i int, l int, r int)

	loop = func(i int, l int, r int) {
		if L > r || R < l {
			return
		}

		if L <= l && r <= R {
			v[i]--
			lazy[i]--
			return
		}

		this.relax(i)

		m := (l + r) / 2
		loop(2*i+1, l, m)
		loop(2*i+2, m+1, r)
		v[i] = min(v[2*i+1], v[2*i+2])
		if v[i] == v[2*i+1] {
			id[i] = id[2*i+1]
		} else {
			id[i] = id[2*i+2]
		}
	}

	loop(0, 0, this.sz-1)
}

func (this *SegTree) FindMin(L, R int) Pair {

	var loop func(i int, l int, r int) Pair
	loop = func(i int, l int, r int) Pair {
		if l > R || r < L {
			return Pair{INF, -1}
		}
		if L <= l && r <= R {
			return Pair{this.v[i], this.id[i]}
		}
		this.relax(i)
		m := (l + r) / 2
		a := loop(2*i+1, l, m)
		b := loop(2*i+2, m+1, r)
		if a.first <= b.first {
			return a
		}
		return b
	}

	return loop(0, 0, this.sz-1)
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
