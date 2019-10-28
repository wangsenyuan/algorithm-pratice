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

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	tc := readNum(scanner)

	for tc > 0 {
		tc--
		scanner.Scan()
		r := solve(scanner.Text())
		if r {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}

func solve(s string) bool {
	// 1 * 10 * 20 * 10 * 20 == s
	n := len(s)
	i := n - 1
	for i >= 0 && s[i] == '0' {
		i--
	}
	if i < 0 {
		// 0
		return false
	}
	if s[i] == '1' {
		// 100000
		return i == 0
	}

	cnt := n - i - 1

	x := uint64(s[i] - '0')
	base := uint64(10)
	for i > 0 {
		y := uint64(s[i-1] - '0')
		x = y*base + x
		base *= 10
		i--
	}
	for x != 1 && cnt > 0 {
		if x&1 == 1 {
			return false
		}

		cnt--
		x >>= 1
	}
	return x == 1
}
