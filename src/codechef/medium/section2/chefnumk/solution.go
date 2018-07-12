package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		N, Q := readTwoNums(scanner)

		A := readNNums(scanner, N)

		queries := make([][]int, Q)
		for i := 0; i < Q; i++ {
			queries[i] = readNNums(scanner, 3)
		}

		res := solve(N, Q, A, queries)

		for _, ans := range res {
			fmt.Printf("%d\n", ans)
		}
	}
}

func readInt(bytes []byte, from int, val *int) int {
	i := from
	tmp := 0
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp
	return i
}

func readNum(scanner *bufio.Scanner) (a int) {
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &a)
	return
}

func readTwoNums(scanner *bufio.Scanner) (a int, b int) {
	res := readNNums(scanner, 2)
	a, b = res[0], res[1]
	return
}

func readNNums(scanner *bufio.Scanner, n int) []int {
	res := make([]int, n)
	x := 0
	scanner.Scan()
	for i := 0; i < n; i++ {
		for x < len(scanner.Bytes()) && scanner.Bytes()[x] == ' ' {
			x++
		}
		x = readInt(scanner.Bytes(), x, &res[i])
	}
	return res
}

func solve(N, Q int, A []int, queries [][]int) []int {
	B := make([]int, N)
	copy(B, A)
	sort.Ints(B)
	indexMap := make(map[int]int)
	var idx int
	for i := 1; i <= N; i++ {
		if i < N && B[i] == B[i-1] {
			continue
		}
		indexMap[B[i-1]] = idx
		idx++
	}
	for i := 0; i < N; i++ {
		A[i] = indexMap[A[i]]
	}

	cnt := make([]int, idx)
	above := make([]int, N+1)

	qs := make(QS, Q)
	for i := 0; i < Q; i++ {
		qs[i] = Query{queries[i][0] - 1, queries[i][1] - 1, queries[i][2], i}
	}

	sort.Sort(qs)

	add := func(val int) {
		above[cnt[val]]++
		cnt[val]++
	}

	remove := func(val int) {
		cnt[val]--
		above[cnt[val]]--
	}

	res := make([]int, Q)

	currentL, currentR := 0, 0

	add(A[0])
	// [L..R] is the range in the window
	for _, q := range qs {
		for currentL < q.L {
			remove(A[currentL])
			currentL++
		}

		for currentL > q.L {
			currentL--
			add(A[currentL])
		}

		for currentR < q.R {
			currentR++
			add(A[currentR])
		}

		for currentR > q.R {
			remove(A[currentR])
			currentR--
		}

		res[q.idx] = above[0] - above[q.K]
	}

	return res
}

type Query struct {
	L, R, K int
	idx     int
}

type QS []Query

func (this QS) Len() int {
	return len(this)
}

func (this QS) Less(i, j int) bool {
	return this[i].L < this[j].L || this[i].L == this[j].L && this[i].R < this[j].R
}

func (this QS) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
