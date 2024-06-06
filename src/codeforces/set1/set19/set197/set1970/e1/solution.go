package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	m, n := readTwoNums(reader)
	s := readNNums(reader, m)
	l := readNNums(reader, m)

	res := solve(n, s, l)

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

const mod = 1e9 + 7

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func sub(a, b int) int {
	return add(a, mod-b)
}

func mul(a, b int) int {
	return a * b % mod
}

func solve(n int, s []int, l []int) int {
	m := len(s)

	dp := make([]int, m)
	// dp[i] = 在当前处于位置i时的状态数
	dp[0] = 1
	sp := mul(1, s[0])
	lp := mul(1, l[0])

	fp := make([]int, m)

	for d := 0; d < n; d++ {
		for i := 0; i < m; i++ {
			// dp[i] = dp[j] * (((l[j] + s[j]) * s[i] + s[j] * (l[i] + s[i]))
			tmp1 := mul(sp, add(l[i], s[i]))
			tmp2 := mul(lp, s[i])
			fp[i] = add(tmp1, tmp2)
		}

		sp = 0
		lp = 0
		for i := 0; i < m; i++ {
			dp[i] = fp[i]
			sp = add(sp, mul(dp[i], s[i]))
			lp = add(lp, mul(dp[i], l[i]))
			fp[i] = 0
		}
	}

	var res int

	for i := 0; i < m; i++ {
		res = add(res, dp[i])
	}

	return res
}
