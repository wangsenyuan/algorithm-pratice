package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := 1

	for tc > 0 {
		tc--
		n, m, k := readThreeNums(reader)
		books := make([][]int, n)
		for i := 0; i < n; i++ {
			books[i] = readNNums(reader, 3)
		}

		tot, ans := solve(n, m, k, books)
		fmt.Println(tot)
		if tot > 0 {
			var buf bytes.Buffer
			for i := 0; i < m; i++ {
				buf.WriteString(fmt.Sprintf("%d ", ans[i]))
			}
			fmt.Println(buf.String())
		}
	}
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

func solve(n, m, k int, arr [][]int) (int64, []int) {
	times := make([][]Pair, 4)

	for i := 0; i < 4; i++ {
		times[i] = make([]Pair, 0, n)
	}

	for i, book := range arr {
		t, a, b := book[0], book[1], book[2]
		times[a*2+b] = append(times[a*2+b], Pair{t, i})
	}

	sums := make([][]int64, 4)

	for i := 0; i < 4; i++ {
		sort.Sort(Pairs(times[i]))
		sums[i] = make([]int64, 0, len(times[i])+1)
		var cur int64
		sums[i] = append(sums[i], cur)
		for j := 0; j < len(times[i]); j++ {
			cur += int64(times[i][j].first)
			sums[i] = append(sums[i], cur)
		}
	}

	var ans int64 = math.MaxInt64
	var pos = math.MaxInt32

	var st *Node
	for iter := 0; iter < 2; iter++ {
		//first iter find the best, second find the result corresponding to best
		st = nil
		var fr *Node
		var sum int64 = 0
		var start int

		for k-start >= len(sums[1]) || k-start >= len(sums[2]) || m-start-(k-start)*2 < 0 {
			start++
		}
		if start >= len(sums[3]) {
			return -1, nil
		}

		need := m - start - (k-start)*2

		for i := 0; i < 3; i++ {
			var x int
			if i != 0 {
				x = k - start
			}
			for p := len(times[i]) - 1; p >= x; p-- {
				fr = Insert(fr, times[i][p])
			}
		}
		st, fr = update(st, fr, &sum, need)
		var x = len(sums[3])
		if iter == 1 {
			x = pos
		}
		for cnt := start; cnt < x; cnt++ {
			if k-cnt >= 0 {
				if cnt+(k-cnt)*2+st.Size() == m {
					tmp := sums[3][cnt] + sums[1][k-cnt] + sums[2][k-cnt] + sum
					if tmp < ans {
						ans = tmp
						pos = cnt + 1
					}
				}
			} else {
				if cnt+st.Size() == m {
					tmp := sums[3][cnt] + sum
					if tmp < ans {
						ans = tmp
						pos = cnt + 1
					}
				}
			}
			if iter == 1 && cnt+1 == pos {
				break
			}
			need--
			if k-cnt > 0 {
				need += 2
				fr = Insert(fr, times[1][k-cnt-1])
				fr = Insert(fr, times[2][k-cnt-1])
			}
			st, fr = update(st, fr, &sum, need)
		}
	}

	res := make([]int, 0, m)

	for i := 0; i+1 < pos; i++ {
		res = append(res, times[3][i].second+1)
	}

	for i := 0; i <= k-pos; i++ {
		res = append(res, times[1][i].second+1)
		res = append(res, times[2][i].second+1)
	}

	Visit(st, func(item Pair) {
		res = append(res, item.second+1)
	})

	return ans, res
}

func update(st *Node, fr *Node, sum *int64, need int) (*Node, *Node) {
	need = max(0, need)
	ok := true
	for ok {
		ok = false
		for st.Size() > need {
			x := MaxValueNode(st)
			*sum -= int64(x.item.first)
			st = Delete(st, x.item)
			fr = Insert(fr, x.item)
			ok = true
		}

		for st.Size() < need && fr.Size() > 0 {
			x := MinValueNode(fr)
			*sum += int64(x.item.first)
			fr = Delete(fr, x.item)
			st = Insert(st, x.item)
			ok = true
		}

		for st.Size() > 0 && fr.Size() > 0 {
			x := MinValueNode(fr)
			y := MaxValueNode(st)
			if !x.item.Less(y.item) {
				break
			}
			*sum += int64(x.item.first)
			*sum -= int64(y.item.first)
			fr = Delete(fr, x.item)
			fr = Insert(fr, y.item)
			st = Delete(st, y.item)
			st = Insert(st, x.item)
			ok = true
		}
	}

	return st, fr
}

type Pair struct {
	first, second int
}

func (this Pair) Less(that Pair) bool {
	if this.first < that.first {
		return true
	}
	if this.first == that.first && (this.second < that.second) {
		return true
	}
	return false
}

type Pairs []Pair

func (ps Pairs) Len() int {
	return len(ps)
}

func (ps Pairs) Less(i, j int) bool {
	return ps[i].Less(ps[j])
}

func (ps Pairs) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

/**
* this is a AVL tree
 */
type Node struct {
	item        Pair
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

func NewNode(item Pair) *Node {
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

func FindEqualOrGreater(node *Node, item Pair) *Node {
	if node == nil {
		return nil
	}
	if node.item.Less(item) {
		return FindEqualOrGreater(node.right, item)
	}

	res := FindEqualOrGreater(node.left, item)
	if res == nil {
		return node
	}
	return res
}

func FindEqualOrLess(node *Node, item Pair) *Node {
	if node == nil {
		return nil
	}
	if item.Less(node.item) {
		return FindEqualOrLess(node.left, item)
	}

	res := FindEqualOrLess(node.right, item)
	if res == nil {
		return node
	}
	return res
}

func Insert(node *Node, item Pair) *Node {
	if node == nil {
		return NewNode(item)
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

func Visit(node *Node, fn func(item Pair)) {
	if node == nil {
		return
	}
	fn(node.item)
	Visit(node.left, fn)
	Visit(node.right, fn)
}

func Delete(root *Node, item Pair) *Node {
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
