package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, k := readTwoNums(reader)
	items := make([][]int, n)
	for i := 0; i < n; i++ {
		items[i] = readNNums(reader, 3)
	}
	res := solve(k, items)
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

const INF_X = 1000000007
const NNF_X = -INF_X
const INF = 2e18
const UNIT = -INF

func solve(k int, items [][]int) int64 {
	n := len(items)
	pp := make([]int, 2*n)
	events := make([]Event, n)

	for i := 0; i < n; i++ {
		cur := items[i]
		pp[2*i] = cur[0]
		pp[2*i+1] = cur[1]
		events[i] = Event{cur[0], cur[1], cur[2]}
	}

	sort.Ints(pp)

	var m int
	for i := 1; i <= 2*n; i++ {
		if i == 2*n || pp[i] > pp[i-1] {
			pp[m] = pp[i-1]
			m++
		}
	}

	sort.Sort(Events(events))

	tree := NewSegmentTree(m)

	var ans int64

	var cur int

	K := int64(k)

	for r := 0; r < m; r++ {
		tree.Add(r, r, INF+K*int64(pp[r]))

		for cur < n && events[cur].y <= pp[r] {
			if events[cur].v > 0 {
				id := search(m, func(i int) bool {
					return pp[i] >= events[cur].x
				})
				tree.Add(0, id, int64(events[cur].v))
			}
			cur++
		}

		ans = max(ans, tree.arr[1]-K*int64(pp[r]))
	}

	return ans
}

func search(n int, fn func(int) bool) int {
	left, right := 0, n
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

type SegmentTree struct {
	arr  []int64
	lazy []int64
	n    int
}

func NewSegmentTree(n int) *SegmentTree {
	arr := make([]int64, 4*n)
	lazy := make([]int64, 4*n)
	for i := 0; i < len(arr); i++ {
		arr[i] = -INF
	}
	return &SegmentTree{arr, lazy, n}
}

func (tree *SegmentTree) change(id int, v int64) {
	tree.arr[id] += v
	tree.lazy[id] += v
}

func (tree *SegmentTree) push(id int) {
	tree.change(id<<1, tree.lazy[id])
	tree.change((id<<1)|1, tree.lazy[id])
	tree.lazy[id] = 0
}

func (tree *SegmentTree) Add(L, R int, V int64) {
	var loop func(id int, ll, rr int)
	loop = func(id int, ll, rr int) {
		if R < ll || rr < L {
			return
		}
		if L <= ll && rr <= R {
			tree.change(id, V)
			return
		}
		tree.push(id)
		mid := (ll + rr) / 2
		loop(id*2, ll, mid)
		loop(id*2+1, mid+1, rr)
		tree.arr[id] = max(tree.arr[id*2], tree.arr[id*2+1])
	}

	loop(1, 0, tree.n-1)
}

func solve1(k int, items [][]int) int64 {
	n := len(items)
	events := make([]Event, n)
	for i := 0; i < n; i++ {
		l, r := items[i][0], items[i][1]
		events[i] = Event{l, r, items[i][2]}
	}

	K := int64(k)
	sort.Sort(Events(events))

	tr := NewNode(NNF_X, INF_X)

	for _, cur := range events {
		tr.Set(cur.x, cur.x+1, K*int64(cur.x))
		tr.Set(cur.y, cur.y+1, K*int64(cur.y))
	}

	var best int64

	for _, cur := range events {
		tr.Add(NNF_X, cur.x+1, int64(cur.v))
		best = max(best, tr.Query(NNF_X, cur.y+1)-K*int64(cur.y))
	}

	return best
}

type Node struct {
	lo, hi      int
	left, right *Node
	mset        int64
	madd        int64
	val         int64
}

func NewNode(lo, hi int) *Node {
	return &Node{lo, hi, nil, nil, UNIT, 0, UNIT}
}

func (node *Node) Query(L, R int) int64 {
	if R <= node.lo || node.hi <= L {
		return UNIT
	}
	if L <= node.lo && node.hi <= R {
		return node.val
	}
	node.push()

	return max(node.left.Query(L, R), node.right.Query(L, R))
}

func (node *Node) push() {
	if node.left == nil {
		mid := node.lo + (node.hi-node.lo)/2
		node.left = NewNode(node.lo, mid)
		node.right = NewNode(mid, node.hi)
	}
	if node.mset != UNIT {
		node.left.Set(node.lo, node.hi, node.mset)
		node.right.Set(node.lo, node.hi, node.mset)
		node.mset = UNIT
	} else if node.madd != 0 {
		node.left.Add(node.lo, node.hi, node.madd)
		node.right.Add(node.lo, node.hi, node.madd)
		node.madd = 0
	}
}

func (node *Node) Set(L, R int, x int64) {
	if R <= node.lo || node.hi <= L {
		return
	}
	if L <= node.lo && node.hi <= R {
		node.mset = x
		node.val = x
		node.madd = 0
		return
	}
	node.push()
	node.left.Set(L, R, x)
	node.right.Set(L, R, x)
	node.val = max(node.left.val, node.right.val)
}

func (node *Node) Add(L, R int, x int64) {
	if R <= node.lo || node.hi <= L {
		return
	}
	if L <= node.lo && node.hi <= R {
		if node.mset != UNIT {
			node.mset += x
		} else {
			node.madd += x
		}
		node.val += x
		return
	}
	node.push()
	node.left.Add(L, R, x)
	node.right.Add(L, R, x)
	node.val = max(node.left.val, node.right.val)
}

type Event struct {
	x, y, v int
}

func (this Event) Less(that Event) bool {
	return this.y < that.y
}

type Events []Event

func (this Events) Len() int {
	return len(this)
}

func (this Events) Less(i, j int) bool {
	return this[i].Less(this[j])
}

func (this Events) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func max(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}
