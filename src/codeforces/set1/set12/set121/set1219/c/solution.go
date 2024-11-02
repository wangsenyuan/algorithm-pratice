package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	L := readNum(reader)
	A := readString(reader)
	res := solve(L, A)

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

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(reader *bufio.Reader) (a int) {
	bs, _ := reader.ReadBytes('\n')
	readInt(bs, 0, &a)
	return
}

func readTwoNums(reader *bufio.Reader) (a int, b int) {
	res := readNNums(reader, 2)
	a, b = res[0], res[1]
	return
}

func readThreeNums(reader *bufio.Reader) (a int, b int, c int) {
	res := readNNums(reader, 3)
	a, b, c = res[0], res[1], res[2]
	return
}

func readNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, n)
	x := 0
	bs, _ := reader.ReadBytes('\n')
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

var digits = "123456789"

func solve(L int, A string) string {
	if L == 1 {
		for i := 0; i < 9; i++ {
			s := strings.Repeat(digits[i:i+1], len(A))
			if s > A {
				return s
			}
		}
		return strings.Repeat("1", len(A)+1)
	}
	// L > 1, 所以可以1000的循环
	lx := (len(A) + L - 1) / L * L

	if lx > len(A) {
		return special(lx, L)
	}
	_9s := strings.Repeat("9", lx)
	if _9s == A {
		// 全部是9
		return special(lx+L, L)
	}
	pref := A[:L]
	s := strings.Repeat(pref, lx/L)
	if s > A {
		return s
	}
	// s <= A
	for i := L - 1; i >= 0; i-- {
		if pref[i] < '9' {
			buf := []byte(pref)
			buf[i]++
			for j := i + 1; j < L; j++ {
				buf[j] = '0'
			}
			pref = string(buf)
			return strings.Repeat(pref, lx/L)
		}
	}
	return ""
}

func special(n int, p int) string {
	buf := make([]byte, p)
	buf[0] = '1'
	for i := 1; i < p; i++ {
		buf[i] = '0'
	}
	res := make([]byte, n)
	for i := 0; i < n/p; i++ {
		copy(res[i*p:], buf)
	}
	return string(res)
}
