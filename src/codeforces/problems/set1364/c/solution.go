package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

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

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	A := readNNums(reader, n)
	B := solve(n, A)
	if len(B) == 0 {
		fmt.Println(-1)
		return
	}
	var buf bytes.Buffer
	for i := 0; i < n; i++ {
		buf.WriteString(strconv.Itoa(B[i]))
		buf.WriteByte(' ')
	}
	fmt.Println(buf.String())
}

const MAX = 1000000

func solve(n int, A []int) []int {
	ex := make([]bool, MAX+1)
	B := make([]int, n)
	for i := 0; i < n; i++ {
		B[i] = -1
	}
	for i := 1; i < n; i++ {
		if A[i] != A[i-1] {
			B[i] = A[i-1]
			ex[B[i]] = true
		}
	}

	ex[A[n-1]] = true
	var mex int
	for i := 0; i < n; i++ {
		for ex[mex] {
			mex++
		}
		if B[i] == -1 {
			B[i] = mex
			ex[mex] = true
		}
	}

	return B
}
