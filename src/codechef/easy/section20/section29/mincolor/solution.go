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
		A := readNNums(reader, 2)
		B := readNNums(reader, 2)
		res := solve(n, m, A, B)
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				buf.WriteString(fmt.Sprintf("%d ", res[i][j]))
			}
			buf.WriteByte('\n')
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
		if s[i] == '\n' {
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

func solve(n int, m int, A []int, B []int) [][]int {
	a := (A[0] + A[1]) & 1
	b := (B[0] + B[1]) & 1

	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, m)
	}

	if a != b {
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				x := (i + 1 + j + 1) & 1
				if x == a {
					res[i][j] = 1
				} else {
					res[i][j] = 2
				}
			}
		}
	} else {
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				x := (i + 1 + j + 1) & 1
				if x == a {
					res[i][j] = 1
				} else {
					res[i][j] = 3
				}
				if i+1 == B[0] && j+1 == B[1] {
					res[i][j] = 2
				}
			}
		}
	}
	return res
}
