package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n, x := readTwoNums(reader)
		res := solve(n, x)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

func solve(n int, x int) int {
	var res int

	for a := 1; a < min(n, x); a++ {
		for b := 1; b*a < n && a+b < x; b++ {
			c := min((n-a*b)/(a+b), x-(a+b))
			res += c
		}
	}

	return res
}

func solve1(n int, x int) int {
	sn := int(math.Sqrt(float64(n) / 3))
	var res int
	for a := 1; a <= sn; a++ {
		// a * b < n
		for b := a; a*b < n && a+b < x; b++ {

			for c := b; a+b+c <= x && a*b+b*c+a*c <= n; c++ {
				// 这里res要计算一下
				if a < b && b < c {
					// a, b, c, a, c, b.... 六种
					res += 6
				} else if a == b && b < c {
					// a, a, c, a, c, a, c, a, a
					res += 3
				} else if a < b && b == c {
					res += 3
				} else {
					// a = b, b = c
					res++
				}
			}
		}
	}

	return res
}
