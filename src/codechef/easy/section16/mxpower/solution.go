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
		n := readNum(scanner)
		E := make([][]int64, n)
		for i := 0; i < n; i++ {
			E[i] = readNInt64Nums(scanner, n)
		}
		fmt.Println(solve(n, E))
	}
}

func solve(N int, E [][]int64) int64 {
	dp := make([][][]int64, N)

	for i := 0; i < N; i++ {
		dp[i] = make([][]int64, N)
		for j := 0; j < N; j++ {
			dp[i][j] = make([]int64, N+1)
		}
	}

	var res int64 = math.MinInt64

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			dp[i][j][1] = E[i][j]
			res = max(res, dp[i][j][1])
		}
	}

	leftDiag := make([][]int64, 2*N-1)
	rightDiag := make([][]int64, 2*N-1)

	for i := 0; i < len(leftDiag); i++ {
		leftDiag[i] = make([]int64, N)
		var j int

		if i >= N {
			j = i - N + 1
		}

		for j < N && j <= i {
			leftDiag[i][j] = E[j][i-j]
			if j > 0 {
				leftDiag[i][j] += leftDiag[i][j-1]
			}
			j++
		}
	}

	for i := 0; i < len(rightDiag); i++ {
		rightDiag[i] = make([]int64, N)
		if i < N {
			for j := 0; j <= i; j++ {
				rightDiag[i][j] = E[N-1-i+j][j]
				if j > 0 {
					rightDiag[i][j] += rightDiag[i][j-1]
				}
			}
		} else {
			for j := i - N + 1; j < N; j++ {
				rightDiag[i][j] = E[j-(i-N+1)][j]
				if j > 0 {
					rightDiag[i][j] += rightDiag[i][j-1]
				}
			}
		}
	}

	for k := 3; k <= N; k += 2 {
		h := k / 2
		for i := 0; i < N; i++ {
			for j := 0; j < N; j++ {
				if i-h >= 0 && i+h < N && j-h >= 0 && j+h < N {
					dp[i][j][k] = dp[i][j][k-2]
					a := i - h
					b := i + h
					c := j - h
					d := j + h

					l1 := a + j
					dp[i][j][k] += leftDiag[l1][i]
					if a > 0 {
						dp[i][j][k] -= leftDiag[l1][a-1]
					}

					l2 := b + j
					dp[i][j][k] += leftDiag[l2][b] - leftDiag[l2][i-1]

					r1 := c + i
					dp[i][j][k] += rightDiag[r1][j]
					if c > 0 {
						dp[i][j][k] -= rightDiag[r1][c-1]
					}

					r2 := d + i
					dp[i][j][k] += rightDiag[r2][d] - rightDiag[r2][j-1]

					dp[i][j][k] -= E[a][j]
					dp[i][j][k] -= E[b][j]
					dp[i][j][k] -= E[i][c]
					dp[i][j][k] -= E[i][d]
				}

				res = max(res, dp[i][j][k])
			}
		}
	}

	return res
}

func max(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}
