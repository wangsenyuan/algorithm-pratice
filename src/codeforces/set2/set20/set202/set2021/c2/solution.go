package main

import (
	"bufio"
	"bytes"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer
	tc := readNum(reader)
	for range tc {
		res := process(reader)
		for _, x := range res {
			if x {
				buf.WriteString("YA")
			} else {
				buf.WriteString("TIDAK")
			}
			buf.WriteByte('\n')
		}
	}
	buf.WriteTo(os.Stdout)
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

func process(reader *bufio.Reader) []bool {
	n, m, k := readThreeNums(reader)
	a := readNNums(reader, n)
	b := readNNums(reader, m)
	updates := make([][]int, k)
	for i := range k {
		updates[i] = readNNums(reader, 2)
	}
	return solve(a, b, updates)
}

func solve(a []int, b []int, updates [][]int) []bool {
	// reorder a to be 1.2.3..
	n := len(a)
	m := len(b)
	pos := make([]int, n+1)
	for i, x := range a {
		pos[x] = i
	}
	// pos[b[i]]
	trs := make([]*Node, n)

	for i := range n {
		trs[i] = Insert(trs[i], m)
	}

	for i, x := range b {
		j := pos[x]
		b[i] = j
		trs[j] = Insert(trs[j], i)
	}

	var count int
	for i := range n - 1 {
		u := MinValueNode(trs[i]).key
		v := MinValueNode(trs[i+1]).key
		if u <= v {
			count++
		}
	}

	updateCount := func(i int, d int) {
		if i >= 0 && i+1 < n {
			u := MinValueNode(trs[i]).key
			v := MinValueNode(trs[i+1]).key
			if u <= v {
				count += d
			}
		}
	}

	remove := func(i int, j int) {
		// 先更新count
		updateCount(i-1, -1)
		updateCount(i, -1)
		trs[i] = Delete(trs[i], j)
		updateCount(i-1, 1)
		updateCount(i, 1)
	}

	add := func(i int, j int) {
		updateCount(i-1, -1)
		updateCount(i, -1)
		trs[i] = Insert(trs[i], j)
		updateCount(i-1, 1)
		updateCount(i, 1)
	}

	update := func(i int, j int) {
		// b[i] = j
		old_j := b[i]
		if old_j == j {
			return
		}
		// 从old_j上删除掉i
		remove(old_j, i)
		// 在j上增加i
		add(j, i)
		b[i] = j
	}

	// good = count == n - 1
	ans := make([]bool, len(updates)+1)
	ans[0] = count == n-1

	for i, cur := range updates {
		update(cur[0]-1, pos[cur[1]])
		ans[i+1] = count == n-1
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
	cur := root

	for cur.left != nil {
		cur = cur.left
	}

	return cur
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
