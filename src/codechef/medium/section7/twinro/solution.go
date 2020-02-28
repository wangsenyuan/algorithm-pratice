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

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	n := readNum(scanner)
	board := make([][]int, n)

	for i := 0; i < n; i++ {
		board[i] = readNNums(scanner, n)
	}

	fmt.Println(solve(n, board))
}

func solve(N int, board [][]int) int64 {

	dp := make([][]int64, N)
	for i := 0; i < N; i++ {
		dp[i] = make([]int64, N)
	}

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			value := int64(board[i][j]) + int64(board[j][N-1-i])

			if i == 0 && j == 0 {
				dp[i][j] = value
			} else if i == 0 {
				dp[i][j] = dp[i][j-1] + value
			} else if j == 0 {
				dp[i][j] = dp[i-1][j] + value
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i-1][j]) + value
			}
		}
	}

	return dp[N-1][N-1]
}

func max(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}
