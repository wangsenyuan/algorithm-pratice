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
	sign := int64(1)
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
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
		A := readNInt64Nums(scanner, n)

		res := solve(n, A)

		for i := 0; i < n; i++ {
			if i < n-1 {
				fmt.Printf("%d ", res[i])
			} else {
				fmt.Printf("%d\n", res[i])
			}
		}
	}
}

const INF = 1 << 30

func solve(n int, A []int64) []int64 {
	res := make([]int64, n)

	res[0] = -INF
	var sum int64

	xp := make([]int64, n)

	for i := 0; i < n; i++ {
		sum += A[i]
		res[0] = max(res[0], sum)
		if sum < 0 {
			sum = 0
		}
		xp[i] = res[0]
	}

	// dp[i] the max sum from 0 <= i
	dp := make([]int64, n)
	sum = A[0]
	dp[0] = A[0]

	for i := 1; i < n; i++ {
		sum += A[i]
		dp[i] = max(dp[i-1], sum)
	}

	fp := make([]int64, n)
	fp[n-1] = A[n-1]
	sum = A[n-1]

	for i := n - 2; i >= 0; i-- {
		sum += A[i]
		fp[i] = max(fp[i+1], sum)
	}

	yp := make([]int64, n)
	var best int64 = -INF
	sum = 0

	for i := n - 1; i >= 0; i-- {
		sum += A[i]
		best = max(best, sum)
		if sum < 0 {
			sum = 0
		}
		yp[i] = best
	}

	for i := 1; i < n; i++ {
		res[i] = fp[i] + dp[i-1]
		res[i] = max(res[i], yp[i])
		res[i] = max(res[i], xp[i-1])
	}

	return res
}

func max(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}
