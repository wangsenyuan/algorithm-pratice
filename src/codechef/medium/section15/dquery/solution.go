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
		A := readNNums(reader, n)
		m := readNum(reader)
		Q := make([][]int, m)
		for i := 0; i < m; i++ {
			Q[i] = readNNums(reader, 2)
		}
		res := solve(A, Q)

		for _, ans := range res {
			buf.WriteString(fmt.Sprintf("%d\n", ans))
		}
	}

	fmt.Print(buf.String())
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

const MAX_M = 100010

var pf [][]int

func init() {

	pf = make([][]int, MAX_M)

	for i := 2; i < MAX_M; i++ {
		pf[i] = make([]int, 0, 2)
	}

	set := make([]bool, MAX_M)

	for i := 2; i < MAX_M; i++ {
		if !set[i] {
			for j := i; j < MAX_M; j += i {
				pf[j] = append(pf[j], i)
				set[j] = true
			}
		}
	}
}

func solve(A []int, Q [][]int) []int64 {
	n := len(A)

	pref := make([]int64, n)

	for i := 0; i < n; i++ {
		pref[i] = int64(A[i])
		if i > 0 {
			pref[i] += pref[i-1]
		}
	}

	arr := make(map[int][]int)
	val := make(map[int][]int)
	process := func(i int, num int) {
		for _, p := range pf[num] {
			arr[p] = append(arr[p], i)
			val[p] = append(val[p], num)
		}
	}

	for i, num := range A {
		process(i, num)
	}
	sets := make(map[int]*Set)
	for k, a := range arr {
		v := val[k]
		sets[k] = NewSet(a, v)
	}

	ans := make([]int64, len(Q))
	for i, cur := range Q {
		p, k := cur[0], cur[1]
		k--
		ans[i] = pref[k]
		if sets[p] != nil {
			ans[i] += sets[p].Replace(k)
		}
	}

	return ans
}

type Set struct {
	arr  []int
	sum1 []int64
	sum2 []int64
}

func NewSet(arr []int, val []int) *Set {
	n := len(arr)
	sum1 := make([]int64, n)

	for i := 0; i < n; i++ {
		sum1[i] = int64(val[i])
		if i > 0 {
			sum1[i] += sum1[i-1]
		}
	}

	sum2 := make([]int64, n)
	sort.Ints(val)
	for i, j := n-1, 0; i >= 0; i-- {
		sum2[j] = int64(val[i])
		if j > 0 {
			sum2[j] += sum2[j-1]
		}
		j++
	}
	return &Set{arr, sum1, sum2}
}

func (s *Set) Replace(k int) int64 {
	i := search(len(s.arr), func(i int) bool {
		return s.arr[i] > k
	})

	if i == 0 {
		return 0
	}
	return s.sum2[i-1] - s.sum1[i-1]
}

func search(n int, f func(int) bool) int {
	l, r := 0, n
	for l < r {
		mid := (l + r) / 2
		if f(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}

func solve1(A []int, Q [][]int) []int64 {
	pos_tree := make(map[int]*Node)
	max_tree := make(map[int]*Node)

	pref_sum := make([]int64, len(A))

	for i, num := range A {
		for _, p := range pf[num] {
			pos_tree[p] = Insert(pos_tree[p], Pair{i, num})
			max_tree[p] = Insert(max_tree[p], Pair{-num, num})
		}

		pref_sum[i] = int64(num)
		if i > 0 {
			pref_sum[i] += pref_sum[i-1]
		}
	}

	ans := make([]int64, len(Q))

	for i, cur := range Q {
		p, k := cur[0], cur[1]

		tmp := pref_sum[k-1]

		// 找出在k之前的节点
		c, s := FindCountAndSumLessThanKey(pos_tree[p], k-1)

		if c == 0 || c == pos_tree[p].Count() {
			ans[i] = tmp
			continue
		}
		tmp += FirstKSum(max_tree[p], c) - s

		ans[i] = tmp
	}

	return ans
}

type Pair struct {
	first  int
	second int
}

/**
* this is a AVL tree
 */
type Node struct {
	key         int
	height      int
	cnt         int
	val         int
	sum         int64
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

func (node *Node) Count() int {
	if node == nil {
		return 0
	}
	return node.cnt
}

func (node *Node) Update() {
	node.height = max(node.left.Height(), node.right.Height()) + 1
	node.sum = node.left.Sum() + node.right.Sum() + int64(node.val)
	node.cnt = node.left.Count() + node.right.Count() + 1
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func NewNode(key Pair) *Node {
	node := new(Node)
	node.key = key.first
	node.height = 1
	node.cnt = 1
	node.val = key.second
	node.sum = int64(key.second)
	return node
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

func Insert(node *Node, key Pair) *Node {
	if node == nil {
		return NewNode(key)
	}

	if node.key >= key.first {
		node.left = Insert(node.left, key)
	} else {
		node.right = Insert(node.right, key)
	}

	node.Update()

	balance := node.GetBalance()

	if balance > 1 && key.first <= node.left.key {
		return rightRotate(node)
	}

	if balance < -1 && key.first >= node.right.key {
		return leftRotate(node)
	}

	if balance > 1 && key.first >= node.left.key {
		node.left = leftRotate(node.left)
		return rightRotate(node)
	}

	if balance < -1 && key.first <= node.right.key {
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

// 计算node.key <= key 的节点的数量及总和
func FindCountAndSumLessThanKey(root *Node, key int) (int, int64) {
	if root == nil {
		return 0, 0
	}

	if root.key > key {
		return FindCountAndSumLessThanKey(root.left, key)
	}
	// root.key <= key
	cnt, sum := root.left.Count(), root.left.Sum()
	cnt++
	sum += int64(root.val)

	a, b := FindCountAndSumLessThanKey(root.right, key)
	cnt += a
	sum += b
	return cnt, sum
}

func FirstKSum(root *Node, k int) int64 {
	if root == nil || k == 0 {
		return 0
	}
	if k <= root.left.Count() {
		return FirstKSum(root.left, k)
	}
	sum := root.left.Sum() + int64(root.val)
	// k > root.left.Count
	k -= root.left.Count()
	// current root
	k--
	if k == 0 {
		return sum
	}

	return sum + FirstKSum(root.right, k)
}
