package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
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

func process(reader *bufio.Reader) []int {
	n, w := readTwoNums(reader)
	table := make([][]int32, n)
	for i := range n {
		s, _ := reader.ReadBytes('\n')
		var m int
		pos := readInt(s, 0, &m) + 1
		table[i] = make([]int32, m)
		for j := 0; j < m; j++ {
			var v int
			pos = readInt(s, pos, &v) + 1
			table[i][j] = int32(v)
		}
	}
	return solve(w, table)
}

func solve(w int, table [][]int32) []int {
	ans := make([]int, w+1)

	for _, row := range table {
		m := len(row)
		if 2*m <= w {
			var pref, suf int32
			for i, v := range row {
				pref = max(pref, v)
				ans[i] += int(pref)
				ans[i+1] -= int(pref)
				suf = max(suf, row[m-i-1])
				ans[w-1-i] += int(suf)
				ans[w-i] -= int(suf)
			}
			ans[m] += int(suf)
			ans[w-m] -= int(suf)
		} else {
			sz := w - m + 1
			var q []int
			for i, v := range row {
				for len(q) > 0 && v >= row[q[len(q)-1]] {
					q = q[:len(q)-1]
				}
				q = append(q, i)
				if q[0] <= i-sz {
					q = q[1:]
				}
				x := row[q[0]]
				if x < 0 && i < w-m {
					x = 0
				}
				ans[i] += int(x)
				ans[i+1] -= int(x)
			}
			var suf int32
			for i := 0; i < w-m; i++ {
				suf = max(suf, row[m-1-i])
				ans[w-1-i] += int(suf)
				ans[w-i] -= int(suf)
			}
		}
	}

	for i := 1; i <= w; i++ {
		ans[i] += ans[i-1]
	}

	return ans[:w]
}

func solve1(w int, table [][]int32) []int {
	n := len(table)

	type pair struct {
		first  int32
		second int32
	}

	add := make([][]pair, w)
	rem := make([][]pair, w)

	for i := 0; i < n; i++ {
		row := table[i]
		m := len(row)
		for j := 0; j < m; j++ {
			add[j] = append(add[j], pair{int32(row[j]), int32(i)})
			rem[w-(m-j)] = append(rem[w-(m-j)], pair{int32(row[j]), int32(i)})
		}
		if m < w {
			add[0] = append(add[0], pair{0, int32(i)})
			rem[w-(m+1)] = append(rem[w-(m+1)], pair{0, int32(i)})
			add[m] = append(add[m], pair{0, int32(i)})
			rem[w-1] = append(rem[w-1], pair{0, int32(i)})
		}
	}

	nodes := make([]*Node, n)

	ans := make([]int, w)

	for i := 0; i < w; i++ {
		for _, cur := range add[i] {
			idx := cur.second
			val := cur.first
			if nodes[idx] != nil {
				ans[i] -= int(MaxValueNode(nodes[idx]).key)
			}
			nodes[idx] = Insert(nodes[idx], val)
			ans[i] += int(MaxValueNode(nodes[idx]).key)
		}
		if i+1 < w {
			ans[i+1] = ans[i]
			for _, cur := range rem[i] {
				idx := cur.second
				val := cur.first
				ans[i+1] -= int(MaxValueNode(nodes[idx]).key)
				nodes[idx] = Delete(nodes[idx], val)
				if nodes[idx] != nil {
					ans[i+1] += int(MaxValueNode(nodes[idx]).key)
				}
			}
		}
	}

	return ans
}

/**
* this is a AVL tree
 */
type Node struct {
	key         int32
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

func NewNode(key int32) *Node {
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

func (node *Node) GetBalance() int32 {
	if node == nil {
		return 0
	}
	return node.left.Height() - node.right.Height()
}

func Insert(node *Node, key int32) *Node {
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

func MaxValueNode(root *Node) *Node {
	cur := root

	for cur.right != nil {
		cur = cur.right
	}

	return cur
}

func (root *Node) MinValue() int32 {
	if root == nil {
		return 0
	}
	node := MinValueNode(root)

	return node.key
}

func Delete(root *Node, key int32) *Node {
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
