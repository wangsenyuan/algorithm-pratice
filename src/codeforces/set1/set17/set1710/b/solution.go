package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var buf bytes.Buffer

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n, m := readTwoNums(reader)
		X := make([]int, n)
		P := make([]int, n)
		for i := 0; i < n; i++ {
			X[i], P[i] = readTwoNums(reader)
		}
		res := solve(n, X, P, m)
		buf.WriteString(res)
		buf.WriteByte('\n')
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

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
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

func solve(n int, X []int, P []int, m int) string {

	diff := make([]Pair, 0, 3*n)

	for i := 0; i < n; i++ {
		diff = append(diff, Pair{2 * int64(X[i]-P[i]), 1})
		diff = append(diff, Pair{2 * int64(X[i]), -2})
		diff = append(diff, Pair{2 * int64(X[i]+P[i]), 1})
	}

	sort.Slice(diff, func(i, j int) bool {
		return diff[i].first < diff[j].first || diff[i].first == diff[j].first && diff[i].second < diff[j].second
	})

	var a, d int64
	var last int64

	M := 2 * int64(m)

	key := Pair{0, -0x3f3f3f3f3f3f3f3f}

	for _, p := range diff {
		if p.first != last {
			a = a + (p.first-last)*d
			last = p.first
			if a > M {
				key = intercection(key, Pair{p.first, a - M})
			}
		}
		d += p.second
	}

	res := make([]byte, n)
	for i := 0; i < n; i++ {
		cur := Pair{2 * int64(X[i]), 2 * int64(P[i])}
		tmp := intercection(key, cur)
		if tmp == cur {
			res[i] = '1'
		} else {
			res[i] = '0'
		}
	}
	return string(res)
}

func intercection(a, b Pair) Pair {
	x := max(a.first+a.second, b.first+b.second)
	y := max(a.second-a.first, b.second-b.first)
	return Pair{(x - y) / 2, (x + y) / 2}
}

type Pair struct {
	first  int64
	second int64
}

func solve1(n int, X []int, P []int, m int) string {
	rs := make([]RainDay, n)

	for i := 0; i < n; i++ {
		rs[i] = RainDay{X[i], P[i], i}
	}

	sort.Slice(rs, func(i, j int) bool {
		return rs[i].pos < rs[j].pos
	})

	V := NewSegTree(n)
	arr := make([]int64, n)

	cnt := make([]int, n+1)

	for i := n - 1; i >= 0; i-- {
		cnt[i] += cnt[i+1]
		cur := rs[i]
		j := search(n, func(j int) bool {
			return rs[j].pos >= cur.pos-cur.vol
		})
		V.Update(j, i+1, int64(cur.vol-cur.pos))
		cnt[i]++
		if j > 0 {
			cnt[j-1]--
		}
		arr[i] += int64(cnt[i])*int64(cur.pos) + V.Get(i)
	}

	V.Reset()
	for i := 0; i <= n; i++ {
		cnt[i] = 0
	}

	for i := 0; i < n; i++ {
		if i > 0 {
			cnt[i] += cnt[i-1]
		}
		cur := rs[i]
		j := search(n, func(j int) bool {
			return rs[j].pos > cur.pos+cur.vol
		})
		// (i+1, v-1), (i + 2, v - 2), (i + 3, v - 3)... (j, 0)
		// k - i + v
		// 是在(i + 1, j) 之间 使用 v- i
		V.Update(i, j, int64(cur.vol+cur.pos))
		cnt[i]++
		cnt[j]--

		arr[i] += int64(cnt[i])*int64(-cur.pos) + V.Get(i)
	}

	for i := 0; i < n; i++ {
		arr[i] -= int64(rs[i].vol)
	}

	arr2 := make([]int64, 2*n)
	copy(arr2[n:], arr)

	for i := n - 1; i > 0; i-- {
		arr2[i] = max(arr2[i*2], arr2[i*2+1])
	}

	getMax := func(l, r int) int64 {
		l += n
		r += n
		var res int64

		for l < r {
			if l&1 == 1 {
				res = max(res, arr2[l])
				l++
			}
			if r&1 == 1 {
				r--
				res = max(res, arr2[r])
			}
			l >>= 1
			r >>= 1
		}
		return res
	}

	for i := 0; i < n; i++ {
		arr[i] += int64(rs[i].pos)
	}
	// 取消某个i， 相当于从  arr[a...b] 中减去 v - i

	root := BuildTree(arr)

	M := int64(m)

	buf := make([]byte, n)

	for i := 0; i < n; i++ {
		cur := rs[i]

		root.Add(i, i, -2*int64(cur.pos))

		v := cur.vol
		a := search(n, func(j int) bool {
			return rs[j].pos >= cur.pos-v
		})

		b := search(n, func(j int) bool {
			return rs[j].pos > cur.pos+v
		})
		// -(v - (i - j)) = -v + i - j
		root.Add(a, i, -int64(v)+int64(cur.pos))

		if i+1 <= b-1 {
			// -(v - (j - i)) = -(v - j + i) = -v + j - i
			root.Add(i+1, b-1, -int64(v)-int64(cur.pos))
		}

		tmp := root.Get(a, b-1)

		if a > 0 {
			tmp = max(tmp, getMax(0, a))
		}
		if b < n {
			tmp = max(tmp, getMax(b, n))
		}

		if tmp > M {
			buf[cur.id] = '0'
		} else {
			buf[cur.id] = '1'
		}
		// add back
		root.Add(a, i, int64(v)-int64(cur.pos))
		if i+1 <= b-1 {
			root.Add(i+1, b-1, int64(v)+int64(cur.pos))
		}
	}
	return string(buf)
}

