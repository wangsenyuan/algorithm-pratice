package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n := readNum(reader)
		x, y := solve(n)
		buf.WriteString(fmt.Sprintf("%d %d\n", x, y))
	}
	fmt.Print(buf.String())
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

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
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

func solve(n int) (x int, y int) {
	// x + y = n
	// and sum_digits(x) - sum_digits(y) <= 1

	if n%2 == 0 {
		x, y = n/2, n/2
		return
	}
	// n is odd
	// 3, 5, 7, 9
	num := n
	// a for x, and b for y
	var a, b int
	base := 1
	for num > 0 {
		last := num % 10
		if last%2 == 0 {
			// divides it equals
			x += num / 2 * base
			y += num / 2 * base
			return
		}
		if a >= b {
			// more one for b
			a += last / 2
			b += (last + 1) / 2
			x += last / 2 * base
			y += (last + 1) / 2 * base
		} else {
			// more one for a
			a += (last + 1) / 2
			b += last / 2
			x += (last + 1) / 2 * base
			y += last / 2 * base
		}
		num /= 10
		base *= 10
	}

	return
}
