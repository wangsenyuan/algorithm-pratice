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
		m, n, x := readThreeNums(reader)
		A := make([][]int, m)
		for i := 0; i < m; i++ {
			A[i] = readNNums(reader, n)
		}
		B := make([][]int, m)
		for i := 0; i < m; i++ {
			B[i] = readNNums(reader, n)
		}
		res := solve(x, A, B)

		if res {
			buf.WriteString("Yes\n")
		} else {
			buf.WriteString("No\n")
		}
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

func solve(X int, A, B [][]int) bool {
	C := make([][]int, X)
	for i := 0; i < X; i++ {
		C[i] = make([]int, X)
	}
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A[i]); j++ {
			C[i%X][j%X] += A[i][j]
			C[i%X][j%X] -= B[i][j]
		}
	}

	for i := 1; i < X; i++ {
		delta := C[i][0] - C[0][0]
		for j := 1; j < X; j++ {
			if C[i][j]-C[0][j] != delta {
				return false
			}
		}
	}
	return true
}
