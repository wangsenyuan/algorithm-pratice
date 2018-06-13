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

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	t := readNum(scanner)

	for t > 0 {
		n := readNum(scanner)

		A := readNNums(scanner, n)
		B := readNNums(scanner, n)
		res := solve(n, A, B)
		for i := 0; i < n; i++ {
			if i < n-1 {
				fmt.Printf("%d ", res[i])
			} else {
				fmt.Printf("%d\n", res[i])
			}
		}
		t--
	}
}

func solve(n int, A []int, B []int) []uint64 {
	items := make(Items, n)

	for i := 0; i < n; i++ {
		items[i] = Item{uint64(A[i]), uint64(B[i])}
	}

	sort.Sort(items)

	f := make([]uint64, n+1)

	for i := 0; i < n; i++ {
		for j := i; j >= 0; j-- {
			f[j+1] = max(f[j+1], f[j]+uint64(j)*items[i].b+items[i].a)
		}
	}

	return f[1:]
}

func max(a, b uint64) uint64 {
	if a > b {
		return a
	}
	return b
}

type Item struct {
	a, b uint64
}

type Items []Item

func (items Items) Len() int {
	return len(items)
}

func (items Items) Less(i, j int) bool {
	return items[i].b > items[j].b
}

func (items Items) Swap(i, j int) {
	items[i], items[j] = items[j], items[i]
}
