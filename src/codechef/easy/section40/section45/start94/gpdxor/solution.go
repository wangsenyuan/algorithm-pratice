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
		n, m := readTwoNums(reader)
		A := make([][]int, n)
		for i := 0; i < n; i++ {
			A[i] = readNNums(reader, m)
		}
		res := solve(A)
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

const H = 30

func solve(A [][]int) int64 {
	n := len(A)
	m := len(A[0])
	// 是 sum，
	// sum(a ^ b)
	sum := make([]int, H)
	row := make([][]int, n)
	col := make([][]int, m)
	for i := 0; i < n; i++ {
		row[i] = make([]int, H)
	}
	for j := 0; j < m; j++ {
		col[j] = make([]int, H)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			x := A[i][j]
			for d := 0; d < H; d++ {
				sum[d] += (x >> d) & 1
				row[i][d] += (x >> d) & 1
				col[j][d] += (x >> d) & 1
			}
		}
	}

	var best int64

	N := int64(n)
	M := int64(m)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			var tmp int64
			for d := 0; d < H; d++ {
				cnt1 := sum[d] - row[i][d] - col[j][d]
				if (A[i][j]>>d)&1 == 1 {
					cnt1++
					tmp += (N*M - N - M + 1 - int64(cnt1)) * (1 << d)
				} else {
					tmp += int64(cnt1) * (1 << d)
				}
			}
			best = max(best, tmp)
		}
	}
	return best
}

func max(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}
