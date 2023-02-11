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

	tc := 1

	for tc > 0 {
		tc--
		var n int
		var x uint64
		s, _ := reader.ReadBytes('\n')
		pos := readInt(s, 0, &n)
		readUint64(s, pos+1, &x)
		res := solve(n, x)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			return s[:i]
		}
	}
	return s
}

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
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

const INF = 1 << 30

func solve(n int, x uint64) int {
	// len(x * y)= n
	// p = a * b => len(p) <= len(a) + len(b)
	// 100 * 100 = 10000

	var a uint64 = 1
	for i := 0; i < n; i++ {
		a *= 10
	}
	a /= 10
	b := a * 10

	dp := make(map[uint64]int)

	best := INF

	var dfs func(x uint64, w int) int

	dfs = func(x uint64, w int) int {
		if x >= b {
			return INF
		}

		if v, ok := dp[x]; ok {
			best = min(best, w+v)
			return v
		}

		if x >= a && x < b {
			dp[x] = 0
			return 0
		}
		dp[x] = INF
		tmp := x
		for tmp > 0 {
			r := tmp % 10
			if r > 1 {
				dp[x] = min(dp[x], dfs(x*r, w+1)+1)
			}
			tmp /= 10
		}

		return dp[x]
	}

	res := dfs(x, 0)

	if res < INF {
		return res
	}
	return -1
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
