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
		n, k := readTwoNums(reader)
		x := readString(reader)[:n]
		Q := make([][]int, k)
		for i := 0; i < k; i++ {
			s, _ := reader.ReadBytes('\n')
			var tp int
			pos := readInt(s, 0, &tp) + 1
			if tp == 1 {
				Q[i] = make([]int, 4)
			} else {
				Q[i] = make([]int, 3)
			}
			Q[i][0] = tp

			for j := 1; j < len(Q[i]); j++ {
				pos = readInt(s, pos, &Q[i][j]) + 1
			}
		}
		res := solve(x, Q)
		for _, b := range res {
			if b {
				buf.WriteString("YES\n")
			} else {
				buf.WriteString("NO\n")
			}
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

func solve(s string, Q [][]int) []bool {
	n := len(s)

	arr := make([]int, n+1)

	add := func(p int, v int) {
		v = (v%26 + 26) % 26
		p++
		for p <= n {
			arr[p] = (arr[p] + v) % 26
			p += p & -p
		}
	}

	get := func(p int) int {
		p++
		var res int
		for p > 0 {
			res += arr[p]
			p -= p & -p
		}
		return res % 26
	}

	var m2 *Node
	var m3 *Node

	for i := 0; i < n; i++ {
		if i > 0 {
			add(i, int(s[i]-'a')-int(s[i-1]-'a'))
		} else {
			add(0, int(s[0]-'a'))
		}
		if i+1 < n && s[i] == s[i+1] {
			m2 = Insert(m2, i)
		}
		if i+2 < n && s[i] == s[i+2] {
			m3 = Insert(m3, i)
		}
	}

	relax := func(l int, r int) {
		l = max(0, l)
		r = min(n-1, r)

		for i := l; i+1 <= r; i++ {
			c1 := get(i)
			c2 := get(i + 1)
			if c1 == c2 {
				m2 = Insert(m2, i)
			} else {
				m2 = Delete(m2, i)
			}
			if i+2 <= r {
				c3 := get(i + 2)
				if c1 == c3 {
					m3 = Insert(m3, i)
				} else {
					m3 = Delete(m3, i)
				}
			}
		}
	}

	update := func(l int, r int, x int) {
		add(l, x)
		relax(l-5, l+5)
		add(r, 26-x)
		relax(r-5, r+5)
	}

	query := func(l int, r int) bool {
		L := LowerBound(m2, l)
		if L != nil && L.key+1 < r {
			return false
		}
		L = LowerBound(m3, l)
		if L != nil && L.key+2 < r {
			return false
		}
		return true
	}

	var res []bool

	for _, cur := range Q {
		if cur[0] == 1 {
			// update
			l, r, x := cur[1], cur[2], cur[3]
			l--
			update(l, r, x%26)
		} else {
			l, r := cur[1], cur[2]
			l--
			res = append(res, query(l, r))
		}
	}

	return res
}

/**
* this is a AVL tree
 */
type Node struct {
	key         int
	height      int
	cnt         int
	sum         int64
	sz          int
	left, right *Node
}

func (node *Node) Height() int {
	if node == nil {
		return 0
	}
	return node.height
}

func (node *Node) Sum() int64 {
	if node == nil {
		return 0
	}
	return node.sum
}

func (node *Node) Size() int {
	if node == nil {
		return 0
	}
	return node.sz
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func NewNode(key int) *Node {
	node := new(Node)
	node.key = key
	node.height = 1
	node.cnt = 1
	node.sum = int64(key)
	node.sz = 1
	return node
}

func (node *Node) update() {
	node.height = max(node.left.Height(), node.right.Height()) + 1
	node.sum = node.left.Sum() + node.right.Sum() + int64(node.key)*int64(node.cnt)
	node.sz = node.left.Size() + node.right.Size() + node.cnt
}

func rightRotate(y *Node) *Node {
	x := y.left

	t2 := x.right

	x.right = y
	y.left = t2
	y.update()
	x.update()

	return x
}

func leftRotate(x *Node) *Node {
	y := x.right
	t2 := y.left

	y.left = x
	x.right = t2

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

func Insert(node *Node, key int) *Node {
	if node == nil {
		return NewNode(key)
	}
	if node.key == key {
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
		root.sum -= int64(key)
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
			//root.sum = tmp.key
			// make sure tmp node deleted after call delete on root.right
			tmp.cnt = 1
			root.right = Delete(root.right, tmp.key)
		}
	}

	if root == nil {
		return root
	}

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

func LowerBound(node *Node, key int) *Node {
	if node == nil {
		return nil
	}
	if node.key >= key {
		res := LowerBound(node.left, key)
		if res != nil {
			return res
		}
		return node
	}
	return LowerBound(node.right, key)
}
