package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, k := readTwoNums(reader)
	res := solve(n, k)
	fmt.Println(res)
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

func solve(n, k int) int64 {

	dp := make([][][][]int64, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][][]int64, k+1)
		for j := 0; j <= k; j++ {
			dp[i][j] = make([][]int64, 4)
			for x := 0; x < 4; x++ {
				// 0 for -1, 1 for + 1
				dp[i][j][x] = make([]int64, 4)
			}
		}
	}

	for a := 0; a < 4; a++ {
		for b := a + 1; b < 4; b++ {
			dp[1][0][a][b]++
		}
	}

	for i := 1; i+1 < n; i++ {
		for j := 0; j <= k; j++ {
			for a := 0; a < 4; a++ {
				for b := 0; b < 4; b++ {
					if a == b || dp[i][j][a][b] == 0 {
						continue
					}
					if j+1 <= k {
						for c := b + 1; c < 4; c++ {
							dp[i+1][j][b][c] += dp[i][j][a][b]
						}
					}
					if b > 0 {
						if a < b && j+1 <= k {
							for c := b - 1; c >= 0; c-- {
								dp[i+1][j+1][b][c] += dp[i][j][a][b]
							}
						}
						if a > b {
							for c := b - 1; c >= 0; c-- {
								dp[i+1][j][b][c] += dp[i][j][a][b]
							}
						}
					}
				}
			}
		}
	}

	var res int64

	for a := 1; a < 4; a++ {
		for b := a - 1; b >= 0; b-- {
			res += dp[n-1][k][a][b]
		}
	}
	return res
}
