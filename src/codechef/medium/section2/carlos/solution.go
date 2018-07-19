package main

import (
	"bufio"
	"fmt"
	"math"
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

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for tc > 0 {
		M, K, N := readThreeNums(scanner)
		T := make([][]int, K)
		for i := 0; i < K; i++ {
			T[i] = readNNums(scanner, 2)
		}
		W := readNNums(scanner, N)
		res := solve(N, M, K, T, W)
		fmt.Println(res)
		tc--
	}
}

const MAX_N = 200000
const MAX_M = 200
const INF = math.MaxInt32

var G [][]bool

var dp [][]int

func init() {
	G = make([][]bool, MAX_M+1)
	for i := 0; i <= MAX_M; i++ {
		G[i] = make([]bool, MAX_M)
	}
	dp = make([][]int, MAX_N+1)
	for i := 0; i <= MAX_N; i++ {
		dp[i] = make([]int, MAX_M+1)
	}
}

func solve(N, M, K int, T [][]int, W []int) int {
	for i := 0; i <= M; i++ {
		for j := 0; j <= M; j++ {
			G[i][j] = false
		}
		G[i][i] = true
	}
	for _, t := range T {
		x, y := t[0], t[1]
		G[x][y] = true
		G[y][x] = true
	}

	for k := 1; k <= M; k++ {
		for i := 1; i <= M; i++ {
			for j := 1; j <= M; j++ {
				G[i][j] = G[i][j] || (G[i][k] && G[k][j])
			}
		}
	}

	// dp[i][j] = maxmim number j ending at i

	for i := 1; i <= N; i++ {
		x := W[i-1]
		dp[i][0] = INF
		for y := 1; y <= M; y++ {
			var cnt int
			if x != y {
				cnt++
			}
			dp[i][y] = dp[i][y-1]
			if G[x][y] && dp[i-1][y] < INF {
				dp[i][y] = min(dp[i][y], dp[i-1][y]+cnt)
			}
		}
	}

	if dp[N][M] >= INF {
		return -1
	}
	return dp[N][M]
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
