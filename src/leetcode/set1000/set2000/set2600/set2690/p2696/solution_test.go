package p2696

import (
	"container/heap"
	"testing"
)

func runSample(t *testing.T, n int, edges [][]int, source, dest int, target int, expect bool) {
	res := modifiedGraphEdges(n, edges, source, dest, target)

	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}

	if !expect {
		return
	}

	g := NewGraph(n, 2*len(res))

	for _, e := range res {
		u, v, w := e[0], e[1], e[2]
		g.AddEdge(u, v, w)
		g.AddEdge(v, u, w)
	}

	items := make([]*Item, n)
	pq := make(PriorityQueue, n)

	for i := 0; i < n; i++ {
		it := new(Item)
		it.id = i
		it.value = 1 << 30
		it.index = i
		items[i] = it
		pq[i] = it
	}

	items[source].value = 0
	heap.Init(&pq)

	for pq.Len() > 0 {
		cur := heap.Pop(&pq).(*Item)
		u := cur.id
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if g.val[i] < 0 {
				continue
			}
			if items[v].value > items[u].value+g.val[i] {
				pq.update(items[v], items[u].value+g.val[i])
			}
		}
	}

	if items[dest].value != target {
		t.Fatalf("Sample result %v, get wrong min-dist %d", res, items[dest].value)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	edges := [][]int{
		{4, 1, -1}, {2, 0, -1}, {0, 3, -1}, {4, 3, -1},
	}
	source := 0
	destination := 1
	target := 5
	expect := true
	runSample(t, n, edges, source, destination, target, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	edges := [][]int{
		{0, 1, -1}, {0, 2, 5},
	}
	source := 0
	destination := 2
	target := 6
	expect := false
	runSample(t, n, edges, source, destination, target, expect)
}

func TestSample3(t *testing.T) {
	n := 4
	edges := [][]int{
		{1, 0, 4}, {1, 2, 3}, {2, 3, 5}, {0, 3, -1},
	}
	source := 0
	destination := 2
	target := 6
	expect := true
	runSample(t, n, edges, source, destination, target, expect)
}

func TestSample4(t *testing.T) {
	n := 5
	edges := [][]int{
		{4, 1, -1}, {2, 0, -1}, {0, 3, -1}, {4, 3, -1},
	}
	source := 0
	destination := 1
	target := 6
	expect := true
	runSample(t, n, edges, source, destination, target, expect)
}
func TestSample5(t *testing.T) {
	n := 4
	edges := [][]int{
		{1, 0, 3}, {0, 2, -1}, {0, 3, 10}, {2, 3, 4}, {1, 3, 7}, {2, 1, 4},
	}
	source := 0
	destination := 2
	target := 2
	expect := true
	runSample(t, n, edges, source, destination, target, expect)
}

func TestSample6(t *testing.T) {
	n := 4
	edges := [][]int{
		{0, 1, -1}, {2, 0, 2}, {3, 2, 6}, {2, 1, 10}, {3, 0, -1},
	}
	source := 1
	destination := 3
	target := 12
	expect := true
	runSample(t, n, edges, source, destination, target, expect)
}

func TestSample7(t *testing.T) {
	n := 3
	edges := [][]int{
		{1, 0, 3}, {1, 2, 4},
	}
	source := 0
	destination := 2
	target := 8
	expect := false
	runSample(t, n, edges, source, destination, target, expect)
}

func TestSample8(t *testing.T) {
	n := 4
	edges := [][]int{
		{3, 0, -1}, {1, 2, -1}, {2, 3, -1}, {1, 3, 9}, {2, 0, 5},
	}
	source := 0
	destination := 1
	target := 7
	expect := true
	runSample(t, n, edges, source, destination, target, expect)
}

func TestSample9(t *testing.T) {
	n := 5
	edges := [][]int{
		{3, 0, -1}, {2, 1, 6}, {4, 1, -1}, {4, 3, -1},
	}
	source := 0
	destination := 1
	target := 9
	expect := true
	runSample(t, n, edges, source, destination, target, expect)
}
