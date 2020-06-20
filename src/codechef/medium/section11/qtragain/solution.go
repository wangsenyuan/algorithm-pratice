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

	tc := readNum(reader)

	for tc > 0 {
		tc--

		N, Q := readTwoNums(reader)

		E := make([][]int, N-1)
		for i := 0; i < N-1; i++ {
			E[i] = readNNums(reader, 2)
		}
		solver := NewSolver(N, E)
		var t, x, y int

		for Q > 0 {
			Q--
			bs, _ := reader.ReadSlice('\n')
			pos := readInt(bs, 0, &t)
			pos = readInt(bs, pos+1, &x)
			if t == 2 {
				fmt.Println(solver.Query(x))
			} else {
				readInt(bs, pos+1, &y)
				solver.Update(x, y)
			}
		}
	}
}

type Solver struct {
	lazy   [][]int64
	value  []int64
	parent []int
	n      int
}

const H = 63
const MOD = 1000000007

var PW []int64

func init() {
	PW = make([]int64, H)
	PW[0] = 1
	for i := 1; i < H; i++ {
		PW[i] = 2 * PW[i-1]
		if PW[i] >= MOD {
			PW[i] -= MOD
		}
	}
}

func NewSolver(n int, E [][]int) Solver {
	lazy := make([][]int64, n+1)
	for i := 0; i <= n; i++ {
		lazy[i] = make([]int64, H)
	}
	value := make([]int64, n+1)

	degree := make([]int, n+1)

	for _, e := range E {
		u, v := e[0], e[1]
		degree[u]++
		degree[v]++
	}

	adj := make([][]int, n+1)

	for i := 0; i <= n; i++ {
		adj[i] = make([]int, 0, degree[i])
	}

	for _, e := range E {
		u, v := e[0], e[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	parent := make([]int, n+1)
	var dfs func(p, u int)

	dfs = func(p, u int) {
		parent[u] = p
		for _, v := range adj[u] {
			if v == p {
				continue
			}
			dfs(u, v)
		}
	}

	dfs(-1, 1)

	return Solver{lazy, value, parent, n}
}

func (solver *Solver) Update(x int, y int) {
	px, py := x, y

	for px > 0 && py >= 0 {
		solver.value[px] += PW[py]
		px = solver.parent[px]
		py--
	}

	for i := 1; i <= y; i++ {
		solver.lazy[x][i] += PW[y-i]
	}
}

func (solver Solver) Query(x int) int {
	var res int64 = solver.value[x]

	px, py := solver.parent[x], 1

	for px > 0 && py < H {
		res += solver.lazy[px][py]
		if res >= MOD {
			res -= MOD
		}
		px = solver.parent[px]
		py++
	}

	return int(res)
}
