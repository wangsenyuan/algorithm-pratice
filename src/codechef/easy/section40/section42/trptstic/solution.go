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
		n, m, k := readThreeNums(reader)
		A := make([][]int, n)
		for i := 0; i < n; i++ {
			A[i] = readNNums(reader, m)
		}
		res := solve(A, k)
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
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func solve(A [][]int, k int) int {
	// first of all, mentor needs to locate at the center

	n := len(A)
	m := len(A[0])

	sum := make([][]int64, n)
	for i := 0; i < n; i++ {
		sum[i] = make([]int64, m)
		for j := 0; j < m; j++ {
			sum[i][j] = int64(A[i][j])
			if i > 0 {
				sum[i][j] += sum[i-1][j]
			}
			if j > 0 {
				sum[i][j] += sum[i][j-1]
			}
			if i > 0 && j > 0 {
				sum[i][j] -= sum[i-1][j-1]
			}
		}
	}
	if sum[n-1][m-1] < int64(k+1) {
		return -1
	}

	checkDistance := func(u int, v int, d int) bool {
		i := max(0, u-d)
		j := max(0, v-d)
		x := min(n-1, u+d)
		y := min(m-1, v+d)
		tmp := sum[x][y]
		if i > 0 {
			tmp -= sum[i-1][y]
		}
		if j > 0 {
			tmp -= sum[x][j-1]
		}
		if i > 0 && j > 0 {
			tmp += sum[i-1][j-1]
		}
		return tmp > int64(k)
	}

	check := func(u int, v int) int {
		if A[u][v] == 0 {
			return -1
		}
		// u, v is the mentor room

		l, r := 0, max(m, n)

		if !checkDistance(u, v, r) {
			return -1
		}

		for l < r {
			mid := (l + r) / 2
			if checkDistance(u, v, mid) {
				r = mid
			} else {
				l = mid + 1
			}
		}

		return r
	}

	res := max(n, m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			tmp := check(i, j)
			if tmp >= 0 {
				res = min(res, tmp)
			}
		}
	}

	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
