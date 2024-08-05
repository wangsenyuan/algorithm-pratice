package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	readString(reader)
	s1 := readString(reader)
	s2 := readString(reader)
	res := solve(s1, s2)

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

const mod = 998244353

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

func inverse(num int) int {
	return pow(num, mod-2)
}

func solve(s1 string, s2 string) int {
	n := len(s1)

	F := make([]int, n+1)
	F[0] = 1
	for i := 1; i <= n; i++ {
		F[i] = mul(i, F[i-1])
	}

	I := make([]int, n+1)
	I[n] = inverse(F[n])
	for i := n - 1; i >= 0; i-- {
		I[i] = mul(i+1, I[i+1])
	}

	a := make([]int, 26)
	for i := 0; i < len(s1); i++ {
		a[int(s1[i]-'A')]++
	}

	b := make([]int, 26)
	for i := 0; i < len(s2); i++ {
		b[int(s2[i]-'A')]++
	}

	dp := []int{1}

	for i := 25; i >= 0; i-- {
		cur := make([]int, b[i]+1)
		for j := 0; j <= b[i]; j++ {
			if b[i]-j <= a[i] {
				// x是i+1的数量
				x := a[i] - (b[i] - j)
				if x < len(dp) {
					u := mul(I[x], I[b[i]-j])
					cur[j] = mul(dp[x], u)
				}
			}
		}
		for j := len(cur) - 2; j >= 0; j-- {
			cur[j] = add(cur[j], cur[j+1])
		}
		dp = cur
	}

	return mul(dp[0], F[n])
}
