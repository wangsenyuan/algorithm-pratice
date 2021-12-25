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
		N, M := readTwoNums(scanner)

		grid := make([][]int, N)

		for i := 0; i < N; i++ {
			grid[i] = readNNums(scanner, M)
		}

		fmt.Println(solve(N, M, grid))
	}
}

func solve(N int, M int, grid [][]int) int64 {
	row := make([]int, M)
	col := make([]int, N)
	bak := make([]int, 2*max(M, N)+3)
	res := make([]int, max(N, M))
	dp := make([][]int, N)

	for i := 0; i < N; i++ {
		dp[i] = make([]int, M)
	}

	for i := 0; i < N; i++ {
		copy(row, grid[i])
		manacher(row, res, bak)

		for j := 0; j < M; j++ {
			dp[i][j] = res[j]
		}
	}

	for j := 0; j < M; j++ {
		for i := 0; i < N; i++ {
			col[i] = grid[i][j]
		}
		manacher(col, res, bak)

		for i := 0; i < N; i++ {
			dp[i][j] = min(dp[i][j], res[i])
		}
	}

	var ans int64

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			x := dp[i][j]
			ans += int64(x)
		}
	}

	return ans
}

func manacher(nums []int, res []int, bak []int) {
	n := len(nums)

	bak[0] = 0

	for i := 0; i < n; i++ {
		bak[2*i+1] = -1
		bak[2*i+2] = nums[i]
	}

	bak[2*n+1] = -1
	bak[2*n+2] = -2

	m := 2*n + 3

	P := make([]int, m)
	var C, R int

	for i := 1; i < m-1; i++ {
		j := 2*C - i

		if R > i {
			P[i] = min(R-i, P[j])
		}

		for bak[i+1+P[i]] == bak[i-1-P[i]] {
			P[i]++
		}

		if i+P[i] > R {
			R = i + P[i]
			C = i
		}
	}

	for i := 0; i < n; i++ {
		j := 2*i + 2
		res[i] = (P[j] + 1) / 2
	}
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
