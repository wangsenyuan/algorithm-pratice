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
		n := readNum(reader)
		D := make([]int64, 2*n)
		var pos int
		var tmp uint64
		s, _ := reader.ReadBytes('\n')
		for i := 0; i < 2*n; i++ {
			pos = readUint64(s, pos, &tmp) + 1
			D[i] = int64(tmp)
		}

		res := solve(n, D)

		if res {
			buf.WriteString("YES\n")
		} else {
			buf.WriteString("NO\n")
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

func solve(n int, A []int64) bool {
	// don't care the ordering,
	sort.Slice(A, func(i, j int) bool {
		return A[i] > A[j]
	})

	B := make([]int64, n)

	for i := 0; i < n; i++ {
		if A[2*i] != A[2*i+1] {
			return false
		}
		B[i] = A[2*i]
	}

	D := make([]int64, n)
	for i := 1; i < n; i++ {
		if B[i-1] == B[i] || (B[i-1]-B[i])%(2*int64(n-i)) > 0 {
			return false
		}
		D[i] = (B[i-1] - B[i]) / 2 / int64(n-i)
	}

	for i := 1; i < n; i++ {
		B[n-1] -= 2 * int64(i) * D[i]
	}

	if B[n-1] <= 0 || B[n-1]%(2*int64(n)) > 0 {
		return false
	}
	return true
}
