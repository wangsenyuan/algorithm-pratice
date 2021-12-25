package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func fillNNums(scanner *bufio.Scanner, n int, res []int) {
	x := 0
	scanner.Scan()
	for i := 0; i < n; i++ {
		for x < len(scanner.Bytes()) && scanner.Bytes()[x] == ' ' {
			x++
		}
		x = readInt(scanner.Bytes(), x, &res[i])
	}
}

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var tmp int64
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp
	return i
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

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		n := readNum(scanner)
		A := readNNums(scanner, n)
		k := readNum(scanner)
		x := readNum(scanner)
		res := solve(n, A, k, x)
		fmt.Println(res)
	}
}

func solve(n int, A []int, k int, x int) int64 {
	X := int64(x)

	B := make([]Pair, n)

	for i := 0; i < n; i++ {
		a := int64(A[i])
		b := a ^ X
		B[i] = Pair{a, b}
	}

	var a int64
	var b int64

	for i := 0; i < n; i++ {
		a += B[i].first
		b += B[i].second
	}

	if k == n {

		return max(a, b)
	}

	if k%2 == 1 {
		for i := 0; i < n; i++ {
			tmp := B[i].second - B[i].first
			if tmp > 0 {
				a += tmp
			}
		}

		return a
	}

	sort.Sort(Pairs(B))

	for i := 0; 2*i+1 < n; i++ {
		tmp := B[2*i].second - B[2*i].first + B[2*i+1].second - B[2*i+1].first
		if tmp > 0 {
			a += tmp
		}
	}

	return a
}

type Pair struct {
	first  int64
	second int64
}

type Pairs []Pair

func (this Pairs) Len() int {
	return len(this)
}

func (this Pairs) Less(i, j int) bool {
	a := this[i].first - this[i].second
	b := this[j].first - this[j].second
	return a < b
}

func (this Pairs) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func max(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}
