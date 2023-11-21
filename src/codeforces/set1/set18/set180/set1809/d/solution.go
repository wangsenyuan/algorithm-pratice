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

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')

	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
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

const cost = 1e12

func solve(s string) int {
	// 假设公共进行了x次操作，其中b次第二类操作，那么所有的cost = cost * x + b
	n := len(s)
	var front int
	for front < n && s[front] == '0' {
		front++
	}
	s = s[front:]
	tail := len(s) - 1
	for tail >= 0 && s[tail] == '1' {
		tail--
	}
	s = s[:tail+1]

	if len(s) == 0 {
		return 0
	}

	// 假设 s[i] = 1, s[i+1] = 0
	// 那么交换它们，并把i前面的1都删掉, i+1后面的0都删掉
	n = len(s)
	dp := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		if s[i] == '0' {
			dp[i]++
		}
		dp[i] += dp[i+1]
	}
	// 把少的那方全部删掉
	best := min(dp[0], n-dp[0]) * (cost + 1)
	var cnt int
	for i := 0; i < n; i++ {
		if i+1 < n && s[i] == '1' && s[i+1] == '0' {
			// 一次交换
			tmp := (cnt+dp[i+2])*(cost+1) + cost
			best = min(best, tmp)
		}
		// 删除它前面的1，和它后面的0
		tmp := (cnt + dp[i]) * (cost + 1)
		best = min(best, tmp)

		if s[i] == '1' {
			cnt++
		}
	}

	return best
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
