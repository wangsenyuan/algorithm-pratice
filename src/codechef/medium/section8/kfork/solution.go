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

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n, m := readTwoNums(reader)

		queens := make([][]int, m)

		for i := 0; i < m; i++ {
			queens[i] = readNNums(reader, 2)
		}

		fmt.Println(solve(n, m, queens))
	}
}

func solve(n int, m int, queens [][]int) int {
	board := make([][]int, n)

	for i := 0; i < n; i++ {
		board[i] = make([]int, n)
	}

	row := make([]bool, n)
	col := make([]bool, n)
	diag1 := make([]bool, 2002)
	diag2 := make([]bool, 2002)

	getDiag1 := func(x, y int) int {
		return x + y
	}

	getDiag2 := func(x, y int) int {
		return x - y + 1001
	}

	for _, queen := range queens {
		a, b := queen[0], queen[1]
		a--
		b--
		board[a][b] = 1
		row[a] = true
		col[b] = true
		diag1[getDiag1(a, b)] = true
		diag2[getDiag2(a, b)] = true
	}

	get := func(x, y int) int {
		if x < 0 || x >= n || y < 0 || y >= n {
			return 0
		}
		return board[x][y]
	}

	safe := func(x, y int) bool {
		if row[x] || col[y] || diag1[getDiag1(x, y)] || diag2[getDiag2(x, y)] {
			return false
		}
		return true
	}

	var res int

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			// is pos (i, j) a fork position
			if !safe(i, j) {
				continue
			}
			count := get(i+1, j+2) + get(i+1, j-2) + get(i-1, j+2) + get(i-1, j-2) +
				get(i+2, j+1) + get(i+2, j-1) + get(i-2, j+1) + get(i-2, j-1)

			if count >= 2 {
				res++
			}
		}
	}

	return res
}
