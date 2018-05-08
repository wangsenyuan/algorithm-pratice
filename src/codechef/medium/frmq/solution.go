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

	n := readNum(scanner)

	A := readNNums(scanner, n)

	second := readNNums(scanner, 3)

	m := second[0]
	x, y := second[1], second[2]

	fmt.Println(solve(n, m, A, x, y))
}

func solve(n int, m int, A []int, x, y int) int64 {
	rmq := ConstructRMQ(n, A)
	var res int64
	for i := 0; i < m; i++ {
		a := min(x, y)
		b := max(x, y)
		res += int64(rmq.QueryRange(a, b))
		x = (x + 7) % (n - 1)
		y = (y + 11) % n
	}
	return res
}

type RMQ struct {
	tree []int
	size int
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func ConstructRMQ(n int, arr []int) *RMQ {
	m := 4 * n
	tree := make([]int, m)

	var init func(i int, l, r int)
	init = func(i int, l, r int) {
		if l == r {
			tree[i] = arr[l]
			return
		}
		mid := l + (r-l)/2
		init(2*i+1, l, mid)
		init(2*i+2, mid+1, r)
		tree[i] = max(tree[2*i+1], tree[2*i+2])
	}
	init(0, 0, n-1)
	return &RMQ{tree, n}
}

func (rmq RMQ) QueryRange(left, right int) int {
	var loop func(i int, l, r int) int
	loop = func(i int, l, r int) int {
		if r < left || right < l {
			return 0
		}
		if left <= l && r <= right {
			return rmq.tree[i]
		}
		mid := l + (r-l)/2
		return max(loop(2*i+1, l, mid), loop(2*i+2, mid+1, r))
	}

	return loop(0, 0, rmq.size-1)
}
