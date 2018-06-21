package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

// 605293
const N = 1000000

var cycleLength = make([]int, N+1)

var st SegTree

func init() {

	for i := 0; 1<<uint(i) <= N; i++ {
		cycleLength[1<<uint(i)] = i + 1
	}

	var f func(n int) int

	f = func(n int) int {
		if n <= N && cycleLength[n] > 0 {
			return cycleLength[n]
		}
		var res int
		if n&1 == 1 {
			res = 1 + f(n*3+1)
		} else {
			res = 1 + f(n>>1)
		}
		if n <= N {
			cycleLength[n] = res
		}
		return res
	}

	for i := 1; i <= N; i++ {
		cycleLength[i] = f(i)
	}

	st = CreateTree(cycleLength, N+1)
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
	scanner.Scan()
	x := readInt(scanner.Bytes(), 0, &a)
	readInt(scanner.Bytes(), x+1, &b)
	return
}

func readNNums(scanner *bufio.Scanner, n int) []int {
	res := make([]int, n)
	x := -1
	scanner.Scan()
	var i int
	for i < n {
		x = readInt(scanner.Bytes(), x+1, &res[i])
		i++
		if x == len(scanner.Bytes()) {
			break
		}
	}
	return res[:i]
}

func main() {
	f, err := os.OpenFile("./input.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("open file error: %v", err)
		return
	}

	defer f.Close()

	// reader := fromConsol()
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		var a, b int
		pos := readInt(scanner.Bytes(), 0, &a)
		readInt(scanner.Bytes(), pos+1, &b)
		fmt.Printf("%d %d %d\n", a, b, solve(a, b))
	}
}

func fromConsol() io.Reader {
	return os.Stdin
}

func solve(a, b int) int {
	if a > b {
		a, b = b, a
	}
	return st.Query(a, b)
}

type SegTree struct {
	tree []int
	size int
}

func CreateTree(arr []int, n int) SegTree {
	tree := make([]int, 4*n)

	var loop func(i int, start, end int)

	loop = func(i int, start, end int) {
		if start == end {
			tree[i] = arr[start]
			return
		}
		mid := (start + end) / 2
		loop(2*i+1, start, mid)
		loop(2*i+2, mid+1, end)
		tree[i] = max(tree[2*i+1], tree[2*i+2])
	}

	loop(0, 0, n-1)
	return SegTree{tree, n}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (st SegTree) Query(left, right int) int {
	tree := st.tree

	var loop func(i int, start, end int) int

	loop = func(i int, start, end int) int {
		if end < left || start > right {
			// MIN
			return 0
		}
		if left <= start && end <= right {
			return tree[i]
		}
		mid := (start + end) / 2
		return max(loop(2*i+1, start, mid), loop(2*i+2, mid+1, end))
	}

	return loop(0, 0, st.size-1)
}
