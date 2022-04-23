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
		x, y := readTwoNums(reader)
		ok, s1, s2 := solve(x, y)
		if !ok {
			buf.WriteString("-1\n")
		} else {
			buf.WriteString(s1)
			buf.WriteByte('\n')
			buf.WriteString(s2)
			buf.WriteByte('\n')
		}
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

func solve(x, y int) (bool, string, string) {
	if x&1 == 1 && y&1 == 1 {
		return false, "", ""
	}

	if x == 1 || y == 1 {
		return false, "", ""
	}

	var a, b byte = 'a', 'b'
	if y&1 == 1 {
		a, b = b, a
		x, y = y, x
	}

	if x&1 == 0 && y&1 == 0 {
		s1 := construct(a, x, b, y)
		s2 := construct(b, y, a, x)
		return true, s1, s2
	}
	// x & 1 == 1
	s1 := construct(a, x-1, b, y)
	s2 := construct(b, y, a, x-1)
	n := x + y
	s1 = insertAtMid(s1, a, n)
	s2 = insertAtMid(s2, a, n)
	return true, s1, s2
}

func insertAtMid(s string, a byte, n int) string {
	var buf bytes.Buffer
	buf.WriteString(s[:n/2])
	buf.WriteByte(a)
	buf.WriteString(s[n/2:])
	return buf.String()
}

func construct(first byte, first_cnt int, second byte, second_cnt int) string {
	res := make([]byte, first_cnt+second_cnt)
	i, j := 0, len(res)-1
	for i < first_cnt/2 {
		res[i] = first
		res[j] = first
		i++
		j--
	}
	for i < j {
		res[i] = second
		res[j] = second
		i++
		j--
	}
	return string(res)
}
