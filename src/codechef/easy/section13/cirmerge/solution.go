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

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var tmp int64
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp
	return i
}

func readNInt64Nums(scanner *bufio.Scanner, n int) []int64 {
	res := make([]int64, n)
	x := -1
	scanner.Scan()
	for i := 0; i < n; i++ {
		x = readInt64(scanner.Bytes(), x+1, &res[i])
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		n := readNum(scanner)
		A := readNNums(scanner, n)
		fmt.Println(solve(n, A))
	}
}

const INF = 1 << 50

func solve(n int, A []int) int64 {
	if n < 2 {
		return 0
	}

	if n == 2 {
		return int64(A[0] + A[1])
	}

	B := make([]int, 2*n)
	copy(B, A)
	copy(B[n:], A)

	N := 2 * n
	dp := make([][]int64, N)
	sum := make([][]int64, N)
	for i := 0; i < N; i++ {
		dp[i] = make([]int64, N)
		sum[i] = make([]int64, N)
		for j := 0; j < N; j++ {
			dp[i][j] = INF
		}
		dp[i][i] = int64(A[i%n])
		sum[i][i] = int64(A[i%n])
	}

	for j := 1; j < N; j++ {
		for i := j - 1; i >= 0 && i > j-n; i-- {
			sum[i][j] = int64(A[i%n]) + sum[i+1][j]
			for k := i; k < j; k++ {
				dp[i][j] = min(dp[i][j], dp[i][k]+dp[k+1][j])
			}
			dp[i][j] += sum[i][j]
		}
	}

	var best int64 = INF

	for i := 0; i < n; i++ {
		for j := i; j < i+n; j++ {
			a := dp[i][j]
			b := dp[j+1][i-1+n]
			best = min(best, a+b)
		}
	}

	return best
}

func min(a, b int64) int64 {
	if a <= b {
		return a
	}
	return b
}
