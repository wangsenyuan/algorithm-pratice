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
		A := readNNums(reader, n+1)
		res := solve(A)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

const MOD = 1000000007

const MAX_N = 100010

var P []int

func init() {
	P = make([]int, MAX_N)
	P[0] = 1
	for i := 1; i < MAX_N; i++ {
		P[i] = modMul(2, P[i-1])
	}
}

func solve(A []int) int {
	n := len(A)
	suf := make([]int, n+1)

	for i := n; i >= 1; i-- {
		suf[i] = modMul(A[i-1], P[n-i])
		if i < n {
			suf[i] = modAdd(suf[i], suf[i+1])
		}
	}

	var res int
	p2 := 1

	for i := 1; i < n; i++ {
		count := modMul(A[i-1], p2)
		count = modMul(count, suf[i+1])
		res = modAdd(res, count)
		if i >= 2 {
			p2 = modMul(p2, 2)
		}
	}

	return modMul(2, res)
}

func modAdd(a, b int) int {
	a += b
	if a >= MOD {
		a -= MOD
	}
	return a
}

func modMul(a, b int) int {
	A := int64(a) * int64(b)
	A %= MOD
	return int(A)
}

func bruteForce(A []int) int {
	n := len(A)
	buf := make([]int, 2*n)

	var loop func(i int, front, end int, cur int) int
	loop = func(i int, front, end int, cur int) int {
		if i == n-1 {
			// it is the last one
			res := A[i] * buf[front]
			res += A[i] * buf[end]

			return res + cur
		}
		// i < n - 1
		buf[front-1] = A[i]
		res := loop(i+1, front-1, end, cur+A[i]*buf[front])
		buf[end+1] = A[i]
		res += loop(i+1, front, end+1, cur+A[i]*buf[end])
		return res
	}
	buf[n] = A[0]
	return loop(1, n, n, 0)
}
