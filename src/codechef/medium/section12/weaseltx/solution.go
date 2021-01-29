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

	E := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		E[i] = readNNums(reader, 2)
	}
	s, _ := reader.ReadBytes('\n')
	X := make([]uint64, n)
	var pos int
	for i := 0; i < n; i++ {
		pos = readUint64(s, pos, &X[i]) + 1
	}
	solver := NewSolver(n, X, E)

	var buf bytes.Buffer
	for q > 0 {
		q--
		var d uint64
		s, _ := reader.ReadBytes('\n')
		readUint64(s, 0, &d)
		res := solver.Query(d)
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

const SZ = 400040

var node [SZ]int
var next [SZ]int
var to [SZ]int
var dep [SZ]int

const H = 20

var dp [H][SZ]uint64

type Solver struct {
	dp [SZ]uint64
	N  uint64
}

func NewSolver(n int, X []uint64, E [][]int) *Solver {
	var cur int
	ins := func(u, v int) {
		cur++
		next[cur] = node[u]
		node[u] = cur
		to[cur] = v
	}

	var dfs func(p, u int)
	dfs = func(p, u int) {
		for j := node[u]; j > 0; j = next[j] {
			if to[j] != p {
				dep[to[j]] = dep[u] + 1
				dfs(u, to[j])
			}
		}
		dp[0][dep[u]] ^= X[u]
	}

	for _, e := range E {
		ins(e[0], e[1])
		ins(e[1], e[0])
	}

	dfs(-1, 0)

	var i uint

	for 1<<i <= n {
		i++
	}

	var N uint64 = 1 << i

	for ii := 0; ii < int(i); ii++ {
		for j := 0; j < int(N); j++ {
			dp[ii+1][j] = dp[ii][j^(1<<uint(ii))]
			if (j>>uint(ii))&1 == 0 {
				dp[ii+1][j] ^= dp[ii][j]
			}
		}
	}

	return &Solver{dp[i], N}
}

func (solver *Solver) Query(d uint64) uint64 {
	return solver.dp[int((d-1)&(solver.N-1))]
}
