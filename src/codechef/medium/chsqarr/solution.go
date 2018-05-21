package main

import (
	"bufio"
	"fmt"
	"os"
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

	N, M := readTwoNums(scanner)
	MTR := make([][]int, N)

	for i := 0; i < N; i++ {
		MTR[i] = readNNums(scanner, M)
	}
	Q := readNum(scanner)

	queries := make([][]int, Q)
	for i := 0; i < Q; i++ {
		queries[i] = readNNums(scanner, 2)
	}
	res := solve(N, M, MTR, Q, queries)
	for _, ans := range res {
		fmt.Println(ans)
	}
}

func solve(N, M int, MTR [][]int, Q int, queries [][]int) []int {
	sum := make([][]int, N+1)
	for i := 0; i <= N; i++ {
		sum[i] = make([]int, M+1)
	}

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			sum[i+1][j+1] = sum[i+1][j] + sum[i][j+1] - sum[i][j] + MTR[i][j]
		}
	}

	res := make([]int, Q)

	mx := make([][]int, N)
	for i := 0; i < N; i++ {
		mx[i] = make([]int, M)
	}

	front, rear := make([]int, M), make([]int, M)

	dq := make([][]int, M)
	for i := 0; i < M; i++ {
		dq[i] = make([]int, M)
	}

	ddq := make([]int, M)

	pp := func(a, b int) {
		for j := 0; j < M; j++ {
			front[j] = 0
			rear[j] = 0
		}

		for i := 0; i < a; i++ {
			for j := 0; j < M; j++ {
				for front[j] < rear[j] && dq[j][rear[j]-1] < MTR[i][j] {
					rear[j]--
				}
				dq[j][rear[j]] = MTR[i][j]
				rear[j]++
			}
		}

		for i := a - 1; i < N; i++ {
			f, r := 0, 0
			for j := 0; j < b; j++ {
				for f < r && ddq[r-1] < dq[j][front[j]] {
					r--
				}
				ddq[r] = dq[j][front[j]]
				r++
			}
			for j := b - 1; j < M; j++ {
				mx[i][j] = ddq[f]
				if ddq[f] == dq[j-b+1][front[j-b+1]] {
					f++
				}
				if j+1 < M {
					for f < r && ddq[r-1] < dq[j+1][front[j+1]] {
						r--
					}
					ddq[r] = dq[j+1][front[j+1]]
					r++
				}
			}
			for j := 0; j < M; j++ {
				if dq[j][front[j]] == MTR[i-a+1][j] {
					front[j]++
				}
			}
			if i+1 < N {
				for j := 0; j < M; j++ {
					for front[j] < rear[j] && dq[j][rear[j]-1] < MTR[i+1][j] {
						rear[j]--
					}
					dq[j][rear[j]] = MTR[i+1][j]
					rear[j]++
				}
			}
		}

	}

	for q := 0; q < Q; q++ {
		a, b := queries[q][0], queries[q][1]

		pp(a, b)
		best := 1 << 30
		for i := a; i <= N; i++ {
			for j := b; j <= M; j++ {
				total := sum[i][j] - sum[i][j-b] - sum[i-a][j] + sum[i-a][j-b]
				maxnum := mx[i-1][j-1]
				cnt := maxnum*a*b - total
				best = min(cnt, best)
			}
		}
		res[q] = best
	}

	return res
}

func solve1(N, M int, MTR [][]int, Q int, queries [][]int) []int {
	sum := make([][]int, N+1)
	for i := 0; i <= N; i++ {
		sum[i] = make([]int, M+1)
	}

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			sum[i+1][j+1] = sum[i+1][j] + sum[i][j+1] - sum[i][j] + MTR[i][j]
		}
	}

	sts := make([]*SegTree, N)

	for i := 0; i < N; i++ {
		sts[i] = InitSegTree(M, MTR[i])
	}

	res := make([]int, Q)
	col := make([]int, N)

	for q := 0; q < Q; q++ {
		a, b := queries[q][0], queries[q][1]
		best := 1 << 30
		for j := 0; j+b <= M; j++ {
			for i := 0; i < N; i++ {
				col[i] = sts[i].Query(j, j+b-1)
			}
			cur := InitSegTree(N, col)

			for i := 0; i+a <= N; i++ {
				total := sum[i+a][j+b] - sum[i][j+b] - sum[i+a][j] + sum[i][j]
				maxnum := cur.Query(i, i+a-1)
				cnt := maxnum*a*b - total
				best = min(cnt, best)
			}
		}

		res[q] = best
	}

	return res
}

type SegTree struct {
	tree []int
	size int
}

func InitSegTree(n int, arr []int) *SegTree {
	tree := make([]int, 4*n)

	var loop func(i int, left, right int)

	loop = func(i int, left, right int) {
		if left == right {
			tree[i] = arr[left]
			return
		}
		mid := (left + right) / 2
		loop(2*i+1, left, mid)
		loop(2*i+2, mid+1, right)
		tree[i] = max(tree[2*i+1], tree[2*i+2])
	}

	loop(0, 0, n-1)
	return &SegTree{tree, n}
}

func (st SegTree) Query(left, right int) int {
	var loop func(i int, start, end int) int
	loop = func(i int, start, end int) int {
		if end < left || right < start {
			return 0
		}
		if left <= start && end <= right {
			return st.tree[i]
		}

		mid := (start + end) / 2
		a := loop(2*i+1, start, mid)
		b := loop(2*i+2, mid+1, end)
		return max(a, b)
	}

	return loop(0, 0, st.size-1)
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
