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

	n, q := readTwoNums(reader)

	H := readNNums(reader, n)
	S := readNNums(reader, n)

	solver := NewSolver(n, H, S)

	for q > 0 {
		q--
		a, b, c := readThreeNums(reader)
		if a == 2 {
			fmt.Println(solver.Query(b, c))
			continue
		}
		solver.Update(b, c)
	}
}

type Solver struct {
	right *FlattenedTree
	left  *FlattenedTree
}

func NewSolver(n int, H []int, S []int) Solver {
	stack := make([]int, n+1)

	process := func(i int, p []int, j int) int {
		for j > 0 && H[stack[j-1]-1] <= H[i-1] {
			j--
		}
		if j > 0 {
			p[i] = stack[j-1]
		} else {
			p[i] = 0
		}
		stack[j] = i
		j++
		return j
	}
	pRight := make([]int, n+1)
	pRight[0] = -1
	var p int
	for i := 1; i <= n; i++ {
		p = process(i, pRight, p)
	}
	pLeft := make([]int, n+1)
	pLeft[0] = -1
	p = 0
	for i := n; i > 0; i-- {
		p = process(i, pLeft, p)
	}

	X := make([]int, n+1)
	copy(X[1:], S)

	flatRight := NewFlattenedTree(pRight, X)
	flatLeft := NewFlattenedTree(pLeft, X)

	return Solver{&flatRight, &flatLeft}
}

func (solver *Solver) Update(i int, v int) {
	solver.right.update(i, v)
	solver.left.update(i, v)
}

func (solver Solver) Query(u, v int) int {
	if u <= v {
		return solver.right.query(u, v)
	}
	return solver.left.query(u, v)
}

type FlattenedTree struct {
	tree *SegTree
	n    int
	entr []int
	ext  []int
}

func NewFlattenedTree(p []int, s []int) FlattenedTree {
	n := len(p)
	adj := make([][]int, n)

	for i := 0; i < n; i++ {
		adj[i] = make([]int, 0, 10)
	}

	for u := 0; u < n; u++ {
		v := p[u]
		if v >= 0 {
			adj[u] = append(adj[u], v)
			adj[v] = append(adj[v], u)
		}
	}
	entr := make([]int, n)
	ext := make([]int, n)

	var dfs func(p, u int, time *int)

	dfs = func(p, u int, time *int) {
		entr[u] = *time
		*time++
		for _, v := range adj[u] {
			if v == p {
				continue
			}
			dfs(u, v, time)
		}
		ext[u] = *time
		*time++
	}

	dfs(-1, 0, new(int))

	arr := make([]int, 2*n)

	for i := 0; i < n; i++ {
		arr[entr[i]] = s[i]
		arr[ext[i]] = -s[i]
	}
	tree := NewSegTree(arr)

	return FlattenedTree{&tree, n, entr, ext}
}

func (tree *FlattenedTree) update(i int, v int) {
	tree.tree.update(tree.entr[i], v)
	tree.tree.update(tree.ext[i], -v)
}

func (tree FlattenedTree) isAncestor(u, v int) bool {
	return tree.entr[u] <= tree.entr[v] && tree.ext[v] <= tree.ext[u]
}

func (tree FlattenedTree) query(u, v int) int {
	if !tree.isAncestor(u, v) {
		return -1
	}
	return tree.tree.Get(tree.entr[u], tree.entr[v]+1)
}

type SegTree struct {
	arr []int
	n   int
}

func NewSegTree(arr []int) SegTree {
	n := len(arr)
	tree := make([]int, 2*n)

	for i := 0; i < n; i++ {
		tree[i+n] = arr[i]
	}

	for i := n - 1; i > 0; i-- {
		tree[i] = tree[i<<1] + tree[i<<1|1]
	}

	return SegTree{tree, n}
}

func (tree *SegTree) update(p int, v int) {
	p += tree.n
	tree.arr[p] = v

	for p > 0 {
		tree.arr[p>>1] = tree.arr[p] + tree.arr[p^1]
		p >>= 1
	}
}

func (tree SegTree) Get(l int, r int) int {
	l += tree.n
	r += tree.n
	var res int
	for l < r {
		if l&1 == 1 {
			res += tree.arr[l]
			l++
		}
		if r&1 == 1 {
			r--
			res += tree.arr[r]
		}
		l >>= 1
		r >>= 1
	}
	return res
}
