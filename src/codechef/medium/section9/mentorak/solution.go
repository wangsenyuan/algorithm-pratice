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

	line := readNNums(reader, 3)
	n, m, k := line[0], line[1], line[2]

	obstacles := make([][]int, k)
	for i := 0; i < k; i++ {
		obstacles[i] = readNNums(reader, 2)
	}
	solver := NewSolver(n, m, k, obstacles)

	q := readNum(reader)

	for q > 0 {
		q--
		a, b := readTwoNums(reader)
		fmt.Println(solver.Query(a, b))
	}
}

type Solver struct {
	ans [][]int
}

func NewSolver(n int, m int, k int, obstacles [][]int) Solver {
	mat := make([][]int, n)
	for i := 0; i < n; i++ {
		mat[i] = make([]int, m)
	}

	for _, obstacle := range obstacles {
		a, b := obstacle[0], obstacle[1]
		a--
		b--
		mat[a][b] = 1
	}

	up := make([][]int, n)
	left := make([][]int, n)
	right := make([][]int, n)

	for i := 0; i < n; i++ {
		up[i] = make([]int, m)
		left[i] = make([]int, m)
		right[i] = make([]int, m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if mat[i][j] == 1 {
				up[i][j] = 0
			} else {
				up[i][j] = 1
				if i > 0 {
					up[i][j] += up[i-1][j]
				}
			}
		}

		for j := 0; j < m; j++ {
			left[i][j] = j - 1
			for left[i][j] >= 0 && up[i][left[i][j]] >= up[i][j] {
				left[i][j] = left[i][left[i][j]]
			}
		}

		for j := m - 1; j >= 0; j-- {
			right[i][j] = j + 1
			for right[i][j] < m && up[i][right[i][j]] >= up[i][j] {
				right[i][j] = right[i][right[i][j]]
			}
		}
	}

	c := make([][]int, n)
	for i := 0; i < n; i++ {
		c[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			c[i-up[i][j]+1][i] = max(c[i-up[i][j]+1][i], right[i][j]-left[i][j]-1)
		}
	}

	for k := n; k > 1; k-- {
		for l := 0; l+k <= n; l++ {
			r := l + k - 1
			if l+1 < n {
				c[l+1][r] = max(c[l+1][r], c[l][r])
			}
			if r-1 >= 0 {
				c[l][r-1] = max(c[l][r-1], c[l][r])
			}
		}
	}

	d := make([][]int, n)

	for i := 0; i < n; i++ {
		d[i] = make([]int, n)
	}

	for k := 1; k <= n; k++ {
		for l := 0; l+k <= n; l++ {
			r := l + k - 1
			d[l][r] = c[l][r] * k
			if l+1 < n {
				d[l][r] = max(d[l][r], d[l+1][r])
			}
			if r-1 >= 0 {
				d[l][r] = max(d[l][r], d[l][r-1])
			}
		}
	}

	return Solver{d}
}

func (solver Solver) Query(l, h int) int {
	l--
	h--
	return solver.ans[l][h]
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
