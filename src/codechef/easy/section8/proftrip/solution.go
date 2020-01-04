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

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		n, r := readTwoNums(scanner)
		roads := make([][]int, r)

		for i := 0; i < r; i++ {
			roads[i] = readNNums(scanner, 3)
		}

		F := readNNums(scanner, n)
		P, Q := readTwoNums(scanner)

		fmt.Println(solve(n, roads, F, P, Q))
	}
}

func solve(N int, roads [][]int, F []int, P int, Q int) int64 {
	const INF = math.MaxInt64

	dp := make([][]int64, N)

	for i := 0; i < N; i++ {
		dp[i] = make([]int64, N)
		for j := 0; j < N; j++ {
			dp[i][j] = INF
		}
		dp[i][i] = 0
	}

	for _, road := range roads {
		u, v, w := road[0]-1, road[1]-1, road[2]
		dp[u][v] = min(int64(w), dp[u][v])
		dp[v][u] = min(int64(w), dp[v][u])
	}

	for k := 0; k < N; k++ {
		for u := 0; u < N; u++ {
			if dp[u][k] >= INF {
				continue
			}
			for v := 0; v < N; v++ {
				if dp[k][v] < INF {
					dp[u][v] = min(dp[u][v], dp[u][k]+dp[k][v])
				}
			}
		}
	}

	// fp[u][v] = min cost from u, to v
	fp := make([][]int64, N)
	for i := 0; i < N; i++ {
		fp[i] = make([]int64, N)

		for j := 0; j < N; j++ {
			if dp[i][j] < INF {
				fp[i][j] = int64(F[i]) * dp[i][j]
			} else {
				fp[i][j] = INF
			}
		}
	}

	for k := 0; k < N; k++ {
		for u := 0; u < N; u++ {
			if fp[u][k] == INF {
				continue
			}
			for v := 0; v < N; v++ {
				if fp[k][v] < INF && fp[u][k]+fp[k][v] < fp[u][v] {
					fp[u][v] = fp[u][k] + fp[k][v]
				}
			}
		}
	}

	if fp[P-1][Q-1] == INF {
		return -1
	}

	return fp[P-1][Q-1]
}

func min(a, b int64) int64 {
	if a <= b {
		return a
	}
	return b
}
