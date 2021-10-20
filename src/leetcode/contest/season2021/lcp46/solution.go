package lcp46

func volunteerDeployment(finalCnt []int, totalNum int64, edges [][]int, plans [][]int) []int {
	n := len(finalCnt) + 1

	g := NewGraph(n, len(edges)*2+2)
	for _, e := range edges {
		u, v := e[0], e[1]
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	A := make([]int, n)
	copy(A[1:], finalCnt)
	B := make([]int, n)
	B[0] = 1

	for i := len(plans) - 1; i >= 0; i-- {
		plan := plans[i]
		p := plan[1]
		if plan[0] == 1 {
			A[p] *= 2
			B[p] *= 2
		} else if plan[0] == 2 {
			for j := g.nodes[p]; j > 0; j = g.next[j] {
				v := g.to[j]
				A[v] -= A[p]
				B[v] -= B[p]
			}
		} else {
			for j := g.nodes[p]; j > 0; j = g.next[j] {
				v := g.to[j]
				A[v] += A[p]
				B[v] += B[p]
			}
		}
	}

	// A[0] * B[0] + A[1] * B[1] ==== totalsum
	var x int64
	for i := 0; i < n; i++ {
		x += int64(B[i])
		totalNum -= int64(A[i])
	}

	o := int(totalNum / x)
	for i := 0; i < n; i++ {
		A[i] += o * B[i]
	}

	return A
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
	next := make([]int, e)
	to := make([]int, e)
	return &Graph{nodes, next, to, 0}
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}
