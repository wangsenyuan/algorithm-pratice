package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(scanner *bufio.Scanner) (a int) {
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &a)
	return
}

func readTwoNums(scanner *bufio.Scanner) (a int, b int) {
	res := readNNums(scanner, 2)
	a, b = res[0], res[1]
	return
}

func readNNums(scanner *bufio.Scanner, n int) []int {
	res := make([]int, n)
	x := 0
	scanner.Scan()
	for i := 0; i < n; i++ {
		for x < len(scanner.Bytes()) && scanner.Bytes()[x] == ' ' {
			x++
		}
		x = readInt(scanner.Bytes(), x, &res[i])
	}
	return res
}

func fillNNums(scanner *bufio.Scanner, n int, res []int) {
	x := 0
	scanner.Scan()
	for i := 0; i < n; i++ {
		for x < len(scanner.Bytes()) && scanner.Bytes()[x] == ' ' {
			x++
		}
		x = readInt(scanner.Bytes(), x, &res[i])
	}
}

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var tmp int64
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp
	return i
}

func readNInt64Nums(scanner *bufio.Scanner, n int) []int64 {
	res := make([]int64, n)
	x := -1
	scanner.Scan()
	for i := 0; i < n; i++ {
		x = readInt64(scanner.Bytes(), x+1, &res[i])
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		n, q := readTwoNums(scanner)
		edges := make([][]int, n-1)
		for i := 0; i < n-1; i++ {
			edges[i] = readNNums(scanner, 2)
		}

		solver := NewSolver(n, edges)

		for q > 0 {
			q--
			ques := readNNums(scanner, 4)
			res := solver.Query(ques[0], ques[1], ques[2], ques[3])
			fmt.Println(res)
		}

	}
}

type Solver struct {
	dp     [][]Item
	parent [][]int
	level  []int
}

const MAX_H = 20

func NewSolver(n int, edges [][]int) Solver {
	conns := make([][]int, n)

	for i := 0; i < n; i++ {
		conns[i] = make([]int, 0, 10)
	}

	for _, edge := range edges {
		u, v := edge[0]-1, edge[1]-1
		conns[u] = append(conns[u], v)
		conns[v] = append(conns[v], u)
	}

	parent := make([][]int, n)

	for i := 0; i < n; i++ {
		parent[i] = make([]int, MAX_H)
	}

	level := make([]int, n)

	var dfs1 func(p, u int)

	dfs1 = func(p, u int) {
		parent[u][0] = p

		if p >= 0 {
			level[u] = level[p] + 1
		}

		for _, v := range conns[u] {
			if p == v {
				continue
			}
			dfs1(u, v)
		}
	}

	dfs1(-1, 0)

	for h := 1; h < MAX_H; h++ {
		for i := 0; i < n; i++ {
			p := parent[i][h-1]
			if p >= 0 {
				parent[i][h] = parent[p][h-1]
			}
		}
	}

	dp := make([][]Item, n)

	for i := 0; i < n; i++ {
		dp[i] = make([]Item, 3)
	}

	var dfs2 func(p, u int)

	dfs2 = func(p, u int) {
		dp[u][0] = Item{-1, -1}
		dp[u][1] = Item{-1, -1}
		dp[u][2] = Item{-1, -1}

		if len(conns[u]) == 1 && p >= 0 {
			dp[u][0] = Item{0, u}
			return
		}

		for _, v := range conns[u] {
			if p == v {
				continue
			}
			dfs2(u, v)
			dp[u][2] = max(dp[u][2], dp[v][0].inc())
			if dp[u][2].dist > dp[u][1].dist {
				dp[u][1], dp[u][2] = dp[u][2], dp[u][1]
			}
			if dp[u][1].dist > dp[u][0].dist {
				dp[u][0], dp[u][1] = dp[u][1], dp[u][0]
			}
		}
	}

	dfs2(-1, 0)

	var dfs3 func(p, u int, toPar Item)

	dfs3 = func(p, u int, toPar Item) {
		for _, v := range conns[u] {
			if p == v {
				continue
			}

			if dp[u][0].index == v {
				dfs3(u, v, max(dp[u][1], toPar).inc())
			} else {
				dfs3(u, v, max(dp[u][0], toPar).inc())
			}
		}

		dp[u][2] = max(dp[u][2], toPar)
		if dp[u][2].dist > dp[u][1].dist {
			dp[u][1], dp[u][2] = dp[u][2], dp[u][1]
		}
		if dp[u][1].dist > dp[u][0].dist {
			dp[u][0], dp[u][1] = dp[u][1], dp[u][0]
		}
	}

	dfs3(-1, 0, Item{0, 0})

	return Solver{dp, parent, level}
}

func (solver *Solver) Query(x, dx, y, dy int) int {
	x--
	y--
	ddx := dx
	ddy := dy

	level := solver.level
	parent := solver.parent

	lca := func(x, y int) int {
		if level[x] > level[y] {
			x, y = y, x
		}

		for h := MAX_H - 1; h >= 0; h-- {
			if level[y]-(1<<uint(h)) >= level[x] {
				y = parent[y][h]
			}
		}

		if x == y {
			return x
		}

		for h := MAX_H - 1; h >= 0; h-- {
			if parent[x][h] != parent[y][h] {
				x = parent[x][h]
				y = parent[y][h]
			}
		}
		return parent[x][0]
	}

	distance := func(x, y int) int {
		ll := lca(x, y)
		return level[x] + level[y] - 2*level[ll]
	}

	// ll := lca(x, y)
	dis := distance(x, y)

	if (dis & 1) != ((dx + dy) & 1) {
		return -1
	}

	jump := func(x int, K int) int {
		node := x

		for j := 0; j < MAX_H; j++ {
			if K&(1<<uint(j)) > 0 {
				node = parent[node][j]
			}
		}
		return node
	}

	Kth := func(x, y, K int) int {
		ll := lca(x, y)

		a := level[x] - level[ll] + 1
		if a >= K {
			return jump(x, K-1)
		}
		K -= a
		return jump(y, level[y]-level[ll]-K)
	}

	center := Kth(x, y, (dis+dx-dy)/2+1)

	dcx := distance(x, center)
	dcy := distance(y, center)
	dx -= dcx
	dy -= dcy

	if dx != dy || dx < 0 {
		return -1
	}

	if dx == 0 {
		return center + 1
	}

	dp := solver.dp

	for k := 0; k < 3; k++ {
		other := dp[center][k].index
		dd := dp[center][k].dist

		if other < 0 || dd < dx {
			continue
		}

		tmp := Kth(other, center, dd-dx+1)

		if distance(tmp, x) == ddx && distance(tmp, y) == ddy {
			return tmp + 1
		}
	}

	return -1
}

type Item struct {
	dist  int
	index int
}

func (item Item) inc() Item {
	return Item{item.dist + 1, item.index}
}

func max(a, b Item) Item {
	if a.dist >= b.dist {
		return a
	}
	return b
}
