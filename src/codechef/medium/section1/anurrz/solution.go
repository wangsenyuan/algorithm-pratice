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

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	t := readNum(scanner)

	for t > 0 {
		n := readNum(scanner)
		B := readNNums(scanner, n)
		fmt.Println(solve(n, B))
		t--
	}
}

const MOD = 1000000007

func solve(n int, B []int) int64 {
	dp := make([][]int64, n+2)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int64, n+2)
	}

	dp[0][0] = 1

	for i := 1; i <= n; i++ {
		for j := B[i-1]; j > 0; j-- {
			dp[i][j] = dp[i-1][j-1] + dp[i][j+1]
			if dp[i][j] >= MOD {
				dp[i][j] -= MOD
			}
		}
		dp[i][0] = dp[i][1]
	}

	return dp[n][0]
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
