package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math/rand"
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
		heights := readNNums(reader, n)
		solver := NewSolver(n, E, heights)
		for q > 0 {
			q--
			u, v := readTwoNums(reader)
			res := solver.Query(u, v)
			if res {
				buf.WriteByte('1')
			} else {
				buf.WriteByte('0')
			}
		}
		fmt.Println(buf.String())
		buf.Reset()
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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
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
	P    [][]int
	inc  []int
	dec  []int
	D    []int
	root int
}

const H = 20

func NewSolver(n int, E [][]int, heights []int) *Solver {
	G := getGraph(n, E)

	P := make([][]int, n)
	for i := 0; i < n; i++ {
		P[i] = make([]int, H)
	}

	inc := make([]int, n)
	dec := make([]int, n)

	D := make([]int, n)

	var dfs func(p, u int)
	dfs = func(p, u int) {
		P[u][0] = p
		for i := 1; i < H; i++ {
			P[u][i] = P[P[u][i-1]][i-1]
		}
		inc[u] = u
		dec[u] = u
		if heights[p] > heights[u] {
			inc[u] = inc[p]
		} else if heights[p] < heights[u] {
			dec[u] = dec[p]
		}

		for _, v := range G[u] {
			if p != v {
				D[v] = D[u] + 1
				dfs(u, v)
			}
		}
	}

	root := rand.Intn(n)

	dfs(root, root)

	return &Solver{P, inc, dec, D, root}
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

func (solver *Solver) lca(u, v int) int {
	D := solver.D
	if D[u] < D[v] {
		u, v = v, u
	}
	P := solver.P

	for i := H - 1; i >= 0; i-- {
		if D[u]-(1<<uint(i)) >= D[v] {
			u = P[u][i]
		}
	}
	if u == v {
		return u
	}
	for i := H - 1; i >= 0; i-- {
		if P[u][i] != P[v][i] {
			u = P[u][i]
			v = P[v][i]
		}
	}
	return P[u][0]
}

func (solver *Solver) Query(u, v int) bool {
	u--
	v--
	c := solver.lca(u, v)

	inc := solver.inc
	dec := solver.dec

	if inc[u] == inc[c] {
		// not increase from u to lca
		if inc[v] == inc[c] {
			// increase from u to c, then decrease to v
			return true
		}

		if dec[inc[v]] == dec[c] {
			// increase from u to c, then increase to inv[v], then decrease to v
			return true
		}
	}

	if inc[v] == inc[c] && dec[inc[u]] == dec[c] {
		return true
	}

	return false
}
