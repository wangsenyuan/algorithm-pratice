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

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		scanner.Scan()
		s := scanner.Text()
		fmt.Println(solve(s))
	}
}

func solve(s string) int {
	n := len(s)

	if n&1 == 0 {
		var j int
		for i := 0; i < n/2-1; i++ {
			cache[j] = s[i]
			j++
		}
		a2 := string(cache[:j])
		j = 0
		for i := n - 1; i >= n/2; i-- {
			cache[j] = s[i]
			j++
		}
		a1 := string(cache[:j])

		j = 0
		for i := 0; i < n/2; i++ {
			cache[j] = s[i]
			j++
		}
		b1 := string(cache[:j])
		j = 0
		for i := n - 1; i > n/2; i-- {
			cache[j] = s[i]
			j++
		}
		b2 := string(cache[:j])
		return getAnswer(a1, a2) + getAnswer(b1, b2) + checkPalindrom(s)
	} else {
		var j int
		for i := 0; i <= n/2; i++ {
			cache[j] = s[i]
			j++
		}
		a1 := string(cache[:j])

		j = 0
		for i := n - 1; i > n/2; i-- {
			cache[j] = s[i]
			j++
		}
		a2 := string(cache[:j])
		j = 0
		for i := 0; i < n/2; i++ {
			cache[j] = s[i]
			j++
		}
		b2 := string(cache[:j])
		j = 0
		for i := n - 1; i >= n/2; i-- {
			cache[j] = s[i]
			j++
		}
		b1 := string(cache[:j])

		return getAnswer(a1, a2) + getAnswer(b1, b2)
	}
}

const MX = 1e5 + 10

var pre [MX]bool
var suf [MX]bool

var cache [MX]byte

func getAnswer(a, b string) int {
	n := len(b)
	pre[0] = true
	suf[n+1] = true

	for i := 0; i < n; i++ {
		pre[i+1] = pre[i] && a[i] == b[i]
	}
	for i := n - 1; i >= 0; i-- {
		suf[i+1] = suf[i+2] && b[i] == a[i+1]
	}

	var ret int
	for i := -1; i < n; i++ {
		if pre[i+1] && suf[i+2] {
			ret++
		}
	}
	return ret
}

func checkPalindrom(s string) int {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return 0
		}
	}
	return 1
}
