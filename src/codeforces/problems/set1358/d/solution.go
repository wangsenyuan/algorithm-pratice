package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, k := readTwoNums(reader)
	A := readNNums(reader, n)

	fmt.Println(solve(n, k, A))
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

func solve(n int, k int, a []int) int64 {
	A := make([]int64, 2*n)
	for i := 0; i < n; i++ {
		A[i] = int64(a[i])
		A[n+i] = int64(a[i])
	}
	n *= 2

	B := make([]int64, n+1)
	for i := 0; i < n; i++ {
		B[i+1] = B[i] + int64(A[i])
	}

	C := make([]int64, n+1)
	for i := 0; i < n; i++ {
		C[i+1] = C[i] + A[i]*(A[i]+1)/2
	}

	var ans int64
	K := int64(k)
	for i := 0; i < n; i++ {
		if B[i+1] >= K {
			z := sort.Search(len(B), func(j int) bool {
				return B[j] > B[i+1]-K
			})

			cnt := C[i+1] - C[z]
			days := B[i+1] - B[z]
			too := K - days

			cnt += A[z-1] * (A[z-1] + 1) / 2
			cnt -= (A[z-1] - too) * (A[z-1] - too + 1) / 2
			ans = max(ans, cnt)
		}
	}

	return ans
}

func max(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}
