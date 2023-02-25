package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n := readNum(reader)
	C := readNNums(reader, n)
	A := make([][]int, n)
	for i := 0; i < n; i++ {
		A[i] = readNNums(reader, C[i])
	}
	res := solve(A)
	var buf bytes.Buffer

	buf.WriteString(fmt.Sprintf("%d\n", len(res)))
	for _, op := range res {
		for j := 0; j < len(op); j++ {
			buf.WriteString(fmt.Sprintf("%d ", op[j]))
		}
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

func solve(A [][]int) [][]int {
	// s = sum(len(A[i]))
	// 1， 2， 3， 4
	// 5， 6， 7
	// 。。
	// s
	var s int
	for _, a := range A {
		s += len(a)
	}
	pos := make([][]int, s+1)
	for i, row := range A {
		for j := 0; j < len(row); j++ {
			pos[row[j]] = []int{i, j}
		}
	}
	var res [][]int
	r, c := 0, 0
	for i := 1; i <= s; i++ {
		if A[r][c] != i {
			x, y := pos[i][0], pos[i][1]
			res = append(res, []int{r + 1, c + 1, x + 1, y + 1})
			A[x][y], A[r][c] = A[r][c], A[x][y]
			pos[A[x][y]] = []int{x, y}
		}
		c++
		if c == len(A[r]) {
			r++
			c = 0
		}
	}

	return res
}
