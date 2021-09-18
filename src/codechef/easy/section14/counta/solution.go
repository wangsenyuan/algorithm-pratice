package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var buf bytes.Buffer

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n := readNum(reader)
		B := readNNums(reader, n-1)
		res := solve(B)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	fmt.Print(buf.String())
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

const M = 1000000007

const X = 100000

func solve(B []int) int {
	n := len(B)
	/**
	denote a as ways filling A[..i], ending by A[i] = B[i]
	denote b as ways filling A[..i+1], ending by A[i+1] = B[i]
	*/
	var a int64 = 1
	var b int64 = int64(X - B[0])

	for i := 1; i < n; i++ {
		var c int64
		// if B[i] >= B[i-1], then can set A[i] = B[i], else, if B[i] < B[i-1], then A[i-1] = B[i-1], A[i] = B[i], then B[i-1] = min(A[i-1], A[i]), which is invalid
		if B[i] >= B[i-1] {
			c = a
		}
		if B[i] == B[i-1] {
			// if B[i] == B[i-1], then A[i] = B[i-1] is valid (which is b), and set A[i] = B[i] is also valid
			c = modAdd(c, b)
		}
		// let try to set A[i+1] = B[i]
		// if A[i-1] = B[i-1], A[i] >= A[i-1] & A[i] > A[i+1]
		var d int64 = a * int64(X-max(B[i-1], B[i]+1)+1) % M
		// if A[i] = B[i-1], and B[i] < B[i-1]
		if B[i] < B[i-1] {
			d = modAdd(d, b)
		}
		a, b = c, d
	}

	res := int64(X-B[n-1]+1)*a%M + b
	return int(res % M)
}

func modAdd(a, b int64) int64 {
	a += b
	if a >= M {
		a -= M
	}
	return a
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
