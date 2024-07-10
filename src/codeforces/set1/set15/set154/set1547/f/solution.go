package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n := readNum(reader)
		a := readNNums(reader, n)
		res := solve(a)
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
func solve(a []int) int {
	n := len(a)

	var g int
	for _, num := range a {
		g = gcd(g, num)
	}

	h := 32 - bits.LeadingZeros32(uint32(n))
	dp := make([][]int, 2*n)

	for i := 0; i < 2*n; i++ {
		dp[i] = make([]int, h+1)
		dp[i][0] = a[i%n]
		for j := 1; j <= h && i >= 1<<(j-1); j++ {
			dp[i][j] = gcd(dp[i][j-1], dp[i-(1<<(j-1))][j-1])
		}
	}

	get := func(l int, r int) int {
		var res int
		for d := h; d >= 0; d-- {
			if (r-l+1)&(1<<d) > 0 {
				res = gcd(res, dp[r][d])
				r -= 1 << d
			}
		}
		return res
	}

	l, r := 0, 0
	for r < 2*n && get(l, r) > g {
		r++
	}
	// get(l, r) == g
	res := r - l + 1
	r++

	for r < 2*n {
		for get(l, r) == g {
			l++
		}
		l--
		res = max(res, r-l+1)
		r++
	}

	return res - 1
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
