package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, q := readTwoNums(reader)
	W := readNNums(reader, n)
	E := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		E[i] = readNNums(reader, 2)
	}
	solver := NewSolver(n, E, W)

	var buf bytes.Buffer

	for q > 0 {
		q--
		u, v := readTwoNums(reader)
		res := solver.Query(u, v)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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
	n       int
	W       []uint
	dot     []uint
	level   []int
	verList [][]int
	parent  []int
	off     []int
	id      []int
	mem     []bool
	ans     []uint
	toAdd   []int
}

const P = 40

func NewSolver(n int, E [][]int, w []int) *Solver {
	W := make([]uint, n)
	for i := 0; i < n; i++ {
		W[i] = uint(w[i])
	}
	G := getGraph(n, E)

	parent := make([]int, n)
	level := make([]int, n)
	levelCnt := make([]int, n)

	dot := make([]uint, n)

	var dfs func(p, u int) int

	dfs = func(p, u int) int {
		levelCnt[level[u]]++
		parent[u] = p
		var maxDepth int
		dot[u] += W[u] * W[u]

		for _, v := range G[u] {
			if p == v {
				continue
			}
			dot[v] = dot[u]
			level[v] = level[u] + 1
			tmp := dfs(u, v)
			if tmp > maxDepth {
				maxDepth = tmp
			}
		}
		return maxDepth + 1
	}

	d := dfs(-1, 0)
	verList := make([][]int, d)

	for i := 0; i < d; i++ {
		verList[i] = make([]int, 0, levelCnt[i])
	}
	for i := 0; i < n; i++ {
		verList[level[i]] = append(verList[level[i]], i)
	}

	off := make([]int, d)

	id := make([]int, n)

	var offCount int

	for i := 0; i < d; i += P {
		var best = i
		for j := 0; j < P && i+j < d; j++ {
			if len(verList[i+j]) < len(verList[best]) {
				best = i + j
			}
		}

		var t int

		for _, v := range verList[best] {
			id[v] = t
			t++
		}

		off[best] = offCount
		offCount += t * (t - 1) / 2
	}

	return &Solver{n, W, dot, level, verList, parent, off, id, make([]bool, n), make([]uint, n), make([]int, d)}
}

func (solver *Solver) Query(u, v int) uint {
	u--
	v--
	var ret uint
	var it int

	for u != v {
		curLevel := solver.level[u]

		if solver.off[curLevel] > 0 {
			pu, pv := solver.id[u], solver.id[v]
			if pu < pv {
				pu, pv = pv, pu
			}
			//size := len(solver.verList[curLevel])
			place := solver.off[curLevel] + pu*(pu-1)/2 + pv

			if solver.mem[place] {
				ret += solver.ans[place]
				for i := 0; i < it; i++ {
					solver.ans[solver.toAdd[i]] += ret
				}
				return ret
			}
			solver.mem[place] = true
			solver.ans[place] = -ret
			solver.toAdd[it] = place
			it++
			ret += solver.W[u] * solver.W[v]
		} else {
			ret += solver.W[u] * solver.W[v]
		}
		u = solver.parent[u]
		v = solver.parent[v]
	}
	ret += solver.dot[u]
	for i := 0; i < it; i++ {
		solver.ans[solver.toAdd[i]] += ret
	}
	return ret
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
