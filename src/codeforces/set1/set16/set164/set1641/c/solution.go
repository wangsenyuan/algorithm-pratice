package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m := readTwoNums(reader)

	Q := make([][]int, m)

	for i := 0; i < m; i++ {
		s, _ := reader.ReadBytes('\n')
		var c int
		pos := readInt(s, 0, &c)
		if c == 0 {
			var l, r, x int
			pos = readInt(s, pos+1, &l)
			pos = readInt(s, pos+1, &r)
			readInt(s, pos+1, &x)
			Q[i] = []int{0, l, r, x}
		} else {
			var id int
			readInt(s, pos+1, &id)
			Q[i] = []int{1, id}
		}
	}

	res := solve(n, Q)

	var buf bytes.Buffer

	for _, yn := range res {
		buf.WriteString(yn)
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
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

func solve(n int, Q [][]int) []string {
	// l, r, 1 => at least one person sick between l..r
	// l, r, 0 => no one sick
	// 如果某个人属于no sick seg答案是no
	// 如果某个人属于sick seg, 且只有1个人，
	var unknown *Node

	for i := 0; i < n; i++ {
		unknown = Insert(unknown, i)
	}

	b := NewBIT(n)

	health := make([]bool, n)

	var res []string

	for _, cur := range Q {
		if cur[0] == 0 {
			l, r, x := cur[1], cur[2], cur[3]
			l--
			if x == 0 {
				// health
				for l < r {
					node := Search(unknown, l)
					if node == nil || node.key >= r {
						break
					}
					health[node.key] = true
					l = node.key
					unknown = Delete(unknown, l)
				}
			} else {
				// sick
				b.Set(l, r)
			}
		} else {
			id := cur[1]
			id--
			if health[id] {
				res = append(res, "NO")
			} else {
				// not health, so id must be in unknown
				l, r := 0, n
				it := SearchLess(unknown, id)
				if it != nil {
					l = it.key + 1
				}

				it = SearchGreater(unknown, id)
				// 要找到id后面的那个节点
				if it != nil {
					r = it.key
				}
				y := b.Get(l, n)
				if y <= r {
					// sick
					res = append(res, "YES")
				} else {
					res = append(res, "N/A")
				}
			}
		}
	}

	return res
}

const INF = 1e9

type SegTree struct {
	arr []int
	sz  int
}

func NewBIT(n int) *SegTree {
	arr := make([]int, 2*n)
	for i := 1; i < 2*n; i++ {
		arr[i] = INF
	}
	return &SegTree{arr, n}
}

func (s *SegTree) Set(p int, v int) {
	p += s.sz
	s.arr[p] = min(s.arr[p], v)
	for p > 1 {
		s.arr[p>>1] = min(s.arr[p], s.arr[p^1])
		p >>= 1
	}
}

func (s *SegTree) Get(l, r int) int {
	var res int = INF
	l += s.sz
	r += s.sz
	for l < r {
		if l&1 == 1 {
			res = min(res, s.arr[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = min(res, s.arr[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
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

func Search(node *Node, key int) *Node {
	// find the node where node.key >= key
	if node == nil {
		return nil
	}
	if node.key >= key {
		// a candidate
		tmp := Search(node.left, key)
		if tmp != nil {
			return tmp
		}
		return node
	}
	return Search(node.right, key)
}

func SearchLess(node *Node, key int) *Node {
	if node == nil {
		return nil
	}
	if node.key < key {
		tmp := SearchLess(node.right, key)
		if tmp != nil {
			return tmp
		}
		return node
	}
	// node.key >= key
	return SearchLess(node.left, key)
}

func SearchGreater(node *Node, key int) *Node {
	if node == nil {
		return nil
	}
	if node.key > key {
		tmp := SearchGreater(node.left, key)
		if tmp != nil {
			return tmp
		}
		return node
	}
	return SearchGreater(node.right, key)
}
