package p1765

const H = 51

var coprimes [H][]int

func init() {
	for i := 1; i < H; i++ {
		coprimes[i] = make([]int, 0, H)
		for j := 1; j < H; j++ {
			if gcd(i, j) == 1 {
				coprimes[i] = append(coprimes[i], j)
			}
		}
	}
}

func getCoprimes(nums []int, edges [][]int) []int {
	n := len(edges) + 1
	pos := make([]int, H)
	ii := make([][]int, H)
	for i := 1; i < H; i++ {
		ii[i] = make([]int, n)
	}

	g := NewGraph(n, 2*n)

	for _, e := range edges {
		u, v := e[0], e[1]
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	ans := make([]int, n)
	ans[0] = -1
	D := make([]int, n)
	var dfs func(p, u int)
	dfs = func(p, u int) {
		ans[u] = -1
		num := nums[u]
		for _, x := range coprimes[num] {
			if pos[x] > 0 {
				j := ii[x][pos[x]-1]
				if ans[u] < 0 || D[ans[u]] < D[j] {
					ans[u] = j
				}
			}
		}
		ii[num][pos[num]] = u
		pos[num]++

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				D[v] = D[u] + 1
				dfs(u, v)
			}
		}
		pos[num]--
	}

	dfs(-1, 0)

	return ans
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

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
