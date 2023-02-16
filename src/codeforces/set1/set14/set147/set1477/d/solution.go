package main

import (
	"bufio"
	"bytes"
	"container/heap"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n, m := readTwoNums(reader)
		L, R := make([]int, m), make([]int, m)
		for i := 0; i < m; i++ {
			L[i], R[i] = readTwoNums(reader)
		}
		p, q := solve(n, m, L, R)
		for i := 0; i < n; i++ {
			buf.WriteString(fmt.Sprintf("%d ", p[i]))
		}
		buf.WriteByte('\n')
		for i := 0; i < n; i++ {
			buf.WriteString(fmt.Sprintf("%d ", q[i]))
		}
		buf.WriteByte('\n')
	}

	fmt.Print(buf.String())
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

func solve(n, m int, L, R []int) ([]int, []int) {
	g := make([]map[int]bool, n)
	t := make([]map[int]bool, n)
	for i := 0; i < n; i++ {
		g[i] = make(map[int]bool)
		t[i] = make(map[int]bool)
	}

	for i := 0; i < m; i++ {
		g[L[i]-1][R[i]-1] = true
		g[R[i]-1][L[i]-1] = true
	}
	var rv *Node
	for i := 0; i < n; i++ {
		rv = Insert(rv, i)
	}

	var dfs func(u int)

	dfs = func(u int) {
		rv = Delete(rv, u)

		var v = -1
		for {
			tmp := FindUpperBound(rv, v)
			if tmp == nil {
				break
			}
			v = tmp.key
			if g[u][v] {
				continue
			}
			t[u][v] = true
			t[v][u] = true
			dfs(v)
		}
	}

	for rv.Size() > 0 {
		dfs(rv.MinValue())
	}

	rem := n

	d := make(PriorityQueue, 0, n)
	perm1 := make([]int, n)
	perm2 := make([]int, n)
	deg := make([]int, n)
	for i := 0; i < n; i++ {
		deg[i] = len(t[i])
		if deg[i] > 0 {
			heap.Push(&d, NewItem(i, deg[i]))
		} else {
			perm1[i] = rem
			perm2[i] = rem
			rem--
		}
	}

	stars := make([]Star, 0, n)
	for d.Len() > 0 {
		idx := d[0].value
		if d[0].priority != deg[idx] {
			heap.Pop(&d)
			continue
		}
		f := getFirst(t[idx])
		deg[f] = 0
		var leaves []int
		for c := range t[f] {
			deg[c]--
			if deg[c] == 0 {
				leaves = append(leaves, c)
			} else {
				heap.Push(&d, NewItem(c, deg[c]))
				delete(t[c], f)
			}
		}
		stars = append(stars, NewStar(f, leaves))
	}

	var l, r int
	for _, c := range stars {
		l++
		perm1[c.center] = l
		for _, x := range c.leaves {
			l++
			perm1[x] = l
			r++
			perm2[x] = r
		}
		r++
		perm2[c.center] = r
	}

	return perm1, perm2
}

type Star struct {
	center int
	leaves []int
}

func NewStar(c int, l []int) Star {
	return Star{c, l}
}

func NewItem(v, p int) *Item {
	item := new(Item)
	item.value = v
	item.priority = p
	return item
}

func getFirst(mem map[int]bool) int {
	for k := range mem {
		return k
	}
	return -1
}

// An Item is something we manage in a priority queue.
type Item struct {
	value    int // The value of the item; arbitrary.
	priority int // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, priority int) {
	item.priority = priority
	heap.Fix(pq, item.index)
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

func (node *Node) Size() int {
	if node == nil {
		return 0
	}
	return node.cnt
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

func FindUpperBound(node *Node, key int) *Node {
	if node == nil {
		return nil
	}
	if node.key > key {
		tmp := FindUpperBound(node.left, key)
		if tmp == nil {
			return node
		}
		return tmp
	}
	// node.key <= key
	return FindUpperBound(node.right, key)
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
