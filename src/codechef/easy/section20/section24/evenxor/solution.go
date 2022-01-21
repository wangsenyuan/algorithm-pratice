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
		res := solve(n)
		for i := 0; i < n; i++ {
			buf.WriteString(fmt.Sprintf("%d ", res[i]))
		}
		buf.WriteByte('\n')
	}
	fmt.Println(buf.String())
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

func solve(n int) []int {
	// A[i] = 0 or A[i] = 1 << x | 1 << y
	// A[i] ^ A[j] =  (1 << x | 1 << y) | (1 << a | a << b)
	// C(19, 4) = 19 * 18 * 17 * 16 / 4 * 3 * 2
	A := make([]int, n)
	A[0] = 15
	for i := 1; i < n; i++ {
		A[i] = nextNum(A[i-1])
	}
	return A
}

func nextNum(x int) int {
	// right most set bit
	rightOne := x & -x

	// reset the pattern and set next higher bit
	// left part of x will be here
	nextHigherOneBit := x + rightOne

	// nextHigherOneBit is now part [D] of the above explanation.

	// isolate the pattern
	rightOnesPattern := x ^ nextHigherOneBit

	// right adjust pattern
	rightOnesPattern = (rightOnesPattern) / rightOne

	// correction factor
	rightOnesPattern >>= 2

	// rightOnesPattern is now part [A] of the above explanation.

	// integrate new pattern (Add [D] and [A])
	next := nextHigherOneBit | rightOnesPattern

	return next
}
