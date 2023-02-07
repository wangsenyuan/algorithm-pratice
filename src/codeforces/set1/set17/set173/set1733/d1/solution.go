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
		n, x, y := readThreeNums(reader)
		a := readString(reader)[:n]
		b := readString(reader)[:n]
		res := solve(x, y, a, b)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp
	return i
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
		if s[i] == '\n' || s[i] == '\r' {
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

const INF = 1 << 60

func solve(x int, y int, a string, b string) int64 {
	if a == b {
		return 0
	}

	n := len(a)
	// n <= 3000
	// x >= y
	// 一次操作 a[i] = 1 - a[i], a[j] = 1 - a[j] 同时改变两个，不变的是？ xor值是不变的
	// 00 => 11, 01 => 10
	ax := xor(a)
	bx := xor(b)

	if ax != bx {
		return -1
	}

	if n == 2 {
		return int64(x)
	}

	if n == 3 {
		// 001 vs 100
		if a[0] != b[0] {
			if a[2] != b[2] {
				return int64(y)
			}
		}
		// 010  001
		return int64(x)
	}

	// 然后我们来操作使得a = b,
	// count of diff % 2 == 0
	// 1100 vs 0000 => min(x, 2 * y)
	var diff []int

	for i := 0; i < n; i++ {
		if a[i] != b[i] {
			diff = append(diff, i)
		}
	}
	var res int64

	if len(diff) > 2 || diff[1] > diff[0]+1 {
		res = int64(len(diff)/2) * int64(y)
		// 0 pair half
		// 1 pair half + 1
		// ...
	} else {
		res = int64(min(2*y, x))
	}

	return res
}

func xor(a string) int {
	var res int
	for i := 0; i < len(a); i++ {
		res ^= int(a[i] - '0')
	}
	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
