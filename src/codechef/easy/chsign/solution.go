package main

import (
	"bufio"
	"os"
	"fmt"
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

func readNNums(scanner *bufio.Scanner, n int) []int {
	res := make([]int, n)
	x := -1
	scanner.Scan()
	for i := 0; i < n; i++ {
		x = readInt(scanner.Bytes(), x+1, &res[i])
	}
	return res
}

const INF = 1 << 60
const MAX_N = 100002

var A []int
var dp [][]int64
var B []int

func init() {
	A = make([]int, MAX_N)
	B = make([]int, MAX_N)
	dp = make([][]int64, MAX_N)
	for i := 0; i < MAX_N; i++ {
		dp[i] = make([]int64, 3)
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	t := readNum(scanner)

	for t > 0 {
		n := readNum(scanner)
		scanner.Scan()
		pos := -1

		for i := 0; i < n; i++ {
			pos = readInt(scanner.Bytes(), pos+1, &A[i])
		}

		res := solve(n, A)

		fmt.Printf("%d", res[0])
		for i := 1; i < n; i++ {
			fmt.Printf(" %d", res[i])
		}
		fmt.Println()
		t--
	}
}

func solve(n int, A []int) []int {

	dp[1][0] = int64(A[0] + A[1])
	if A[0]-A[1] > 0 {
		dp[1][1] = int64(A[0] - A[1])
	} else {
		dp[1][1] = INF
	}
	if A[1]-A[0] > 0 {
		dp[1][2] = int64(A[1] - A[0])
	} else {
		dp[1][2] = INF
	}

	for i := 2; i < n; i++ {
		dp[i][0] = INF
		dp[i][1] = INF
		dp[i][2] = INF
		// 0 -> ++
		// 1-> +-
		// 2 -> -+
		dp[i][0] = min(dp[i-1][0], dp[i-1][2]) + int64(A[i])

		if A[i-1]-A[i] > 0 {
			dp[i][1] = dp[i-1][0] - int64(A[i])
		}

		if -A[i-2]+A[i-1]-A[i] > 0 {
			dp[i][1] = min(dp[i][1], dp[i-1][2]-int64(A[i]))
		}

		if -A[i-1]+A[i] > 0 {
			dp[i][2] = dp[i-1][1] + int64(A[i])
		}
	}

	//B := make([]int, n)
	copy(B, A[:n])

	var j int

	for k := 1; k < 3; k++ {
		if dp[n-1][k] < dp[n-1][j] {
			j = k
		}
	}

	for i := n - 1; i > 0; i-- {
		if j == 1 {
			B[i] = -B[i]
		}
		if i == 1 && j == 2 {
			B[0] = -B[0]
		}
		if j == 0 || j == 1 {
			if j == 0 && dp[i][j] == dp[i-1][0]+int64(A[i]) {
				j = 0
			} else if j == 1 && dp[i][j] == dp[i-1][0]-int64(A[i]) {
				j = 0
			} else {
				j = 2
			}
		} else {
			j = 1
		}
	}

	return B
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func solve1(n int, A []int) []int {
	B := make([]int, n+2)
	copy(B[1:], A)
	B[0] = 1 << 30
	B[n+1] = 1 << 30
	C := make([]int, n+2)
	i := 1
	for i <= n {
		if B[i] < B[i-1] && B[i] < B[i+1] {
			j := i
			i += 2
			for i <= n && B[i] < B[i-1] && B[i] < B[i+1] && B[i]+B[i-2] >= B[i-1] {
				i += 2
			}
			a := 0
			b := B[j]
			k := j + 2
			for k < i {
				if a+B[k] >= b {
					C[k] = a + B[k]
				} else {
					C[k] = b
				}
				a, b = b, C[k]
				k += 2
			}

			k -= 2

			for k >= j {
				if k == j || C[k] > C[k-2] {
					B[k] = -B[k]
					k -= 4
				} else {
					k -= 2
				}
			}

		} else {
			i++
		}

	}

	return B[1 : n+1]
}
