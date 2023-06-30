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
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func solve(A []int) bool {
	n := len(A)

	if n%2 == 0 {
		return solveEven(A)
	}
	return solveOdd(A)
}

func solveOdd(A []int) bool {
	n := len(A)
	var sum int64 = int64(A[0])
	for i := 1; i < n; i++ {
		if i&1 == 1 {
			sum += int64(A[i])
		} else {
			sum -= int64(A[i])
		}
	}

	if sum < 0 || sum%2 == 1 {
		return false
	}
	if sum/2 > int64(A[0]) {
		return false
	}
	cur := int(sum / 2)

	for i := 1; i < n; i++ {
		cur = A[i] - cur
		if cur < 0 {
			return false
		}
	}

	return true
}

func solveEven(A []int) bool {
	n := len(A)

	cnt := make([]int64, 2)
	l, r := 0, A[0]
	cnt[0] = int64(A[0])
	var sum int
	for i := 1; i < n; i++ {
		cnt[i&1] += int64(A[i])
		if i&1 == 1 {
			sum += A[i]
			// odd position
			r = min(r, sum)
		} else {
			sum -= A[i]
			l = max(l, sum)
		}
	}

	return cnt[0] == cnt[1] && r >= 0 && l <= r
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	return a + b - min(a, b)
}
