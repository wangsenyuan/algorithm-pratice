package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	n, q := readTwoNums(reader)
	A := readNNums(reader, n)
	cmds := make([][]int, q)
	for i := 0; i < q; i++ {
		line, _ := reader.ReadBytes('\n')
		var op int
		pos := readInt(line, 0, &op) + 1
		if op == 1 {
			cmds[i] = make([]int, 3)
			cmds[i][0] = 1
			pos = readInt(line, pos, &cmds[i][1]) + 1
			readInt(line, pos, &cmds[i][2])
		} else {
			cmds[i] = make([]int, 2)
			cmds[i][0] = 2
			readInt(line, pos, &cmds[i][1])
		}
	}
	res := solve(n, A, cmds)
	var buf bytes.Buffer
	for _, num := range res {
		buf.WriteString(fmt.Sprintf("%d\n", num))
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

func solve(n int, a []int, cmds [][]int) []int64 {
	A := make([]int, n+1)
	copy(A[1:], a)

	each := make([]*Node, n+1)

	for i := 1; i <= n; i++ {
		each[A[i]] = Insert(each[A[i]], i)
	}

	seg := NewSegmentTree(n)
	seg.Build(1, 1, n)

	for i := 1; i <= n; i++ {
		each[i] = Insert(each[i], 0)
		each[i] = Insert(each[i], n+1)

		back := MinValueNode(each[i])
		front := FindNextNode(each[i], back)
		// front.key >= 1
		for front != nil {
			seg.Add(1, back.key+1, front.key-1, int64(back.key))
			if front.key <= n {
				seg.Add(1, front.key, front.key, int64(front.key))
			}
			back = front
			front = FindNextNode(each[i], back)
		}
	}

	remove := func(x int, pos int) {
		each[x] = Delete(each[x], pos)
		r := FindGreaterOrEqual(each[x], pos)
		l := FindPrevNode(each[x], r)
		seg.Add(1, pos, r.key-1, -int64(pos-l.key))
	}

	add := func(x int, pos int) {
		r := FindGreaterOrEqual(each[x], pos)
		l := FindPrevNode(each[x], r)
		seg.Add(1, pos, r.key-1, int64(pos-l.key))
		each[x] = Insert(each[x], pos)
	}

	res := make([]int64, 0, len(cmds))

	for _, cmd := range cmds {
		op := cmd[0]
		if op == 1 {
			x := cmd[1]
			v := cmd[2]
			// set A[x] = y
			remove(A[x], x)
			add(v, x)
			A[x] = v
		} else {
			k := cmd[1]
			res = append(res, seg.Get(1, 1, k))
		}
	}

	return res
}

type Item struct {
	l, r int
	sum  int64
	add  int64
}

type SegmentTree struct {
	nodes []*Item
	size  int
}

func NewSegmentTree(n int) *SegmentTree {
	tree := new(SegmentTree)
	tree.nodes = make([]*Item, 4*n+10)
	tree.size = n
	return tree
}

func (tree *SegmentTree) pushup(x int) {
	tree.nodes[x].sum = tree.nodes[2*x].sum + tree.nodes[2*x+1].sum
}

func (tree *SegmentTree) tag(x int, add int64) {
	tree.nodes[x].sum += int64(tree.nodes[x].r-tree.nodes[x].l+1) * add
	tree.nodes[x].add += add
}

func (tree *SegmentTree) pushdown(x int) {
	if tree.nodes[x].add != 0 {
		tree.tag(2*x, tree.nodes[x].add)
		tree.tag(2*x+1, tree.nodes[x].add)
		tree.nodes[x].add = 0
	}
}

func (tree *SegmentTree) Build(x int, l, r int) {
	tree.nodes[x] = new(Item)
	tree.nodes[x].l = l
	tree.nodes[x].r = r
	if l != r {
		mid := (l + r) / 2
		tree.Build(2*x, l, mid)
		tree.Build(2*x+1, mid+1, r)
		tree.pushup(x)
	}
}

func (tree *SegmentTree) Add(x int, l int, r int, v int64) {
	if l <= r {
		if l <= tree.nodes[x].l && tree.nodes[x].r <= r {
			tree.tag(x, v)
		} else {
			tree.pushdown(x)

			mid := (tree.nodes[x].l + tree.nodes[x].r) / 2

			if l <= mid {
				tree.Add(2*x, l, r, v)
			}

			if mid < r {
				tree.Add(2*x+1, l, r, v)
			}
			tree.pushup(x)
		}
	}
}

func (tree *SegmentTree) Get(x int, l, r int) int64 {
	if l <= tree.nodes[x].l && tree.nodes[x].r <= r {
		return tree.nodes[x].sum
	}
	tree.pushdown(x)
	mid := (tree.nodes[x].l + tree.nodes[x].r) / 2

	var ans int64

	if l <= mid {
		ans += tree.Get(2*x, l, r)
	}

	if mid < r {
		ans += tree.Get(2*x+1, l, r)
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
	if root == nil {
		return nil
	}

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

func FindGreaterOrEqual(node *Node, key int) *Node {
	if node == nil {
		return nil
	}
	if node.key >= key {
		tmp := FindGreaterOrEqual(node.left, key)
		if tmp != nil {
			return tmp
		}
		return node
	}
	return FindGreaterOrEqual(node.right, key)
}

func FindPrevNode(node *Node, target *Node) *Node {
	if node == target {
		return MaxValueNode(node.left)
	}
	if node.right == target {
		tmp := MaxValueNode(target.left)
		if tmp != nil {
			return tmp
		}
		return node
	}
	if node.key < target.key {
		return FindPrevNode(node.right, target)
	}
	return FindPrevNode(node.right, target)
}

func MaxValueNode(node *Node) *Node {
	if node == nil {
		return nil
	}
	if node.right != nil {
		return MaxValueNode(node.right)
	}
	return node
}

func FindNextNode(node *Node, target *Node) *Node {
	if node == target {
		return MinValueNode(node.right)
	}
	if node.left == target {
		tmp := MinValueNode(target.right)
		if tmp != nil {
			return tmp
		}
		return node
	}
	if node.key < target.key {
		return FindNextNode(node.right, target)
	}
	return FindNextNode(node.left, target)
}
