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

	n, m := readTwoNums(scanner)
	edges := make([][]int, m)

	for i := 0; i < m; i++ {
		edges[i] = readNNums(scanner, 3)
	}

	solver := NewSolver(n, edges)

	q := readNum(scanner)

	for q > 0 {
		q--
		x := readNum(scanner)
		fmt.Println(solver.Query(x))
	}
}

type Solver struct {
	ans []int
}

const MAX_X = 100

func NewSolver(n int, edges [][]int) Solver {
	dp := make([][][]int, MAX_X+1)

	for i := 0; i <= MAX_X; i++ {
		dp[i] = make([][]int, 5)
		for j := 0; j < 5; j++ {
			dp[i][j] = make([]int, n)
		}
	}

	for _, edge := range edges {
		u, v, w := edge[0], edge[1], edge[2]
		u--
		v--
		dp[w][1][u]++
		dp[w][1][v]++
	}

	for j := 2; j < 5; j++ {
		for _, edge := range edges {
			u, v, w := edge[0], edge[1], edge[2]
			u--
			v--
			for x := 1; x <= MAX_X; x++ {
				y := max(x, w)
				dp[y][j][u] += dp[x][j-1][v]
				dp[y][j][v] += dp[x][j-1][u]
			}
		}
	}

	ans := make([]int, MAX_X+1)
	for x := 1; x <= MAX_X; x++ {
		for i := 0; i < n; i++ {
			ans[x] += dp[x][4][i]
		}
	}

	return Solver{ans}
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func (solver Solver) Query(x int) int {
	return solver.ans[x]
}
