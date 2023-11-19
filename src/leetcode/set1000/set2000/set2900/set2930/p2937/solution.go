package p2937

import (
	"container/heap"
	"sort"
)

func leftmostBuildingQueries(heights []int, queries [][]int) []int {
	ans := make([]int, len(queries))
	for i := 0; i < len(queries); i++ {
		ans[i] = -1
	}
	n := len(heights)
	qs := make([][]Query, n)

	for i, cur := range queries {
		l, r := cur[0], cur[1]
		if l > r {
			l, r = r, l
		}
		if l == r || heights[l] < heights[r] {
			ans[i] = r
			continue
		}
		if len(qs[r]) == 0 {
			qs[r] = make([]Query, 0, 1)
		}
		qs[r] = append(qs[r], Query{i, heights[l]})
	}
	pq := make(PQ, 0, n)
	for i, h := range heights {
		for pq.Len() > 0 && pq[0].value < h {
			it := heap.Pop(&pq).(Query)
			ans[it.id] = i
		}
		for _, p := range qs[i] {
			heap.Push(&pq, p)
		}
	}

	return ans
}

type Query struct {
	id    int
	value int
}

type PQ []Query

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Less(i, j int) bool {
	return pq[i].value < pq[j].value
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PQ) Push(x interface{}) {
	it := x.(Query)
	*pq = append(*pq, it)
}

func (pq *PQ) Pop() interface{} {
	old := *pq
	n := len(old)
	res := old[n-1]
	*pq = old[:n-1]
	return res
}

func leftmostBuildingQueries1(heights []int, queries [][]int) []int {
	n := len(heights)

	m := len(queries)

	id := make([]int, m)
	for i := 0; i < m; i++ {
		if queries[i][0] > queries[i][1] {
			queries[i][0], queries[i][1] = queries[i][1], queries[i][0]
		}
		id[i] = i
	}
	// 找到在区间[L, R]中，且val > v的的最左边的值
	sort.Slice(id, func(i, j int) bool {
		return queries[id[i]][1] > queries[id[j]][1]
	})

	nodes := make([]Node, 4*n)

	var set func(i int, l int, r int, p int, v int)
	set = func(i int, l int, r int, p int, v int) {
		if l == r {
			nodes[i] = Node{p, v}
			return
		}
		mid := (l + r) / 2
		if p <= mid {
			set(i*2+1, l, mid, p, v)
		} else {
			set(i*2+2, mid+1, r, p, v)
		}
		nodes[i] = maxNode(nodes[i*2+1], nodes[i*2+2])
	}

	var get func(i int, l int, r int, v int) Node
	get = func(i int, l int, r int, v int) Node {
		if l == r {
			return nodes[i]
		}
		mid := (l + r) / 2
		if nodes[2*i+1].val >= v {
			return get(i*2+1, l, mid, v)
		}
		return get(i*2+2, mid+1, r, v)
	}

	ans := make([]int, m)
	p := n
	for _, i := range id {
		cur := queries[i]
		l, r := cur[0], cur[1]
		for p > r {
			set(0, 0, n-1, p-1, heights[p-1])
			p--
		}
		if l == r || heights[r] > heights[l] {
			ans[i] = r
			continue
		}
		v := max(heights[l], heights[r])
		if nodes[0].val <= v {
			ans[i] = -1
			continue
		}
		node := get(0, 0, n-1, v+1)
		ans[i] = node.pos
	}

	return ans
}

type Node struct {
	pos int
	val int
}

func maxNode(a, b Node) Node {
	if a.val > b.val || a.val == b.val && a.pos < b.pos {
		return a
	}
	return b
}
