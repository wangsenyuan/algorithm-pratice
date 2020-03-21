package main

import (
	"bufio"
	"fmt"
	"os"
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

func readThreeNums(scanner *bufio.Scanner) (a int, b int, c int) {
	res := readNNums(scanner, 3)
	a, b, c = res[0], res[1], res[2]
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

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	n := readNum(scanner)

	var res int

	for n > 0 {
		n--
		scanner.Scan()
		if check(scanner.Text()) {
			res++
		}
	}

	fmt.Println(res)
}

func solve(ss []string) int {
	var cnt int
	for _, s := range ss {
		if check(s) {
			cnt++
		}
	}
	return cnt
}

func check(s string) bool {
	i := strings.IndexByte(s, '+')
	A := strings.TrimSpace(s[:i])
	j := strings.IndexByte(s, '=')
	B := strings.TrimSpace(s[i+1 : j])
	C := strings.TrimSpace(s[j+1:])

	x := 2
	for i := 0; i < len(s); i++ {
		if s[i] == '+' || s[i] == '=' {
			continue
		}
		y := getDigit(s[i])
		x = max(x, y)
	}

	x++
	for b := x; b <= 16; b++ {
		aa := toDecimal(A, b)
		bb := toDecimal(B, b)
		cc := toDecimal(C, b)
		if aa+bb == cc {
			return true
		}
	}

	return false
}

func toDecimal(s string, b int) int64 {
	B := int64(b)
	var res int64
	for i := 0; i < len(s); i++ {
		x := int64(getDigit(s[i]))
		res = res*B + x
	}
	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func getDigit(c byte) int {
	if c >= 0 && c <= '9' {
		return int(c - '0')
	}
	return 10 + int(c-'a')
}
