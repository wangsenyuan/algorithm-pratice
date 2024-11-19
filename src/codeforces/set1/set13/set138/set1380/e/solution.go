package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d\n", x))
	}

	fmt.Print(buf.String())
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

func process(reader *bufio.Reader) []int {
	n, m := readTwoNums(reader)
	a := readNNums(reader, n)
	queries := make([][]int, m-1)
	for i := 0; i < m-1; i++ {
		queries[i] = readNNums(reader, 2)
	}
	return solve(a, m, queries)
}

func solve(discs []int, m int, queries [][]int) (ans []int) {
	// n := len(discs)
	g := NewGraph(m*2, m*2)
	set := make([]int, m)
	for i := 0; i < m; i++ {
		set[i] = i
	}

	for i, cur := range queries {
		u, v := cur[0]-1, cur[1]-1
		// u = set[u]
		// v = set[v]
		g.AddEdge(m+i, set[u])
		g.AddEdge(m+i, set[v])
		set[u] = m + i
	}
	h := bits.Len(uint(2 * m))

	fa := make([][]int, 2*m)
	for i := 0; i < 2*m; i++ {
		fa[i] = make([]int, h)
	}
	dep := make([]int, 2*m)
	var dfs func(u int, p int)
	dfs = func(u int, p int) {
		fa[u][0] = p
		for j := 1; j < h; j++ {
			fa[u][j] = fa[fa[u][j-1]][j-1]
		}
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			dep[v] = dep[u] + 1
			dfs(v, u)
		}
	}

	dfs(2*m-2, 2*m-2)

	lca := func(u int, v int) int {
		if dep[u] < dep[v] {
			u, v = v, u
		}
		for i := h - 1; i >= 0; i-- {
			if dep[u]-1<<i >= dep[v] {
				u = fa[u][i]
			}
		}
		if u == v {
			return u
		}
		for i := h - 1; i >= 0; i-- {
			if fa[u][i] != fa[v][i] {
				u = fa[u][i]
				v = fa[v][i]
			}
		}
		return fa[u][0]
	}

	ans = make([]int, m)

	n := len(discs)
	for i := range n - 1 {
		t := lca(discs[i]-1, discs[i+1]-1)
		if t < m {
			ans[0]++
		} else {
			ans[t-m+1]++
		}
	}

	for i := 1; i < m; i++ {
		ans[i] += ans[i-1]
	}
	for i := 0; i < m; i++ {
		ans[i] = n - 1 - ans[i]
	}
	return ans
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
	e++
	next := make([]int, e)
	to := make([]int, e)
	return &Graph{nodes, next, to, 0}
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}

type pair struct {
	first  int
	second int
}

func solve1(discs []int, m int, queries [][]int) (ans []int) {
	n := len(discs)

	set := make([]int, n+m)
	for i := 0; i < m; i++ {
		set[i] = i
	}
	for i := 0; i < n; i++ {
		discs[i]--
		set[i+m] = discs[i]
	}

	var find func(u int) int
	find = func(u int) int {
		if set[u] != u {
			set[u] = find(set[u])
		}
		return set[u]
	}

	var tr *Node

	head := make([]int, m)
	sz := make([]int, m)
	for i := 0; i < m; i++ {
		head[i] = -1
	}
	next := make([]int, n)
	prev := make([]int, n)

	for i := 0; i < n; i++ {
		next[i] = n
		prev[i] = -1
	}

	for i := 0; i < n; {
		j := i
		for i < n && find(i+m) == find(j+m) {
			i++
		}
		t := find(j + m)

		if head[t] >= 0 {
			next[head[t]] = j
			prev[j] = head[t]
		}
		head[t] = j
		sz[t]++

		tr = Insert(tr, j)
	}

	remove := func(x int, a int) {
		if head[a] == x {
			head[a] = prev[x]
		}
		l, r := prev[x], next[x]
		if l >= 0 {
			next[l] = r
		}
		if r < n {
			prev[r] = l
		}
		prev[x] = -1
		next[x] = n
		sz[a]--
	}

	merge := func(a, b int) {
		a = find(a)
		b = find(b)
		if sz[a] < sz[b] {
			a, b = b, a
		}

		for x := head[b]; x >= 0; x = head[b] {
			remove(x, b)
			if x > 0 {
				l := tr.GetLt(x - 1)
				if l != nil && find(l.key+m) == a {
					y := l.key
					// delete 的时候会改变l的值
					tr = Delete(tr, x)
					x = y
				}
			}

			if x+1 < n {
				r := tr.GetLe(x + 1)
				if r != nil && find(r.key+m) == a {
					remove(r.key, a)
					tr = Delete(tr, r.key)
				}
			}

			if find(x+m) != a {
				// 如果这个区间和前面的没有merge在一起
				if head[a] >= 0 {
					next[head[a]] = x
					prev[x] = head[a]
				}
				head[a] = x
				next[x] = n
				set[x+m] = a

				sz[a]++
			}
		}

		set[b] = a
	}

	ans = append(ans, tr.Size()-1)

	for _, cur := range queries {
		a, b := cur[0], cur[1]
		merge(a-1, b-1)
		ans = append(ans, tr.Size()-1)
	}

	return ans
}

