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

func readThreeNums(scanner *bufio.Scanner) (a int, b int, c int) {
	res := readNNums(scanner, 3)
	a, b, c = res[0], res[1], res[2]
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
		fmt.Println(solve(n, A))
	}
}

const MOD = 1000000007

var C [][]int64

func init() {
	C = make([][]int64, 201)
	for i := 0; i < 201; i++ {
		C[i] = make([]int64, 201)
		C[i][0] = 1
		C[i][i] = 1
		for j := 1; j < i; j++ {
			C[i][j] = (C[i-1][j-1] + C[i-1][j]) % MOD
		}
	}
}

func solve(n int, A []int) int {
	dp := make([][]int64, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int64, 201)
	}
	dp[0][0] = 1

	var total int = 1

	for x := 1; x <= n; x++ {
		for ij := 1; ij <= A[x-1]; ij++ {
			for y := 0; y <= total; y++ {
				if dp[x-1][y] > 0 {
					for j := 0; j <= min(ij, y); j++ {
						i := ij - j
						z := total - y
						dp[x][y-j+A[x-1]-ij] += ((dp[x-1][y] * C[y][j]) % MOD * C[z][i] % MOD) * C[A[x-1]-1][ij-1] % MOD
						dp[x][y-j+A[x-1]-ij] %= MOD
					}
				}
			}
		}
		total += A[x-1]
	}

	return int(dp[n][0])
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func solve1(n int, A []int) int {
	dp := make([]int64, 201)
	dp[0] = 1

	next := make([]int64, 201)
	var all int = 1
	for k := 0; k < n; k++ {
		for i := 0; i < len(next); i++ {
			next[i] = 0
		}
		for i := 0; i < 201; i++ {
			if dp[i] > 0 {
				for b := 1; b <= A[k]; b++ {
					mul := dp[i] * C[A[k]-1][b-1] % MOD
					for r := 0; r <= b; r++ {
						add := (mul * C[i][r] % MOD) * C[all-i][b-r] % MOD
						if add > 0 {
							next[i+A[k]-b-r] += add
						}
					}
				}
			}
		}

		for i := 0; i < 201; i++ {
			dp[i] = next[i] % MOD
		}

		all += A[k]
	}

	return int(dp[0])
}
