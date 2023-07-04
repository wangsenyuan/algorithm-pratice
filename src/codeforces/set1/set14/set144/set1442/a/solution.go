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
		A := readNNums(reader, n)
		res := solve(A)
		if res {
			buf.WriteString("YES\n")
		} else {
			buf.WriteString("NO\n")
		}
	}
	fmt.Print(buf.String())
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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

func solve(A []int) bool {
	// 始终处在两头大，中间小的状态?
	// 也不一定的，比如  [1, 2, 1]
	// 假设在i处，操作1进行了x次
	// 操作2进行了y次，那么 x + y = A[i]
	// 同时 X[i-1] >= x
	// 如果 A[i-1] > A[i]， 那么在i-1处，至少要进行 A[i-1] - A[i]次操作1
	// 同理，A[i-1] < A[i], 那么在i处，至少要进行 A[i] - A[i-1]次操作2
	n := len(A)
	x := make([]int, n)
	for i := 0; i+1 < n; i++ {
		x[i] = max(0, A[i]-A[i+1])
	}

	for i := n - 2; i >= 0; i-- {
		x[i] += x[i+1]
	}

	y := make([]int, n)
	for i := n - 1; i > 0; i-- {
		y[i] = max(0, A[i]-A[i-1])
	}

	for i := 1; i < n; i++ {
		y[i] += y[i-1]
	}

	for i := 0; i < n; i++ {
		A[i] -= (x[i] + y[i])
	}
	if A[0] < 0 {
		return false
	}
	for i := 1; i < n; i++ {
		if A[i] != A[0] {
			return false
		}
	}
	return true
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
