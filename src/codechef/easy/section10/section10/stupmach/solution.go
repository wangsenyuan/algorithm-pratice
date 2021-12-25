package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
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
	scanner := bufio.NewReader(os.Stdin)

	tc := readNum(scanner)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n := readNum(scanner)
		A := readNNums(scanner, n)
		buf.WriteString(fmt.Sprintf("%d\n", solve(n, A)))
	}
	fmt.Print(buf.String())
}

func solve(n int, A []int) int64 {
	x := make([]int, n)
	x[0] = A[0]

	for i := 1; i < n; i++ {
		x[i] = x[i-1]
		if A[i] < x[i] {
			x[i] = A[i]
		}
	}

	var res int64
	var prev int
	for i := n - 1; i >= 0; i-- {
		res += int64(i+1) * int64(x[i]-prev)
		prev = x[i]
	}

	return res
}
