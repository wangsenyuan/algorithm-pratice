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

func fillNNums(bs []byte, n int, res []int) {
	x := 0
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	N := readNum(reader)

	G := make([][]int, N)

	for i := 0; i < N; i++ {
		G[i] = readNNums(reader, N)
	}

	solver := NewSolver(N, G)

	m := readNum(reader)

	for m > 0 {
		m--
		k, x := readTwoNums(reader)

		res := solver.Solve(k, x)

		fmt.Println(len(res))
		if len(res) == 0 {
			fmt.Println("-1")
		} else {
			for i := 0; i < len(res); i++ {
				fmt.Printf("%d ", res[i])
			}
			fmt.Println()
		}
	}
}

type Solver struct {
	N int
	G [][][]int
}

const K = 1000000001

func NewSolver(N int, G [][]int) Solver {
	var h int
	for 1<<uint(h) <= K {
		h++
	}

	gg := make([][][]int, h)
	gg[0] = G

	for i := 1; i < h; i++ {
		gg[i] = multipleMatrix(gg[i-1], gg[i-1], N)
	}

	return Solver{N, gg}
}

func (solver Solver) Solve(k int, x int) []int {
	x--

	res := UnitMatrix(solver.N)

	var i int

	for k > 0 {
		if k&1 == 1 {
			res = multipleMatrix(res, solver.G[i], solver.N)
		}
		i++
		k >>= 1
	}
	ans := make([]int, 0, solver.N)

	for j := 0; j < solver.N; j++ {
		if res[x][j] == 1 {
			ans = append(ans, j+1)
		}
	}
	return ans
}

func UnitMatrix(n int) [][]int {
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
		res[i][i] = 1
	}
	return res
}

func multipleMatrix(a, b [][]int, n int) [][]int {
	c := make([][]int, n)
	for i := 0; i < n; i++ {
		c[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				c[i][j] += a[i][k] * b[k][j]
				if c[i][j] > 0 {
					break
				}
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if c[i][j] > 0 {
				c[i][j] = 1
			}
		}
	}
	return c
}
