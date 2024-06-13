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
		B := make([][]int, n)
		for i := 0; i < n; i++ {
			B[i] = readNNums(reader, m)
		}
		res := solve(n, m, A, B)

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

func solve(n int, m int, A [][]int, B [][]int) bool {
	pos := make([]int, n*m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			pos[A[i][j]-1] = i*m + j
		}
	}
	row := make([]int, n)
	col := make([]int, m)
	for i := 0; i < n; i++ {
		row[i] = -1
	}
	for j := 0; j < m; j++ {
		col[j] = -1
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			x := B[i][j] - 1
			u, v := pos[x]/m, pos[x]%m
			if row[u] != -1 && row[u] != i {
				return false
			}
			row[u] = i
			if col[v] != -1 && col[v] != j {
				return false
			}
			col[v] = j
		}
	}

	return true
}
