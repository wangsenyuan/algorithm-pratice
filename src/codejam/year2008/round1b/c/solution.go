package main

import (
	"fmt"
	"os"
)

func main() {
	// f, err := os.Open("./A-small-practice.in")
	f, err := os.Open("./C-large-practice.in")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	// scanner := bufio.NewScanner(f)

	// tc := readNum(scanner)
	var tc int
	fmt.Fscanf(f, "%d", &tc)
	for i := 1; i <= tc; i++ {
		var K int
		fmt.Fscanf(f, "%d", &K)
		var N int
		fmt.Fscanf(f, "%d", &N)
		queries := make([]int, N)
		for j := 0; j < N; j++ {
			fmt.Fscanf(f, "%d", &queries[j])
		}
		res := solve(K, N, queries)
		fmt.Printf("Case #%d: ", i)
		for j := 0; j < N; j++ {
			if j < N-1 {
				fmt.Printf("%d ", res[j])
			} else {
				fmt.Printf("%d\n", res[j])
			}
		}
	}
}

func solve(K int, N int, queries []int) []int {
	st := InitSegTree(K)
	// NO of avalible slots in the left side
	var X int
	// Total NO of avaliable slots
	T := K
	ans := make([]int, K)
	for i := 1; i <= K; i++ {
		// need to find ith avalible slot after X
		Y := X + i
		Y %= T
		if Y == 0 {
			Y = T
		}
		j := st.Query(Y)
		ans[j] = i
		st.Mark(j)
		X = Y - 1
		T--
	}

	res := make([]int, N)
	for i := 0; i < N; i++ {
		res[i] = ans[queries[i]-1]
	}
	return res
}

type SegTree struct {
	tree []int
	size int
}

func InitSegTree(n int) SegTree {
	tree := make([]int, 4*n)
	var loop func(i int, start, end int)
	loop = func(i int, start, end int) {
		if start == end {
			tree[i] = 1
			return
		}

		mid := (start + end) >> 1
		loop(2*i+1, start, mid)
		loop(2*i+2, mid+1, end)
		tree[i] = tree[2*i+1] + tree[2*i+2]
	}
	loop(0, 0, n-1)
	return SegTree{tree, n}
}

func (st SegTree) Query(val int) int {
	tree := st.tree
	var loop func(i int, start, end int) int

	loop = func(i int, start, end int) int {
		if start == end {
			return start
		}
		mid := (start + end) >> 1
		if val <= tree[2*i+1] {
			return loop(2*i+1, start, mid)
		}
		val -= tree[2*i+1]
		return loop(2*i+2, mid+1, end)
	}
	return loop(0, 0, st.size-1)
}

func (st *SegTree) Mark(pos int) {
	tree := st.tree
	var loop func(i int, start, end int)
	loop = func(i int, start, end int) {
		tree[i]--
		if start == end {
			return
		}
		mid := (start + end) >> 1
		if pos <= mid {
			loop(2*i+1, start, mid)
		} else {
			loop(2*i+2, mid+1, end)
		}
	}
	loop(0, 0, st.size-1)
}
