package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var tmp int64
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp
	return i
}

func readNInt64Nums(scanner *bufio.Scanner, n int) []int64 {
	res := make([]int64, n)
	x := -1
	scanner.Scan()
	for i := 0; i < n; i++ {
		x = readInt64(scanner.Bytes(), x+1, &res[i])
	}
	return res
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	line, _ := reader.ReadString('\n')

	tc, _ := strconv.Atoi(strings.TrimSpace(line))

	for tc > 0 {
		tc--
		s, _ := reader.ReadString('\n')
		if len(s) == 0 {
			fmt.Println("no input read")
		} else {
			fmt.Println(solve(strings.TrimSpace(s)))
		}
	}
}

func solve(s string) string {
	var l, r int = 0, len(s) - 1
	for l <= r && s[l] == s[r] {
		l++
		r--
	}
	if l > r {
		return s
	}
	// l < r
	x := s[l : r+1]

	a := findLongestPrefixPalindrome(x)
	b := findLongestPrefixPalindrome(reverse(x))

	if len(a) < len(b) {
		a, b = b, a
	}

	return s[:l] + a + s[r+1:]
}

const BASE = 26
const MOD = 1000000007

func findLongestPrefixPalindrome(s string) string {
	n := len(s)
	var m int
	var a int64
	var b int64
	var X int64 = 1

	for i := 0; i < n; i++ {
		a = (a*BASE + int64(s[i]-'a')) % MOD
		b = (X*int64(s[i]-'a') + b) % MOD
		X = (X * BASE) % MOD
		if a != b {
			continue
		}

		j, k := 0, i

		for j < k && s[j] == s[k] {
			j++
			k--
		}

		if j >= k {
			m = i + 1
		}
	}
	return s[:m]
}

func reverse(s string) string {
	buf := []byte(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		buf[i], buf[j] = buf[j], buf[i]
	}
	return string(buf)
}
