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
		s := readString(reader)
		res := solve(s)
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

func solve(s string) int {
	n := len(s)

	// dp[i][j] 表示从i, j开始的最长的字符串的长度, 以及最后一个字符的在第一段中的结尾
	// 在长度相同时，选择结尾更小的

	comp := func(x, y byte) bool {
		if x == y {
			return true
		}
		return x == '?' || y == '?'
	}

	// dp[i][j] 表示从(i, j)开始的字符是否一直相等
	// s[i] == s[j], s[i+1] == s[j+1] ....
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, n+1)
	}
	for j := n - 1; j >= 0; j-- {
		for i := j - 1; i >= 0; i-- {
			if comp(s[i], s[j]) {
				if i+1 == j {
					dp[i][j] = 1
				} else {
					tmp := min(dp[i+1][j+1]+1, j-i)
					dp[i][j] = max(dp[i][j], tmp)
				}
			} else {
				dp[i][j] = 0
			}
		}
	}

	var res int
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			k := dp[i][j]
			if i+k == j {
				res = max(res, k)
			}
		}
	}
	return 2 * res
}
