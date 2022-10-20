package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var buf bytes.Buffer

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n := readNum(reader)
		R := readNNums(reader, n)
		B := readNNums(reader, n)
		res := solve(n, R, B)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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

func solve(n int, R []int, B []int) int {
	// R[i] <= 200, n<= 100
	// sum(R[i]) <= 20000 also for B
	// 2-d dp[r][b] will not work, as 20000 * 20000 = 400000000
	// dp[r][b] = dp[r - r[i]][b] || dp[r][b - b[i]]
	// 假设 min(x, y) = x
	// 假设给定x，如何分配，得到y，那么找到最大的x即可，y >= x
	// 感觉贪心就可以了，如果 R[i] > B[i], 则优先作为red
	//                    R[i] < B[i],          blue

	var red int
	for i := 0; i < n; i++ {
		red += R[i]
	}
	dp := make([]int, red+1)

	for i := 0; i <= red; i++ {
		dp[i] = -1
	}

	dp[0] = 0

	for i := 0; i < n; i++ {
		for j := red; j >= 0; j-- {
			tmp := -1
			if dp[j] >= 0 {
				tmp = dp[j] + B[i]
			}
			if j >= R[i] && dp[j-R[i]] >= 0 {
				tmp = max(tmp, dp[j-R[i]])
			}
			dp[j] = tmp
		}
	}

	var res int

	for i := 0; i <= red; i++ {
		if dp[i] >= 0 {
			res = max(res, min(i, dp[i]))
		}
	}

	return res
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
