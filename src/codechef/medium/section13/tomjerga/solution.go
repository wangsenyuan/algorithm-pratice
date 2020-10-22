package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n, q := readTwoNums(reader)
		E := make([][]int, n-1)
		for i := 0; i < n-1; i++ {
			E[i] = readNNums(reader, 2)
		}
		solver := NewSolver(n, E)

		for q > 0 {
			q--
			x, y := readTwoNums(reader)
			buf.WriteString(fmt.Sprintf("%d\n", solver.Solve(y, x)))
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

type Solver struct {
	jerry func(x, y int) int
	tom   func(x, y int) int
	dist  func(x, y int) int
	far   [][]Pair
}

func (solver Solver) Solve(x, y int) int {
	if x == y {
		return 0
	}
	x--
	y--
	jerry := solver.jerry(x, y)
	tom := solver.tom(x, y)
	ans := solver.dist(x, jerry) + (1 - solver.dist(x, y)%2)

	if solver.dist(solver.far[jerry][0].second, tom) <= 1 {
		ans += solver.far[jerry][1].first
	} else {
		ans += solver.far[jerry][0].first
	}

	return ans + 1
}

const H = 20

func NewSolver(n int, E [][]int) Solver {
	G := getGraph(n, E)

	pp := make([][]int, H)
	for i := 0; i < H; i++ {
		pp[i] = make([]int, n)
	}

	depth := make([]int, n)

	var dfs1 func(p int, u int)

	dfs1 = func(p int, u int) {
		for i := 1; i < H; i++ {
			if pp[i-1][u] >= 0 {
				pp[i][u] = pp[i-1][pp[i-1][u]]
			} else {
				pp[i][u] = -1
			}
		}

		for _, v := range G[u] {
			if p == v {
				continue
			}
			pp[0][v] = u
			depth[v] = depth[u] + 1
			dfs1(u, v)
		}
	}

	dfs1(-1, 0)

	far := make([][]Pair, n)

	var dfs2 func(p, u int)

	dfs2 = func(p, u int) {
		far[u] = make([]Pair, 2)
		for j := 0; j < 2; j++ {
			far[u][j] = Pair{0, -1}
		}
		for _, v := range G[u] {
			if p == v {
				continue
			}
			dfs2(u, v)
			if far[v][0].first+1 > far[u][0].first {
				far[u][1] = far[u][0]
				far[u][0] = Pair{far[v][0].first + 1, v}
			} else if far[v][0].first+1 > far[u][1].first {
				far[u][1] = Pair{far[v][0].first + 1, v}
			}
		}
	}

	dfs2(-1, 0)

	var dfs3 func(p int, u int, fromParent int)

	dfs3 = func(p int, u int, fromParent int) {
		if fromParent > far[u][0].first {
			far[u][1] = far[u][0]
			far[u][0] = Pair{fromParent, p}
		} else if fromParent > far[u][1].first {
			far[u][1] = Pair{fromParent, p}
		}

		for _, v := range G[u] {
			if p == v {
				continue
			}
			if far[u][0].second == v {
				dfs3(u, v, max(fromParent+1, far[u][1].first+1))
			} else if far[u][1].second == v {
				dfs3(u, v, max(fromParent+1, far[u][0].first+1))
			}
		}
	}

	dfs3(-1, 0, 0)

	lca := func(x, y int) int {
		if depth[x] < depth[y] {
			x, y = y, x
		}

		for i := H - 1; i >= 0; i-- {
			if depth[x]-(1<<uint(i)) >= depth[y] {
				x = pp[i][x]
			}
		}
		if x == y {
			return x
		}

		for i := H - 1; i >= 0; i-- {
			if pp[i][x] != pp[i][y] {
				x = pp[i][x]
				y = pp[i][y]
			}
		}
		return pp[0][x]
	}

	getDistance := func(x, y int) int {
		p := lca(x, y)
		return depth[x] + depth[y] - 2*depth[p]
	}

	getKthAncestor := func(x, k int) int {
		for i := H - 1; i >= 0; i-- {
			if k-(1<<uint(i)) >= 0 {
				x = pp[i][x]
				k -= 1 << uint(i)
			}
		}
		return x
	}

	getKthNode := func(x, y int, k int) int {
		p := lca(x, y)

		if depth[x]-depth[p]+1 >= k {
			return getKthAncestor(x, k-1)
		}
		k -= depth[x] - depth[p] + 1
		k++
		return getKthAncestor(y, depth[y]-depth[p]+1-k)
	}

	whereJerry := func(x, y int) int {
		d := getDistance(x, y)
		req := (d + 1) / 2
		return getKthNode(x, y, req)
	}

	whereTom := func(x, y int) int {
		d := getDistance(x, y)
		req := (d + 1) / 2
		return getKthNode(x, y, req+1+(1-d%2))
	}

	return Solver{whereJerry, whereTom, getDistance, far}
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

type Pair struct {
	first, second int
}

func getGraph(n int, E [][]int) [][]int {
	degree := make([]int, n)
	for _, e := range E {
		u, v := e[0]-1, e[1]-1
		degree[u]++
		degree[v]++
	}
	adj := make([][]int, n)
	for i := 0; i < n; i++ {
		adj[i] = make([]int, 0, degree[i])
	}

	for _, e := range E {
		u, v := e[0]-1, e[1]-1
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	return adj
}
