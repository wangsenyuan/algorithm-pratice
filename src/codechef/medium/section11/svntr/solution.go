package main

import (
	"bufio"
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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
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
		n, m, s := readThreeNums(reader)
		A := make([][]int, n)
		for i := 0; i < n; i++ {
			A[i] = readNNums(reader, m)
		}
		fmt.Println(solve(n, m, s, A))
	}
}

func solve(N, M, s int, A [][]int) int64 {
	S := int64(s + 1)

	F := make([][]int64, N)
	for i := 0; i < N; i++ {
		F[i] = make([]int64, M)
	}

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			F[i][j] = int64(A[i][j])
			if i > 0 {
				F[i][j] += F[i-1][j]
			}
			if j > 0 {
				F[i][j] += F[i][j-1]
			}
			if i > 0 && j > 0 {
				F[i][j] -= F[i-1][j-1]
			}
		}
	}

	findSum := func(i, ii, j int) int64 {
		res := F[ii][j]
		if i > 0 {
			res -= F[i-1][j]
		}
		return res
	}

	B := make([]int64, M+1)
	C := make([]int64, M+1)
	var res int64

	var mergeSort func(l, r int)

	mergeSort = func(l, r int) {
		if l >= r {
			return
		}
		mid := (l + r) / 2
		mergeSort(l, mid)
		mergeSort(mid+1, r)
		var h = l

		for i := mid + 1; i <= r; i++ {
			for h <= mid && B[i]-B[h] >= S {
				h++
			}
			res += int64(h - l)
		}
		var p1, p2 = l, mid + 1

		for i := l; i <= r; i++ {
			if p2 > r || (p1 <= mid && B[p1] < B[p2]) {
				C[i] = B[p1]
				p1++
			} else {
				C[i] = B[p2]
				p2++
			}
		}

		for i := l; i <= r; i++ {
			B[i] = C[i]
		}
	}

	for i := 0; i < N; i++ {
		for ii := i; ii < N; ii++ {
			B[0] = 0
			for j := 0; j < M; j++ {
				B[j+1] = findSum(i, ii, j)
			}

			mergeSort(0, M)
		}
	}

	return int64(N)*int64(N+1)/2*int64(M)*int64(M+1)/2 - res
}
