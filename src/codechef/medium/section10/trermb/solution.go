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

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n, q := readTwoNums(reader)
		E := make([][]int, n-1)
		for i := 0; i < n-1; i++ {
			E[i] = readNNums(reader, 2)
		}

		C := readNNums(reader, n)
		A := readNNums(reader, n)
		solver := NewSolver(n, E, C, A)

		for q > 0 {
			q--
			a, x := readTwoNums(reader)
			res := solver.ChangeColor(a, x)
			fmt.Println(res)
		}
	}
}

type Solver struct {
	n      int
	C      []int
	A      []int64
	trees  []*Node
	D      []int
	DS     []int
	SD     []int
	P      [][]int
	answer int64
}

func NewSolver(n int, E [][]int, C []int, a []int) Solver {
	A := make([]int64, n)
	for i := 0; i < n; i++ {
		A[i] = int64(a[i])
	}

	conns := make([][]int, n)
	for i := 0; i < n; i++ {
		conns[i] = make([]int, 0, 5)
	}

	for _, e := range E {
		u, v := e[0], e[1]
		u--
		v--
		conns[u] = append(conns[u], v)
		conns[v] = append(conns[v], u)
	}

	P := make([][]int, n)
	for i := 0; i < n; i++ {
		P[i] = make([]int, 20)
	}

	D := make([]int, n)
	DS := make([]int, n)
	SD := make([]int, n)
	var dfs func(p, u int)

	var dt int

	dfs = func(p, u int) {
		DS[u] = dt
		SD[dt] = u
		dt++
		P[u][0] = p

		for _, v := range conns[u] {
			if v == p {
				continue
			}
			D[v] = D[u] + 1
			A[v] += A[u]
			dfs(u, v)
		}
	}

	dfs(-1, 0)

	for h := 1; h < 20; h++ {
		for i := 0; i < n; i++ {
			if P[i][h-1] >= 0 {
				P[i][h] = P[P[i][h-1]][h-1]
			}
		}
	}

	trees := make([]*Node, n)

	solver := Solver{n, C, A, trees, D, DS, SD, P, 0}

	for i := 0; i < n; i++ {
		solver.answer += solver.Update(i)
		x := C[i] - 1
		trees[x] = Insert(trees[x], DS[i])
	}

	return solver
}

func lca(P [][]int, D []int, u, v int) int {
	if D[u] < D[v] {
		u, v = v, u
	}
	for i := 19; i >= 0; i-- {
		if D[u]-(1<<uint(i)) >= D[v] {
			u = P[u][i]
		}
	}
	if u == v {
		return u
	}

	for i := 19; i >= 0; i-- {
		if P[u][i] != P[v][i] {
			u = P[u][i]
			v = P[v][i]
		}
	}
	return P[u][0]
}

func (solver *Solver) ChangeColor(i int, newColor int) int64 {
	i--

	oldColor := solver.C[i]
	solver.trees[oldColor-1] = Delete(solver.trees[oldColor-1], solver.DS[i])
	solver.answer -= solver.Update(i)

	solver.C[i] = newColor

	solver.answer += solver.Update(i)
	solver.trees[newColor-1] = Insert(solver.trees[newColor-1], solver.DS[i])

	return solver.answer
}

func (solver *Solver) Update(i int) int64 {
	var w = -1
	c := solver.C[i] - 1
	t1 := FindEqualOrLess(solver.trees[c], solver.DS[i])
	t2 := FindEqualOrGreater(solver.trees[c], solver.DS[i])
	if t1 != nil {
		u := solver.SD[t1.key]
		w = lca(solver.P, solver.D, u, i)
	}
	if t2 != nil {
		u := solver.SD[t2.key]
		w2 := lca(solver.P, solver.D, u, i)
		if w < 0 || solver.D[w2] > solver.D[w] {
			w = w2
		}
	}

	y := solver.A[i]
	if w >= 0 {
		y -= solver.A[w]
	}
	return y
}

/**
* this is a AVL tree
 */
