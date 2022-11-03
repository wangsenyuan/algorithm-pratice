package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer
	for tc > 0 {
		tc--
		s := readString(reader)
		res := solve(s)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

	for len(bs) == 0 || bs[0] == '\n' || bs[0] == '\r' {
		bs, _ = reader.ReadBytes('\n')
	}

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

func solve(s string) int64 {
	// flip once
	// 考虑在不做操作的情况下，A[i] = 1的贡献是 (i + 1) * (n - i)
	// 010  => 2 * 2 = 4, 01 10 010, 1
	// 然后将其flip后，它的贡献被取消，
	// A[j] = 0的贡献是 (j + 1) (n - j)
	// 假设选定了[L...R]
	// s[L...R] = 其中1的贡献
	// x[L...R] = 其中0的贡献
	// S[i] = 到i为止的1的贡献 s[l...r] = S[r] - S[l-1]
	// X[i]                  x[l...r] = X[r] - X[l-1]
	// 给定l的情况下，
	// 需要找到 + x[l...r] - s[l...r] 最大的值
	//        + (X[r] - X[l-1]) - (S[r] - S[l-1])
	//        + (X[r] - S[r]) - (X[l-1] - S[l-1])
	n := len(s)
	S := make([]int64, n)
	X := make([]int64, n)
	for i := 0; i < n; i++ {
		if s[i] == '1' {
			S[i] += int64(i+1) * int64(n-i)
		} else {
			X[i] += int64(i+1) * int64(n-i)
		}
		if i > 0 {
			S[i] += S[i-1]
			X[i] += X[i-1]
		}
	}

	// no operation
	res := S[n-1]

	var best int64 = math.MinInt64

	for l := n - 1; l >= 0; l-- {
		// fix l
		best = max(best, X[l]-S[l])

		var cur = S[n-1] + best
		if l > 0 {
			cur -= (X[l-1] - S[l-1])
		}

		res = max(res, cur)
	}

	return res
}

func max(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}
