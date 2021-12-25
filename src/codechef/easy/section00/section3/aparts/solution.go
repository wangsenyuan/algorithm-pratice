package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	tmp := 0
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp
	return i
}

func readNum(scanner *bufio.Scanner) (a int) {
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &a)
	return
}

func readTwoNums(scanner *bufio.Scanner) (a int, b int) {
	scanner.Scan()
	x := readInt(scanner.Bytes(), 0, &a)
	readInt(scanner.Bytes(), x+1, &b)
	return
}

func readNNums(scanner *bufio.Scanner, n int) []int {
	res := make([]int, n)
	x := -1
	scanner.Scan()
	for i := 0; i < n; i++ {
		x = readInt(scanner.Bytes(), x+1, &res[i])
	}
	return res
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

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		n, m := readTwoNums(scanner)
		grid := make([][]int, n)
		for i := 0; i < n; i++ {
			grid[i] = readNNums(scanner, m)
		}
		res := solve(n, m, grid)
		for i := 0; i < n; i++ {
			fmt.Println(res[i])
		}
	}
}

func solve(n int, m int, grid [][]int) []string {
	dp := make([][]int, n)

	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
	}

	res := make([][]byte, n)

	for i := 0; i < n; i++ {
		res[i] = make([]byte, m)
	}

	for j := 0; j < m; j++ {
		res[0][j] = '1'
	}

	for i := 1; i < n; i++ {
		for j := 0; j < m; j++ {
			dp[i][j] = grid[i-1][j]
			if j > 0 {
				dp[i][j] = max(dp[i][j], dp[i-1][j-1])
			}
			if j < m-1 {
				dp[i][j] = max(dp[i][j], dp[i-1][j+1])
			}
			if dp[i][j] < grid[i][j] {
				res[i][j] = '1'
			} else {
				res[i][j] = '0'
			}
			dp[i][j] = max(dp[i][j], grid[i][j])
		}
	}

	ans := make([]string, n)

	for i := 0; i < n; i++ {
		ans[i] = string(res[i])
	}

	return ans
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
