package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
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

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	n := readNum(scanner)

	A := readNNums(scanner, n)

	m := readNum(scanner)

	queries := make([][]int, m)

	for i := 0; i < m; i++ {
		scanner.Scan()
		bs := scanner.Bytes()
		var pos int
		var kind int
		pos = readInt(bs, pos, &kind)
		if kind == 0 {
			queries[i] = make([]int, 4)
			for j := 1; j < 4; j++ {
				pos = readInt(bs, pos+1, &queries[i][j])
			}
		} else {
			queries[i] = make([]int, 3)
			queries[i][0] = 1
			for j := 1; j < 3; j++ {
				pos = readInt(bs, pos+1, &queries[i][j])
			}
		}
	}

	// fmt.Printf("[debug] %d %v %d %v\n", n, A, m, queries)
	res := solve(n, A, m, queries)

	for _, ans := range res {
		fmt.Println(ans)
	}
}

func solve(n int, A []int, m int, queries [][]int) []int {
	st := NewSegTree(n, A)

	res := make([]int, m)
	var j int

	for i := 0; i < m; i++ {
		query := queries[i]
		if query[0] == 0 {
			left, right, force := query[1], query[2], query[3]
			st.Update(left, right, force)
		} else {
			left, right := query[1], query[2]
			res[j] = st.Query(left, right)
			j++
		}
	}
	return res[:j]
}

type SegTree struct {
	array [][]int
	p     []int
	size  int
}

func NewSegTree(n int, input []int) SegTree {

	array := make([][]int, 4*n)
	p := make([]int, 4*n)

	var loop func(root int, start int, end int)
	loop = func(root int, start int, end int) {
		if start == end {
			array[root] = make([]int, 12)
			for j := 0; j < 12; j++ {
				array[root][j] = rotateLeft(input[start], j)
			}
			return
		}
		mid := (start + end) >> 1
		loop(root<<1, start, mid)
		loop(root<<1|1, mid+1, end)
		array[root] = merge(array[root<<1], array[root<<1|1])
	}

	loop(1, 0, n-1)

	return SegTree{array, p, n}
}

func rotateLeft(x int, num int) int {
	digits := make([]int, 4)
	i := 4
	for x > 0 {
		i--
		digits[i] = x % 10
		x /= 10
	}
	digits = digits[i:]
	n := len(digits)
	num %= n

	for j := 0; j < n; j++ {
		x = x*10 + digits[(j+num)%n]
	}
	return x
}

func merge(a, b []int) []int {
	c := make([]int, 12)

	for i := 0; i < 12; i++ {
		c[i] = max(a[i], b[i])
	}
	return c
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func (st *SegTree) Update(left, right int, force int) {
	array := st.array
	p := st.p

	var loop func(root int, start, end int)

	loop = func(root int, start, end int) {
		if end < left || start > right {
			return
		}
		if p[root] != 0 {
			a := array[root]
			b := make([]int, 12)
			copy(b, a)

			for i := 0; i < 12; i++ {
				a[i] = b[(i+p[root])%12]
			}
			if start != end {
				p[root<<1] = (p[root<<1] + p[root]) % 12
				p[root<<1|1] = (p[root<<1|1] + p[root]) % 12
			}
			p[root] = 0
		}

		if left <= start && end <= right {
			return
		}

		mid := (start + end) >> 1
		loop(root<<1, start, mid)
		loop(root<<1|1, mid+1, end)
	}
	p[1] = (p[1] + force) % 12

	loop(1, 0, st.size-1)
}

func (st *SegTree) Query(left, right int) int {
	var loop func(root int, start, end int) int
	p := st.p
	array := st.array
	loop = func(root int, start, end int) int {
		if end < left || start > right {
			return 0
		}
		if p[root] != 0 {
			a := array[root]
			b := make([]int, 12)
			copy(b, a)

			for i := 0; i < 12; i++ {
				a[i] = b[(i+p[root])%12]
			}
			if start != end {
				p[root<<1] = (p[root<<1] + p[root]) % 12
				p[root<<1|1] = (p[root<<1|1] + p[root]) % 12
			}

			p[root] = 0
		}

		if left <= start && end <= right {
			return array[root][0]
		}
		mid := (start + end) >> 1
		a := loop(root<<1, start, mid)
		b := loop(root<<1|1, mid+1, end)
		return max(a, b)
	}

	return loop(1, 0, st.size-1)
}
