package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	tc := readNum(reader)
	var buf bytes.Buffer
	for tc > 0 {
		tc--
		s, _ := reader.ReadBytes('\n')
		var n int
		pos := readInt(s, 0, &n) + 1
		var k uint64
		readUint64(s, pos, &k)
		A := make([][]int, n)
		for i := 0; i < n; i++ {
			A[i] = readNNums(reader, n)
		}
		res := solve(n, A, int64(k))
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

func solve(n int, A [][]int, k int64) int {
	N := n * n
	V := make([]int, N)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			V[i*n+j] = A[i][j]
		}
	}
	sort.Ints(V)

	var minCost int64

	for i, j := n/2, 0; j < n; i, j = i+n/2+1, j+1 {
		minCost += int64(V[i])
	}

	if minCost > k {
		return -1
	}
	lo := n / 2
	hi := N - 1

	second := n - n/2 - 1

	for lo < hi {
		// hi can reach
		md := (lo + hi + 1) / 2
		var cost int64
		less := md
		rpos := N
		var taken int
		for i := md; taken < n && i < rpos; i++ {
			if less >= n/2 {
				cost += int64(V[i])
				less -= n / 2
				rpos -= second
				taken++
				if rpos <= i {
					cost = k + 1
					break
				}
			} else {
				less++
			}
		}

		if taken < n {
			cost = k + 1
		}
		if cost <= k {
			lo = md
		} else {
			hi = md - 1
		}
	}

	return V[lo]
}
