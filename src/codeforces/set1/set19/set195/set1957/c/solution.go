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
		n, k := readTwoNums(reader)

		already := make([][]int, k)
		for i := 0; i < k; i++ {
			already[i] = readNNums(reader, 2)
		}
		res := solve(n, already)

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

const N = 3*100000 + 10
const mod = 1e9 + 7

var F [N]int
var I [N]int

func init() {
	F[0] = 1

	for i := 1; i < N; i++ {
		F[i] = mul(i, F[i-1])
	}

	I[N-1] = pow(F[N-1], mod-2)

	for i := N - 2; i >= 0; i-- {
		I[i] = mul(i+1, I[i+1])
	}

}
func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func mul(a, b int) int {
	return a * b % mod
}

func pow(a, b int) int {
	r := 1
	for b > 0 {
		if b&1 == 1 {
			r = mul(r, a)
		}
		a = mul(a, a)
		b >>= 1
	}
	return r
}

func nCr(n int, r int) int {
	if n < r {
		return 0
	}
	return mul(F[n], mul(I[r], I[n-r]))
}

func reverse(num int) int {
	return pow(num, mod-2)
}

func solve(n int, already [][]int) int {
	m := n
	for _, cur := range already {
		if cur[0] == cur[1] {
			m--
		} else {
			m -= 2
		}
	}

	var res int

	for i := 0; i <= m; i++ {
		j := m - i
		if j%2 == 0 {
			x := nCr(m, i)
			y := F[m-i]
			z := F[(m-i)/2]
			res = add(res, mul(x, mul(y, reverse(z))))
		}
	}
	return res
}

func solve1(n int, already [][]int) int {
	m := n
	for _, cur := range already {
		if cur[0] == cur[1] {
			m--
		} else {
			m -= 2
		}
	}
	dp := make([]int, m+3)

	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= m; i++ {
		// place at (i, i)
		tmp := mul(i-1, dp[i-2])
		tmp = mul(2, tmp)
		dp[i] = add(dp[i-1], tmp)
	}

	return dp[m]
}
