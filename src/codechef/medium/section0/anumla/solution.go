package main

import (
	"bufio"
	"container/heap"
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
		n := readNum(scanner)

		m := 1 << uint(n)

		S := readNNums(scanner, m)
		R := solve(n, S)
		fmt.Printf("%d", R[0])

		for i := 1; i < n; i++ {
			fmt.Printf(" %d", R[i])
		}

		fmt.Println()

		t--
	}
}

func solve(n int, S []int) []int {
	sort.Ints(S)
	m := 1 << uint(n)
	a := make([]int, n)
	c := make([]int, m)

	h := &IntHeap{}
	heap.Init(h)
	var fptr int
	var ptr int
	for i := 1; i < m; i++ {
		if h.Len() > 0 && (*h)[0] == S[i] {
			heap.Pop(h)
		} else {
			a[fptr] = S[i]
			tptr := ptr
			for j := 0; j < tptr; j++ {
				c[ptr] = c[j] + a[fptr]
				heap.Push(h, c[ptr])
				ptr++
			}
			c[ptr] = a[fptr]
			ptr++
			fptr++
		}
	}

	return a
}

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
