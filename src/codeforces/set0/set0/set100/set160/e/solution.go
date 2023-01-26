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
	n, m := readTwoNums(reader)
	buses := make([][]int, n)
	for i := 0; i < n; i++ {
		buses[i] = readNNums(reader, 3)
	}
	people := make([][]int, m)
	for i := 0; i < m; i++ {
		people[i] = readNNums(reader, 3)
	}
	res := solve(buses, people)
	var buf bytes.Buffer
	for i := 0; i < m; i++ {
		buf.WriteString(fmt.Sprintf("%d ", res[i]))
	}
	buf.WriteByte('\n')
	fmt.Print(buf.String())
}

func readFloat64(bs []byte, pos int, r *float64) int {
	var first float64
	var second float64
	exp := 1.0
	var dot bool
	for pos < len(bs) && (bs[pos] == '.' || isDigit(bs[pos])) {
		if bs[pos] == '.' {
			dot = true
		} else if !dot {
			first = first*10 + float64(bs[pos]-'0')
		} else {
			second = second*10 + float64(bs[pos]-'0')
			exp *= 10
		}
		pos++
	}
	*r = first + second/exp
	//fmt.Fprintf(os.Stderr, "%.6f ", *r)
	return pos
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
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
		if s[i] == '\n' || s[i] == '\r' {
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

func solve(buses [][]int, people [][]int) []int {
	events := make([]Event, len(buses)+len(people))

	for i := 0; i < len(buses); i++ {
		events[i] = Event{1, i, buses[i][2]}
	}
	for i := 0; i < len(people); i++ {
		events[i+len(buses)] = Event{0, i, people[i][2]}
	}

	sort.Slice(events, func(i, j int) bool {
		return events[i].Less(events[j])
	})

	var persons *Node

	ans := make([]int, len(people))

	for i := 0; i < len(people); i++ {
		ans[i] = -1
	}

	for i := 0; i < len(events); i++ {
		cur := events[i]
		id := cur.id
		if cur.tp == 0 {
			// a person, use r as key, and l as value
			persons = Insert(persons, Item{people[id][1], people[id][0], id})
		} else {
			// a bus
			// persons is awaiting before bus time,
			for {
				// find the node having maximum value with key <= f
				p := Find(persons, buses[id][1])
				if p.id < 0 || p.val < buses[id][0] {
					break
				}
				ans[p.id] = id + 1
				persons = Delete(persons, p)
			}
		}
	}

	return ans
}

type Event struct {
	tp   int
	id   int
	time int
}

func (this Event) Less(that Event) bool {
	return this.time < that.time || this.time == that.time && this.tp < that.tp
}

type Item struct {
	key int
	val int
	id  int
}

func (this Item) Less(that Item) bool {
	return this.key < that.key || this.key == that.key && this.id < that.id
}

func (this Item) LessValue(that Item) bool {
	return this.val < that.val
}

/**
* this is a AVL tree
 */
type Node struct {
	key         Item
	val         Item
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

func NewNode(key Item) *Node {
	node := new(Node)
	node.key = key
	node.val = key
	node.height = 1
	node.cnt = 1
	return node
}

func (node *Node) Value() Item {
	if node == nil {
		return Item{-1, -1, -1}
	}
	return node.val
}

func (node *Node) update() {
	node.height = max(node.left.Height(), node.right.Height()) + 1
	node.val = node.key

	if node.val.LessValue(node.left.Value()) {
		node.val = node.left.Value()
	}
	if node.val.LessValue(node.right.Value()) {
		node.val = node.right.Value()
	}
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

func Insert(node *Node, key Item) *Node {
	if node == nil {
		return NewNode(key)
	}

	if key.Less(node.key) {
		node.left = Insert(node.left, key)
	} else {
		node.right = Insert(node.right, key)
	}

	node.update()
	balance := node.GetBalance()

	if balance > 1 && key.Less(node.left.key) {
		return rightRotate(node)
	}

	if balance < -1 && node.right.key.Less(key) {
		return leftRotate(node)
	}

	if balance > 1 && node.left.key.Less(key) {
		node.left = leftRotate(node.left)
		return rightRotate(node)
	}

	if balance < -1 && key.Less(node.right.key) {
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

func Delete(root *Node, key Item) *Node {
	if root == nil {
		return nil
	}

	if key.Less(root.key) {
		root.left = Delete(root.left, key)
	} else if root.key.Less(key) {
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
			root.val = tmp.key
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

func Find(node *Node, f int) Item {
	if node == nil {
		return Item{-1, -1, -1}
	}
	if node.key.key > f {
		return Find(node.left, f)
	}
	// node.key.key <= f
	// 唯一确定的是左子树都满足条件
	a := node.left.Value()
	if a.LessValue(node.key) {
		a = node.key
	}
	b := Find(node.right, f)
	if b.LessValue(a) {
		return a
	}
	return b
}
