package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	T := readNNums(reader, n)
	X := readNNums(reader, n)
	Y := readNNums(reader, n)
	A := readNNums(reader, n)
	res := solve(n, T, X, Y, A)
	fmt.Println(res)
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

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
}

const N = 8005
const N1 = N / 2

func solve(n int, T []int, X []int, Y []int, A []int) int64 {
	V := make([][]int, N1)

	for i := 0; i < n; i++ {
		if len(V[X[i]]) == 0 {
			V[X[i]] = make([]int, 0, 10)
		}
		V[X[i]] = append(V[X[i]], i)
	}

	best := make([][]int, N)

	for i := 0; i < N; i++ {
		best[i] = make([]int, N1)
	}

	seg := make([]Car, 2*n)

	for i := 1; i < N1; i++ {
		if len(V[i]) == 0 {
			continue
		}
		var k int
		for _, p := range V[i] {
			seg[k] = Car{T[p], 1, A[p]}
			k++
			seg[k] = Car{T[p] + Y[p] + 1, -1, A[p]}
			k++
		}

		sort.Sort(Cars(seg[:k]))
		var root *Node
		var cur int
		for j := 1; j < N; j++ {
			for cur < k && seg[cur].first == j {
				tp := seg[cur].second
				if tp == 1 {
					root = Insert(root, seg[cur].tip)
				} else {
					root = Delete(root, seg[cur].tip)
				}
				cur++
			}
			if root != nil {
				best[j][i] = MaxValueNode(root).key
			}
		}
	}
	dp := make([]int64, N)
	for i := 0; i < N; i++ {
		dp[i] = -1
	}
	var dfs func(pos int) int64
	dfs = func(pos int) int64 {
		if pos >= N {
			return 0
		}
		if dp[pos] >= 0 {
			return dp[pos]
		}
		dp[pos] = 0

		for j := 1; j < N1; j++ {
			if best[pos][j] > 0 {
				dp[pos] = max2(dp[pos], int64(best[pos][j])+dfs(pos+j))
			}
		}

		dp[pos] = max2(dp[pos], dfs(pos+1))

		return dp[pos]
	}

	return dfs(1)
}

func max2(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}

type Car struct {
	first, second int
	tip           int
}

type Cars []Car

func (ps Cars) Len() int {
	return len(ps)
}

func (ps Cars) Less(i, j int) bool {
	return ps[i].first < ps[j].first || ps[i].first == ps[j].first && ps[i].second < ps[j].second
}

func (ps Cars) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
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

func Search(node *Node, key int) *Node {
	if node == nil || node.key == key {
		return node
	}
	if key < node.key {
		return Search(node.left, key)
	}
	return Search(node.right, key)
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

func MaxValueNode(root *Node) *Node {
	cur := root

	for cur.right != nil {
		cur = cur.right
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
