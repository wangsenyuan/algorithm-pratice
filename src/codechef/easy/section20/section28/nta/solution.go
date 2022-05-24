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
		n := readNum(reader)
		E := make([][]int, n-1)
		for i := 0; i < n-1; i++ {
			E[i] = readNNums(reader, 2)
		}

		A := readNInt64s(reader, n)

		res := solve(n, A, E)

		for _, num := range res {
			buf.WriteString(fmt.Sprintf("%d ", num))
		}
		buf.WriteByte('\n')
	}

	fmt.Print(buf.String())
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

func solve(n int, A []int64, E [][]int) []int64 {

	g := NewGraph(n, len(E)*2)

	for _, e := range E {
		u, v := e[0], e[1]
		u--
		v--
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	size := make([]int, n)
	big := make([]int, n)
	var dfs_size func(p int, u int)

	dfs_size = func(p int, u int) {
		big[u] = -1
		size[u] = 1
		for i := g.node[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dfs_size(u, v)
				size[u] += size[v]
				if big[u] < 0 || size[big[u]] < size[v] {
					big[u] = v
				}
			}
		}
	}

	dfs_size(0, 0)

	ans := make([]int64, n)

	findAndUpdate := func(node *Node, val int64, ans *int64) *Node {
		tmp := FindEqOrLessNode(node, val)
		if tmp != nil {
			if *ans < 0 || tmp.key^val < *ans {
				*ans = val ^ tmp.key
			}
		}
		tmp = FindEqOrGreaterNode(node, val)
		if tmp != nil {
			if *ans < 0 || tmp.key^val < *ans {
				*ans = val ^ tmp.key
			}
		}
		return Insert(node, val)
	}

	forest := make([]*Node, n)

	arr := make([][]int64, n)

	var dfs func(p, u int)

	dfs = func(p, u int) {
		ans[u] = -1
		arr[u] = make([]int64, 0, 1)
		if big[u] >= 0 {
			dfs(u, big[u])
			arr[u] = arr[big[u]]
			forest[u] = forest[big[u]]
			ans[u] = ans[big[u]]
		}

		forest[u] = findAndUpdate(forest[u], A[u], &ans[u])
		arr[u] = append(arr[u], A[u])

		for i := g.node[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p == v || big[u] == v {
				continue
			}
			dfs(u, v)
			for _, x := range arr[v] {
				forest[u] = findAndUpdate(forest[u], x, &ans[u])
				arr[u] = append(arr[u], x)
			}
		}
	}

	dfs(0, 0)

	return ans
}

type Graph struct {
	node []int
	next []int
	to   []int
	cur  int
}

func NewGraph(n int, e int) *Graph {
	g := new(Graph)
	g.node = make([]int, n)
	g.next = make([]int, e+1)
	g.to = make([]int, e+1)
	return g
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.node[u]
	g.node[u] = g.cur
	g.to[g.cur] = v
}

/**
* this is a AVL tree
 */
type Node struct {
	key         int64
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

func NewNode(key int64) *Node {
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

func Search(node *Node, key int64) *Node {
	if node == nil || node.key == key {
		return node
	}
	if key < node.key {
		return Search(node.left, key)
	}
	return Search(node.right, key)
}

func Insert(node *Node, key int64) *Node {
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

func MaxValueNode(root *Node) *Node {
	cur := root

	for cur.right != nil {
		cur = cur.right
	}
	return cur
}
func Delete(root *Node, key int64) *Node {
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

func FindEqOrLessNode(node *Node, key int64) *Node {
	if node == nil {
		return nil
	}
	if node.key > key {
		return FindEqOrLessNode(node.left, key)
	}
	// node.key <= key, if any better
	ans := FindEqOrLessNode(node.right, key)
	if ans != nil {
		return ans
	}
	return node
}

func FindEqOrGreaterNode(node *Node, key int64) *Node {
	if node == nil {
		return nil
	}
	if node.key < key {
		return FindEqOrGreaterNode(node.right, key)
	}
	ans := FindEqOrGreaterNode(node.left, key)
	if ans != nil {
		return ans
	}
	return node
}
