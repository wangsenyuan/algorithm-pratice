package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	s := readString(reader)
	x := readString(reader)
	n := readNum(reader)
	costs := make([]string, n)
	for i := 0; i < n; i++ {
		costs[i] = readString(reader)
	}
	res, s := solve(s, x, costs)
	fmt.Println(res)
	if res >= 0 {
		fmt.Println(s)
	}
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

func solve(s string, x string, costs []string) (int, string) {
	if len(s) != len(x) {
		return -1, ""
	}
	dp := make([][]int, 26)
	for i := 0; i < 26; i++ {
		dp[i] = make([]int, 26)
		for j := 0; j < 26; j++ {
			dp[i][j] = inf
		}
		dp[i][i] = 0
	}

	for _, cur := range costs {
		a, b := int(cur[0]-'a'), int(cur[2]-'a')
		var c int
		readInt([]byte(cur), 4, &c)
		dp[a][b] = min(dp[a][b], c)
	}

	for k := 0; k < 26; k++ {
		for i := 0; i < 26; i++ {
			for j := 0; j < 26; j++ {
				dp[i][j] = min(dp[i][j], dp[i][k]+dp[k][j])
			}
		}
	}

	buf := make([]byte, len(s))
	var res int
	for i := 0; i < len(s); i++ {
		a := int(s[i] - 'a')
		b := int(x[i] - 'a')
		var best int
		for c := 0; c < 26; c++ {
			if dp[a][best]+dp[b][best] > dp[a][c]+dp[b][c] {
				best = c
			}
		}
		res += dp[a][best] + dp[b][best]
		if res >= inf {
			return -1, ""
		}
		buf[i] = byte(best + 'a')
	}
	return res, string(buf)
}
