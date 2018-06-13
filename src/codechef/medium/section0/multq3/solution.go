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
	queries := make([][]int, m)
	for i := 0; i < m; i++ {
		queries[i] = readNNums(scanner, 3)
	}
	res := solve(n, queries)
	for _, num := range res {
		fmt.Println(num)
	}
}

func solve(n int, queries [][]int) []int {
	st := ConstructST(n)

	res := make([]int, 0, len(queries))

	for _, query := range queries {
		a, l, r := query[0], query[1], query[2]
		if a == 0 {
			st.Update(l, r)
		} else {
			res = append(res, st.Query(l, r))
		}
	}

	return res
}

type ST struct {
	n    int
	rem0 []int
	rem1 []int
	rem2 []int
	add  []int
}

func ConstructST(n int) *ST {
	m := 4*n + 1
	rem0, rem1, rem2, add := make([]int, m), make([]int, m), make([]int, m), make([]int, m)

	var build func(i int, l int, r int)

	build = func(i int, l int, r int) {
		if l == r {
			rem0[i] = 1
			return
		}
		mid := (l + r) / 2
		build(2*i+1, l, mid)
		build(2*i+2, mid+1, r)
		rem0[i] = rem0[2*i+1] + rem0[2*i+2]
	}

	build(0, 0, n-1)

	return &ST{n, rem0, rem1, rem2, add}
}

func (st *ST) Update(left, right int) {

	change := func(i int) {
		st.rem0[i], st.rem1[i], st.rem2[i] = st.rem2[i], st.rem0[i], st.rem1[i]
	}

	var update func(i int, l int, r int, a int, b int)
	update = func(i int, l int, r int, a int, b int) {
		if l == a && r == b {
			st.add[i]++
			change(i)
			return
		}
		mid := (l + r) / 2
		if b <= mid {
			update(2*i+1, l, mid, a, b)
		} else if a > mid {
			update(2*i+2, mid+1, r, a, b)
		} else {
			update(2*i+1, l, mid, a, mid)
			update(2*i+2, mid+1, r, mid+1, b)
		}

		st.rem0[i] = st.rem0[2*i+1] + st.rem0[2*i+2]
		st.rem1[i] = st.rem1[2*i+1] + st.rem1[2*i+2]
		st.rem2[i] = st.rem2[2*i+1] + st.rem2[2*i+2]

		for j := 0; j < st.add[i]%3; j++ {
			change(i)
		}
	}

	update(0, 0, st.n-1, left, right)
}

func (st ST) Query(left, right int) int {
	var query func(i int, l, r, a, b, added int) int

	query = func(i, l, r, a, b, added int) int {
		if l == a && r == b {
			if added%3 == 0 {
				return st.rem0[i]
			} else if added%3 == 1 {
				return st.rem2[i]
			}
			return st.rem1[i]
		}
		added += st.add[i]
		mid := (l + r) / 2
		if b <= mid {
			return query(2*i+1, l, mid, a, b, added)
		} else if a > mid {
			return query(2*i+2, mid+1, r, a, b, added)
		}
		return query(2*i+1, l, mid, a, mid, added) + query(2*i+2, mid+1, r, mid+1, b, added)
	}

	return query(0, 0, st.n-1, left, right, 0)
}
