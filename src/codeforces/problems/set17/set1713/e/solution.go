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

		A := make([][]int, n)

		for i := 0; i < n; i++ {
			A[i] = readNNums(reader, n)
		}

		solve(A)

		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				buf.WriteString(fmt.Sprintf("%d ", A[i][j]))
			}
			buf.WriteByte('\n')
		}
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

func solve(A [][]int) {

	n := len(A)

	par := make([]int, n+1)

	for i := 1; i <= n; i++ {
		par[i] = i
	}

	var find func(u int) int

	find = func(u int) int {
		if u < 0 {
			return -find(-u)
		}

		if u == par[u] {
			return u
		}
		par[u] = find(par[u])
		return par[u]
	}

	join := func(u int, v int) {
		u = find(u)
		v = find(v)
		if abs(u) != abs(v) {
			if u > 0 {
				par[u] = v
			} else {
				par[-u] = -v
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if A[i][j] < A[j][i] {
				join(i+1, j+1)
			}
			if A[i][j] > A[j][i] {
				join(i+1, -(j + 1))
			}
		}
	}

	for i := 1; i <= n; i++ {
		if find(i) < 0 {
			continue
		}
		for j := 1; j <= n; j++ {
			A[i-1][j-1], A[j-1][i-1] = A[j-1][i-1], A[i-1][j-1]
		}
	}
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
