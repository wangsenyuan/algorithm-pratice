package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m := readTwoNums(reader)

	s, _ := reader.ReadBytes('\n')

	var pos int
	P := make([]float64, n)
	for i := 0; i < n; i++ {
		pos = readFloat64(s, pos, &P[i]) + 1
	}

	Q := make([][]int, m)
	U := make([]float64, m)

	for i := 0; i < m; i++ {
		pos = 0
		s, _ := reader.ReadBytes('\n')
		Q[i] = make([]int, 3)
		for j := 0; j < 3; j++ {
			pos = readInt(s, pos, &Q[i][j]) + 1
		}
		if Q[i][0] == 1 {
			readFloat64(s, pos, &U[i])
		}
	}

	res := solve(P, Q, U)

	var buf bytes.Buffer
	for _, cur := range res {
		buf.WriteString(fmt.Sprintf("%.10f\n", cur))
	}

	fmt.Print(buf.String())
}

func readFloat64(bytes []byte, from int, val *float64) int {
	i := from

	var real int64

	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		real = real*10 + int64(bytes[i]-'0')
		i++
	}

	if i == len(bytes) || bytes[i] != '.' {
		*val = float64(real)
		return i
	}

	// bytes[i] == '.'
	i++

	var fraq float64
	var base float64 = 0.1
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		fraq += base * float64(bytes[i]-'0')
		base /= 10
		i++
	}

	*val = float64(real) + fraq

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

func solve(P []float64, Q [][]int, UP []float64) []float64 {
	tree := Build(P)

	var res []float64

	for i := 0; i < len(Q); i++ {
		l, r := Q[i][1], Q[i][2]
		l--
		r--
		if Q[i][0] == 0 {
			// query
			res = append(res, tree.Get(l, r))
		} else {
			tree.Update(l, r, UP[i])
		}
	}

	return res
}

const LI = 100

type Node struct {
	tay []float64
}

func merge(a, b, c *Node) {
	for i := 0; i < LI; i++ {
		a.tay[i] = b.tay[i] + c.tay[i]
	}
}

type Tree struct {
	tree    []*Node
	temp    []*Node
	size    int
	updates []float64
}

func Build(arr []float64) *Tree {
	n := len(arr)
	tree := make([]*Node, 4*n+2)

	for u := 0; u < len(tree); u++ {
		tree[u] = new(Node)
		tree[u].tay = make([]float64, LI)
	}

	updates := make([]float64, 4*n+2)
	var loop func(u int, l int, r int)

	loop = func(u int, l int, r int) {
		updates[u] = 1
		if l == r {

			tree[u].tay[0] = arr[l]
			for i := 1; i < LI; i++ {
				tree[u].tay[i] = tree[u].tay[i-1] * arr[l]
			}
		} else {
			mid := (l + r) >> 1
			loop(u<<1, l, mid)
			loop((u<<1)|1, mid+1, r)
			merge(tree[u], tree[u<<1], tree[(u<<1)|1])
		}
	}
	loop(1, 0, n-1)
	temp := make([]*Node, 200)
	for i := 0; i < 200; i++ {
		temp[i] = new(Node)
		temp[i].tay = make([]float64, LI)
	}
	return &Tree{tree, temp, n, updates}
}

func (tree *Tree) pushUpdateAtNode(u int, l int, r int) {
	if 1-tree.updates[u] <= 1e-7 {
		return
	}
	var cur float64 = 1
	p := tree.updates[u]
	for i := 0; i < LI; i++ {
		cur *= p
		tree.tree[u].tay[i] *= cur
	}
	if l != r {
		tree.updates[u<<1] *= p
		tree.updates[(u<<1)|1] *= p
	}
	tree.updates[u] = 1
}

func (tree *Tree) Update(l, r int, up float64) {

	var loop func(u int, L int, R int, l int, r int)
	loop = func(u int, L int, R int, l int, r int) {
		tree.pushUpdateAtNode(u, L, R)
		if L == l && R == r {
			tree.updates[u] *= up
			tree.pushUpdateAtNode(u, L, R)
			return
		}
		m := (L + R) >> 1
		if r <= m {
			loop(u<<1, L, m, l, r)
			tree.pushUpdateAtNode((u<<1)|1, m+1, R)
		} else if l > m {
			tree.pushUpdateAtNode(u<<1, L, m)
			loop((u<<1)|1, m+1, R, l, r)
		} else {
			loop(u<<1, L, m, l, m)
			loop((u<<1)|1, m+1, R, m+1, r)
		}
		merge(tree.tree[u], tree.tree[u<<1], tree.tree[(u<<1)|1])
	}

	loop(1, 0, tree.size-1, l, r)
}

func (tree *Tree) Get(l, r int) float64 {
	temp := tree.temp
	var ptr int
	var loop func(u int, L int, R int, l, r int) *Node
	loop = func(u int, L int, R int, l int, r int) *Node {
		tree.pushUpdateAtNode(u, L, R)
		if L == l && R == r {
			return tree.tree[u]
		}
		m := (L + R) >> 1
		t := temp[ptr]
		ptr++
		if r <= m {
			t = loop(u<<1, L, m, l, r)
			tree.pushUpdateAtNode((u<<1)|1, m+1, R)
		} else if l > m {
			tree.pushUpdateAtNode(u<<1, L, m)
			t = loop((u<<1)|1, m+1, R, l, r)
		} else {
			merge(t, loop(u<<1, L, m, l, m), loop((u<<1)|1, m+1, R, m+1, r))
		}
		return t
	}
	t := loop(1, 0, tree.size-1, l, r)
	var ans float64
	for i := 1; i <= LI; i++ {
		if t.tay[i-1] > 1e-13 {
			ans += t.tay[i-1] / float64(i)
		} else {
			break
		}
	}
	return math.Exp(-ans)
}
