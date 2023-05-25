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

	tc := readNum(reader)
	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n := readNum(reader)
		A := readNNums(reader, n)
		res := solve(A)
		for i := 0; i < n; i++ {
			buf.WriteString(fmt.Sprintf("%d ", res[i]))
		}
		buf.WriteByte('\n')
	}

	fmt.Print(buf.String())
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

func solve(A []int) []int64 {
	n := len(A)
	// 修改 A[i]， 不会影响其他的已经构成inv的部分
	// 先减去目前它贡献的部分，即i后面比它小的部分，已经i前面比它大的部分
	// 在增加A[i]的时候，会减少前面的inv，同时增加后面的inv，也就是一个单调的函数
	inv := countInversions(A)
	var tot int64
	for i := 0; i < n; i++ {
		tot += int64(inv[i])
	}
	tot /= 2
	compress := make(map[int]int)
	for _, a := range A {
		compress[a]++
	}
	arr := make([]int, 0, len(compress))
	for k := range compress {
		arr = append(arr, k)
	}
	sort.Ints(arr)
	var id int = 1
	for _, num := range arr {
		compress[num] = id
		id += 2
	}

	tr := NewNode()

	for i := 1; i < n; i++ {
		tr.Update(compress[A[i]]+1, id-1, 0, id-1, 1)
	}

	res := make([]int64, n)

	for i := 0; i < n; i++ {
		res[i] = tot - int64(inv[i]) + int64(tr.val)
		if i+1 < n {
			tr.Update(compress[A[i+1]]+1, id-1, 0, id-1, -1)
		}
		tr.Update(0, compress[A[i]]-1, 0, id-1, 1)
	}

	return res
}

type Node struct {
	left, right *Node
	val         int
	lazy        int
}

func NewNode() *Node {
	node := new(Node)
	return node
}

func (node *Node) pushValue(v int) {
	node.val += v
	node.lazy += v
}

func (node *Node) push(l int, r int) {
	if l == r {
		return
	}
	if node.left == nil {
		node.left = NewNode()
		node.right = NewNode()
	}
	if node.lazy != 0 {
		node.left.pushValue(node.lazy)
		node.right.pushValue(node.lazy)
		node.lazy = 0
	}
}

func (node *Node) pull() {
	node.val = max(node.left.val, node.right.val)
}

func (node *Node) Update(L int, R int, l int, r int, v int) {
	if R < l || r < L {
		return
	}
	node.push(l, r)
	if L <= l && r <= R {
		node.pushValue(v)
		return
	}
	mid := (l + r) / 2
	node.left.Update(L, R, l, mid, v)
	node.right.Update(L, R, mid+1, r, v)
	node.pull()
}

type Pair struct {
	first  int
	second int
}

func countInversions(A []int) []int {
	n := len(A)

	arr := make([]int, 2*n)

	set := func(p int, v int) {
		p += n
		arr[p] = v
		for p > 1 {
			arr[p>>1] = arr[p] + arr[p^1]
			p >>= 1
		}
	}

	get := func(l int, r int) int {
		l += n
		r += n
		var res int
		for l < r {
			if l&1 == 1 {
				res += arr[l]
				l++
			}
			if r&1 == 1 {
				r--
				res += arr[r]
			}
			l >>= 1
			r >>= 1
		}
		return res
	}

	res := make([]int, n)

	ps := make([]Pair, n)
	for i := 0; i < n; i++ {
		ps[i] = Pair{A[i], i}
	}

	sort.Slice(ps, func(i, j int) bool {
		return ps[i].first < ps[j].first
	})

	for i := 0; i < n; {
		j := i
		// 比 ps[j].first 小的，但是下标在它后面的数有多少个
		for i < n && ps[i].first == ps[j].first {
			res[ps[i].second] += get(ps[i].second, n)
			i++
		}
		for j < i {
			set(ps[j].second, 1)
			j++
		}
	}
	// reset
	for i := 1; i < len(arr); i++ {
		arr[i] = 0
	}

	sort.Slice(ps, func(i, j int) bool {
		return ps[i].first > ps[j].first
	})

	for i := 0; i < n; {
		j := i
		for i < n && ps[i].first == ps[j].first {
			res[ps[i].second] += get(0, ps[i].second)
			i++
		}
		for j < i {
			set(ps[j].second, 1)
			j++
		}
	}

	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
