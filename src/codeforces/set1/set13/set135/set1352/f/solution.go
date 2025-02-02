package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

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

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n0, n1, n2 := readThreeNums(reader)
		res := solve(n0, n1, n2)
		buf.WriteString(res)
		buf.WriteByte('\n')
	}

	fmt.Print(buf.String())
}

func solve(n0, n1, n2 int) string {
	// n2都是11
	// n0都是00
	var buf bytes.Buffer
	if n0 == 0 && n2 == 0 {
		// 101  这里是两个n1, 不是1个
		for i := 0; i <= n1; i++ {
			buf.WriteByte(byte('0' + (i & 1)))
		}

		return buf.String()
	}
	if n0 == 0 {
		// n2 > 0
		for i := 0; i < n2; i++ {
			buf.WriteString("1")
		}
		buf.WriteString("1")
		for i := 0; i < n1; i++ {
			// 010101
			buf.WriteByte(byte('0' + (i & 1)))
		}
		return buf.String()
	}
	if n2 == 0 {
		for i := 0; i < n0; i++ {
			buf.WriteString("0")
		}
		buf.WriteString("0")
		// 10101
		for i := 0; i < n1; i++ {
			if i&1 == 0 {
				buf.WriteByte('1')
			} else {
				buf.WriteByte('0')
			}
		}

		return buf.String()
	}
	// n0 > 0, n2 > 0, n1 > 0
	for i := 0; i < n0; i++ {
		buf.WriteString("0")
	}
	buf.WriteString("0")
	for i := 0; i < n2; i++ {
		buf.WriteString("1")
	}
	buf.WriteString("1")
	n1--
	for i := 0; i < n1; i++ {
		// 0101
		buf.WriteByte(byte('0' + (i & 1)))
	}
	return buf.String()
}
