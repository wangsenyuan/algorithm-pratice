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
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(reader *bufio.Reader) (a int) {
	bs, _ := reader.ReadBytes('\n')
	readInt(bs, 0, &a)
	return
}

func readTwoNums(reader *bufio.Reader) (a int, b int) {
	res := readNNums(reader, 2)
	a, b = res[0], res[1]
	return
}

func readThreeNums(reader *bufio.Reader) (a int, b int, c int) {
	res := readNNums(reader, 3)
	a, b, c = res[0], res[1], res[2]
	return
}

func readNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, n)
	x := 0
	bs, _ := reader.ReadBytes('\n')
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	line := readNNums(reader, 4)
	n, m, k, mod := line[0], line[1], line[2], line[3]
	fmt.Println(solve(n, m, k, mod))
}

func solve(n int, m int, k int, mod int) int {
	n++
	m++
	k++
	N := n * (n + 1)

	M := m*m/2 + 5
	C := make([][]int64, N/2+1)
	for i := 0; i < len(C); i++ {
		C[i] = make([]int64, M+1)
	}

	MOD := int64(mod)

	C[0][0] = 1

	for i := 1; i < len(C); i++ {
		C[i][0] = 1

		for j := 1; j <= i && j <= M; j++ {
			add(&C[i][j], C[i-1][j], MOD)
			add(&C[i][j], C[i-1][j-1], MOD)
		}
	}

	ways := make([][][]int64, N/2+1)
	for usual := 0; usual < len(ways); usual++ {
		ways[usual] = make([][]int64, n+1)
		ways[usual][0] = make([]int64, m+1)
		for special := 1; special <= n; special++ {
			ways[usual][special] = make([]int64, m+1)
			for pick := 1; pick <= m && pick <= usual+special; pick++ {
				ways[usual][special][pick] = countWays(usual, special, pick, C, MOD)
			}
		}
	}

	dp := make([][][]int64, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([][]int64, m+1)
		for j := 0; j <= m; j++ {
			dp[i][j] = make([]int64, k+1)
		}
	}

	dp[0][0][0] = 1

	for pos := 0; pos < n; pos++ {
		for x := 0; x < m; x++ {
			for y := 0; y <= pos && y < k; y++ {
				if dp[pos][x][y] == 0 {
					continue
				}
				for newPos := pos + 1; newPos <= n; newPos++ {
					for more := 1; x+more <= m; more++ {
						usual := pos * (newPos - pos)
						special := newPos - pos
						if newPos == n {
							usual -= pos
							special = 1
						}
						tmp := dp[pos][x][y] * ways[usual][special][more]
						tmp %= MOD
						add(&dp[newPos][x+more][y+1], tmp, MOD)
					}
				}
			}
		}
	}

	return int(dp[n][m][k])
}

func countWays(usual, special, pick int, C [][]int64, MOD int64) int64 {
	var ans int64

	for take := 1; take <= pick; take++ {
		if pick-take <= usual {
			takeUsual := pick - take
			tmp := C[usual][takeUsual] * C[special][take]
			tmp %= MOD
			add(&ans, tmp, MOD)
		}
	}
	return ans
}

func add(a *int64, b int64, MOD int64) {
	*a += b
	if *a >= MOD {
		*a -= MOD
	}
}
