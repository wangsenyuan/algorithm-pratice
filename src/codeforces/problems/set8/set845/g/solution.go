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

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m := readTwoNums(reader)
	E := make([][]int, m)
	for i := 0; i < m; i++ {
		E[i] = readNNums(reader, 3)
	}

	solver := NewSolver(n, E)

	fmt.Println(solver.Ask(1, n))
}

type Solver struct {
	dp    []int64
	basis [][]int64
	vis   []int
}

func NewSolver(n int, E [][]int) Solver {
	degree := make([]int, n)
	for _, e := range E {
		degree[e[0]-1]++
		degree[e[1]-1]++
	}
	adj := make([][]Pair, n)
	for i := 0; i < n; i++ {
		adj[i] = make([]Pair, 0, degree[i])
	}

	for _, e := range E {
		u, v, w := e[0], e[1], e[2]
		u--
		v--
		adj[u] = append(adj[u], Pair{v, w})
		adj[v] = append(adj[v], Pair{u, w})
	}

	dp := make([]int64, n)
	for i := 0; i < n; i++ {
		dp[i] = -1
	}

	var dfs func(u int, color int, basis []int64)

	vis := make([]int, n)
	for i := 0; i < n; i++ {
		vis[i] = -1
	}

	dfs = func(u int, color int, basis []int64) {
		vis[u] = color
		for _, e := range adj[u] {
			if dp[e.first] < 0 {
				dp[e.first] = dp[u] ^ int64(e.second)
				dfs(e.first, color, basis)
			} else {
				// a loop back
				t := dp[u] ^ dp[e.first] ^ int64(e.second)
				for t > 0 {
					x := clzll(t)
					if basis[x] > 0 {
						t ^= basis[x]
					} else {
						basis[x] = t
						break
					}
				}
			}
		}
	}

	var basis [][]int64

	for i := 0; i < n; i++ {
		if vis[i] < 0 {
			dp[i] = 0
			cur := make([]int64, 64)
			basis = append(basis, cur)
			dfs(i, len(basis), cur)
		}
	}

	return Solver{dp, basis, vis}
}

func (solver Solver) Ask(u, v int) int64 {
	u--
	v--
	if solver.vis[u] != solver.vis[v] {
		return -1
	}
	c := solver.vis[u] - 1
	res := solver.dp[u] ^ solver.dp[v]
	for i := 0; i < 64; i++ {
		if (res>>uint(63-i))&1 > 0 {
			res ^= solver.basis[c][i]
		}
	}
	return res
}

func clzll(num int64) int {
	var pos = 64
	for num > 0 {
		num >>= 1
		pos--
	}
	return pos
}

type Pair struct {
	first, second int
}