/**
* this is a AVL tree
 */
type Node struct {
	key         int
	height      int
	cnt         int
	sz          int
	sum         int
	left, right *Node
}

func (node *Node) Height() int {
	if node == nil {
		return 0
	}
	return node.height
}

func NewNode(key int) *Node {
	node := new(Node)
	node.key = key
	node.height = 1
	node.cnt = 1
	node.sz = 1
	node.sum = key
	return node
}

func rightRotate(y *Node) *Node {
	x := y.left

	t2 := x.right

	x.right = y
	y.left = t2
	//y.height = max(y.left.Height(), y.right.Height()) + 1
	y.update()
	//x.height = max(x.left.Height(), x.right.Height()) + 1
	x.update()

	return x
}

func leftRotate(x *Node) *Node {
	y := x.right
	t2 := y.left

	y.left = x
	x.right = t2

	//x.height = max(x.left.Height(), x.right.Height()) + 1
	//y.height = max(y.left.Height(), y.right.Height()) + 1
	x.update()
	y.update()

	return y
}

func (node *Node) GetBalance() int {
	if node == nil {
		return 0
	}
	return node.left.Height() - node.right.Height()
}

func (node *Node) Size() int {
	if node == nil {
		return 0
	}
	return node.sz
}

func (node *Node) Sum() int {
	if node == nil {
		return 0
	}
	return node.sum
}

func (node *Node) update() {
	node.height = max(node.left.Height(), node.right.Height()) + 1
	node.sz = node.left.Size() + node.right.Size() + node.cnt
	node.sum = node.left.Sum() + node.right.Sum() + node.cnt*node.key
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

	//node.height = max(node.left.Height(), node.right.Height()) + 1
	node.update()

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
func (root *Node) MaxValue() int {
	node := root
	for node.right != nil {
		node = node.right
	}
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
		root.sz--
		root.sum -= key
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
			root.sum = tmp.sum
			// make sure tmp node deleted after call delete on root.right
			tmp.cnt = 1
			root.right = Delete(root.right, tmp.key)
		}
	}

	if root == nil {
		return root
	}

	//root.height = max(root.left.Height(), root.right.Height()) + 1
	root.update()

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

func (tr *Node) GetLe(key int) *Node {
	if tr == nil {
		return nil
	}
	if tr.key < key {
		return tr.right.GetLe(key)
	}
	// tr.key >= key
	res := tr.left.GetLe(key)
	if res == nil {
		return tr
	}
	return res
}

func (tr *Node) GetLt(key int) *Node {
	if tr == nil {
		return nil
	}
	if tr.key > key {
		return tr.left.GetLt(key)
	}
	// tr.key <= key
	res := tr.right.GetLt(key)
	if res == nil {
		return tr
	}
	return res
}
