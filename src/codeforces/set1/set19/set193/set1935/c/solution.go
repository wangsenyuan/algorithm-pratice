package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)
	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n, l := readTwoNums(reader)
		messages := make([][]int, n)
		for i := 0; i < n; i++ {
			messages[i] = readNNums(reader, 2)
		}
		res := solve(l, messages)

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

func solve(l int, messages [][]int) int {
	sort.Slice(messages, func(i, j int) bool {
		return messages[i][1] < messages[j][1]
	})
	n := len(messages)
	fp := make([]int, n)
	var ans int
	for i, cur := range messages {
		fp[i] = cur[0] - cur[1]
		if cur[0] <= l {
			ans = 1
		}
	}

	for i := 1; i < n; i++ {
		mn := fp[i-1]
		for j := i; j < n; j++ {
			fp[j], mn = mn+messages[j][0], min(mn, fp[j])
			if fp[j]+messages[j][1] <= l {
				ans = i + 1
			}
		}
	}

	return ans
}

func solve1(l int, messages [][]int) int {
	slices.SortFunc(messages, func(x, y []int) int {
		// 在b相同的情况下，降序排
		if x[1] == y[1] {
			return y[0] - x[0]
		}
		return x[1] - y[1]
	})

	n := len(messages)

	var best int
	for i := 0; i < n; i++ {
		var tr *Node
		for j := i; j < n; j++ {
			// sum + b[j] - b[i] <= l
			// 要知道区间 i...j 中间最小的x个数，它们的sum <= l - (b[j] - b[i])
			// 要找到前x个数，它们的sum
			// 这样子是不是过于复杂了？
			tr = Insert(tr, messages[j][0])

			expect := l - (messages[j][1] - messages[i][1])
			if expect >= 0 {
				tmp := FindFirstK(tr, expect)

				best = max(best, tmp)
			}

		}
	}

	return best
}

/**
* this is a AVL tree
 */
type Node struct {
	key         int
	height      int
	cnt         int
	size        int
	sum         int
	left, right *Node
}

func (node *Node) Height() int {
	if node == nil {
		return 0
	}
	return node.height
}

func (node *Node) Count() int {
	if node == nil {
		return 0
	}
	return node.cnt
}

func (node *Node) Size() int {
	if node == nil {
		return 0
	}
	return node.size
}

func (node *Node) Sum() int {
	if node == nil {
		return 0
	}
	return node.sum
}

func NewNode(key int) *Node {
	node := new(Node)
	node.key = key
	node.height = 1
	node.cnt = 1
	node.size = 1
	node.sum = key
	return node
}

func (node *Node) Update() {
	node.height = max(node.left.Height(), node.right.Height()) + 1
	node.size = node.left.Size() + node.right.Size() + node.cnt
	node.sum = node.left.Sum() + node.right.Sum() + node.cnt*node.key
}

func rightRotate(y *Node) *Node {
	x := y.left

	t2 := x.right

	x.right = y
	y.left = t2
	y.Update()
	x.Update()

	return x
}

func leftRotate(x *Node) *Node {
	y := x.right
	t2 := y.left

	y.left = x
	x.right = t2

	x.Update()

	y.Update()

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
		node.size++
		node.sum += key
		return node
	}

	if node.key > key {
		node.left = Insert(node.left, key)
	} else {
		node.right = Insert(node.right, key)
	}

	node.Update()

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

func FindFirstK(node *Node, expect int) int {
	if node == nil || node.Sum() <= expect {
		return node.Size()
	}
	// node.Sum() > expect
	if node.left.Sum() >= expect {
		return FindFirstK(node.left, expect)
	}

	res := node.left.Size()
	expect -= node.left.Sum()
	if expect <= node.cnt*node.key {
		res += expect / node.key
		return res
	}
	res += node.cnt
	expect -= node.cnt * node.key
	return res + FindFirstK(node.right, expect)
}
