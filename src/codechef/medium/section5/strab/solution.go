package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(scanner *bufio.Scanner) (a int) {
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &a)
	return
}

func readTwoNums(scanner *bufio.Scanner) (a int, b int) {
	res := readNNums(scanner, 2)
	a, b = res[0], res[1]
	return
}

func readNNums(scanner *bufio.Scanner, n int) []int {
	res := make([]int, n)
	x := 0
	scanner.Scan()
	for i := 0; i < n; i++ {
		for x < len(scanner.Bytes()) && scanner.Bytes()[x] == ' ' {
			x++
		}
		x = readInt(scanner.Bytes(), x, &res[i])
	}
	return res
}

func fillNNums(scanner *bufio.Scanner, n int, res []int) {
	x := 0
	scanner.Scan()
	for i := 0; i < n; i++ {
		for x < len(scanner.Bytes()) && scanner.Bytes()[x] == ' ' {
			x++
		}
		x = readInt(scanner.Bytes(), x, &res[i])
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		n, m := readTwoNums(scanner)
		scanner.Scan()
		X := scanner.Text()
		scanner.Scan()
		Y := scanner.Text()
		res := solve(n, m, X, Y)
		fmt.Println(res)
	}
}

const MOD = 1000000007

func pow(a, b int) int64 {
	A := int64(a)
	var r int64 = 1
	for b > 0 {
		if b&1 == 1 {
			r = (r * A) % MOD
		}
		A = (A * A) % MOD
		b >>= 1
	}

	return r
}

func solve(n, m int, X, Y string) int {
	dp := make([]int64, n+1)

	buf := make([]byte, n)

	valid := func(i int, l int) bool {
		if l < m {
			return true
		}

		s := string(buf[i : i+m])

		return s < X || s > Y
	}

	var F func(l int, i int) int64

	// F(s, n, i) = 24 * dp[i] + F(as, n, i - 1) * V(as) + F(bs, n, i - 1) * V(bs)

	F = func(l int, i int) int64 {
		if i < 0 {
			return 1
		}

		res := (24 * dp[i]) % MOD

		buf[i] = 'A'
		if valid(i, min(m, l)) {
			res += F(l+1, i-1)
		}

		buf[i] = 'B'

		if valid(i, min(m, l)) {
			res += F(l+1, i-1)
		}

		return res % MOD
	}

	dp[0] = 1

	for i := 1; i < m; i++ {
		dp[i] = pow(26, i)
	}

	for i := m; i <= n; i++ {
		dp[i] = F(1, i-1)
	}

	return int(dp[n])
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
