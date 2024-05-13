package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m := readTwoNums(reader)

	res := solve(n, m)

	fmt.Println(res)
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

func solve(n int, m int) int {
	// m <= 1000
	dp := make([]int, m)
	// dp[j] = i * i % m = j 的个数
	// 比如 n = 9, m = 3
	// 3, 6, 9
	// 但是余数  = 1 时， (1, 4, 7)
	cnt := n / m
	r := n % m
	for i := 0; i < m; i++ {
		// 在1...n中间， 有多少个数 % m = i
		tmp := cnt
		if r >= i && i > 0 {
			tmp++
		}
		dp[i*i%m] += tmp
	}

	res := dp[0] * dp[0]

	for i := 1; i < m; i++ {
		// i是行 , j 是列
		j := m - i
		res += dp[i] * dp[j]
	}

	return res
}
