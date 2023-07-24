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

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
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

const X = 1_000_000_000

func solve(A []int) int {
	// 最后一个时不能变的
	n := len(A)
	// abs(A[i] - x) <= A[i+1]
	//  A[i] - x <= A[i+1]
	// x - A[i] <= A[i+1]
	// x <= A[i] + A[i+1]
	// x >= A[i] - A[i+1]
	// abs(A[i] - x) <= abs(A[i+1] - x)
	// 如果 A[i] <= A[i+1]
	//   x <= A[i] => A[i] - x <= A[i+1] - x 可以的
	//   x > A[i+1] => -A[i] <= -A[i+1] bad
	// x < A[i+1] > A[i] 呢？ x - A[i] <= A[i+1] - x => 2 * x <= A[i+1] + A[i]
	// 如果 A[i] >= A[i+1] 呢
	//   x >= A[i] => -A[i] <= -A[i+1] 可以的
	//   x < A[i+1] => A[i] < A[i+1] 不可以
	//   x <= A[i] >= A[i+1] => A[i] - x <= x - A[i+1] => 2 * x >= A[i+1] + A[i]
	lo, hi := 0, X
	for i := 0; i+1 < n; i++ {
		if A[i] == A[i+1] {
			// 如果相等，任何数都是ok的
			continue
		}
		if A[i] < A[i+1] {
			hi = min(hi, (A[i]+A[i+1])/2)
		} else {
			lo = max(lo, (A[i]+A[i+1]+1)/2)
		}
		if lo > hi {
			return -1
		}
	}
	return lo
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
