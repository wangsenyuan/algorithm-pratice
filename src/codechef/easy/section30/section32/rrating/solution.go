package main

import (
	"bufio"
	"bytes"
	"container/heap"
	"fmt"
	"os"
)

func main() {
	// hint(105)
	reader := bufio.NewReader(os.Stdin)

	var buf bytes.Buffer
	n := readNum(reader)

	solver := NewSolver(n)

	for i := 0; i < n; i++ {
		s, _ := reader.ReadBytes('\n')
		if s[0] == '1' {
			var x int
			readInt(s, 2, &x)
			solver.Add(x)
		} else {
			// 2
			ans := solver.Report()
			if ans < 0 {
				buf.WriteString(fmt.Sprintln("No reviews yet"))
			} else {
				buf.WriteString(fmt.Sprintln(fmt.Sprintf("%d", ans)))
			}
		}
	}

	fmt.Print(buf.String())
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			return s[:i]
		}
	}
	return s
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

// An IntHeap is a min-heap of ints.
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type Solver struct {
	heap1 *MinHeap
	heap2 *MaxHeap
	sz    int
}

func NewSolver(n int) *Solver {
	heap1 := make(MinHeap, 0, n/3)
	heap2 := make(MaxHeap, 0, n*2/3)
	return &Solver{&heap1, &heap2, 0}
}

func (solver *Solver) Add(x int) {
	solver.sz++
	p := solver.sz / 3

	if solver.heap1.Len() == 0 || (*solver.heap1)[0] <= x {
		heap.Push(solver.heap1, x)

		for solver.heap1.Len() > p {
			y := heap.Pop(solver.heap1).(int)
			heap.Push(solver.heap2, y)
		}
	} else {
		heap.Push(solver.heap2, x)

		for solver.heap1.Len() < p {
			y := heap.Pop(solver.heap2).(int)
			heap.Push(solver.heap1, y)
		}
	}
}

func (solver *Solver) Report() int {
	if solver.heap1.Len() == 0 {
		return -1
	}
	return (*solver.heap1)[0]
}

type Solver1 struct {
	root *Node
}

func NewSolver1() *Solver1 {
	var root *Node
	return &Solver1{root}
}

func (solver *Solver1) Add(x int) {
	solver.root = Insert(solver.root, Item(x))
}

func (solver *Solver1) Report() int {
	sz := solver.root.size

	id := sz / 3

	if id == 0 {
		return -1
	}

	node := FindKth(solver.root, id)

	return int(node.item)
}

type Item int

func (this Item) Less(that Item) bool {
	return this > that
}

/**
* this is a AVL tree
 */
type Node struct {
	item        Item
	height      int
	size        int
	left, right *Node
}

func (node *Node) Height() int {
	if node == nil {
		return 0
	}
	return node.height
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

func NewNode(item Item) *Node {
	node := new(Node)
	node.item = item
	node.height = 1
	node.size = 1
	return node
}

func rightRotate(y *Node) *Node {
	x := y.left

	t2 := x.right

	x.right = y
	y.left = t2
	y.height = max(y.left.Height(), y.right.Height()) + 1
	y.size = y.left.Size() + y.right.Size() + 1
	x.height = max(x.left.Height(), x.right.Height()) + 1
	x.size = x.left.Size() + x.right.Size() + 1

	return x
}

func leftRotate(x *Node) *Node {
	y := x.right
	t2 := y.left

	y.left = x
	x.right = t2

	x.height = max(x.left.Height(), x.right.Height()) + 1
	x.size = x.left.Size() + x.right.Size() + 1
	y.height = max(y.left.Height(), y.right.Height()) + 1
	y.size = y.left.Size() + y.right.Size() + 1

	return y
}

func (node *Node) GetBalance() int {
	if node == nil {
		return 0
	}
	return node.left.Height() - node.right.Height()
}

func Insert(node *Node, item Item) *Node {
	if node == nil {
		return NewNode(item)
	}
	if node.item == item {
		return node
	}
	if item.Less(node.item) {
		node.left = Insert(node.left, item)
	} else {
		node.right = Insert(node.right, item)
	}

	node.height = max(node.left.Height(), node.right.Height()) + 1
	node.size = node.left.Size() + node.right.Size() + 1
	balance := node.GetBalance()

	if balance > 1 && item.Less(node.left.item) {
		return rightRotate(node)
	}

	if balance < -1 && node.right.item.Less(item) {
		return leftRotate(node)
	}

	if balance > 1 && node.left.item.Less(item) {
		node.left = leftRotate(node.left)
		return rightRotate(node)
	}

	if balance < -1 && item.Less(node.right.item) {
		node.right = rightRotate(node.right)
		return leftRotate(node)
	}

	return node
}

func Delete(root *Node, item Item) *Node {
	if root == nil {
		return nil
	}

	if item.Less(root.item) {
		root.left = Delete(root.left, item)
	} else if root.item.Less(item) {
		root.right = Delete(root.right, item)
	} else {
		if root.left == nil || root.right == nil {
			tmp := root.left
			if root.left == nil {
				tmp = root.right
			}
			root = tmp
		} else {
			tmp := MinValueNode(root.right)

			root.item = tmp.item
			// make sure tmp node deleted after call delete on root.right
			root.right = Delete(root.right, tmp.item)
		}
	}

	if root == nil {
		return root
	}

	root.height = max(root.left.Height(), root.right.Height()) + 1
	root.size = root.left.Size() + root.right.Size() + 1

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

func PreviouseNode(root *Node, cur Item) *Node {
	// find previouse node, that has item < current
	if root == nil {
		return nil
	}

	if root.item.Less(cur) {
		tmp := PreviouseNode(root.right, cur)
		if tmp != nil {
			return tmp
		}
		return root
	}
	return PreviouseNode(root.left, cur)
}

func NextNode(root *Node, cur Item) *Node {
	if root == nil {
		return nil
	}
	if cur.Less(root.item) {
		// root is an candidate
		tmp := NextNode(root.left, cur)
		if tmp != nil {
			return tmp
		}
		return root
	}

	return NextNode(root.right, cur)
}

func FindKth(node *Node, k int) *Node {
	if node.left.Size() >= k {
		return FindKth(node.left, k)
	}
	if node.left.Size()+1 == k {
		return node
	}
	return FindKth(node.right, k-node.left.Size()-1)
}
