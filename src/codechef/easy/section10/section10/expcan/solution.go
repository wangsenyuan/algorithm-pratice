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

	tc := readNum(scanner)

	for tc > 0 {
		tc--

		n := readNum(scanner)
		A := readNNums(scanner, n)

		res := solve(n, A)

		fmt.Printf("%.7f\n", res)
	}
}

func solve(n int, A []int) float64 {
	fp := make([][]float64, n)

	for i := 0; i < n; i++ {
		fp[i] = make([]float64, n)
		var sum int
		for j := i; j < n; j++ {
			sum += A[j]
			// cnt++
			fp[i][j] = float64(sum)
		}
	}

	dp := make([][]float64, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]float64, n)
		dp[i][i] = float64(A[i])
	}

	for j := 0; j < n; j++ {
		for i := j - 1; i >= 0; i-- {
			//dp[i][j] = 0.5 * A[i] + fp[i+1][j] - dp[i+1][j] or 0.5 * A[j] + fp[i][j-1] - dp[i][j-1]
			dp[i][j] = 0.5 * (float64(A[i]) + fp[i+1][j] - dp[i+1][j] + float64(A[j]) + fp[i][j-1] - dp[i][j-1])
		}
	}

	return dp[0][n-1]
}

func max(a, b float64) float64 {
	if a >= b {
		return a
	}
	return b
}
