package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)

	E := make([][]int, n)

	for i := 0; i < n-1; i++ {
		E[i] = readNNums(reader, 2)
	}

	solver := NewSolver(n, E)

	q := readNum(reader)

	for q > 0 {
		q--
		u, v := readTwoNums(reader)
		fmt.Println(solver.InsertPath(u, v))
	}
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
	on, under *Fenwick
	pp        [][]int
	n         int
	tin       []int
	tout      []int
	in        []int
}

const H = 20

func NewSolver(n int, E [][]int) Solver {

	degree := make([]int, n)

	for _, e := range E {
		u, v := e[0], e[1]
		u--
		v--
		degree[u]++
		degree[v]++
	}

	adj := make([][]int, n)
	for i := 0; i < n; i++ {
		adj[i] = make([]int, 0, degree[i])
	}

	for _, e := range E {
		u, v := e[0], e[1]
		u--
		v--
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	tin := make([]int, n)
	tout := make([]int, n)
	var time int

	pp := make([][]int, n)
	for i := 0; i < n; i++ {
		pp[i] = make([]int, H)
		for j := 0; j < H; j++ {
			pp[i][j] = -1
		}
	}

	var dfs func(p, u int)

	dfs = func(p, u int) {
		tin[u] = time
		time++
		pp[u][0] = p

		for _, v := range adj[u] {
			if p == v {
				continue
			}
			dfs(u, v)
		}
		tout[u] = time
		time++
	}

	dfs(-1, 0)

	for j := 1; j < H; j++ {
		for i := 0; i < n; i++ {
			p := pp[i][j-1]
			if p >= 0 {
				pp[i][j] = pp[p][j-1]
			}
		}
	}

	on := NewFenwick(time)
	under := NewFenwick(time)

	in := make([]int, n)

	return Solver{&on, &under, pp, n, tin, tout, in}
}

func isAnc(in, out []int, u int, v int) bool {
	if u <= 0 {
		return true
	}
	return in[u] <= in[v] && out[v] <= out[u]
}

func lca(pp [][]int, in []int, out []int, u, v int) int {
	if isAnc(in, out, u, v) {
		return u
	}
	if isAnc(in, out, v, u) {
		return v
	}

	for j := H - 1; j >= 0; j-- {
		if !isAnc(in, out, pp[u][j], v) {
			u = pp[u][j]
		}
	}
	return pp[u][0]
}

func (solver *Solver) InsertPath(u, v int) int {
	u--
	v--

	g := lca(solver.pp, solver.tin, solver.tout, u, v)

	on, under := solver.on, solver.under
	tin, tout := solver.tin, solver.tout
	in := solver.in

	res := on.Sum(tin[u]) + on.Sum(tin[v]) - 2*on.Sum(tin[g])
	res += under.Sum(tout[g]) - under.Sum(tin[g]-1)
	res += in[g]

	on.Update(tin[g], 1)
	on.Update(tout[g], -1)

	under.Update(tin[u], 1)
	under.Update(tin[v], 1)
	under.Update(tin[g], -2)

	in[g]++

	return res
}

type Fenwick struct {
	arr []int
	n   int
}

func NewFenwick(n int) Fenwick {
	arr := make([]int, n+1)
	return Fenwick{arr, n}
}

func (fen *Fenwick) Update(p int, v int) {
	p++
	for p <= fen.n {
		fen.arr[p] += v
		p += p & (-p)
	}
}

func (fen Fenwick) Sum(p int) int {
	p++
	var res int
	for p > 0 {
		res += fen.arr[p]
		p -= p & -p
	}
	return res
}
