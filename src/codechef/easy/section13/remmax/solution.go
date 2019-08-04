package main

import (
	"bufio"
	"bytes"
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
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)
	for tc > 0 {
		tc--
		scanner.Scan()
		s := scanner.Text()
		fmt.Println(solve(s))
	}
}

func solve(s string) string {
	n := len(s)

	if n <= 1 {
		return s
	}

	// 1024 => 1019 => 9101
	// 1125 => 1119 => 9111
	// 110 => 109 => 901
	// 111 => 109 => 901
	// 1123 => 1099 => 9901
	var buf bytes.Buffer

	if s[0] > '1' {
		i := 1

		for i < n && s[i] == '9' {
			i++
		}

		if i == n {
			return s
		}

		buf.WriteByte(s[0] - 1)
		for i := 0; i < n-1; i++ {
			buf.WriteByte('9')
		}
		return buf.String()
	}

	carry := false
	buf.WriteByte('1')

	for i := 1; i < n; i++ {
		if s[i] == '0' {
			buf.WriteByte('0')
		} else {
			carry = true
			j := i + 1
			for j < n && s[j] == '9' {
				j++
			}
			if j == n {
				// all 0
				buf.WriteByte(s[i])
			} else {
				buf.WriteByte(s[i] - 1)
			}
			j = i + 1
			for j < n {
				buf.WriteByte('9')
				j++
			}
			break
		}
	}

	if !carry {
		buf.Reset()
		for i := 1; i < n; i++ {
			buf.WriteByte('9')
		}

	}

	return buf.String()
}
