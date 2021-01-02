package main

import (
	"bufio"
	"fmt"
	"math"
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

	for tc > 0 {
		tc--
		n := readNum(scanner)
		A := readNNums(scanner, n)

		fmt.Println(solve(n, A))
	}
}
func solve(n int, A []int) int {
	m := A[n-1]
	pivot := n - 1

	for i := n - 2; i >= 0; i-- {
		if A[i] <= m {
			m = A[i]
			pivot = i
		}
	}

	return pivot + 1
}
func solve1(n int, A []int) int {
	pref := make([]int64, n)

	for i := 0; i < n; i++ {
		pref[i] = int64(A[i])
		if i > 0 {
			pref[i] += pref[i-1]
		}
	}

	pivot := n - 1
	var best int64 = math.MaxInt64

	var suf int64

	for i := n - 1; i >= 0; i-- {
		suf += int64(A[i])
		if pref[i]+suf <= best {
			pivot = i
			best = pref[i] + suf
		}
	}

	return pivot + 1
}
