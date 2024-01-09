package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := 1
	var buf bytes.Buffer
	for tc > 0 {
		tc--
		board := make([]string, 2)
		board[0] = readString(reader)
		board[1] = readString(reader)
		res := solve(board)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	fmt.Print(buf.String())
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\r' || s[i] == '\n' {
			return s[:i]
		}
	}
	return s
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

func solve(board []string) int {
	n := len(board[0])
	if n < 2 {
		return 0
	}

	dp := make([]int, 4)
	// dp[0] = 表示row[0][i] and row[1][i] both free
	// dp[1] = row[0][i] is not free, but row[1][i] is free
	// dp[2] = row[0][i] is free, but row[1][i]
	// dp[3] = both not free
	for i := 0; i < 4; i++ {
		dp[i] = -inf
	}
	dp[3] = 0
	if board[0][0] == '0' {
		dp[2] = 0
	}
	if board[1][0] == '0' {
		dp[1] = 0
	}
	if board[0][0] == '0' && board[1][0] == '0' {
		dp[0] = 0
	}

	for i := 1; i < n; i++ {
		fp := make([]int, 4)
		for j := 0; j < 4; j++ {
			fp[j] = -inf
		}

		// 完全舍弃这一列
		fp[3] = max(dp...)

		// 舍弃其中一个格子
		if board[0][i] == '0' || board[1][i] == '0' {
			fp[3] = max(fp[3], dp[0]+1)
		}

		if board[0][i] == '0' {
			fp[2] = dp[3]
		}
		if board[1][i] == '0' {
			fp[1] = dp[3]
		}

		if board[0][i] == '0' && board[1][i] == '0' {
			fp[0] = max(dp...)
			fp[1] = max(fp[1], dp[0]+1)
			fp[2] = max(fp[2], dp[0]+1)
			fp[3] = max(fp[3], max(dp[1], dp[2])+1)
		}
		copy(dp, fp)
	}

	return max(dp...)
}

func max(arr ...int) int {
	res := arr[0]
	for i := 1; i < len(arr); i++ {
		if res < arr[i] {
			res = arr[i]
		}
	}
	return res
}
