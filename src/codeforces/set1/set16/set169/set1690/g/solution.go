package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var buf bytes.Buffer

	tc := readNum(reader)

	for tc > 0 {
		tc--
		var s []byte
		for {
			s, _ = reader.ReadBytes('\n')
			if len(s) == 0 || s[0] == '\r' || s[0] == '\n' {
				continue
			}
			var n, m int
			pos := readInt(s, 0, &n)
			readInt(s, pos+1, &m)
			A := readNNums(reader, n)
			Q := make([][]int, m)
			for i := 0; i < m; i++ {
				Q[i] = readNNums(reader, 2)
			}
			res := solve(A, Q)

			for i := 0; i < m; i++ {
				buf.WriteString(fmt.Sprintf("%d ", res[i]))
			}
			buf.WriteByte('\n')
			break
		}
	}

	fmt.Print(buf.String())
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
}

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
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

func solve(A []int, Q [][]int) []int {
	// 考虑将a[i] - k
	// 如果左边的最小值a[j] <= a[i] - k, 那么i 只能更在a[j]的后面， 同时(j...i)已经跟在了j的后面
	// 同时i右边，原来跟在它后面的继续跟在它后面
	// 同时，不跟在它后面的（但紧随它的）>= a[i] - k 的， 需要调整跟在它后面
	// 直到遇到一个 a[k] < a[i] - k 的
	// 当然还存在一种情况 a[i] - k 的时候，出现一辆新的train
	// 这里要两棵树
	// 一棵树，表示i跟随的train 的 head的编号
	// 当a[i] - k 产生一辆新的train时，需要更新[i...k)
	// 还需要能快速的找到k

	var set *Node

	n := len(A)

	for i := 0; i < n; i++ {
		if set.Size() == 0 || A[set.MaxValue()] > A[i] {
			set = Insert(set, i)
		}
	}
	ans := make([]int, len(Q))
	for i, cur := range Q {
		j, k := cur[0], cur[1]
		j--
		A[j] -= k
		it := set.LessOrEqual(j)
		// 要不要把j放入set呢？
		if it.key != j && A[it.key] > A[j] {
			set = Insert(set, j)
		}

		for {
			it = set.UpperBound(j)
			if it == nil || A[it.key] < A[j] {
				break
			}
			set = Delete(set, it.key)
		}

		ans[i] = set.Size()
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
	node.sz = 1
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

func (node *Node) update() {
	node.height = max(node.left.Height(), node.right.Height()) + 1
	node.sz = node.left.Size() + node.right.Size() + node.cnt
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

func (node *Node) UpperBound(k int) *Node {
	if node == nil {
		return nil
	}
	if node.key > k {
		tmp := node.left.UpperBound(k)
		if tmp != nil {
			return tmp
		}
		return node
	}
	return node.right.UpperBound(k)
}

func (node *Node) LessOrEqual(k int) *Node {
	if node == nil {
		return nil
	}
	if node.key <= k {
		tmp := node.right.LessOrEqual(k)
		if tmp != nil {
			return tmp
		}
		return node
	}
	return node.left.LessOrEqual(k)
}
