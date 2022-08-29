package p2392

func buildMatrix(k int, rowConditions [][]int, colConditions [][]int) [][]int {
	// row[abovei] < row[belowi]
	// col[lefti] < col[righti]
	// k <= 400
	gr := NewGraph(k, len(rowConditions))

	for i := 0; i < len(rowConditions); i++ {
		a, b := rowConditions[i][0], rowConditions[i][1]
		a--
		b--
		gr.AddEdge(a, b)
	}

	ok, row := topoSort(k, gr)

	if !ok {
		return nil
	}

	pr := make([]int, k)

	for i := 0; i < len(row); i++ {
		pr[row[i]] = i + 1
	}

	gc := NewGraph(k, len(colConditions))

	for i := 0; i < len(colConditions); i++ {
		l, r := colConditions[i][0], colConditions[i][1]
		l--
		r--
		gc.AddEdge(l, r)
	}

	ko, col := topoSort(k, gc)

	if !ko {
		return nil
	}

	res := make([][]int, k)

	for i := 0; i < k; i++ {
		res[i] = make([]int, k)
	}

	m := len(row)

	for j := 0; j < len(col); j++ {
		x := col[j]
		if pr[x] > 0 {
			res[pr[x]-1][j] = x + 1
			pr[x] = -1
		} else {
			res[m][j] = x + 1
		}
	}
	m = len(col)
	for i := 0; i < len(row); i++ {
		x := row[i]
		if pr[x] > 0 {
			res[i][m] = x + 1
			pr[x] = -1
		}
	}
	// 如果没有处理完，最后一行或者最后一列肯定没有放置数字
	p := k - 1
	for x := 0; x < k; x++ {
		if pr[x] == 0 {
			// those not placed
			res[p][k-1] = x + 1
			p--
		}
	}

	return res
}

func topoSort(n int, g *Graph) (bool, []int) {

	vis := make([]int, n)

	in_path := make([]bool, n)

	var dfs func(u int, flag int) bool

	stack := make([]int, 0, n)

	dfs = func(u int, flag int) bool {
		if in_path[u] {
			// a loop
			return false
		}
		if vis[u] > 0 {
			// already visited
			return true
		}
		in_path[u] = true
		vis[u] = flag

		for i := g.node[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if !dfs(v, flag) {
				return false
			}
		}
		stack = append(stack, u)
		in_path[u] = false
		return true
	}

	for u := 0; u < n; u++ {
		if vis[u] == 0 {
			if !dfs(u, u+1) {
				return false, nil
			}
		}
	}

	reverse(stack)

	return true, stack

}

type Graph struct {
	node []int
	next []int
	to   []int
	cur  int
}

func NewGraph(n int, e int) *Graph {
	g := new(Graph)
	g.node = make([]int, n)
	g.to = make([]int, e+1)
	g.next = make([]int, e+1)
	g.cur = 0
	return g
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.node[u]
	g.node[u] = g.cur

	g.to[g.cur] = v
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
