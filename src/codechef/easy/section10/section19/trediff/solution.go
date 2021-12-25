package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	for tc > 0 {
		tc--
		N, Q := readTwoNums(reader)
		A := readNNums(reader, N)
		E := make([][]int, N-1)
		for i := 0; i < N-1; i++ {
			E[i] = readNNums(reader, 2)
		}
		solver := NewSolver(N, E, A)

		for Q > 0 {
			Q--
			a, b := readTwoNums(reader)
			fmt.Println(solver.Query(a, b))
		}
	}
}

const MAX_A = 100

type Solver struct {
	n  int
	A  []int
	PP [][]int
	D  []int
}

func NewSolver(n int, E [][]int, A []int) *Solver {
	var h int
	for 1<<uint(h) <= n {
		h++
	}

	PP := make([][]int, h)

	for i := 0; i < h; i++ {
		PP[i] = make([]int, n)
	}

	conns := make([][]int, n)
	for i := 0; i < n; i++ {
		conns[i] = make([]int, 0, 3)
	}

	for _, e := range E {
		u, v := e[0], e[1]
		u--
		v--
		conns[u] = append(conns[u], v)
		conns[v] = append(conns[v], u)
	}

	D := make([]int, n)

	var dfs func(p, u int)

	dfs = func(p, u int) {
		PP[0][u] = p
		for _, v := range conns[u] {
			if v == p {
				continue
			}
			D[v] = D[u] + 1
			dfs(u, v)
		}
	}

	dfs(-1, 0)

	for i := 1; i < h; i++ {
		for j := 0; j < n; j++ {
			k := PP[i-1][j]
			if k >= 0 {
				PP[i][j] = PP[i-1][k]
			}
		}
	}

	return &Solver{n, A, PP, D}
}

func (solver *Solver) FindLca(a, b int) int {
	D := solver.D
	PP := solver.PP
	if D[a] < D[b] {
		a, b = b, a
	}

	for i := len(PP) - 1; i >= 0; i-- {
		if D[a]-(1<<uint(i)) >= D[b] {
			a = PP[i][a]
		}
	}

	if a == b {
		return a
	}

	for i := len(PP) - 1; i >= 0; i-- {
		if PP[i][a] != PP[i][b] {
			a = PP[i][a]
			b = PP[i][b]
		}
	}
	return PP[0][a]
}

func (solver *Solver) Query(a, b int) int {
	a--
	b--
	c := solver.FindLca(a, b)
	D := solver.D
	n := D[a] + D[b] - 2*D[c] + 1

	if n > MAX_A {
		return 0
	}

	x := make([]int, 0, n)

	PP := solver.PP
	A := solver.A

	for a != c {
		x = append(x, A[a])
		a = PP[0][a]
	}

	for b != c {
		x = append(x, A[b])
		b = PP[0][b]
	}

	x = append(x, A[c])

	return FindMinDiff(x)
}

func FindMinDiff(arr []int) int {
	sort.Ints(arr)

	res := arr[1] - arr[0]

	for i := 2; i < len(arr); i++ {
		tmp := arr[i] - arr[i-1]
		if tmp < res {
			res = tmp
		}
	}
	return res
}
