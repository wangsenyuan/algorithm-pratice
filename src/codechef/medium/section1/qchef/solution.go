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

func readNNums2(bs []byte, n int) []int {
	res := make([]int, n)
	x := -1
	for i := 0; i < n; i++ {
		x = readInt(bs, x+1, &res[i])
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	first := readNNums(scanner, 3)
	N, M, K := first[0], first[1], first[2]
	A := readNNums(scanner, N)
	L, R := make([]int, K), make([]int, K)
	for i := 0; i < K; i++ {
		L[i], R[i] = readTwoNums(scanner)
	}
	res := solve(N, M, K, A, L, R)

	for _, ans := range res {
		fmt.Printf("%d\n", ans)
	}
}

func solve(N, M, K int, A []int, L []int, R []int) []int {
	bLen := 100
	bNum := (N + bLen - 1) / bLen
	mem := make([][]int, bNum+1)
	for i := 0; i <= bNum; i++ {
		mem[i] = make([]int, bNum+1)
	}

	flag := make([]int, M+1)
	index := make([]int, M+1)
	var curFlag int

	for i := 0; i < N; i += bLen {
		curFlag++
		var tmp int
		for j := i; j < N; j++ {
			if j%bLen == 0 {
				mem[i/bLen][j/bLen] = tmp
			}
			if flag[A[j]] != curFlag {
				flag[A[j]] = curFlag
				index[A[j]] = j
			} else {
				tmp = max(tmp, j-index[A[j]])
			}
		}
		mem[i/bLen][bNum] = tmp
	}

	tree := make([][]int, M+1)
	for i := 0; i < N; i++ {
		if tree[A[i]] == nil {
			tree[A[i]] = make([]int, 0, 10)
		}
		tree[A[i]] = append(tree[A[i]], i)
	}

	res := make([]int, K)
	for k := 0; k < K; k++ {
		s, t := L[k]-1, R[k]
		ss := (s + bLen - 1) / bLen * bLen
		tt := t / bLen * bLen
		ans := mem[ss/bLen][tt/bLen]

		for i := s; i < ss && i < N; i++ {
			xx := tree[A[i]]
			// can't be nil

			j := sort.Search(len(xx), func(j int) bool {
				// t is exluded
				return xx[j] >= t
			})
			j--
			if j < 0 {
				// no A[i] before position t
				continue
			}

			ans = max(ans, xx[j]-i)
		}

		for i := tt; i < t && i < N; i++ {
			xx := tree[A[i]]
			j := sort.Search(len(xx), func(j int) bool {
				return xx[j] >= s
			})
			if j == len(xx) {
				// no A[i] after s
				continue
			}
			ans = max(ans, i-xx[j])
		}
		res[k] = ans
	}

	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
