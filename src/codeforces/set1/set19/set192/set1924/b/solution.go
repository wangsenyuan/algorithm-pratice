package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m, k := readThreeNums(reader)
	x := readNNums(reader, m)
	v := readNNums(reader, m)
	queries := make([][]int, k)
	for i := 0; i < k; i++ {
		queries[i] = readNNums(reader, 3)
	}

	res := solve(n, x, v, queries)

	var buf bytes.Buffer

	for _, ans := range res {
		buf.WriteString(fmt.Sprintf("%d\n", ans))
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

func solve(n int, X []int, V []int, queies [][]int) []int {
	val := make([]int, n+1)

	var harbors *Node

	for i, x := range X {
		x--
		harbors = Insert(harbors, x)
		val[x] = V[i]
	}

	ships := Build(n)

	// init
	for i := 0; i < n; i++ {
		x := LowerBound(harbors, i)
		// x 肯定不为nil
		if x.key == i {
			continue
		}
		// x.key = X[i+1]
		y := Less(harbors, i)
		// y 肯定不为nil
		// i 的贡献 = (x.key - i) * val[y.key]
		ships.UpdateValue1(i, i, (x.key-i)*val[y.key])
	}

	var ans []int

	for _, cur := range queies {
		if cur[0] == 1 {
			x, v := cur[1], cur[2]
			x--
			val[x] = v
			nxt := LowerBound(harbors, x)
			bef := Less(harbors, x)
			// 这个用于更新 bef.key + 1, 到 x
			tmp1 := val[bef.key] * (x - nxt.key)

			if bef.key+1 <= x {
				ships.UpdateValue1(bef.key+1, x, tmp1)
			}
			// 这个用于更新 x + 1 到 nxt.key
			tmp1 = (v - val[bef.key]) * nxt.key
			tmp2 := v - val[bef.key]
			if x+1 <= nxt.key-1 {
				ships.UpdateValue1(x+1, nxt.key-1, tmp1)
				ships.UpdateValue2(x+1, nxt.key-1, tmp2)
			}
			// 增加一个港口
			harbors = Insert(harbors, x)
		} else {
			l, r := cur[1], cur[2]
			l--
			r--
			ans = append(ans, ships.Query(l, r))
		}
	}

	return ans
}

type Tree struct {
	left, right *Tree
	l, r        int
	val         int
	lazy1       int
	lazy2       int
}

func (tr *Tree) pushValue1(v int) {
	tr.val += v * (tr.r - tr.l + 1)
	tr.lazy1 += v
}

func (tr *Tree) pushValue2(v int) {
	tr.val -= v * (tr.l + tr.r) * (tr.r - tr.l + 1) / 2
	tr.lazy2 += v
}

func (tr *Tree) push() {
	if tr.left == nil {
		// a leaf
		return
	}
	if tr.lazy1 != 0 {
		tr.left.pushValue1(tr.lazy1)
		tr.right.pushValue1(tr.lazy1)
		tr.lazy1 = 0
	}

	if tr.lazy2 != 0 {
		tr.left.pushValue2(tr.lazy2)
		tr.right.pushValue2(tr.lazy2)
		tr.lazy2 = 0
	}
}

func Build(n int) *Tree {
	var loop func(l int, r int) *Tree
	loop = func(l int, r int) *Tree {
		tr := new(Tree)
		tr.l = l
		tr.r = r
		if l < r {
			mid := (l + r) / 2
			tr.left = loop(l, mid)
			tr.right = loop(mid+1, r)
		}
		return tr
	}
	return loop(0, n-1)
}

func (tr *Tree) UpdateValue1(L int, R int, v int) {
	if tr.r < L || R < tr.l {
		// out of bounds
		return
	}
	tr.push()
	if L <= tr.l && tr.r <= R {
		tr.pushValue1(v)
		return
	}
	tr.left.UpdateValue1(L, R, v)
	tr.right.UpdateValue1(L, R, v)
	tr.val = tr.left.val + tr.right.val
}

func (tr *Tree) UpdateValue2(L int, R int, v int) {
	if tr.r < L || R < tr.l {
		return
	}
	tr.push()
	if L <= tr.l && tr.r <= R {
		tr.pushValue2(v)
		return
	}
	tr.left.UpdateValue2(L, R, v)
	tr.right.UpdateValue2(L, R, v)
	tr.val = tr.left.val + tr.right.val
}

func (tr *Tree) Query(L int, R int) int {
	if tr.r < L || R < tr.l {
		return 0
	}

	if L <= tr.l && tr.r <= R {
		return tr.val
	}
	tr.push()

	a := tr.left.Query(L, R)
	b := tr.right.Query(L, R)
	return a + b
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

func LowerBound(node *Node, key int) *Node {
	if node == nil {
		return nil
	}
	if node.key >= key {
		tmp := LowerBound(node.left, key)
		if tmp != nil {
			return tmp
		}
		return node
	}
	// node.key > key
	return LowerBound(node.right, key)
}

func Less(node *Node, key int) *Node {
	if node == nil {
		return nil
	}
	if node.key < key {
		// 如果右边还有比key小的，用右边的
		tmp := Less(node.right, key)
		if tmp != nil {
			return tmp
		}
		return node
	}
	// node.key >= key
	// 肯定在左边
	return Less(node.left, key)
}
