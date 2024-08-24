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
		n, k, z := readThreeNums(reader)
		a := readNNums(reader, n)
		res := solve(a, k, z)
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

const inf = 1 << 50

func solve(a []int, k int, z int) int {
	n := len(a)

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, z+1)
		for j := 0; j <= z; j++ {
			dp[i][j] = -inf
		}
	}
	dp[0][0] = a[0]
	var best int

	for i := 1; i < n; i++ {
		for j := 0; j <= z; j++ {
			dp[i][j] = dp[i-1][j] + a[i]
			if j > 0 {
				dp[i][j] = max(dp[i][j], dp[i][j-1]+a[i]+a[i-1])
			}
			if i+j*2 == k {
				best = max(best, dp[i][j])
			}
			if i+j*2 == k+1 && j > 0 {
				// 最后一次没有返回到位置i
				best = max(best, dp[i][j]-a[i])
			}
		}
	}

	return best
}
