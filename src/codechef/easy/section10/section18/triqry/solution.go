package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

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

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n, Q := readTwoNums(reader)

		points := make([][]int, n)

		for i := 0; i < n; i++ {
			points[i] = readNNums(reader, 2)
		}

		queries := make([][]int, Q)

		for i := 0; i < Q; i++ {
			queries[i] = readNNums(reader, 2)
		}

		res := solve(n, points, Q, queries)

		for i := 0; i < Q; i++ {
			if i < Q-1 {
				fmt.Printf("%d ", res[i])
			} else {
				fmt.Printf("%d\n", res[i])
			}
		}
	}
}

const INF = 2000000

func solve(n int, points [][]int, Q int, queries [][]int) []int {
	ps := make([]Point, n)

	for i := 0; i < n; i++ {
		ps[i] = Point{points[i][0], points[i][1]}
	}

	sort.Sort(Points(ps))

	qs := make([]Query, Q)

	for i := 0; i < Q; i++ {
		qs[i] = Query{queries[i][0], queries[i][1], i}
	}

	sort.Sort(Queries(qs))

	arr := make([]int, INF+1)

	update := func(pos int, v int) {
		pos++
		for pos <= INF {
			arr[pos] += v
			pos += pos & -pos
		}
	}

	get := func(pos int) int {
		var res int
		pos++
		for pos > 0 {
			res += arr[pos]
			pos -= pos & -pos
		}

		return res
	}

	for i := 0; i < n; i++ {
		cur := ps[i]
		update(cur.x+cur.y, 1)
	}

	ans := make([]int, Q)

	for i, j := 0, 0; i < Q; i++ {
		l, r := qs[i].l, qs[i].r

		for j < len(ps) && ps[j].y-ps[j].x > -l {
			update(ps[j].x+ps[j].y, -1)
			j++
		}

		ans[qs[i].i] = get(r)
	}

	return ans
}

func solve1(n int, points [][]int, Q int, queries [][]int) []int {
	ps := make([]Point, n)

	for i := 0; i < n; i++ {
		ps[i] = Point{points[i][0], points[i][1]}
	}

	sort.Sort(Points(ps))

	qs := make([]Query, Q)

	for i := 0; i < Q; i++ {
		qs[i] = Query{queries[i][0], queries[i][1], i}
	}

	sort.Sort(Queries(qs))

	root := NewNode(0, INF)

	for i := 0; i < len(ps); i++ {
		root.Set(ps[i].x+ps[i].y, 1)
	}

	ans := make([]int, Q)

	for i, j := 0, 0; i < Q; i++ {
		l, r := qs[i].l, qs[i].r

		for j < len(ps) && ps[j].y-ps[j].x > -l {
			root.Set(ps[j].x+ps[j].y, -1)
			j++
		}

		ans[qs[i].i] = root.Get(r)
	}

	return ans
}

type Node struct {
	count       int
	start, end  int
	left, right *Node
}

func NewNode(start, end int) *Node {
	node := new(Node)
	node.start = start
	node.end = end
	return node
}

func (node Node) GetMid() int {
	return (node.start + node.end) / 2
}

func (node *Node) push() {
	if node.start == node.end || node.left != nil {
		return
	}
	mid := node.GetMid()
	node.left = NewNode(node.start, mid)
	node.right = NewNode(mid+1, node.end)
}

func (node *Node) Set(pos int, v int) {
	if pos < node.start || pos > node.end {
		return
	}
	node.count += v

	if node.start == node.end {
		return
	}

	node.push()

	mid := node.GetMid()

	if pos <= mid {
		node.left.Set(pos, v)
	} else {
		node.right.Set(pos, v)
	}
}

/**
 * Get the count less than pos
 */
func (node *Node) Get(pos int) int {
	if node == nil || pos < node.start {
		return 0
	}
	if node.start == node.end {
		return node.count
	}
	mid := node.GetMid()
	if pos <= mid {
		return node.left.Get(pos)
	}
	return node.left.count + node.right.Get(pos)
}

type Query struct {
	l, r int
	i    int
}

type Queries []Query

func (qs Queries) Len() int {
	return len(qs)
}

func (qs Queries) Less(i, j int) bool {
	if qs[i].l < qs[j].l {
		return true
	}

	if qs[i].l == qs[j].l {
		return qs[i].r < qs[j].r
	}

	return false
}

func (qs Queries) Swap(i, j int) {
	qs[i], qs[j] = qs[j], qs[i]
}

type Point struct {
	x, y int
}

type Points []Point

func (ps Points) Len() int {
	return len(ps)
}

func (ps Points) Less(i, j int) bool {
	// y = x + 0
	// y = x + 1
	// y - x = d
	if ps[i].y-ps[i].x > ps[j].y-ps[j].x {
		return true
	}
	if ps[i].y-ps[i].x == ps[j].y-ps[j].x {
		return ps[i].x < ps[j].x
	}

	return false
}

func (ps Points) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}
