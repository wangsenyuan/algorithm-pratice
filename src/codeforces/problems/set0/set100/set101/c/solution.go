package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	A := readNNums(reader, 2)
	B := readNNums(reader, 2)
	C := readNNums(reader, 2)
	res := solve(A, B, C)
	if res {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
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
		if s[i] == '\n' || s[i] == '\r' {
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

func solve(A, B, C []int) bool {
	b1 := int64(B[0])
	b2 := int64(B[1])
	c1 := int64(C[0])
	c2 := int64(C[1])
	check := func(x, y int64) bool {
		if c1*c1+c2*c2 == 0 {
			return x == b1 && y == b2
		}
		return (c1*(b2-y)-c2*(b1-x))%(c1*c1+c2*c2) == 0 && (c1*(b1-x)+c2*(b2-y))%(c1*c1+c2*c2) == 0
	}

	if check(int64(A[0]), int64(A[1])) {
		return true
	}

	if check(-int64(A[1]), int64(A[0])) {
		return true
	}

	if check(int64(A[1]), -int64(A[0])) {
		return true
	}
	if check(-int64(A[0]), -int64(A[1])) {
		return true
	}
	return false
}
