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
		A := readNNums(reader, n)
		res := solve(n, A)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

func solve(n int, A []int) int64 {
	pos := make([]int, n)

	for i := 0; i < n; i++ {
		A[i]--
		pos[A[i]] = i
	}

	var root *Node

	var L int

	insert := func(val Pair) {
		root = Insert(root, val)

		for {
			prev := -1
			if val != MinValueNode(root).item {
				prev = PreviouseNode(root, val).item.second
			}

			next := n + 1
			if val != MaxValueNode(root).item {
				next = NextNode(root, val).item.second
			}

			if prev < val.second && val.second < next {
				break
			}

			v1 := Pair{pos[A[L]], A[L]}
			v2 := Pair{pos[L], L}
			root = Delete(root, v1)
			root = Delete(root, v2)
			L++

			if v1 == val || v2 == val {
				break
			}
		}
	}
	var res int64
	for R := 0; R < n; R++ {
		if L <= A[R] && A[R] <= R {
			insert(Pair{pos[A[R]], A[R]})
		}

		if L <= pos[R] && pos[R] <= R {
			insert(Pair{pos[R], R})
		}
		res += int64(R - L + 1)
	}

	return res
}

type Pair struct {
	first, second int
}

func (this Pair) Less(that Pair) bool {
	return this.first < that.first || this.first == that.first && this.second < that.second
}

/**
* this is a AVL tree
 */
type Node struct {
	item        Pair
	height      int
	size        int
	left, right *Node
}

func (node *Node) Height() int {
	if node == nil {
		return 0
	}
	return node.height
}

func (node *Node) Size() int {
	if node == nil {
		return 0
	}
	return node.size
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func NewNode(item Pair) *Node {
	node := new(Node)
	node.item = item
	node.height = 1
	node.size = 1
	return node
}

func rightRotate(y *Node) *Node {
	x := y.left

	t2 := x.right

	x.right = y
	y.left = t2
	y.height = max(y.left.Height(), y.right.Height()) + 1
	y.size = y.left.Size() + y.right.Size() + 1
	x.height = max(x.left.Height(), x.right.Height()) + 1
	x.size = x.left.Size() + x.right.Size() + 1

	return x
}

func leftRotate(x *Node) *Node {
	y := x.right
	t2 := y.left

	y.left = x
	x.right = t2

	x.height = max(x.left.Height(), x.right.Height()) + 1
	x.size = x.left.Size() + x.right.Size() + 1
	y.height = max(y.left.Height(), y.right.Height()) + 1
	y.size = y.left.Size() + y.right.Size() + 1

	return y
}

func (node *Node) GetBalance() int {
	if node == nil {
		return 0
	}
	return node.left.Height() - node.right.Height()
}

func Insert(node *Node, item Pair) *Node {
	if node == nil {
		return NewNode(item)
	}
	if node.item == item {
		return node
	}
	if item.Less(node.item) {
		node.left = Insert(node.left, item)
	} else {
		node.right = Insert(node.right, item)
	}

	node.height = max(node.left.Height(), node.right.Height()) + 1
	node.size = node.left.Size() + node.right.Size() + 1
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

func Delete(root *Node, item Pair) *Node {
	if root == nil {
		return nil
	}

	if item.Less(root.item) {
		root.left = Delete(root.left, item)
	} else if root.item.Less(item) {
		root.right = Delete(root.right, item)
	} else {
		if root.left == nil || root.right == nil {
			tmp := root.left
			if root.left == nil {
				tmp = root.right
			}
			root = tmp
		} else {
			tmp := MinValueNode(root.right)

			root.item = tmp.item
			// make sure tmp node deleted after call delete on root.right
			root.right = Delete(root.right, tmp.item)
		}
	}

	if root == nil {
		return root
	}

	root.height = max(root.left.Height(), root.right.Height()) + 1
	root.size = root.left.Size() + root.right.Size() + 1

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

func PreviouseNode(root *Node, cur Pair) *Node {
	// find previouse node, that has item < current
	if root == nil {
		return nil
	}

	if root.item.Less(cur) {
		tmp := PreviouseNode(root.right, cur)
		if tmp != nil {
			return tmp
		}
		return root
	}
	return PreviouseNode(root.left, cur)
}

func NextNode(root *Node, cur Pair) *Node {
	if root == nil {
		return nil
	}
	if cur.Less(root.item) {
		// root is an candidate
		tmp := NextNode(root.left, cur)
		if tmp != nil {
			return tmp
		}
		return root
	}

	return NextNode(root.right, cur)
}
