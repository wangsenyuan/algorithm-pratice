package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n := readNum(reader)
	p := readNNums(reader, n)
	res := solve(p)
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

func solve(a []int) int {
	n := len(a)
	ln := bits.Len(uint(n))

	sum := make([]int, ln+1)

	for i, x := 1, n/2; i <= ln; i++ {
		sum[i] = sum[i-1] + x
		x /= 2
	}

	dp := make([][]int, ln+1)
	for i := 0; i <= ln; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = -1
		}
	}

	var f func(cnt int, pos int) (res int)

	f = func(cnt int, pos int) (res int) {
		if dp[cnt][pos] != -1 {
			return dp[cnt][pos]
		}

		defer func() {
			dp[cnt][pos] = res
		}()

		if a[pos] < 0 {
			return
		}
		rem := sum[cnt] - pos
		res = 1 << 30
		if cnt < ln {
			res = f(cnt+1, pos+1) + a[pos]
		}
		if rem > 0 {
			res = min(res, f(cnt, pos+1))
		}
		return
	}

	reverse(a)

	return f(0, 0)
}

func reverse(a []int) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}
