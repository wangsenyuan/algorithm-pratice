package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m := readTwoNums(reader)

	solver := NewSolver(n)

	for m > 0 {
		m--
		s, _ := reader.ReadBytes('\n')
		var t int
		pos := readInt(s, 0, &t)
		if t == 0 {
			var l, r, x int
			pos = readInt(s, pos+1, &l)
			pos = readInt(s, pos+1, &r)
			readInt(s, pos+1, &x)
			solver.Update(l, r, x)
		} else {
			var p int
			readInt(s, pos+1, &p)
			fmt.Println(solver.Query(p))
		}
	}
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

type Solver struct {
	n       int
	appears []*Node
	arr     []int
}

func NewSolver(n int) Solver {
	appears := make([]*Node, n+1)
	arr := make([]int, n+2)

	return Solver{n, appears, arr}
}

func (solver *Solver) modifyCount(p int, v int) {
	arr := solver.arr

	for p < len(arr) {
		arr[p] += v
		p += p & -p
	}
}

func (solver *Solver) Update(l, r int, x int) {
	root := solver.appears[x]

	it := FindEqualOrLess(root, Item{l, 0})

	if it != nil && it.item.val >= l {
		l = it.item.key
	}

	for {
		it = FindEqualOrGreater(root, Item{l, 0})
		if it == nil || it.item.key > r {
			break
		}
		r = max(r, it.item.val)
		solver.modifyCount(it.item.key, -1)
		solver.modifyCount(it.item.val+1, 1)
		root = Delete(root, it.item)
	}

	solver.modifyCount(l, 1)
	solver.modifyCount(r+1, -1)

	solver.appears[x] = Insert(root, Item{l, r})
}

func (solver Solver) Query(p int) int {
	var res int
	for p > 0 {
		res += solver.arr[p]
		p -= p & -p
	}
	return res
}

type Item struct {
	key int
	val int
}

func (this Item) Less(that Item) bool {
	return this.key < that.key || this.key == that.key && this.val < that.val
}

/**
* this is a AVL tree
 */
type Node struct {
	item        Item
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

func NewNode(item Item) *Node {
	node := new(Node)
	node.item = item
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

func FindEqualOrGreater(node *Node, item Item) *Node {
	if node == nil {
		return nil
	}
	if node.item.Less(item) {
		return FindEqualOrGreater(node.right, item)
	}

	res := FindEqualOrGreater(node.left, item)
	if res == nil {
		return node
	}
	return res
}

func FindEqualOrLess(node *Node, item Item) *Node {
	if node == nil {
		return nil
	}
	if item.Less(node.item) {
		return FindEqualOrLess(node.left, item)
	}

	res := FindEqualOrLess(node.right, item)
	if res == nil {
		return node
	}
	return res
}

func Insert(node *Node, item Item) *Node {
	if node == nil {
		return NewNode(item)
	}
	if node.item.key == item.key {
		node.cnt++
		return node
	}

	if item.Less(node.item) {
		node.left = Insert(node.left, item)
	} else {
		node.right = Insert(node.right, item)
	}

	node.height = max(node.left.Height(), node.right.Height()) + 1
	balance := node.GetBalance()

	if balance > 1 && item.Less(node.left.item) {
		return rightRotate(node)
	}

	if balance < -1 && node.right.item.Less(item) {
		return leftRotate(node)
	}

	if balance > 1 && node.left.item.Less(item) {
		node.left = leftRotate(node.left)
		return rightRotate(node)
	}

	if balance < -1 && item.Less(node.right.item) {
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

func Delete(root *Node, item Item) *Node {
	if root == nil {
		return nil
	}

	if item.Less(root.item) {
		root.left = Delete(root.left, item)
	} else if root.item.Less(item) {
		root.right = Delete(root.right, item)
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

			root.item = tmp.item
			root.cnt = tmp.cnt
			// make sure tmp node deleted after call delete on root.right
			tmp.cnt = 1
			root.right = Delete(root.right, tmp.item)
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
