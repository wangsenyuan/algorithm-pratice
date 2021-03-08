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

	x, y := make([]int, n), make([]int, n)

	for i := 0; i < n; i++ {
		x[i], y[i] = readTwoNums(reader)
	}

	res := solve(n, x, y)

	var buf bytes.Buffer

	for i := 0; i < n; i++ {
		buf.WriteString(fmt.Sprintf("%d ", res[i]))
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

func solve(n int, x, y []int) []int {
	perm := make([]int, n)

	for i := 0; i < n; i++ {
		perm[i] = i
		for j := i; j > 1; j-- {
			if int64(x[perm[j]]-x[perm[j-1]])*int64(x[perm[j-1]]-x[perm[j-2]])+int64(y[perm[j]]-y[perm[j-1]])*int64(y[perm[j-1]]-y[perm[j-2]]) >= 0 {
				perm[j-1], perm[j] = perm[j], perm[j-1]
			} else {
				break
			}
		}
	}

	for i := 0; i < n; i++ {
		perm[i]++
	}

	return perm
}
