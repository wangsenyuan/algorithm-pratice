package main

import (
	"sort"
	"bufio"
	"os"
	"fmt"
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

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := int64(0)
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * int64(sign)
	return i
}

func readInt64Num(scanner *bufio.Scanner) (a int64) {
	scanner.Scan()
	readInt64(scanner.Bytes(), 0, &a)
	return
}

func readTwoInt64Nums(scanner *bufio.Scanner) (a int64, b int64) {
	scanner.Scan()
	x := readInt64(scanner.Bytes(), 0, &a)
	readInt64(scanner.Bytes(), x+1, &b)
	return
}

func readNInt64Nums(scanner *bufio.Scanner, n int) []int64 {
	res := make([]int64, n)
	x := -1
	scanner.Scan()
	for i := 0; i < n; i++ {
		x = readInt64(scanner.Bytes(), x+1, &res[i])
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	var n int
	pos := readInt(scanner.Bytes(), 0, &n)
	var K int64
	readInt64(scanner.Bytes(), pos+1, &K)
	A := readNInt64Nums(scanner, n)
	fmt.Println(solve(n, K, A))
}

func solve(n int, K int64, A []int64) int64 {
	S := make([]int64, n)
	for i := 0; i < n; i++ {
		S[i] = A[i] - K
		if i > 0 {
			S[i] += S[i-1]
		}
	}
	B := make([]int64, n)
	copy(B, S)

	sort.Sort(Int64Slice(S))

	idx := make(map[int64]int)

	idx[S[0]] = 0
	for i, j := 1, 1; i < n; i++ {
		if S[i] == S[i-1] {
			continue
		}
		idx[S[i]] = j
		j++
	}

	st := CreateSegTree(len(idx))
	var res int64
	for i := 0; i < n; i++ {
		if B[i] >= 0 {
			res++
		}
		res += st.Get(idx[B[i]])
		st.Add(idx[B[i]])
	}

	return res
}

type SegTree struct {
	tree []int64
	size int
}

func CreateSegTree(n int) *SegTree {
	size := 1
	for size < n {
		size <<= 1
	}

	tree := make([]int64, 2*size)
	return &SegTree{tree, size}
}

func (st *SegTree) Add(idx int) {
	idx += st.size
	st.tree[idx]++
	prev := idx
	idx >>= 1
	for idx > 0 {
		if prev&1 == 0 {
			//from a left child
			st.tree[idx]++
		}
		prev = idx
		idx >>= 1
	}

}

func (st *SegTree) Get(idx int) int64 {
	// get count that before idx
	idx += st.size
	res := st.tree[idx]
	prev := idx
	idx >>= 1

	for idx > 0 {
		if prev&1 == 1 {
			// from a right child
			res += st.tree[idx]
		}
		prev = idx
		idx >>= 1
	}

	return res
}

type Int64Slice []int64

func (this Int64Slice) Len() int {
	return len(this)
}

func (this Int64Slice) Less(i, j int) bool {
	return this[i] < this[j]
}

func (this Int64Slice) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
