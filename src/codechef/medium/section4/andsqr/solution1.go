package main

import "math"

const MAX_N = 100005

var bit [][]int64
var groups *Groups
var backup *Groups
var rs []*Nodes

func init() {
	bit = make([][]int64, 2)
	for i := 0; i < 2; i++ {
		bit[i] = make([]int64, MAX_N)
	}

	groups = NewGroups(100)
	backup = NewGroups(100)

	rs = make([]*Nodes, MAX_N)
}

func update(n int, num int, pos int, v int64) {
	for pos < n {
		bit[num][pos] += v
		pos = (pos | (pos + 1))
	}
}

func get(n int, num int, pos int) int64 {
	var res int64

	for pos >= 0 {
		res += bit[num][pos]
		pos = (pos & (pos + 1)) - 1
	}
	return res
}

func reset(n int) {
	for i := 0; i < n; i++ {
		bit[0][i] = 0
		bit[1][i] = 0
		rs[i] = new(Nodes)
	}
	bit[0][n] = 0
	bit[1][n] = 0
}

func solve1(n int, A []int, Q int, quries [][]int) []int64 {
	reset(n)

	for i := 0; i < Q; i++ {
		l, r := quries[i][0], quries[i][1]
		l--
		r--
		rs[r].Add(Pair{l, i})
	}
	groups.Reset()

	ans := make([]int64, Q)

	for i := 0; i < n; i++ {
		backup.Reset()

		for j := 0; j < groups.Size(); j++ {
			cur := groups.Get(j).first & A[i]
			if j == 0 || cur != backup.Last().first {
				backup.Add(Pair{cur, groups.Get(j).second})
			}
		}

		if backup.Size() == 0 || backup.Last().first != A[i] {
			backup.Add(Pair{A[i], i})
		}

		groups, backup = backup, groups

		for j := 0; j < groups.Size(); j++ {
			cur := groups.Get(j)
			sq := int(math.Sqrt(float64(cur.first)))

			if (sq-1)*(sq-1) == cur.first || sq*sq == cur.first || (sq+1)*(sq+1) == cur.first {
				l := cur.second

				r := i

				if j < groups.Size()-1 {
					r = groups.Get(j+1).second - 1
				}

				update(n, 0, l, int64(r+1))
				update(n, 0, r+1, -int64(r+1))
				update(n, 1, l, 1)
				update(n, 1, r+1, -1)
				update(n, 0, 0, int64(r-l+1))
				update(n, 0, l, -int64(r-l+1))
			}
		}

		if rs[i] != nil {
			rs[i].ForEach(func(pair Pair) {
				j := pair.first
				k := pair.second
				ans[k] = get(n, 0, j)
				ans[k] -= get(n, 1, j) * int64(j)
			})
		}
	}

	return ans
}

type Node struct {
	Pair
	next *Node
}

type Nodes struct {
	head, tail *Node
}

func (nodes *Nodes) Add(pair Pair) {
	node := &Node{Pair{pair.first, pair.second}, nil}
	if nodes.head == nil {
		nodes.head = node
		nodes.tail = nodes.head
		return
	}
	nodes.tail.next = node
	nodes.tail = node
}

func (nodes *Nodes) ForEach(fn func(pair Pair)) {
	cur := nodes.head

	for cur != nil {
		fn(cur.Pair)
		cur = cur.next
	}
}

type Pair struct {
	first  int
	second int
}

type Groups struct {
	pairs []Pair
	cur   int
}

func NewGroups(n int) *Groups {
	pairs := make([]Pair, n)
	return &Groups{pairs, 0}
}

func (groups *Groups) Reset() {
	groups.cur = 0
}

func (groups Groups) Size() int {
	return groups.cur
}

func (groups *Groups) Add(pair Pair) {
	groups.pairs[groups.cur] = pair
	groups.cur++
}

func (groups Groups) Get(i int) Pair {
	return groups.pairs[i]
}

func (groups Groups) Last() Pair {
	return groups.pairs[groups.cur]
}
