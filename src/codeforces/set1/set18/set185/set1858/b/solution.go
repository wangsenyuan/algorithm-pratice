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
		n, m, d := readThreeNums(reader)
		s := readNNums(reader, m)
		res := solve(n, d, s)
		x := fmt.Sprintf("%v", res)
		x = x[1 : len(x)-1]
		buf.WriteString(fmt.Sprintf("%s\n", x))
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

func solve(n int, d int, s []int) []int {
	m := len(s)
	r := make([]int, m+2)
	r[0] = 1 - d
	copy(r[1:], s)
	r[m+1] = n + 1
	m += 2
	ans := 1 << 60
	var res []int
	for i := 1; i < m-1; i++ {
		A := r[i] - r[i-1] - 1
		B := r[i+1] - r[i] - 1
		C := r[i+1] - r[i-1] - 1
		D := C/d - (A/d + B/d)
		if D < ans {
			ans = D
			res = res[:0]
		}
		if D == ans {
			res = append(res, r[i])
		}
	}
	var tmp int
	for i := 1; i < len(r); i++ {
		tmp += (r[i] - r[i-1] - 1) / d
	}

	tmp += len(r) - 2

	return []int{ans + tmp - 1, len(res)}
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
