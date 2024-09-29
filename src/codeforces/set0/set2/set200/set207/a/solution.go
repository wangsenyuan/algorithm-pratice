package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n := readNum(reader)
	scientists := make([][]int, n)
	for i := 0; i < n; i++ {
		scientists[i] = readNNums(reader, 5)
	}

	bad, res := solve(n, scientists)

	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n", bad))
	if len(res) > 0 {
		for _, x := range res {
			buf.WriteString(fmt.Sprintf("%d %d\n", x[0], x[1]))
		}
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

const inf = 1 << 60

func solve(n int, scientists [][]int) (int, [][]int32) {

	var res [][]int32
	var bad int32

	for j, cur := range scientists {
		var cnt int32
		k, a, x, y, m := cur[0], int32(cur[1]), int32(cur[2]), int32(cur[3]), int32(cur[4])
		for i := 0; i < k; i++ {
			if len(res) < 200003 {
				res = append(res, []int32{cnt, a, int32(j + 1)})
			}
			tmp := (int32(int(x)*int(a)%int(m)) + y) % m
			if i+1 < k && a > tmp {
				cnt++
			}
			a = tmp
		}
		bad = max(bad, cnt)
	}

	if len(res) > 200000 {
		res = nil
	}
	if len(res) > 0 {
		slices.SortFunc(res, func(a, b []int32) int {
			if a[0] != b[0] {
				return int(a[0] - b[0])
			}
			if a[1] != b[1] {
				return int(a[1] - b[1])
			}
			return int(a[2] - b[2])
		})
		for i := 0; i < len(res); i++ {
			res[i] = res[i][1:]
		}
	}

	return int(bad), res
}

func solve1(n int, scientists [][]int) (int, [][]int32) {

	var res [][]int32
	var bad int

	var prev = -inf

	pos := make([]int, n)

	var tr *Node

	for i, cur := range scientists {
		tr = Insert(tr, Pair{cur[1], i})
	}

	for tr != nil {
		it := FindEqualOrGreater(tr, Pair{prev, -1})

		if it == nil {
			// 那就找到最小的那个
			it = FindEqualOrGreater(tr, Pair{-inf, -1})
			bad++
		}

		tr = Delete(tr, it.item)

		i := it.item.second
		prev = it.item.first
		res = append(res, []int32{int32(prev), int32(i + 1)})
		pos[i]++
		if pos[i] < scientists[i][0] {
			val := (prev*scientists[i][2] + scientists[i][3]) % scientists[i][4]
			tr = Insert(tr, Pair{val, i})
		}
	}

	if len(res) > 200000 {
		res = nil
	}

	return bad, res
}

type Pair struct {
	first, second int
}

func (this Pair) Less(that Pair) bool {
	if this.first < that.first {
		return true
	}
	if this.first == that.first && (this.second < that.second) {
		return true
	}
	return false
}

/**
* this is a AVL tree
 */
type Node struct {
	item        Pair
	height      int32
	cnt         int32
	left, right *Node
}

func (node *Node) Height() int32 {
	if node == nil {
		return 0
	}
	return node.height
}

func NewNode(item Pair) *Node {
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

func (node *Node) GetBalance() int32 {
	if node == nil {
		return 0
	}
	return node.left.Height() - node.right.Height()
}

func FindEqualOrGreater(node *Node, item Pair) *Node {
	if node == nil {
		return nil
	}
	if node.item == item {
		return node
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

func FindEqualOrLess(node *Node, item Pair) *Node {
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

func Insert(node *Node, item Pair) *Node {
	if node == nil {
		return NewNode(item)
	}

	if node.item == item {
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

func MaxValueNode(root *Node) *Node {
	cur := root

	for cur.right != nil {
		cur = cur.right
	}
	return cur
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
			tmp.cnt = 1

			// make sure tmp node deleted after call delete on root.right
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
