package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m := readTwoNums(reader)
	A := make([][]byte, n)
	for i := 0; i < n; i++ {
		A[i], _ = reader.ReadBytes('\n')
	}

	Q := readNum(reader)
	flips := make([][]int, Q)
	for i := 0; i < Q; i++ {
		flips[i] = readNNums(reader, 4)
	}
	B := solve(n, m, A, Q, flips)
	var buf bytes.Buffer
	for i := 0; i < n; i++ {
		buf.WriteString(string(B[i][:m]))
		buf.WriteByte('\n')
	}
	fmt.Print(buf.String())
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

const MAX_N = 1003

var bit [MAX_N][MAX_N]int

func reset(n, m int) {
	for i := 0; i <= n; i++ {
		for j := 0; j <= n; j++ {
			bit[i][j] = 0
		}
	}
}

func solve(n, m int, A [][]byte, Q int, flips [][]int) [][]byte {
	reset(n, m)

	for i := 0; i < Q; i++ {
		a, b, c, d := flips[i][0], flips[i][1], flips[i][2], flips[i][3]
		a--
		b--
		c--
		d--

		bit[a][b]++
		if c+1 < n && d+1 < m {
			bit[c+1][d+1]++
		}
		if c+1 < n {
			bit[c+1][b]--
		}
		if d+1 < n {
			bit[a][d+1]--
		}
	}

	for j := 0; j < m; j++ {
		for i := 1; i < n; i++ {
			bit[i][j] += bit[i-1][j]
		}
	}

	for i := 0; i < n; i++ {
		for j := 1; j < m; j++ {
			bit[i][j] += bit[i][j-1]
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			x := bit[i][j]
			if x&1 == 1 {
				A[i][j] = flip(A[i][j])
			}
		}
	}
	return A
}

func flip(x byte) byte {
	if x == '0' {
		return '1'
	}
	return '0'
}
