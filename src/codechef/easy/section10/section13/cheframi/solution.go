package main

import (
	"bufio"
	"fmt"
	"math"
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

		n, x := readTwoNums(scanner)
		A := readNNums(scanner, n)

		fmt.Println(solve(n, x, A))
	}
}

func solve(N int, x int, A []int) int64 {
	var sum int64

	for i := 0; i < N; i++ {
		sum += int64(A[i])
	}

	if sum == 0 {
		return 0
	}

	X := int64(x)

	cost := make([][]int64, N)

	for i := 0; i < N; i++ {
		cost[i] = make([]int64, N)
		p := i
		// var tot int64
		left := int64(A[i])
		var right int64
		var tot int64
		for j := i + 1; j < N; j++ {
			tot += int64(j-p) * int64(A[j])
			right += int64(A[j])
			for left < right {
				tot += left - right
				left += int64(A[p+1])
				right -= int64(A[p+1])
				p++
			}
			cost[i][j] = tot
		}
	}

	dp := make([]int64, N+1)
	dp[0] = 0

	for i := 0; i < N; i++ {
		dp[i+1] = math.MaxInt64

		for j := 0; j <= i; j++ {
			dp[i+1] = min(dp[i+1], dp[j]+cost[j][i]+X)
		}
	}
	return dp[N]
}

func min(a, b int64) int64 {
	if a <= b {
		return a
	}
	return b
}
