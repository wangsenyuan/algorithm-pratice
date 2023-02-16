package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)

	A := readNNums(reader, n)

	res := solve(n, A)

	fmt.Printf("%d\n", res)
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

func solve(n int, a []int) int {
	A := make([]int, n+1)
	copy(A[1:], a)
	pos := make([]int, n+1)
	for i := 0; i <= n; i++ {
		pos[i] = n + 1
	}

	next := make([]int, n+1)

	for i := n; i >= 0; i-- {
		next[i] = pos[A[i]]
		pos[A[i]] = i
	}

	var x, y int
	var res int
	for i := 1; i <= n; i++ {
		if A[x] == A[i] {
			res += boolValue(A[i] != A[y])
			y = i
		} else if A[y] == A[i] {
			res += boolValue(A[x] != A[i])
			x = i
		} else if next[x] < next[y] {
			res += boolValue(A[x] != A[i])
			x = i
		} else {
			res += boolValue(A[y] != A[i])
			y = i
		}
	}
	return res
}

func boolValue(v bool) int {
	if v {
		return 1
	}
	return 0
}
