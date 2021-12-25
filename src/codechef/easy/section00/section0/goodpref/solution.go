package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	tmp := 0
	for i < len(bytes) && bytes[i] == ' ' {
		i++
	}
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readOneNum(scanner *bufio.Scanner) (a int) {
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &a)
	return
}

func readTwoNums(scanner *bufio.Scanner) (a int, b int) {
	scanner.Scan()
	x := readInt(scanner.Bytes(), 0, &a)
	readInt(scanner.Bytes(), x+1, &b)
	return
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	t := readOneNum(scanner)

	for t > 0 {
		scanner.Scan()
		bs := scanner.Bytes()
		i := len(bs) - 1
		for i > 0 && bs[i] != ' ' {
			i--
		}

		s := bs[:i]
		var n int
		readInt(bs, i+1, &n)

		fmt.Println(solve(s, n))

		t--
	}
}

func solve(s []byte, n int) int64 {
	var a, b int
	for i := 0; i < len(s); i++ {
		if s[i] == 'a' {
			a++
		} else {
			b++
		}
	}

	if a > b {
		return solve1(s, a, b, n)
	} else if a == b {
		return solve2(s, a, b, n)
	} else {
		return solve3(s, a, b, n)
	}
}

func solve1(s []byte, a, b, n int) int64 {
	// m(a - b) <= len(s)
	m := (len(s) + a - b - 1) / (a - b)

	// after m copies, a always more than b

	var res int64
	if n > m+1 {
		res = int64(n-m-1) * int64(len(s))
	}

	var diff int

	for i := 0; i <= m && i < n; i++ {
		for i := 0; i < len(s); i++ {
			if s[i] == 'a' {
				diff++
			} else {
				diff--
			}
			if diff > 0 {
				res++
			}
		}
	}

	return res
}

func solve2(s []byte, a, b, n int) int64 {
	var res int64
	var x, y int
	for i := 0; i < len(s); i++ {
		if s[i] == 'a' {
			x++
		} else {
			y++
		}
		if x > y {
			res += int64(n)
		}
	}
	return res
}

func solve3(s []byte, a, b, n int) int64 {
	// m(b - a) <= len(s)
	m := len(s) / (b - a)

	var res int64

	var diff int
	for i := 0; i <= m && i < n; i++ {
		for i := 0; i < len(s); i++ {
			if s[i] == 'a' {
				diff--
			} else {
				diff++
			}
			if diff < 0 {
				res++
			}
		}
	}
	return res
}
