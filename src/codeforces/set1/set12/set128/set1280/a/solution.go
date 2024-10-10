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
		x := readNum(reader)
		s := readString(reader)
		res := solve(s, x)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadBytes('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return string(s[:i])
		}
	}
	return string(s)
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

const mod = 1e9 + 7

func add(a, b int) int {
	a %= mod
	b %= mod
	a += b
	if a < 0 {
		a += mod
	}
	return a % mod
}

func mul(a, b int) int {
	a = add(a, 0)
	b = add(b, 0)
	return a * b % mod
}

var pool []byte

func init() {
	pool = make([]byte, 1_000_100)
}

func solve(s string, x int) int {
	copy(pool, s)
	n := len(s)

	ok := n < x

	for p := 0; p < x; p++ {
		v := int(pool[p] - '0')
		m := n
		// 重复这么多次
		n = add(n, mul(m-p-1, v-1))
		var tmp []byte
		var ln int
		if ok {
			// n <= x
			tmp = pool[p+1 : m]
			ln = m - p - 1
		}
		for v > 1 && ok {
			copy(pool[m:], tmp)
			m += ln
			ok = m < x
			v--
		}
	}

	return n
}