type Node struct {
	key         int
	height      int
	cnt         int
	left, right *Node
}

func (node *Node) Height() int {
	if node == nil {
		return 0
	}
	return node.height
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func NewNode(key int) *Node {
	node := new(Node)
	node.key = key
	node.height = 1
	node.cnt = 1
	return node
}

func rightRotate(y *Node) *Node {
	x := y.left

	t2 := x.right

	x.right = y
	y.left = t2
	y.height = max(y.left.Height(), y.right.Height()) + 1
	x.height = max(x.left.Height(), x.right.Height()) + 1

	return x
}

func leftRotate(x *Node) *Node {
	y := x.right
	t2 := y.left

	y.left = x
	x.right = t2

	x.height = max(x.left.Height(), x.right.Height()) + 1
	y.height = max(y.left.Height(), y.right.Height()) + 1

	return y
}

func (node *Node) GetBalance() int {
	if node == nil {
		return 0
	}
	return node.left.Height() - node.right.Height()
}

func FindEqualOrGreater(node *Node, key int) *Node {
	if node == nil {
		return nil
	}
	if node.key < key {
		return FindEqualOrGreater(node.right, key)
	}
	res := FindEqualOrGreater(node.left, key)
	if res == nil {
		return node
	}
	return res
}

func FindEqualOrLess(node *Node, key int) *Node {
	if node == nil {
		return nil
	}
	if node.key > key {
		return FindEqualOrLess(node.left, key)
	}
	res := FindEqualOrLess(node.right, key)
	if res == nil {
		return node
	}
	return res
}

func Insert(node *Node, key int) *Node {
	if node == nil {
		return NewNode(key)
	}
	if node.key == key {
		node.cnt++
		return node
	}

	if node.key > key {
		node.left = Insert(node.left, key)
	} else {
		node.right = Insert(node.right, key)
	}

	node.height = max(node.left.Height(), node.right.Height()) + 1
	balance := node.GetBalance()

	if balance > 1 && key < node.left.key {
		return rightRotate(node)
	}

	if balance < -1 && key > node.right.key {
		return leftRotate(node)
	}

	if balance > 1 && key > node.left.key {
		node.left = leftRotate(node.left)
		return rightRotate(node)
	}

	if balance < -1 && key < node.right.key {
		node.right = rightRotate(node.right)
		return leftRotate(node)
	}

	return node
}

func MinValueNode(root *Node) *Node {
	cur := root

	for cur.left != nil {
		cur = cur.left
	}

	return cur
}

func (root *Node) MinValue() int {
	if root == nil {
		return 0
	}
	node := MinValueNode(root)

	return node.key
}

func Delete(root *Node, key int) *Node {
	if root == nil {
		return nil
	}

	if key < root.key {
		root.left = Delete(root.left, key)
	} else if key > root.key {
		root.right = Delete(root.right, key)
	} else {
		root.cnt--
		if root.cnt > 0 {
			return root
		}
		if root.left == nil || root.right == nil {
			tmp := root.left
			if root.left == nil {
				tmp = root.right
			}
			root = tmp
		} else {
			tmp := MinValueNode(root.right)

			root.key = tmp.key
			root.cnt = tmp.cnt
			// make sure tmp node deleted after call delete on root.right
			tmp.cnt = 1
			root.right = Delete(root.right, tmp.key)
		}
	}

	if root == nil {
		return root
	}

	root.height = max(root.left.Height(), root.right.Height()) + 1
	balance := root.GetBalance()

	if balance > 1 && root.left.GetBalance() >= 0 {
		return rightRotate(root)
	}

	if balance > 1 && root.left.GetBalance() < 0 {
		root.left = leftRotate(root.left)
		return rightRotate(root)
	}

	if balance < -1 && root.right.GetBalance() <= 0 {
		return leftRotate(root)
	}

	if balance < -1 && root.right.GetBalance() > 0 {
		root.right = rightRotate(root.right)
		return leftRotate(root)
	}

	return root
}
