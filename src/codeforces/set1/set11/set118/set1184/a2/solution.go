package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	readString(reader)
	s := readString(reader)
	res := solve(s)
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

func solve(s string) int {
	n := len(s)

	if s == strings.Repeat("0", n) {
		// 0 is ok
		return n
	}

	check := func(k int) bool {
		var sum int
		for i := 0; i < k; i++ {
			for j := i; j < n; j += k {
				sum += int(s[j] - '0')
			}
			if sum%2 != 0 {
				return false
			}
		}
		return true
	}

	var res int
	ok := make([]bool, n)
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			if check(i) {
				res++
				ok[i] = true
			}
			if i > 1 && i*i != n && check(n/i) {
				res++
				ok[n/i] = true
			}
		}
	}

	for i := 2; i < n; i++ {
		if !ok[i] {
			g := gcd(i, n)
			if ok[g] {
				res++
			}
		}
	}

	return res
}

func lcm(a, b int) int {
	g := gcd(a, b)
	return a / g * b
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

func reverse(s string) string {
	buf := []byte(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		buf[i], buf[j] = buf[j], buf[i]
	}
	return string(buf)
}
