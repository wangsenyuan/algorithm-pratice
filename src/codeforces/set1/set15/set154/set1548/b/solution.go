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
	// gcd(l...r) > 1
	n := len(a)
	b := make([]int, n-1)

	for i := 0; i < n-1; i++ {
		b[i] = abs(a[i+1] - a[i])
	}
	// gcd(l...r) > 1
	h := bits.Len(uint(n))
	// dp[i][j] = gcd(i, i + 1 << j - 1)
	dp := make([][]int, n)
	dp[n-1] = make([]int, h)
	for i := n - 2; i >= 0; i-- {
		dp[i] = make([]int, h)
		dp[i][0] = b[i]
		for j := 1; j < h; j++ {
			if i+(1<<(j-1)) < n {
				dp[i][j] = gcd(dp[i][j-1], dp[i+(1<<(j-1))][j-1])
			}
		}
	}

	get := func(l int, r int) int {
		var res int
		for i := h - 1; i >= 0; i-- {
			if l+(1<<i) <= r {
				res = gcd(res, dp[l][i])
				l += 1 << i
			}
		}
		return gcd(res, dp[l][0])
	}

	var best int

	for l, r := n-2, n-2; l >= 0; l-- {
		for r >= l && get(l, r) == 1 {
			r--
		}
		best = max(best, r-l+1)
	}

	return best + 1
}

func abs(num int) int {
	return max(num, -num)
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
