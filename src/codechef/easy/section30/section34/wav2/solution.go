package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, q := readTwoNums(reader)
	A := readNNums(reader, n)
	X := make([]int, q)
	for i := 0; i < q; i++ {
		X[i] = readNum(reader)
	}
	res := solve(A, X)

	var buf bytes.Buffer

	for i := 0; i < q; i++ {
		if res[i] < 0 {
			buf.WriteString("NEGATIVE\n")
		} else if res[i] > 0 {
			buf.WriteString("POSITIVE\n")
		} else {
			buf.WriteString("0\n")
		}
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
		if s[i] == '\n' || s[i] == '\r' {
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

func solve(A []int, X []int) []int {
	sort.Ints(A)

	n := len(A)

	ans := make([]int, len(X))

	for i, x := range X {
		pos := search(n, func(j int) bool {
			return A[j] >= x
		})
		if pos < n && A[pos] == x {
			ans[i] = 0
			continue
		}
		// A[pos] > x, so total pos numbers less than x
		cnt := n - pos
		if cnt&1 == 1 {
			// negative
			ans[i] = -1
		} else {
			// positive
			ans[i] = 1
		}
	}
	return ans
}

func search(n int, f func(int) bool) int {
	l, r := 0, n
	for l < r {
		mid := (l + r) / 2
		if f(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}
