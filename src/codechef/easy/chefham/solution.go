package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func readInt(bs []byte, from int, val *int) int {
	tmp := 0
	i := from
	for i < len(bs) && bs[i] != ' ' {
		tmp = tmp*10 + int(bs[i]-'0')
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

func readNNums(scanner *bufio.Scanner, n int) []int {
	scanner.Scan()
	bs := scanner.Bytes()
	res := make([]int, n)
	x := -1
	for i := 0; i < n; i++ {
		x = readInt(bs, x+1, &res[i])
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	t := readNum(scanner)
	for i := 0; i < t; i++ {
		n := readNum(scanner)
		A := readNNums(scanner, n)
		hd, B := solve(A)
		fmt.Println(hd)
		for j := 0; j < n; j++ {
			fmt.Printf("%d", B[j])
			if j < n-1 {
				fmt.Print(" ")
			} else {
				fmt.Println()
			}
		}
	}
}

func solve(A []int) (int, []int) {
	if len(A) == 1 {
		return 0, A
	}
	n := len(A)
	B := make([]int, n)
	if n == 2 {
		B[0], B[1] = A[1], A[0]
	} else {
		pairs := make([]Pair, n)
		for i, a := range A {
			pairs[i] = Pair{a, i}
		}
		sort.Sort(PairSlice(pairs))
		for i := 0; i < n; i++ {
			B[pairs[i].w] = A[pairs[(i-2+n)%n].w]
		}
	}
	var best int

	for i := 0; i < n; i++ {
		if A[i] != B[i] {
			best++
		}
	}

	return best, B
}

type Pair struct {
	v int
	w int
}

type PairSlice []Pair

func (this PairSlice) Len() int {
	return len(this)
}

func (this PairSlice) Less(i, j int) bool {
	return this[i].v < this[j].v
}

func (this PairSlice) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
