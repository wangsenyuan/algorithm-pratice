package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
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

	tc := readNum(reader)

	for tc > 0 {
		tc--
		N, K := readTwoNums(reader)
		A := readNNums(reader, N)
		res := solve(N, A, K)
		if res {
			fmt.Println("yes")
		} else {
			fmt.Println("no")
		}
	}
}

func solve(N int, A []int, K int) bool {

	if N == K {
		return sort.IntsAreSorted(A)
	}

	if K == 1 {
		return true
	}
	r := N % K

	if r != 0 {
		M := ((N + K - 1) / K) * K
		B := make([]int, M)
		copy(B, A)
		for i := N; i < M; i++ {
			B[i] = math.MaxInt32
		}
		A = B
		N = M
	}

	for i := 0; i < K; i++ {
		ss := &SortSlice{A, K, i}
		sort.Sort(ss)
	}

	return sort.IntsAreSorted(A)
}

type SortSlice struct {
	arr []int
	K   int
	i   int
}

func (this *SortSlice) Len() int {
	n := len(this.arr)
	m := (n + this.K - 1) / this.K
	return m
}

func (this *SortSlice) Less(i, j int) bool {
	K := this.K
	x := this.i
	return this.arr[i*K+x] < this.arr[j*K+x]
}

func (this *SortSlice) Swap(i, j int) {
	K := this.K
	x := this.i
	this.arr[i*K+x], this.arr[j*K+x] = this.arr[j*K+x], this.arr[i*K+x]
}
