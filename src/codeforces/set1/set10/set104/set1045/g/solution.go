package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Println(res)
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

func process(reader *bufio.Reader) int {
	n, k := readTwoNums(reader)
	robots := make([][]int, n)
	for i := range n {
		robots[i] = readNNums(reader, 3)
	}
	return solve(robots, k)
}

type Robot struct {
	x  int
	r  int
	q  int
	id int
}

func solve(robots [][]int, k int) int {
	n := len(robots)
	rs := make([]Robot, n)
	qs := make(map[int]int)
	for i, cur := range robots {
		rs[i] = Robot{cur[0], cur[1], cur[2], i}
		qs[cur[2]]++
	}

	slices.SortFunc(rs, func(a, b Robot) int {
		return b.r - a.r
	})

	arr := make([]int, 0, len(qs))
	for k := range qs {
		arr = append(arr, k)
	}

	sort.Ints(arr)

	for i, q := range arr {
		qs[q] = i
	}

	query := func(root *Node, x int, r int) int {
		res := CountLowerBound(root, x+r)
		res -= CountLowerBound(root, x-r-1)
		return res
	}

	nodes := make([]*Node, len(arr))

	var ans int

	for _, r := range rs {
		// r.r 是最小的，所以，在[x-r, x + r]范围内的，肯定是相互可见的
		// 因为iq比较小， 可以每个iq值一棵树
		// 然后在这棵树上上去找这个范围内的就可以了
		for i := r.q - k; i <= r.q+k; i++ {
			if j, ok := qs[i]; ok {
				ans += query(nodes[j], r.x, r.r)
			}
		}
		nodes[qs[r.q]] = Insert(nodes[qs[r.q]], r.x)
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
		node.sz++
		return node
	}

	if node.key > key {
		node.left = Insert(node.left, key)
	} else {
		node.right = Insert(node.right, key)
	}

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

func CountLowerBound(node *Node, key int) int {
	if node == nil {
		return 0
	}
	// <= key 的数量
	if node.key > key {
		return CountLowerBound(node.left, key)
	}
	// node.key <= key
	return node.left.Size() + node.cnt + CountLowerBound(node.right, key)
}
