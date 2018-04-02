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

	n, m := readTwoNums(scanner)

	ops := make([][]int, m)

	for i := 0; i < m; i++ {
		ops[i] = readNNums(scanner, 3)
	}

	res := solve(n, ops)
	for i := 0; i < len(res); i++ {
		fmt.Println(res[i])
	}
}

func solve(n int, ops [][]int) []int {
	st := ConstructSegTree(n)

	res := make([]int, 0, len(ops))
	for _, op := range ops {
		left, right := op[1], op[2]
		if op[0] == 0 {
			st.Flip(left, right)
		} else {
			res = append(res, st.Query(left, right))
		}
	}

	return res
}

type SegTree struct {
	n     int
	heads []int
	flips []int
}

func ConstructSegTree(n int) *SegTree {
	h := 0
	for 1<<uint(h) < n {
		h++
	}
	sz := 1<<uint(h+1) - 1
	heads := make([]int, sz)
	flips := make([]int, sz)
	return &SegTree{n, heads, flips}
}

func (this *SegTree) Flip(left, right int) {
	var fn func(i, s, e int)

	fn = func(i, s, e int) {
		if s > e || s > right || e < left {
			return
		}
		if s >= left && e <= right {
			this.flips[i]++
			this.heads[i] = e - s + 1 - this.heads[i]
			return
		}

		mid := (s + e) / 2
		fn(2*i+1, s, mid)
		fn(2*i+2, mid+1, e)

		this.heads[i] = this.heads[2*i+1] + this.heads[2*i+2]
		if this.flips[i]%2 == 1 {
			this.heads[i] = e - s + 1 - this.heads[i]
		}
	}

	fn(0, 0, this.n-1)
}

func (this *SegTree) Query(left, right int) int {

	var fn func(i, s, e, flips int) int

	fn = func(i, s, e, flips int) int {
		if s > e || s > right || e < left {
			return 0
		}
		if s >= left && e <= right {
			if flips%2 == 1 {
				return e - s + 1 - this.heads[i]
			}
			return this.heads[i]
		}

		flips += this.flips[i]

		mid := (s + e) / 2

		return fn(2*i+1, s, mid, flips) + fn(2*i+2, mid+1, e, flips)
	}

	return fn(0, 0, this.n-1, 0)
}
