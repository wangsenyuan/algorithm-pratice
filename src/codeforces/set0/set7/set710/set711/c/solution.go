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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
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

	n, m, k := readThreeNums(reader)
	c := readNNums(reader, n)
	p := make([][]int, n)
	for i := 0; i < n; i++ {
		p[i] = readNNums(reader, m)
	}
	res := solve(n, m, k, c, p)
	fmt.Println(res)
}

const inf = 1 << 60

func solve(n int, m int, k int, c []int, p [][]int) int {
	dp := make([][][]int, 2)
	for i := 0; i < 2; i++ {
		dp[i] = make([][]int, m+1)
		for j := 0; j <= m; j++ {
			dp[i][j] = make([]int, k+1)
			for x := 0; x <= k; x++ {
				dp[i][j][x] = inf
			}
		}
	}
	if c[0] == 0 {
		for j := 1; j <= m; j++ {
			dp[0][j][1] = p[0][j-1]
		}
	} else {
		dp[0][c[0]][1] = 0
	}
	var pt int
	for i := 1; i < n; i++ {
		np := 1 ^ pt
		for j := 0; j <= m; j++ {
			for x := 0; x <= k; x++ {
				dp[np][j][x] = inf
			}
		}
		// 但是必须知道前面的其他颜色情况下的最小值
		// 在给定x的情况下，知道其他颜色的最小值
		for x := 1; x <= k && x <= i; x++ {
			best := []int{-1, -1}
			for j := 1; j <= m; j++ {
				if best[0] < 0 || dp[pt][j][x] < dp[pt][best[0]][x] {
					best[1] = best[0]
					best[0] = j
				} else if best[1] < 0 || dp[pt][j][x] < dp[pt][best[1]][x] {
					best[1] = j
				}
			}
			if c[i] == 0 {
				for j := 1; j <= m; j++ {
					// 和前面的颜色一致时，x不变
					dp[np][j][x] = min(dp[np][j][x], dp[pt][j][x]+p[i][j-1])
					// 给当前位涂色j, 且和前面的颜色不一致
					if x+1 <= k {
						if j == best[0] && best[1] > 0 {
							dp[np][j][x+1] = min(dp[np][j][x+1], dp[pt][best[1]][x]+p[i][j-1])
						} else if best[0] > 0 {
							dp[np][j][x+1] = min(dp[np][j][x+1], dp[pt][best[0]][x]+p[i][j-1])
						}
					}
				}

			} else {
				// 不用花费
				dp[np][c[i]][x] = min(dp[np][c[i]][x], dp[pt][c[i]][x])
				if x+1 <= k {
					if c[i] == best[0] && best[1] > 0 {
						// 前面选用 best[1]
						dp[np][c[i]][x+1] = min(dp[np][c[i]][x+1], dp[pt][best[1]][x])
					} else if c[i] != best[0] && best[0] > 0 {
						dp[np][c[i]][x+1] = min(dp[np][c[i]][x+1], dp[pt][best[0]][x])
					}
				}

			}
		}
		pt = np
	}

	ans := inf
	for j := 1; j <= m; j++ {
		ans = min(ans, dp[pt][j][k])
	}

	if ans == inf {
		return -1
	}
	return ans
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