type Node struct {
	left  *Node
	right *Node
	begin int
	end   int
	val   int64
	lazy  int64
}

func NewNode(begin, end int) *Node {
	node := new(Node)
	node.begin = begin
	node.end = end
	return node
}

func BuildTree(arr []int64) *Node {
	var build func(l int, r int) *Node

	build = func(l int, r int) *Node {
		node := NewNode(l, r)
		if l == r {
			node.val = arr[l]
			return node
		}
		mid := (l + r) / 2
		node.left = build(l, mid)
		node.right = build(mid+1, r)
		node.val = max(node.left.val, node.right.val)
		return node
	}

	return build(0, len(arr)-1)
}

func (node *Node) push() {
	if node.lazy != 0 {
		if node.left != nil {
			node.left.lazy += node.lazy
			node.left.val += node.lazy
			node.right.lazy += node.lazy
			node.right.val += node.lazy
		}
		node.lazy = 0
	}
}

func (node *Node) Add(l int, r int, v int64) {
	if node.end < l || r < node.begin {
		return
	}
	node.push()
	if l <= node.begin && node.end <= r {
		node.val += v
		node.lazy += v
		return
	}
	node.left.Add(l, r, v)
	node.right.Add(l, r, v)

	node.val = max(node.left.val, node.right.val)
}

func (node *Node) Get(l int, r int) int64 {
	if node.end < l || r < node.begin {
		return 0
	}

	node.push()

	if l <= node.begin && node.end <= r {
		return node.val
	}

	a := node.left.Get(l, r)
	b := node.right.Get(l, r)

	node.val = max(node.left.val, node.right.val)

	return max(a, b)
}

func max(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}

type SegTree struct {
	arr []int64
	n   int
}

func NewSegTree(n int) *SegTree {
	arr := make([]int64, 2*n)
	return &SegTree{arr, n}
}

func (t *SegTree) Update(l, r int, v int64) {
	l += t.n
	r += t.n
	for l < r {
		if l&1 == 1 {
			t.arr[l] += v
			l++
		}
		if r&1 == 1 {
			r--
			t.arr[r] += v
		}
		l >>= 1
		r >>= 1
	}
}

func (t *SegTree) Get(p int) int64 {
	var res int64
	p += t.n

	for p > 0 {
		res += t.arr[p]
		p >>= 1
	}
	return res
}

func (t *SegTree) Reset() {
	for i := 0; i < len(t.arr); i++ {
		t.arr[i] = 0
	}
}

type RainDay struct {
	pos int
	vol int
	id  int
}

func search(n int, f func(int) bool) int {
	l, r := 0, n
	for l < r {
		mid := (l + r) / 2
		if f(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}
