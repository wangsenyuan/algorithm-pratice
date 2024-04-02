package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n := readNum(reader)
	a := readNNums(reader, n)
	res := solve(a)

	fmt.Println(res)
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

func solve(a []int) int {
	n := len(a)

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		dp[i][i] = 1
	}

	for ln := 2; ln <= n; ln++ {
		for i := 0; i+ln <= n; i++ {
			j := i + ln - 1
			dp[i][j] = max(dp[i+1][j]+checkValue(a[i] == 2), dp[i][j-1]+checkValue(a[j] == 1))
		}
	}

	res := dp[0][n-1]

	suf := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		suf[i] = checkValue(a[i] == 2) + suf[i+1]
	}

	res = max(res, suf[0])

	var prev int
	for i := 0; i < n; i++ {
		res = max(res, prev+suf[i])
		for j := i + 1; j < n; j++ {
			res = max(res, prev+suf[j+1]+dp[i][j])
		}
		prev += checkValue(a[i] == 1)
	}

	res = max(res, prev)

	return res
}

func checkValue(b bool) int {
	if b {
		return 1
	}
	return 0
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
