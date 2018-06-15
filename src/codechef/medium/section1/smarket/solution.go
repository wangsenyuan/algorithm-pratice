package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

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
	scanner.Scan()
	x := readInt(scanner.Bytes(), 0, &a)
	readInt(scanner.Bytes(), x+1, &b)
	return
}

func readNNums(scanner *bufio.Scanner, n int) []int {
	res := make([]int, n)
	x := -1
	scanner.Scan()
	for i := 0; i < n; i++ {
		x = readInt(scanner.Bytes(), x+1, &res[i])
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	t := readNum(scanner)

	for t > 0 {
		n, Q := readTwoNums(scanner)
		A := readNNums(scanner, n)
		queries := make([][]int, Q)
		for i := 0; i < Q; i++ {
			queries[i] = readNNums(scanner, 3)
		}
		res := solve(n, A, Q, queries)

		for _, ans := range res {
			fmt.Printf("%d\n", ans)
		}
		t--
	}
}

func solve(n int, A []int, m int, queries [][]int) []int {
	qs := make(QS, m)
	for i := 0; i < m; i++ {
		qs[i] = new(Q)
		qs[i].i = i
		qs[i].L, qs[i].R, qs[i].K = queries[i][0]-1, queries[i][1]-1, queries[i][2]
		qs[i].SL = fastSqrt(qs[i].L)
	}
	sort.Sort(qs)

	B := make([]int, n)

	for i := 0; i < n; i++ {
		if i == 0 || A[i] != A[i-1] {
			B[i] = 1
		} else {
			B[i] = B[i-1] + 1
		}
	}

	C := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		if i == n-1 || A[i] != A[i+1] {
			C[i] = 1
		} else {
			C[i] = C[i+1] + 1
		}
	}

	res := make([]int, m)

	currentL, currentR := 0, 0

	cnt := make([]int, n+1)

	for _, q := range qs {
		L, R, K := q.L, q.R, q.K
		for currentL < L {
			cnt[C[currentL]]--
			currentL++
		}

		for currentL > L {
			cnt[C[currentL-1]]++
			currentL--
		}

		for currentR <= R {
			cnt[B[currentR]]++
			currentR++
		}

		for currentR > R+1 {
			cnt[B[currentR-1]]--
			currentR--
		}
		res[q.i] = cnt[K]
	}

	return res
}

func fastSqrt(x int) int {
	j := sort.Search(1<<10, func(i int) bool {
		return i*i > x
	})
	return j - 1
}

type Q struct {
	L, R, K int
	SL      int
	i       int
}
type QS []*Q

func (this QS) Len() int {
	return len(this)
}

func (this QS) Less(i, j int) bool {
	return this[i].SL < this[j].SL || (this[i].SL == this[j].SL && this[i].R < this[j].R)
}

func (this QS) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func solve1(n int, A []int, Q int, queries [][]int) []int {
	B := make([]int, n)
	B[0] = 1
	for i := 1; i < n; i++ {
		if A[i] == A[i-1] {
			B[i] = B[i-1] + 1
		} else {
			B[i] = 1
		}
	}

	II := make([][]int, n+1)

	for i := 0; i < n; i++ {
		b := B[i]

		if II[b] == nil {
			II[b] = make([]int, 0, 10)
		}
		II[b] = append(II[b], i)
	}

	res := make([]int, Q)

	insideBlock := func(L, R, K int) int {
		I := II[K]
		if len(I) == 0 {
			return 0
		}

		j := sort.Search(len(I), func(j int) bool {
			return I[j] >= L
		})

		if j == len(I) || I[j] > R {
			return 0
		}

		if I[j] < L+K-1 {
			j++
			if j == len(I) || I[j] > R {
				return 0
			}
		}
		k := sort.Search(len(I), func(k int) bool {
			return I[k] > R
		})
		return k - j
	}

	leftBoundary := func(L, R, K int) int {
		if L > 0 && A[L] == A[L-1] {
			if L+K-1 > R {
				return 0
			}

			if B[L+K-1] == K+B[L-1] {
				return 1
			}

		}
		return 0
	}

	for i := 0; i < Q; i++ {
		query := queries[i]
		L, R, K := query[0]-1, query[1]-1, query[2]

		res[i] = insideBlock(L, R, K) + leftBoundary(L, R, K)
	}

	return res
}
