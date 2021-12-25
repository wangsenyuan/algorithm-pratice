package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n, m := readTwoNums(reader)
		A := readNNums(reader, n)
		res := solve(n, m, A)
		fmt.Printf("%.8f\n", res)
	}
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

func solve(n int, m int, A []int) float64 {
	dp := make([][][]float64, m+1)
	next := make([][][]float64, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([][]float64, n+2)
		next[i] = make([][]float64, n+2)
		for j := 0; j <= n+1; j++ {
			dp[i][j] = make([]float64, n+2)
			next[i][j] = make([]float64, n+2)
		}
	}

	dp[0][0][1] = 1
	sum := make([]float64, m+1)
	for s := 0; s < n; s++ {
		if A[s] == 0 {
			// ni == i
			for i := 1; i <= m; i++ {
				for j := 0; j <= n; j++ {
					for k := 0; k <= n; k++ {
						ni := i
						nj := j + 1
						nk := k
						if nk < nj {
							nk = nj
						}
						next[ni][nj][nk] += dp[i][j][k] / float64(m)
					}
				}
			}
			//ni != i
			for k := 0; k <= n; k++ {
				for i := 0; i <= m; i++ {
					for j := 0; j <= n; j++ {
						sum[i] += dp[i][j][k]
					}
				}
				for i := 0; i <= m; i++ {
					for ni := 1; ni <= m; ni++ {
						if i != ni {
							next[ni][1][k] += sum[i] / float64(m)
						}
					}
					sum[i] = 0
				}
			}
		} else {
			for i := 0; i <= m; i++ {
				for j := 0; j <= n; j++ {
					for k := 0; k <= n; k++ {
						ni := A[s]
						nj := j + 1
						if i != ni {
							nj = 1
						}
						nk := k
						if nk < nj {
							nk = nj
						}
						next[ni][nj][nk] += dp[i][j][k]
					}
				}
			}
		}

		for i := 0; i <= m; i++ {
			for j := 0; j <= n; j++ {
				for k := 0; k <= n; k++ {
					dp[i][j][k] = next[i][j][k]
					next[i][j][k] = 0
				}
			}
		}
	}

	var ans float64

	for k := 1; k <= n; k++ {
		for i := 0; i <= m; i++ {
			for j := 0; j <= n; j++ {
				ans += float64(k) * dp[i][j][k]
			}
		}
	}
	return ans
}

func solve1(n int, m int, A []int) float64 {
	// dp[i][j] is the probability ending at i, with length j
	P := make([][][][]float64, n)
	for i := 0; i < n; i++ {
		P[i] = make([][][]float64, m+1)
		for j := 0; j <= m; j++ {
			P[i][j] = make([][]float64, n+1)
			for k := 0; k <= n; k++ {
				P[i][j][k] = make([]float64, n+1)
			}
		}
	}
	pp := 1 / float64(m)
	if A[0] > 0 {
		P[0][A[0]][1][1] = 1
	} else {
		for x := 1; x <= m; x++ {
			P[0][x][1][1] = pp
		}
	}

	sum := make([][]float64, m+1)
	for i := 0; i <= m; i++ {
		sum[i] = make([]float64, n+1)
	}

	for x := 1; x < n; x++ {
		y := A[x]
		if y > 0 {
			for j := 2; j <= n; j++ {
				P[x][y][j][j] += P[x-1][y][j-1][j-1]
				for k := j; k <= n; k++ {
					P[x][y][j][k] += P[x-1][y][j-1][k]
				}
			}
			for i := 1; i <= m; i++ {
				if i == y {
					continue
				}
				for j := 1; j <= n; j++ {
					for k := 1; k <= n; k++ {
						P[x][y][1][k] += P[x-1][i][j][k]
					}
				}
			}
		} else {
			for i := 1; i <= m; i++ {
				for k := 0; k <= n; k++ {
					sum[i][k] = 0
				}
			}

			for k := 1; k <= n; k++ {
				for i := 1; i <= m; i++ {
					for j := 1; j <= n; j++ {
						sum[i][k] += P[x-1][i][j][k]
					}
				}
			}

			for y := 1; y <= m; y++ {
				for j := 2; j <= n; j++ {
					P[x][y][j][j] += P[x-1][y][j-1][j-1] * pp
					for k := j; k <= n; k++ {
						P[x][y][j][k] += P[x-1][y-1][j-1][k] * pp
					}
				}
				for i := 1; i <= m; i++ {
					if i == y {
						continue
					}
					for k := 1; k <= n; k++ {
						P[x][y][1][k] += sum[i][k] * pp
					}
				}
			}
		}
	}

	var ans float64

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			for k := j; k <= n; k++ {
				ans += float64(k) * P[n-1][i][j][k]
			}
		}
	}

	return ans
}
