package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
)

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

func main() {
	reader := bufio.NewReader(os.Stdin)

	var buf bytes.Buffer

	tc := 1

	for tc > 0 {
		tc--
		n := readNum(reader)
		a := readNNums(reader, n)
		m := readNum(reader)
		queries := make([][]int, m)
		for i := range m {
			queries[i] = readNNums(reader, 2)
		}
		res := solve(a, queries)
		for _, x := range res {
			buf.WriteString(fmt.Sprintf("%d\n", x))
		}
	}

	fmt.Print(buf.String())
}

func solve(a []int, queries [][]int) []int {
	n := len(a)

	type pair struct {
		first  int
		second int
	}

	nums := make([]pair, n)
	for i := range n {
		nums[i] = pair{a[i], i}
	}

	slices.SortFunc(nums, func(a, b pair) int {
		if a.first != b.first {
			return b.first - a.first
		}
		return a.second - b.second
	})

	at := make([][]int, n+1)

	for i, cur := range queries {
		k := cur[0]
		at[k] = append(at[k], i)
	}

	var tr *Node

	ans := make([]int, len(queries))

	for i := 0; i < n; i++ {
		// 把下标放进去
		tr = Insert(tr, nums[i].second)

		for _, j := range at[i+1] {
			pos := queries[j][1]
			node := KthNode(tr, pos)
			ans[j] = a[node.key]
		}
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
	size        int
	left, right *Node
}

func (node *Node) Height() int {
	if node == nil {
		return 0
	}
	return node.height
}

func (node *Node) Count() int {
	if node == nil {
		return 0
	}
	return node.cnt
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

func NewNode(key int) *Node {
	node := new(Node)
	node.key = key
	node.height = 1
	node.cnt = 1
	node.size = 1
	return node
}

func (node *Node) Update() {
	node.height = max(node.left.Height(), node.right.Height()) + 1
	node.size = node.left.Size() + node.right.Size() + node.cnt
}

func rightRotate(y *Node) *Node {
	x := y.left

	t2 := x.right

	x.right = y
	y.left = t2
	y.Update()
	x.Update()

	return x
}

func leftRotate(x *Node) *Node {
	y := x.right
	t2 := y.left

	y.left = x
	x.right = t2

	x.Update()

	y.Update()

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
		node.size++
		return node
	}

	if node.key > key {
		node.left = Insert(node.left, key)
	} else {
		node.right = Insert(node.right, key)
	}

	node.Update()

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

func KthNode(node *Node, k int) *Node {
	if node == nil || k > node.size {
		return nil
	}
	if node.left.Size() >= k {
		return KthNode(node.left, k)
	}
	// node.left.Size() < k
	if node.left.Size()+node.cnt >= k {
		return node
	}
	return KthNode(node.right, k-node.left.Size()-node.cnt)
}