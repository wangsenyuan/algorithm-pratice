package p2642

type Graph struct {
	adj [][]int
}

func Constructor(n int, edges [][]int) Graph {
	adj := make([][]int, n)
	for i := 0; i < n; i++ {
		adj[i] = make([]int, n)
		for j := 0; j < n; j++ {
			adj[i][j] = -1
		}
		adj[i][i] = 0
	}

	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		adj[u][v] = w
	}

	for k := 0; k < n; k++ {
		for u := 0; u < n; u++ {
			for v := 0; v < n; v++ {
				if u == v {
					continue
				}
				a := adj[u][k]
				b := adj[k][v]
				if a >= 0 && b >= 0 && (adj[u][v] < 0 || adj[u][v] > a+b) {
					adj[u][v] = a + b
				}
			}
		}
	}

	g := new(Graph)
	g.adj = adj
	return *g
}

func (this *Graph) AddEdge(edge []int) {
	x, y, w := edge[0], edge[1], edge[2]
	n := len(this.adj)
	for u := 0; u < n; u++ {
		for v := 0; v < n; v++ {
			a := this.adj[u][x]
			b := this.adj[y][v]
			if a >= 0 && b >= 0 && (this.adj[u][v] < 0 || this.adj[u][v] > a+b+w) {
				this.adj[u][v] = a + b + w
			}
		}
	}
}

func (this *Graph) ShortestPath(node1 int, node2 int) int {
	return this.adj[node1][node2]
}

/**
 * Your Graph object will be instantiated and called as such:
 * obj := Constructor(n, edges);
 * obj.AddEdge(edge);
 * param_2 := obj.ShortestPath(node1,node2);
 */
