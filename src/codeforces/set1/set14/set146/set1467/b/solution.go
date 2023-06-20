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
		X := readNNums(reader, n)
		res := solve(X)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

func solve(A []int) int {
	n := len(A)

	if n < 3 {
		return 0
	}

	var cnt int

	isHill := func(i, j, k int) bool {
		return i >= 0 && k < n && A[i] < A[j] && A[j] > A[k]
	}

	isVally := func(i, j, k int) bool {
		return i >= 0 && k < n && A[i] > A[j] && A[j] < A[k]
	}

	for i := 1; i+1 < n; i++ {
		if isVally(i-1, i, i+1) || isHill(i-1, i, i+1) {
			cnt++
		}
	}

	getNeighbors := func(i int) []int {
		var res []int
		if i > 0 {
			res = append(res, A[i-1])
		}
		if i+1 < n {
			res = append(res, A[i+1])
		}
		return res
	}

	var best int

	for i := 0; i < n; i++ {
		var old int
		if isVally(i-1, i, i+1) || isHill(i-1, i, i+1) {
			// change i to between A[i-1] and A[i+1]
			old++
		}

		if isHill(i-2, i-1, i) || isVally(i-2, i-1, i) {
			old++
		}

		if isHill(i, i+1, i+2) || isVally(i, i+1, i+2) {
			old++
		}

		arr := getNeighbors(i)

		for _, x := range arr {
			y := A[i]
			A[i] = x
			var cur int
			if isVally(i-1, i, i+1) || isHill(i-1, i, i+1) {
				cur++
			}

			if isHill(i-2, i-1, i) || isVally(i-2, i-1, i) {
				cur++
			}

			if isHill(i, i+1, i+2) || isVally(i, i+1, i+2) {
				cur++
			}

			if old-cur > best {
				best = old - cur
			}

			A[i] = y
		}

	}

	return cnt - best
}
