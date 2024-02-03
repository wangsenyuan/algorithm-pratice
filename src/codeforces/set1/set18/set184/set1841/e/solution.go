package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n := readNum(reader)
		a := readNNums(reader, n)
		m := readNum(reader)
		res := solve(a, m)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

func solve(a []int, m int) int {
	n := len(a)
	id := make([]int, n)
	for i := 0; i < n; i++ {
		id[i] = i
	}
	sort.Slice(id, func(i, j int) bool {
		return a[id[i]] > a[id[j]]
	})

	var used *Node
	used = Insert(used, Key{0, n, n})

	cnt := make([]int, n+1)

	for i, j := n, 0; i >= 0; i-- {
		for j < n && a[id[j]] >= i {
			it := LowerBound(used, Key{id[j], -1, 0})
			if it != nil {
				l, r, v := it.key.l, it.key.r, it.key.v
				cnt[r-l] += v - i
				used = Delete(used, it.key)
				if l < id[j] {
					used = Insert(used, Key{l, id[j], i})
				}
				if r > id[j]+1 {
					used = Insert(used, Key{id[j] + 1, r, i})
				}
			}
			j++
		}
	}

	var ans int

	for i := n; i > 0 && m > 0; i-- {
		tmp := min(cnt[i], m/i)
		ans += tmp * (i - 1)
		m -= tmp * i
		if tmp != cnt[i] && m > 0 {
			ans += m - 1
			m = 0
		}
	}

	return ans
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

type Key struct {
	l int
	r int
	v int
}

func (this Key) Less(that Key) bool {
	return this.l < that.l
}

/**
* this is a AVL tree
 */
type Node struct {
	key         Key
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

func NewNode(key Key) *Node {
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

func Insert(node *Node, key Key) *Node {
	if node == nil {
		return NewNode(key)
	}
	if node.key.l == key.l {
		node.key = key
		return node
	}

	if key.l < node.key.l {
		node.left = Insert(node.left, key)
	} else {
		node.right = Insert(node.right, key)
	}

	node.height = max(node.left.Height(), node.right.Height()) + 1
	balance := node.GetBalance()

	if balance > 1 && key.l < node.left.key.l {
		return rightRotate(node)
	}

	if balance < -1 && node.right.key.l < key.l {
		return leftRotate(node)
	}

	if balance > 1 && node.left.key.l < key.l {
		node.left = leftRotate(node.left)
		return rightRotate(node)
	}

	if balance < -1 && key.l < node.right.key.l {
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

func Delete(root *Node, key Key) *Node {
	if root == nil {
		return nil
	}

	if key.l < root.key.l {
		root.left = Delete(root.left, key)
	} else if root.key.l < key.l {
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

func LowerBound(node *Node, key Key) *Node {
	if node == nil {
		return nil
	}

	if node.key.l <= key.l {
		tmp := LowerBound(node.right, key)
		if tmp != nil {
			return tmp
		}
		return node
	}
	// node.key > key
	return LowerBound(node.left, key)
}
