package p3213

import (
	"container/heap"
	"index/suffixarray"
)

const inf = 1 << 60

func minimumCost(target string, words []string, costs []int) int {
	minCost := map[string]uint16{}
	for i, w := range words {
		c := uint16(costs[i])
		if minCost[w] == 0 {
			minCost[w] = c
		} else {
			minCost[w] = min(minCost[w], c)
		}
	}

	n := len(target)
	type pair struct{ l, cost uint16 }
	from := make([][]pair, n+1)
	sa := suffixarray.New([]byte(target))
	for w, c := range minCost {
		for _, l := range sa.Lookup([]byte(w), -1) {
			r := l + len(w)
			from[r] = append(from[r], pair{uint16(l), c})
		}
	}

	f := make([]int, n+1)
	for i := 1; i <= n; i++ {
		f[i] = inf
		for _, p := range from[i] {
			f[i] = min(f[i], f[p.l]+int(p.cost))
		}
	}
	if f[n] >= inf {
		return -1
	}
	return f[n]
}

func minimumCost1(target string, words []string, costs []int) int {
	// n 很大
	// words时模式
	// dp[i]是将target的前缀搞定的最小cost
	// dp[j] = dp[i-len(words[k])] + cost(words[k]) if words[k] found
	// 所以要构造图
	tr := NewNode()

	for i, word := range words {
		tr.Add(word, costs[i])
	}

	n := len(target)
	items := make([]*Item, n+1)
	pq := make(PriorityQueue, n+1)
	for i := 0; i <= n; i++ {
		it := new(Item)
		it.id = i
		it.index = i
		it.priority = inf
		items[i] = it
		pq[i] = it
	}
	items[0].priority = 0

	heap.Init(&pq)

	for pq.Len() > 0 {
		it := heap.Pop(&pq).(*Item)
		if it.priority >= inf {
			break
		}
		u := it.id
		node := tr
		for j := u; j < n; j++ {
			x := int(target[j] - 'a')
			node = node.children[x]
			if node == nil {
				break
			}
			if node.val < inf && items[j+1].priority > it.priority+node.val {
				pq.update(items[j+1], it.priority+node.val)
			}
		}
	}

	res := items[n].priority
	if res >= inf {
		return -1
	}

	return res
}

type Trie struct {
	children []*Trie
	val      int
}

func NewNode() *Trie {
	tr := new(Trie)
	tr.children = make([]*Trie, 26)
	tr.val = inf
	return tr
}

func (tr *Trie) Add(s string, v int) {
	if len(s) == 0 {
		tr.val = min(tr.val, v)
		return
	}
	x := int(s[0] - 'a')
	if tr.children[x] == nil {
		tr.children[x] = NewNode()
	}
	tr.children[x].Add(s[1:], v)
}

// An Item is something we manage in a priority queue.
type Item struct {
	id       int
	priority int
	index    int
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	item.index = len(*pq)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) update(it *Item, p int) {
	it.priority = p
	heap.Fix(pq, it.index)
}
