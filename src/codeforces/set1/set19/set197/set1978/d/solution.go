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

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n, c := readTwoNums(reader)
		a := readNNums(reader, n)

		res := solve(c, a)

		for i := 0; i < n; i++ {
			buf.WriteString(fmt.Sprintf("%d ", res[i]))
		}
		buf.WriteByte('\n')
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

func solve(c int, a []int) []int {
	n := len(a)
	if n == 1 {
		return []int{0}
	}

	var r int
	for i, num := range a {
		if num >= a[r] {
			r = i
		}
	}

	var sum int
	ans := make([]int, n)

	var l int
	for i, num := range a {
		ans[i] = n - 1
		if i == 0 {
			if i == r || a[r] <= num+c {
				ans[i] = 0
			}
		} else {
			if a[l] < num && (r < i || a[r] <= num) && a[0]+c < num {
				ans[i] = 0
			}
		}
		tmp := i
		if r > i && a[r] > num+sum+c {
			tmp++
		}
		ans[i] = min(ans[i], tmp)
		if num > a[l] {
			l = i
		}
		sum += num
	}

	return ans
}
func solve2(c int, a []int) []int {
	n := len(a)
	if n == 1 {
		return []int{0}
	}

	items := make([]*Item, n)
	suf := make(PriorityQueue, n)

	for i := 0; i < n; i++ {
		it := new(Item)
		it.id = i
		it.priority = a[i]
		it.index = i
		items[i] = it
		suf[i] = it
	}

	heap.Init(&suf)

	pop := func(i int) {
		suf.update(items[i], 1<<30)
		heap.Pop(&suf)
	}

	ans := make([]int, n)

	pref := make(PriorityQueue, 0, n)
	var sum int

	for i, num := range a {
		pop(i)
		ans[i] = n - 1
		if i == 0 {
			if suf[0].priority <= num+c {
				ans[i] = 0
			}
		} else {
			// i > 0
			if pref[0].priority < num && a[0]+c < num && (suf.Len() == 0 || suf[0].priority <= num) {
				ans[i] = 0
			}
		}
		// 前面的都删掉
		tmp := i
		if suf.Len() > 0 && suf[0].priority > num+sum+c {
			tmp++
		}
		ans[i] = min(ans[i], tmp)
		sum += num
		items[i].priority = a[i]
		heap.Push(&pref, items[i])
	}

	return ans
}

// An Item is something we manage in a priority queue.
type Item struct {
	id       int // The value of the item; arbitrary.
	priority int // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority
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
	// item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

func solve1(c int, a []int) []int {
	n := len(a)

	var suf, pref *Node

	for _, num := range a {
		suf = Insert(suf, num)
	}

	ans := make([]int, n)

	var sum int

	for i, num := range a {
		suf = Delete(suf, num)
		// 其他人都删掉
		ans[i] = n - 1
		// 如果前面有剩余，所有a[j] + c >= a[i]的都应该退选
		// 如果a[i]是suf最大的
		if i == 0 {
			// 无人退选，且他是0号，所以c都支持他
			cnt := CountLowerBound(suf, num+c+1)
			if cnt == 0 {
				ans[i] = 0
			}
		} else {
			// 无人退选的情况下
			cnt := CountLowerBound(pref, num)
			cnt2 := CountLowerBound(suf, num+1)
			if cnt == 0 && cnt2 == 0 && a[0]+c < num {
				// 无人退选，c都去支持0号了， 且没有人比i更受欢迎
				ans[i] = 0
			}
		}
		// 前面的都退选了
		cnt := CountLowerBound(suf, num+sum+c+1)
		// 如果要删除，只删除那个最受欢迎的即可
		ans[i] = min(ans[i], i+min(1, cnt))
		pref = Insert(pref, num)
		sum += num
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
	left, right *Node
}

func NewNode(key int) *Node {
	node := new(Node)
	node.key = key
	node.height = 1
	node.cnt = 1
	node.sz = 1
	return node
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
	return node.sz
}

func (node *Node) update() {
	node.height = max(node.left.Height(), node.right.Height()) + 1
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

func CountLowerBound(node *Node, key int) int {
	if node == nil {
		return 0
	}
	if node.key < key {
		return CountLowerBound(node.right, key)
	}
	// node.key >= key
	return node.right.Size() + node.cnt + CountLowerBound(node.left, key)
}
