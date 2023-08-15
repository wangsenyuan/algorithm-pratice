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

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n, m := readTwoNums(reader)
		mat := make([][]int, n)
		for i := 0; i < n; i++ {
			mat[i] = readNNums(reader, m)
		}
		res, k := solve(mat)

		if len(res) == 0 {
			buf.WriteString("NO\n")
		} else {
			buf.WriteString("YES\n")
			buf.WriteString(fmt.Sprintf("%s %d\n", res, k))
			//buf.WriteByte('\n')
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

func solve(mat [][]int) (string, int) {
	n := len(mat)
	m := len(mat[0])
	ids := make([]int, n)
	for i := 0; i < n; i++ {
		ids[i] = i
	}

	sort.Slice(ids, func(i, j int) bool {
		return mat[ids[i]][0] < mat[ids[j]][0]
	})
	// A[i][j] is the min value in sub mat[i][j] from top-left
	// B[i][j] is the min value in sub mat[i][j] from top-right
	A := make([][]int, n)
	B := make([][]int, n)
	C := make([][]int, n)
	D := make([][]int, n)
	for i := 0; i < n; i++ {
		A[i] = make([]int, m)
		B[i] = make([]int, m)
		C[i] = make([]int, m)
		D[i] = make([]int, m)
	}
	for i, r := range ids {
		for j := 0; j < m; j++ {
			A[i][j] = mat[r][j]
			if i > 0 {
				A[i][j] = max(A[i][j], A[i-1][j])
			}
			if j > 0 {
				A[i][j] = max(A[i][j], A[i][j-1])
			}
		}
		for j := m - 1; j >= 0; j-- {
			B[i][j] = mat[r][j]
			if i > 0 {
				B[i][j] = min(B[i][j], B[i-1][j])
			}
			if j+1 < m {
				B[i][j] = min(B[i][j], B[i][j+1])
			}
		}
	}
	for i := n - 1; i >= 0; i-- {
		r := ids[i]
		for j := 0; j < m; j++ {
			C[i][j] = mat[r][j]
			if i+1 < n {
				C[i][j] = min(C[i][j], C[i+1][j])
			}
			if j > 0 {
				C[i][j] = min(C[i][j], C[i][j-1])
			}
		}

		for j := m - 1; j >= 0; j-- {
			D[i][j] = mat[r][j]
			if i+1 < n {
				D[i][j] = max(D[i][j], D[i+1][j])
			}
			if j+1 < m {
				D[i][j] = max(D[i][j], D[i][j+1])
			}
		}
	}

	color := func(x int) string {
		res := make([]byte, n)
		for i := 0; i < n; i++ {
			res[i] = 'R'
		}
		for i := 0; i <= x; i++ {
			res[ids[i]] = 'B'
		}

		return string(res)
	}

	for i := 0; i+1 < n; i++ {
		for j := 0; j+1 < m; j++ {
			a := A[i][j]
			b := B[i][j+1]
			c := C[i+1][j]
			d := D[i+1][j+1]
			if a < c && b > d {
				return color(i), j + 1
			}
		}
	}

	return "", 0
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
