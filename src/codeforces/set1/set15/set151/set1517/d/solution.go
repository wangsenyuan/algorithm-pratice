package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, m, k := readThreeNums(reader)
	x := make([][]int, n)
	for i := 0; i < n; i++ {
		x[i] = readNNums(reader, m-1)
	}
	y := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		y[i] = readNNums(reader, m)
	}

	res := solve(n, m, k, x, y)
	var buf bytes.Buffer

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			buf.WriteString(fmt.Sprintf("%d ", res[i][j]))
		}
		buf.WriteByte('\n')
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

const inf = 1 << 30

func solve(n int, m int, k int, x [][]int, y [][]int) [][]int {
	dp := make([][]int, n)
	fp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
		fp[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fp[i][j] = -1
		}
	}

	if k%2 == 1 {
		return fp
	}

	k /= 2

	for k1 := 0; k1 < k; k1++ {
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				// 4个方向
				fp[i][j] = inf
				if j > 0 {
					fp[i][j] = min(fp[i][j], dp[i][j-1]+x[i][j-1])
				}
				if j+1 < m {
					fp[i][j] = min(fp[i][j], dp[i][j+1]+x[i][j])
				}
				if i > 0 {
					fp[i][j] = min(fp[i][j], dp[i-1][j]+y[i-1][j])
				}
				if i+1 < n {
					fp[i][j] = min(fp[i][j], dp[i+1][j]+y[i][j])
				}
			}
		}

		dp, fp = fp, dp
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			dp[i][j] *= 2
		}
	}

	return dp
}
