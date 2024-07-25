package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	n := readNum(reader)
	x, c := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		x[i], c[i] = readTwoNums(reader)
	}

	res := solve(x, c)
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

type marble struct {
	pos  int
	cost int
}

func solve(x []int, c []int) int {
	n := len(x)
	marbles := make([]marble, n)
	for i := 0; i < n; i++ {
		marbles[i] = marble{x[i], c[i]}
	}

	sort.Slice(marbles, func(i, j int) bool {
		return marbles[i].pos < marbles[j].pos
	})
	// dp[i] = min cost end at i
	dp := make([]int, n+1)
	dp[0] = 0
	for i := 1; i <= n; i++ {
		dp[i] = 1 << 50
		var sum int
		for j := i - 1; j >= 0; j-- {
			sum += marbles[j].pos
			tmp := sum - (i-j)*marbles[j].pos + marbles[j].cost
			dp[i] = min(dp[i], dp[j]+tmp)
		}
	}

	return dp[n]
}
