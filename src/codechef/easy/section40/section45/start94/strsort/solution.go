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
		cost, res := solve(A)
		if cost < 0 {
			buf.WriteString("NO\n")
		} else {
			buf.WriteString("YES\n")
			buf.WriteString(fmt.Sprintf("%d %d\n", len(res), cost))
			for i := 0; i < len(res); i++ {
				for j := 0; j < len(res[i]); j++ {
					buf.WriteString(fmt.Sprintf("%d ", res[i][j]))
				}
				buf.WriteByte('\n')
			}
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

func solve(A []int) (int, [][]int) {
	n := len(A)
	if A[n-2] < A[n-1] {
		return 0, solve2(A, nil)
	}
	// A[n-1] > A[n]
	// check whether one swap can solve the problem
	for i := 0; i < n-2; i++ {
		if A[i] < A[n-1] && n-2-i > 1 {
			res := [][]int{{1, i + 1, n - 1}}
			A[i], A[n-2] = A[n-2], A[i]
			res = solve2(A, res)
			return 1, res
		} else if A[i] > A[n-2] {
			res := [][]int{{1, i + 1, n}}
			A[i], A[n-1] = A[n-1], A[i]
			res = solve2(A, res)
			return 1, res
		}
	}
	// whether cost of 2 work?
	if n >= 5 {
		var res [][]int
		if A[0] < A[1] {
			res = append(res, []int{1, 1, n - 1})
			res = append(res, []int{1, 2, n})
			A[0], A[n-2] = A[n-2], A[0]
			A[1], A[n-1] = A[n-1], A[1]
		} else {
			res = append(res, []int{1, 1, n})
			res = append(res, []int{1, 2, n - 1})
			A[0], A[n-1] = A[n-1], A[0]
			A[1], A[n-2] = A[n-2], A[1]
		}
		res = solve2(A, res)
		return 2, res
	}

	if n == 3 {
		return -1, nil
	}

	// n == 4
	res := [][]int{{1, 1, 4}, {1, 1, 3}}
	A[0], A[3] = A[3], A[0]
	A[0], A[2] = A[2], A[0]
	return 2, solve2(A, res)
}

func solve2(A []int, res [][]int) [][]int {
	var rem int
	n := len(A)
	pos, truePos := n-2, n-2
	for i := n - 3; i >= 0; i-- {
		if A[i] < A[truePos] {
			pos = i + rem
			truePos = i
		} else {
			res = append(res, []int{2, i + 1, pos - rem + 1, n - rem})
			rem++
		}
	}
	return res
}
