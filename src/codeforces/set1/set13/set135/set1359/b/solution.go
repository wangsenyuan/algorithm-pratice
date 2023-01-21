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
		line := readNNums(reader, 4)
		n, m, x, y := line[0], line[1], line[2], line[3]
		B := make([][]byte, n)
		for i := 0; i < n; i++ {
			B[i], _ = reader.ReadBytes('\n')
		}
		fmt.Println(solve(n, m, B, x, y))
	}
}

func solve(n, m int, board [][]byte, x, y int) int64 {
	// row independent with each other
	X, Y := int64(x), int64(y)
	find := func(r int) int64 {
		// 前两个cell的cost
		var A, B int64

		for j := 0; j < m; j++ {
			if board[r][j] == '*' {
				A = B
				continue
			}
			if j == 0 || board[r][j-1] == '*' {
				// 只能使用x
				A, B = B, B+X
				continue
			}
			//也能使用y
			A, B = B, min(A+Y, B+X)
		}
		return B
	}

	var res int64

	for i := 0; i < n; i++ {
		res += find(i)
	}

	return res
}

func min(a, b int64) int64 {
	if a <= b {
		return a
	}
	return b
}
