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

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for x := 1; x <= tc; x++ {
		scanner.Scan()
		S := scanner.Text()
		res := solve(S)
		fmt.Printf("Case #%d: %04d\n", x, res)
	}
}

const T = "welcome to code jam"
const MOD = 10000

func solve(S string) int {
	n := len(S)
	m := len(T)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
	}

	for i := 0; i <= n; i++ {
		dp[i][0] = 1
	}

	//dp[i][j] = 前i位包含target的计数
	//dp[i][j] = dp[i-1][j] + (dp[i-1][j-1] if S[i] == T[j])
	for j := 1; j <= m; j++ {
		for i := 1; i <= n; i++ {
			dp[i][j] = dp[i-1][j]
			if S[i-1] == T[j-1] {
				dp[i][j] += dp[i-1][j-1]
				if dp[i][j] >= MOD {
					dp[i][j] -= MOD
				}
			}
		}
	}

	return dp[n][m]
}
