package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)

	solver := NewSolver(n)

	str := toString(solver.board)

	fmt.Print(str)

	q := readNum(reader)

	for q > 0 {
		q--
		var x uint64
		s, _ := reader.ReadBytes('\n')
		readUint64(s, 0, &x)
		ans := solver.Query(int64(x))

		fmt.Print(ans)
	}
}

func toString(arr [][]int64) string {
	var buf bytes.Buffer
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			buf.WriteString(fmt.Sprintf("%d ", arr[i][j]))
		}
		buf.WriteByte('\n')
	}
	return buf.String()
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
	n     int
	board [][]int64
}

func NewSolver(n int) Solver {
	board := make([][]int64, n)

	for i := 0; i < n; i++ {
		board[i] = make([]int64, n)

		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if i == n-1 && j == n-1 {
				continue
			}

			if i%2 == 0 {
				board[i][j] = int64(1) << (i + j)
			}
		}
	}

	return Solver{n, board}
}

func (solver Solver) Query(q int64) string {
	var buf bytes.Buffer

	buf.WriteString(fmt.Sprintf("%d %d\n", 1, 1))

	n := solver.n
	B := solver.board
	var x, y int

	for x < n-1 || y < n-1 {
		if x == n-1 {
			y++
		} else if y == n-1 {
			x++
		} else {
			and := max(B[x+1][y], B[x][y+1])

			if q&and > 0 != (B[x+1][y] > 0) {
				y++
			} else {
				x++
			}
		}
		buf.WriteString(fmt.Sprintf("%d %d\n", x+1, y+1))
	}

	return buf.String()
}

func max(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}
