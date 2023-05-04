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
		A := readString(reader)[:n]
		B, C := solve(A)
		buf.WriteString(B)
		buf.WriteByte('\n')
		buf.WriteString(C)
		buf.WriteByte('\n')
	}

	fmt.Print(buf.String())
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

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')

	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
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

const INF = 1 << 30

func solve(A string) (string, string) {
	// A = B - C
	// sum(B) + sum(C) 最小
	// 从低位开始，如果 A[i] = 1, 可以 B[i] = 1, C[i] = 0, 这个的贡献是1
	// 或者 B[i] = 0, C[i] = 1, 但是需要借位  1000 - 1 = 0111
	n := len(A)

	B := make([]byte, n)
	C := make([]byte, n)
	for i := 0; i < n; i++ {
		B[i] = A[i]
		C[i] = '0'
	}

	for l := n - 1; l >= 0; {
		if B[l] == '0' {
			l--
			continue
		}
		r := l
		for l >= 0 && B[l] == '1' {
			l--
		}

		if l < 0 {
			break
		}
		// A[l] = '0'
		if r-l == 1 {
			continue
		}
		B[l] = '1'
		C[r] = '1'
		for i := l + 1; i <= r; i++ {
			B[i] = '0'
		}
	}

	return string(B), string(C)
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
