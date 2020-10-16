package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString('\n')
	var i int
	for i < len(s) && (s[i] >= 'a' && s[i] <= 'z' || s[i] >= 'A' && s[i] <= 'Z' || s[i] >= '0' && s[i] <= '9') {
		i++
	}
	fmt.Println(solve(s[:i]))
}

const P1 = 31

//const P2 = 41
const M1 = 1000000007

//const M2 = 1000000009

func solve(s string) int64 {
	n := len(s)
	var h1 int64
	//var h2 int64
	var h3 int64
	//var h4 int64
	var p1 int64 = 1
	//var p2 int64 = 1

	dp := make([]int64, n)
	var res int64
	for i := 0; i < n; i++ {
		x := int64(s[i]-'a') + 1
		h1 = (h1*P1 + x) % M1
		//h2 = (h2*P2 + x) % M2

		h3 = (h3 + (x*p1)%M1) % M1
		//h4 = (h4 + (x*p2)%M2) % M2

		p1 = (p1 * P1) % M1
		//p2 = (p2 * P2) % M2

		if h1 == h3 {
			dp[i] = 1
			if i > 0 {
				if i&2 == 1 {
					dp[i] += dp[i/2]
				} else {
					dp[i] += dp[(i-1)/2]
				}
			}
		}
		res += dp[i]
	}

	return res
}

func solve1(s string) int64 {
	p := manachers(s)

	n := len(s)

	dp := make([]int64, n)
	dp[0] = 1
	var res int64 = 1

	for i := 2; i <= n; i++ {
		if p[i-1] == 0 {
			dp[i-1] = 1
			//prefix s[0..i] is palindrome
			if p[i/2-1] == 0 {
				dp[i-1] += dp[i/2-1]
			}
			res += dp[i-1]
		}
	}

	return res
}

func manachers(s string) []int {
	ss := modify(s)
	n := len(ss)
	P := make([]int, n)
	var c, r int

	for i := 0; i < n; i++ {
		mirror := 2*c - r

		if i < r {
			P[i] = min(r-i, P[mirror])
		}
		a := i + 1 + P[i]
		b := i - (1 + P[i])
		for a < n && b >= 0 && ss[a] == ss[b] {
			P[i]++
			a++
			b--
		}

		if i+P[i] > r {
			r = i + P[i]
			c = i
		}
	}
	f := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		f[i] = i
	}
	for i := 0; i < n; i++ {
		//l := P[i]
		// #a#a#b#
		a := i + P[i]
		b := i - P[i]
		a--
		b++
		if a > b || i%2 != 0 {
			f[a/2] = min(f[a/2], b/2)
		}
	}
	return f
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func modify(s string) string {
	var buf bytes.Buffer

	for i := 0; i < len(s); i++ {
		buf.WriteByte('#')
		buf.WriteByte(s[i])
	}
	buf.WriteByte('#')
	return buf.String()
}
