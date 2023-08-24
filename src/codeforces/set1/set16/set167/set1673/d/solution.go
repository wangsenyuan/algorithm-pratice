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
		B := readNNums(reader, 3)
		C := readNNums(reader, 3)
		res := solve(B, C)
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

func gcd(a, b int64) int64 {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int64) int64 {
	return a * b / gcd(a, b)
}

const mod = 1000000007

func modMul(a, b int64) int64 {
	return a * b % mod
}

func modAdd(a, b int64) int64 {
	a += b
	return a % mod
}

func solve(B []int, C []int) int {
	// b, q, y
	// c, r, z
	// r % q = 0
	// r % s = 0
	b, q, y := int64(B[0]), int64(B[1]), int64(B[2])
	c, r, z := int64(C[0]), int64(C[1]), int64(C[2])

	e := b + (y-1)*q
	f := c + (z-1)*r

	if c < b || c > e || f < b || f > e || r%q != 0 || (c-b)%q != 0 {
		return 0
	}
	if c-r < b || f+r > e {
		return -1
	}
	var ans int64
	for p := int64(1); p*p <= r; p++ {
		if r%p == 0 {
			if lcm(p, q) == r {
				ans = modAdd(ans, modMul(r/p, r/p))
			}
			if p*p < r && lcm(r/p, q) == r {
				ans = modAdd(ans, modMul(p, p))
			}
		}
	}
	return int(ans)
}
