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

	var buf bytes.Buffer
	tc := readNum(reader)

	for tc > 0 {
		tc--
		n, m, k := readThreeNums(reader)
		A := readNNums(reader, n)
		Q := make([][]int, k)
		for i := 0; i < k; i++ {
			Q[i] = readNNums(reader, 2)
		}
		res := solve(n, m, k, A, Q)

		for i := 0; i < k; i++ {
			buf.WriteString(fmt.Sprintf("%d\n", res[i]))
		}
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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
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

const H = 30
const INF = 1 << 30

func solve(n int, m int, k int, A []int, Q [][]int) []int {
	root := NewNode()

	best := make([]int, n+1)
	for i := 0; i <= n; i++ {
		best[i] = INF
	}

	var cur = 0
	root.Insert(cur, H-1, cur)

	for i := 1; i <= n; i++ {
		cur ^= A[i-1]
		j := root.Query(cur, m, H-1)
		if j >= 0 {
			best[j+1] = min(best[j+1], i)
		}
		root.Insert(cur, H-1, i)
	}

	ends := make([]Pair, k)

	for i := 0; i < k; i++ {
		ends[i] = Pair{Q[i][1], i + 1}
	}

	sort.Sort(Pairs(ends))

	where := make([]int, n+1)

	for i := 0; i < k; i++ {
		where[ends[i].second] = i + 1
	}

	lazy := make([]int, 4*(k+1))
	st := make([]int, 4*(k+1))

	var build func(i int, l, r int)
	build = func(i, l, r int) {
		lazy[i] = INF
		if l == r {
			st[i] = INF
			return
		}
		mid := (l + r) >> 1
		build(i<<1, l, mid)
		build((i<<1)|1, mid+1, r)
		st[i] = INF
	}
	build(1, 1, k)

	push := func(i, l, r int) {
		if lazy[i] < INF {
			if l < r {
				lazy[i<<1] = min(lazy[i<<1], lazy[i])
				lazy[(i<<1)|1] = min(lazy[(i<<1)|1], lazy[i])
			}
			st[i] = min(st[i], lazy[i])
			lazy[i] = INF
		}
	}
	// ll, rr is the range to update
	// l, r is the range for ith node
	var update func(ll, rr, l, r int, i int, t int, v int)

	update = func(ll, rr, l, r int, i int, t int, v int) {
		push(i, l, r)
		if r < ll || rr < l {
			return
		}
		if ll <= l && r <= rr {
			// point update
			if t == 0 {
				st[i] = v
			} else {
				lazy[i] = v
				push(i, l, r)
			}
			return
		}
		mid := (l + r) >> 1
		update(ll, rr, l, mid, i<<1, t, v)
		update(ll, rr, mid+1, r, (i<<1)|1, t, v)
		st[i] = min(st[i<<1], st[(i<<1)|1])
	}

	var query func(pos int, i int, l, r int) int

	query = func(pos int, i, l, r int) int {
		push(i, l, r)
		if pos > r || pos < l {
			return INF
		}
		if l == r {
			return st[i]
		}

		mid := (l + r) >> 1
		if pos <= mid {
			return query(pos, i<<1, l, mid)
		}
		return query(pos, (i<<1)|1, mid+1, r)
	}

	events := make(Events, 0, n+k)

	for i := 1; i <= n; i++ {
		events = append(events, Event{1, i, i})
	}
	for i := 1; i <= k; i++ {
		events = append(events, Event{0, i, Q[i-1][0]})
	}

	sort.Sort(events)

	for _, cur := range events {
		if cur.tp == 0 {
			// activate R
			update(where[cur.idx], where[cur.idx], 1, k, 1, 0, INF)
		} else {
			if best[cur.idx] < INF {
				from := sort.Search(k, func(i int) bool {
					return ends[i].first >= best[cur.idx]
				})
				from++
				update(from, k, 1, k, 1, 1, best[cur.idx]-cur.idx+1)
			}
		}
	}

	ans := make([]int, k)
	for i := 0; i < k; i++ {
		ans[i] = query(where[i+1], 1, 1, k)
		if ans[i] == INF {
			ans[i] = -1
		}
	}

	return ans
}

type Event struct {
	tp  int
	idx int
	pos int
}

type Events []Event

func (evts Events) Len() int {
	return len(evts)
}

func (evts Events) Less(i, j int) bool {
	return evts[i].pos < evts[j].pos || evts[i].pos == evts[j].pos && evts[i].tp < evts[j].tp
}

func (evts Events) Swap(i, j int) {
	evts[i], evts[j] = evts[j], evts[i]
}

type Pair struct {
	first, second int
}

type Pairs []Pair

func (ps Pairs) Len() int {
	return len(ps)
}

func (ps Pairs) Less(i, j int) bool {
	return ps[i].first < ps[j].first
}

func (ps Pairs) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

type Trie struct {
	children [2]*Trie
	maxIndex int
}

func NewNode() *Trie {
	node := new(Trie)
	node.maxIndex = -1
	return node
}

func (this *Trie) Insert(num int, pos int, index int) {
	this.maxIndex = max(this.maxIndex, index)
	if pos < 0 {
		return
	}
	d := (num >> uint(pos)) & 1
	if this.children[d] == nil {
		this.children[d] = NewNode()
	}
	this.children[d].Insert(num, pos-1, index)
}

func (this Trie) Query(a, b int, pos int) int {
	if pos < 0 {
		return this.maxIndex
	}
	d := (a >> uint(pos)) & 1
	dm := (b >> uint(pos)) & 1
	var ans = -1
	if dm == 0 && this.children[1-d] != nil {
		ans = this.children[1-d].maxIndex
	}
	if dm == 0 && this.children[d] != nil {
		ans = max(ans, this.children[d].Query(a, b, pos-1))
	} else if dm == 1 && this.children[1-d] != nil {
		ans = max(ans, this.children[1-d].Query(a, b, pos-1))
	}
	return ans
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	return a + b - max(a, b)
}
