package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	// hint(105)
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)
	var buf bytes.Buffer
	for tc > 0 {
		tc--
		x, y := readTwoNums(reader)
		res := solve(x, y)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	fmt.Print(buf.String())
}

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			return s[:i]
		}
	}
	return s
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

const UINT_MAX = 32767

const MASK = (1 << 19) - 1

const BITS = 19

func mul(x, y, m int64) int64 {
	if x <= UINT_MAX && y <= UINT_MAX {
		return x * y % m
	}
	rslt := y * (x & MASK) % m

	x >>= BITS
	for x > 0 {
		y = y << BITS
		rslt = (rslt + y%m) * (x & MASK) % m
		x >>= BITS
	}
	return rslt
}

func isPrime(n int64) bool {
	if n < 29 {
		return n == 2 || n == 3 || n == 5 || n == 7 || n == 11 || n == 13 || n == 17 || n == 19 || n == 23
	}

	if !((n%2 > 0) && (n%3 > 0) && (n%5 > 0) && (n%7 > 0) && (n%11 > 0) && (n%13 > 0) && (n%17 > 0) && (n%19 > 0) && (n%23 > 0)) {
		return false
	}

	return test(n, 2) && test(n, 1215) && (n == 17431 || test(n, 34862) && (n == 3281359 || test(n, 574237825)))
}

func test(n int64, a int64) bool {
	n1 := n - 1
	s := n1
	var r int
	for (s & 1) == 0 {
		s >>= 1
		r++
	}
	t := a
	for s >>= 1; s > 0; s >>= 1 {
		a = mul(a, a, n)
		if s&1 == 1 {
			t = mul(t, a, n)
		}
	}
	if t == 1 || t == n1 {
		return true
	}
	r--
	for r > 0 {
		t = mul(t, t, n)
		if t == n1 {
			return true
		}
	}
	return false
}

func solve(x int, y int) int {

	get_num := func(x, y int) int64 {
		// need a fast way to get the number
		if abs(x) > abs(y) {
			if x > 0 {
				return 4*int64(x)*int64(x) - 3*int64(x) + int64(y)
			}
			return 4*int64(x)*int64(x) - int64(x) - int64(y)
		}

		if y > 0 {
			return 4*int64(y)*int64(y) - int64(y) - int64(x)
		}
		return 4*int64(y)*int64(y) - 3*int64(y) + int64(x)
	}

	num := get_num(x, y)

	if isPrime(num) {
		return 0
	}

	for dist := 1; ; dist++ {
		for u := 1; u <= dist; u++ {
			v := dist - u
			if isPrime(get_num(u+x, v+y)) {
				return dist
			}
		}
		for u := 1 - dist; u <= 0; u++ {
			v := dist + u
			if isPrime(get_num(u+x, v+y)) {
				return dist
			}
		}

		for u := -dist; u < 0; u++ {
			v := -dist - u
			if isPrime(get_num(u+x, v+y)) {
				return dist
			}
		}

		for u := 0; u < dist; u++ {
			v := -dist + u
			if isPrime(get_num(u+x, v+y)) {
				return dist
			}
		}
	}
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
