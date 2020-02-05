package main

import (
	"bufio"
	"fmt"
	"os"
)

const MAX_W = 1000

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

	firstLine := readNNums(scanner, 3)
	n, s, m := firstLine[0], firstLine[1], firstLine[2]

	R := readNNums(scanner, n-1)
	W := readNNums(scanner, n)
	S := readNNums(scanner, s)
	solver := NewSolver(n, s, R, W, S)

	for m > 0 {
		m--
		line := readNNums(scanner, 3)
		res := solver.Query(line[0], line[1], line[2])
		fmt.Println(res)
	}

}

type Solver struct {
	can [][][]bool
	s   int
}

func NewSolver(n, s int, R []int, W []int, S []int) Solver {
	can := make([][][]bool, s)

	for i := 0; i < s; i++ {
		can[i] = make([][]bool, n+1)
		for j := 0; j <= n; j++ {
			can[i][j] = make([]bool, MAX_W+1)
		}
	}

	conns := make([][]int, n+1)

	for i := 0; i <= n; i++ {
		conns[i] = make([]int, 0, 10)
	}

	for i := 1; i < n; i++ {
		v := R[i-1]
		conns[v] = append(conns[v], i+1)
		conns[i+1] = append(conns[i+1], v)
	}

	var dfs func(j int, u int, p int)

	dfs = func(j int, u int, p int) {
		for w := 0; w <= MAX_W; w++ {
			if can[j][p][w] {
				can[j][u][w] = true
				if w+W[u-1] <= MAX_W {
					can[j][u][w+W[u-1]] = true
				}
			}
		}
		for _, v := range conns[u] {
			if p == v {
				continue
			}
			dfs(j, v, u)
		}
	}

	for j := 0; j < s; j++ {
		can[j][0][0] = true
		dfs(j, S[j], 0)
	}

	return Solver{can, s}
}

func (solver Solver) Query(u, v, w int) int {
	can := solver.can

	var ans int

	for j := 0; j < solver.s; j++ {
		for i := w; i >= 0; i-- {
			if can[j][u][i] && can[j][v][i] {
				ans = max(ans, 2*i)
				break
			}
		}
	}

	return ans
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
