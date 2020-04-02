package main

import (
	"bufio"
	"fmt"
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

func fillNNums(bs []byte, n int, res []int) {
	x := 0
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	N, M := readTwoNums(reader)

	A := readNNums(reader, N)

	L, R := make([]int, M), make([]int, M)
	for i := 0; i < M; i++ {
		L[i], R[i] = readTwoNums(reader)
	}
	K := readNum(reader)

	X := make([]int, K)
	for i := 0; i < K; i++ {
		X[i] = readNum(reader)
	}
	Y := solve(N, M, A, L, R, K, X)

	for i := 0; i < K; i++ {
		fmt.Println(Y[i])
	}
}

const INF = 1 << 30

func solve(N int, M int, A []int, L []int, R []int, K int, X []int) []int {

	arr := make([]int, 4*M+3)

	counts := make([]int, N+1)

	V := make([]*Pair, 0, M)

	for i := 0; i < M; i++ {
		l, r := L[i], R[i]
		l--
		r--
		counts[l]++
		V = append(V, &Pair{l, r})
	}

	for i := 1; i <= N; i++ {
		counts[i] += counts[i-1]
	}

	sort.Sort(PairSlice(V))

	var build func(i int, left int, right int)

	build = func(i int, left int, right int) {
		if left == right {
			arr[i] = V[left].w
			return
		}
		mid := (left + right) / 2
		build(2*i+1, left, mid)
		build(2*i+2, mid+1, right)
		arr[i] = min(arr[2*i+1], arr[2*i+2])
	}

	build(0, 0, M-1)

	var query func(i int, left int, right int, low int, high int, R int) bool

	query = func(i int, left int, right int, low int, high int, R int) bool {
		if high < left || low > right {
			return false
		}

		if low <= left && right <= high && arr[i] > R {
			return false
		}
		if left == right {
			arr[i] = INF
			// V[left].w = INF
			return true
		}
		mid := (left + right) / 2

		if query(2*i+1, left, mid, low, high, R) {
			arr[i] = min(arr[2*i+1], arr[2*i+2])
			return true
		}

		if query(2*i+2, mid+1, right, low, high, R) {
			arr[i] = min(arr[2*i+1], arr[2*i+2])
			return true
		}
		return false
	}

	left := make([]int, N+1)
	right := make([]int, N+1)
	for i := 0; i < N; i++ {
		left[i] = i
		right[i] = i
	}

	var findLeft func(i int) int
	var findRight func(i int) int

	findLeft = func(i int) int {
		if left[i] != i {
			left[i] = findLeft(left[i])
		}
		return left[i]
	}

	findRight = func(i int) int {
		if right[i] != i {
			right[i] = findRight(right[i])
		}
		return right[i]
	}

	var ans int
	process := func(at int) {
		if at > 0 && A[at-1] == 0 {
			left[at] = left[at-1]
		}
		if at+1 < N && A[at+1] == 0 {
			left[at+1] = at
		}

		if at+1 < N && A[at+1] == 0 {
			right[at] = at + 1
		}
		if at > 0 && A[at-1] == 0 {
			right[at-1] = at
		}

		l := findLeft(at)
		r := findRight(at)

		var low int
		if l > 0 {
			low = counts[l-1]
		}
		var high int = counts[r] - 1

		for query(0, 0, M-1, low, high, r) {
			ans++
		}
	}

	for i := 0; i < N; i++ {
		if A[i] == 0 {
			process(i)
		}
	}

	Y := make([]int, K)

	for i := 0; i < K; i++ {
		x := X[i] - 1
		x += ans
		A[x]--

		if A[x] > 0 {
			Y[i] = ans
			continue
		}
		process(x)
		Y[i] = ans
	}
	return Y
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

type Pair struct {
	v int
	w int
}

type PairSlice []*Pair

func (this PairSlice) Len() int {
	return len(this)
}

func (this PairSlice) Less(i, j int) bool {
	return this[i].v < this[j].v
}

func (this PairSlice) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
