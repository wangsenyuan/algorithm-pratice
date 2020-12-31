package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	tc := readNum(reader)
	var buf bytes.Buffer
	for tc > 0 {
		tc--
		n, m := readTwoNums(reader)
		mat := make([][]int, n)
		for i := 0; i < n; i++ {
			mat[i] = readNNums(reader, m)
		}
		solver := NewSolver(n, m, mat)

		q := readNum(reader)

		for q > 0 {
			q--
			x, y, v := readThreeNums(reader)
			res := solver.Change(x, y, v)

			if res {
				buf.WriteString("Yes\n")
			} else {
				buf.WriteString("No\n")
			}
		}
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
	n, m  int
	mat   [][]int
	count int
}

func NewSolver(n, m int, mat [][]int) *Solver {
	cur := make([]bool, n+m)

	for i := 0; i < n+m; i++ {
		cur[i] = true
	}
	var count int
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if mat[i][j] != mat[i-1][j-1] {
				count++
			}
		}
	}

	return &Solver{n, m, mat, count}
}

func (this *Solver) Change(x, y, v int) bool {
	x--
	y--
	if min(x, y) > 0 && this.mat[x][y] != this.mat[x-1][y-1] {
		this.count--
	}
	if x < this.n-1 && y < this.m-1 && this.mat[x][y] != this.mat[x+1][y+1] {
		this.count--
	}

	this.mat[x][y] = v

	if min(x, y) > 0 && this.mat[x][y] != this.mat[x-1][y-1] {
		this.count++
	}
	if x < this.n-1 && y < this.m-1 && this.mat[x][y] != this.mat[x+1][y+1] {
		this.count++
	}

	return this.count == 0
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
