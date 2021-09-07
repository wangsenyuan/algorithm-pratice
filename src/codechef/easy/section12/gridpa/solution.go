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
		S := make([]string, n)
		for i := 0; i < n; i++ {
			S[i], _ = reader.ReadString('\n')
		}
		A := make([][]int, n)
		for i := 0; i < n; i++ {
			A[i] = readNNums(reader, n)
		}
		res := solve(n, k, S, A)
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

func solve(n int, k int, S []string, A [][]int) int64 {
	dp := make([][][]int64, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][]int64, n)
		for j := 0; j < n; j++ {
			dp[i][j] = make([]int64, k+1)
			for x := 0; x <= k; x++ {
				dp[i][j][x] = -1
			}
		}
	}

	dp[0][0][0] = int64(A[0][0])

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			// try to find min cost reach i, j, x
			if S[i][j] == '.' {
				if i > 0 && dp[i-1][j][0] >= 0 {
					dp[i][j][0] = max(dp[i][j][0], dp[i-1][j][0]+int64(A[i][j]))
				}
				if j > 0 && dp[i][j-1][0] >= 0 {
					dp[i][j][0] = max(dp[i][j][0], dp[i][j-1][0]+int64(A[i][j]))
				}
			}
			for x := 1; x <= k; x++ {
				if i > 0 && dp[i-1][j][x-1] >= 0 {
					dp[i][j][x] = max(dp[i][j][x], dp[i-1][j][x-1]+int64(A[i][j]))
				}
				if j > 0 && dp[i][j-1][x-1] >= 0 {
					dp[i][j][x] = max(dp[i][j][x], dp[i][j-1][x-1]+int64(A[i][j]))
				}
			}
		}
	}
	var res int64 = -1

	if S[n-1][n-1] == '#' {
		// it has to be in flip path
		for x := 1; x <= k; x++ {
			res = max(res, dp[n-1][n-1][x])
		}
		return res
	}

	fp := make([][]int64, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fp[i] = make([]int64, n)
			for j := 0; j < n; j++ {
				fp[i][j] = -1
			}
		}
	}
	fp[n-1][n-1] = int64(A[n-1][n-1])
	for i := n - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if S[i][j] == '#' {
				continue
			}
			if i+1 < n && fp[i+1][j] >= 0 {
				fp[i][j] = max(fp[i][j], fp[i+1][j]+int64(A[i][j]))
			}
			if j+1 < n && fp[i][j+1] >= 0 {
				fp[i][j] = max(fp[i][j], fp[i][j+1]+int64(A[i][j]))
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for x := 0; x <= k; x++ {
				if dp[i][j][x] >= 0 {
					if i+1 < n && fp[i+1][j] >= 0 {
						res = max(res, dp[i][j][x]+fp[i+1][j])
					}
					if j+1 < n && fp[i][j+1] >= 0 {
						res = max(res, dp[i][j][x]+fp[i][j+1])
					}
				}
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
