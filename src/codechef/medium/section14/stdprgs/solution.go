package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"math/rand"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	A := readNNums(reader, n)
	m := readNum(reader)
	Q := make([][]int, m)
	for i := 0; i < m; i++ {
		Q[i] = readNNums(reader, 2)
	}
	res := solve(n, A, Q)

	var buf bytes.Buffer

	for i := 0; i < m; i++ {
		buf.WriteString(fmt.Sprintf("%d\n", res[i]))
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

var BIN_SIZE = 200

func solve(n int, A []int, Q [][]int) []int64 {
	var mx int

	for i := 0; i < n; i++ {
		if A[i] > mx {
			mx = A[i]
		}
	}
	if mx <= 100 {
		return solveTask3(n, A, Q, mx)
	}

	BIN_SIZE = int(math.Sqrt(float64(n)) + 1)

	qq := make([]Query, len(Q))

	for i := 0; i < len(Q); i++ {
		qq[i] = Query{Q[i][0] - 1, Q[i][1] - 1, i}
	}
	sort.Sort(Queries(qq))

	var ans int64
	var t0 *Node

	remove := func(num int) {
		// remove num from the tree, and cancel its contribution
		var t1, t2, t3 *Node

		// find where num is
		i := countLess(t0, num)
		split(t0, &t1, &t2, i)
		split(t2, &t2, &t3, 1)
		//t2 holds the num now
		a := maxNode(t1)
		b := minNode(t3)

		if a != nil {
			ans -= square(num - a.value)
		}
		if b != nil {
			ans -= square(b.value - num)
		}
		if a != nil && b != nil {
			ans += square(b.value - a.value)
		}
		merge(&t0, t1, t3)
	}

	add := func(num int) {
		a := findLess(t0, num)  //a.value < num
		b := findLower(t0, num) // b.value >= num
		if a != nil && b != nil {
			ans -= square(b.value - a.value)
		}
		if a != nil {
			ans += square(num - a.value)
		}
		if b != nil {
			ans += square(b.value - num)
		}
		c := NewNode(num)
		tmp := countLess(t0, num)
		var t1, t2 *Node
		split(t0, &t1, &t2, tmp)
		//t1 hold values < num while t2 holds values >= num
		merge(&t0, t1, c)
		merge(&t0, t0, t2)
	}

	res := make([]int64, len(qq))
	L, R := 0, 0

	for _, q := range qq {

		for R <= q.right {
			add(A[R])
			R++
		}
		for R > q.right+1 {
			R--
			remove(A[R])
		}

		for L < q.left {
			remove(A[L])
			L++
		}

		for q.left < L {
			L--
			add(A[L])
		}

		res[q.pos] = ans
	}

	return res
}

type Query struct {
	left, right int
	pos         int
}

type Queries []Query

func (qs Queries) Len() int {
	return len(qs)
}

func (qs Queries) Less(i, j int) bool {
	return qs[i].left/BIN_SIZE < qs[j].left/BIN_SIZE || qs[i].left/BIN_SIZE == qs[j].left/BIN_SIZE && qs[i].right < qs[j].right
}

func (qs Queries) Swap(i, j int) {
	qs[i], qs[j] = qs[j], qs[i]
}

type Node struct {
	priority int
	cnt      int
	value    int

	left, right *Node
}

func NewNode(value int) *Node {
	node := new(Node)
	node.value = value
	node.priority = rand.Int()
	node.left = nil
	node.right = nil
	node.cnt = 1
	return node
}

func (node *Node) Count() int {
	if node == nil {
		return 0
	}
	return node.cnt
}

func push(node *Node) {
	if node == nil {
		return
	}

	node.cnt = node.left.Count() + node.right.Count() + 1
}

func countLess(t *Node, key int) int {
	var res int
	cur := t

	for cur != nil {
		if cur.value < key {
			res += cur.left.Count() + 1
			cur = cur.right
		} else {
			cur = cur.left
		}
	}
	return res
}

func findLess(t *Node, key int) *Node {
	if t == nil {
		return nil
	}
	if t.value > key {
		return findLess(t.left, key)
	}
	// t.value <= key
	tmp := findLess(t.right, key)

	if tmp == nil {
		return t
	}
	return tmp
}

func findLower(t *Node, key int) *Node {
	if t == nil {
		return nil
	}
	if t.value < key {
		return findLower(t.right, key)
	}
	// t >= key
	tmp := findLower(t.left, key)
	if tmp == nil {
		return t
	}
	return tmp
}

func maxNode(t *Node) *Node {
	if t == nil {
		return nil
	}
	res := t
	for res.right != nil {
		res = res.right
	}
	return res
}

func minNode(t *Node) *Node {
	if t == nil {
		return nil
	}
	res := t
	for res.left != nil {
		res = res.left
	}
	return res
}

func split(t *Node, l **Node, r **Node, at int) {
	push(t)
	if t == nil {
		*l = nil
		*r = nil
		return
	}
	if at <= t.left.Count() {
		split(t.left, l, &(t.left), at)
		*r = t
	} else {
		split(t.right, &(t.right), r, at-t.left.Count()-1)
		*l = t
	}
	push(*l)
	push(*r)
}

func merge(t **Node, l *Node, r *Node) {
	push(l)
	push(r)
	if l == nil || r == nil {
		if l != nil {
			*t = l
		} else {
			*t = r
		}
	} else if l.priority > r.priority {
		merge(&(l.right), l.right, r)
		*t = l
	} else {
		merge(&(r.left), l, r.left)
		*t = r
	}
	push(*t)
}

func square(num int) int64 {
	return int64(num) * int64(num)
}

func solveTask3(n int, A []int, Q [][]int, mx int) []int64 {
	qq := make([][]int, mx+1)
	for i := 0; i <= mx; i++ {
		qq[i] = make([]int, 0, 100)
	}

	for i := 0; i < n; i++ {
		cur := A[i]
		qq[cur] = append(qq[cur], i)
	}

	res := make([]int64, len(Q))

	for i, q := range Q {
		l, r := q[0], q[1]
		l--
		r--
		prev := -1
		for num := 0; num <= mx; num++ {
			li := search(len(qq[num]), func(li int) bool {
				return qq[num][li] >= l
			})
			if li < len(qq[num]) && qq[num][li] <= r {
				// num has some i >= l && i <= r
				if prev >= 0 {
					res[i] += square(num - prev)
				}
				prev = num
			}
		}
	}
	return res
}

func search(n int, fn func(int) bool) int {
	var left, right = 0, n
	for left < right {
		mid := (left + right) / 2
		if fn(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return right
}
