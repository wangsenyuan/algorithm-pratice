package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)
	var buf bytes.Buffer
	for tc > 0 {
		tc--
		n, k := readTwoNums(reader)
		s, _ := reader.ReadString('\n')
		res := solve(n, k, s)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	fmt.Print(buf.String())
}

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

const K = 32

const N_INF = -(1 << 30)

var dp [2][K + 2][K + 2][K + 2][K + 2]int

func reset(p int, k int) {
	for i := 0; i <= k+1; i++ {
		for j := 0; j <= k+1; j++ {
			for u := 0; u <= k+1; u++ {
				for v := 0; v <= k+1; v++ {
					dp[p][i][j][u][v] = N_INF
				}
			}
		}
	}
}

func solve(n, k int, s string) int {
	reset(0, k)
	dp[0][0][0][0][0] = 0

	for cur := 1; cur <= n; cur++ {
		reset(cur&1, k)
		x := s[cur-1]

		for i := 0; i <= k+1; i++ {
			for j := 0; j <= k+1; j++ {
				for u := 0; u <= k+1; u++ {
					for v := 0; v <= k+1; v++ {
						tmp := dp[(cur-1)&1][i][j][u][v]
						if tmp < 0 {
							continue
						}
						dp[cur&1][i][j][u][v] = max(tmp, dp[cur&1][i][j][u][v])
						if x == 'c' {
							dp[cur&1][i][j][u][min(k+1, v+1)] = max(dp[cur&1][i][j][u][min(k+1, v+1)], tmp+1)
						}
						if x == 'h' {
							dp[cur&1][i][j][min(k+1, u+v)][v] = max(dp[cur&1][i][j][min(k+1, u+v)][v], tmp+1)
						}
						if x == 'e' {
							dp[cur&1][i][min(k+1, j+u)][u][v] = max(dp[cur&1][i][min(k+1, j+u)][u][v], tmp+1)
						}
						if x == 'f' {
							dp[cur&1][min(k+1, i+j)][j][u][v] = max(dp[cur&1][min(k+1, i+j)][j][u][v], tmp+1)
						}
					}
				}
			}
		}
	}
	var res = -1

	for j := 0; j <= k+1; j++ {
		for u := 0; u <= k+1; u++ {
			for v := 0; v <= k+1; v++ {
				if dp[n&1][k][j][u][v] < 0 {
					continue
				}
				if res < dp[n&1][k][j][u][v] {
					res = dp[n&1][k][j][u][v]
				}
			}
		}
	}
	if res < 0 {
		return -1
	}
	return n - res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
