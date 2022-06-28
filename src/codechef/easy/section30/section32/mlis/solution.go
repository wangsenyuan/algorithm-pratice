package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	// hint(105)
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)
	var buf bytes.Buffer
	for tc > 0 {
		tc--

		n := readNum(reader)
		A := readNNums(reader, n)
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

func solve(A []int) int {
	// n <= 2 * 10e5
	// LIS is easy, but after change [L...R] ?
	// L[i] = 到i为止最长的升子序列
	// R[i] = 从后到i为止最长的降子序列
	// 如果只是把L[i] 和 R[i+1] 连起来，是否可以cover所有的case呢?
	// 考虑一种情况 1, 2, 3, 2,  3, 2, 5, 3, 4, 7, 8, 9
	// 应该是可以的，因为假设中间段增加了x，使得它可以和L对接起来，但是无法和R对接，或者对接到了某个后续的R[j]
	// 则可以直接认为它是从i开始的R对接了
	n := len(A)
	stack := make([]int, n)
	var p int
	L := make([]int, n)
	for i := 0; i < n; i++ {
		j := search(p, func(j int) bool {
			return stack[j] >= A[i]
		})
		L[i] = j + 1
		if i > 0 {
			L[i] = max(L[i], L[i-1])
		}
		stack[j] = A[i]
		if p == j {
			p++
		}
	}

	best := p

	p = 0

	var R int

	for i := n - 1; i > 0; i-- {
		j := search(p, func(j int) bool {
			return stack[j] <= A[i]
		})
		R = max(R, j+1)
		best = max(best, R+L[i-1])
		stack[j] = A[i]
		if j == p {
			p++
		}
	}

	return best
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func search(n int, fn func(int) bool) int {
	l, r := 0, n
	for l < r {
		mid := (l + r) / 2
		if fn(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}
