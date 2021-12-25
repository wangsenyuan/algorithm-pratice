package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n, m := readTwoNums(reader)
		W := readNNums(reader, n)
		E := make([][]int, m)
		for i := 0; i < m; i++ {
			E[i] = readNNums(reader, 2)
		}
		solver := NewSolver(n, m, W, E)

		q := readNum(reader)

		for q > 0 {
			q--
			a, k := readTwoNums(reader)
			fmt.Println(solver.Query(a, k))
		}
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
	best []int
}

func NewSolver(n int, m int, W []int, E [][]int) Solver {
	set := make([]int, n)
	cnt := make([]int, n)

	for i := 0; i < n; i++ {
		set[i] = i
		cnt[i] = 1
	}

	D := make([]int, n)

	for i := 0; i < n; i++ {
		D[i] = digitCount(W[i]) & 1
	}

	var find func(x int) int

	find = func(x int) int {
		if set[x] != x {
			set[x] = find(set[x])
		}
		return set[x]
	}

	union := func(x, y int) {
		px := find(x)
		py := find(y)
		if px == py {
			return
		}

		if cnt[px] < cnt[py] {
			px, py = py, px
		}

		set[py] = px
		cnt[px] += cnt[py]
	}

	for _, e := range E {
		u, v := e[0]-1, e[1]-1

		if D[u] == D[v] {
			union(u, v)
		}
	}

	best := make([]int, 2)

	for i := 0; i < n; i++ {
		p := find(i)
		if cnt[p] > best[D[p]] {
			best[D[p]] = cnt[p]
		}
	}

	return Solver{best}
}

func (solver Solver) Query(kind int, K int) int {
	d := digitCount(K)

	if kind == 1 {
		// need odd
		if d&1 == 1 {
			return solver.best[0]
		}
		return solver.best[1]
	}
	// need even
	if d&1 == 1 {
		return solver.best[1]
	}
	return solver.best[0]
}

func digitCount(num int) int {
	var res int

	for num > 0 {
		res += num & 1
		num >>= 1
	}
	return res
}
