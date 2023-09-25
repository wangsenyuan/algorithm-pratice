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
		n := readNum(reader)
		s := readString(reader)[:n]
		a := readNNums(reader, n)
		res := solve(s, a)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	fmt.Print(buf.String())
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
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

func solve(s string, a []int) int64 {
	n := len(s)
	dp := make([][]int64, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int64, 2)
	}
	if s[0] == '1' {
		dp[0][0] = 0
		dp[0][1] = int64(a[0])
	}

	for i := 1; i < n; i++ {
		if s[i] == '0' {
			// it would soak
			dp[i][0] = max(dp[i-1][0], dp[i-1][1])
			dp[i][1] = 0
		} else {
			dp[i][1] = max(dp[i-1][0], dp[i-1][1]) + int64(a[i])
			dp[i][0] = dp[i-1][0] + int64(a[i-1])
		}
	}

	return max(dp[n-1][0], dp[n-1][1])
}

func max(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}
