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

func readThreeNums(scanner *bufio.Scanner) (a int, b int, c int) {
	scanner.Scan()
	x := readInt(scanner.Bytes(), 0, &a)
	x = readInt(scanner.Bytes(), x+1, &b)
	readInt(scanner.Bytes(), x+1, &c)
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
		n, k, p := readThreeNums(scanner)
		A := make([][]int, k)
		for i := 0; i < k; i++ {
			A[i] = readNNums(scanner, 2)
		}

		P := make([][]int, p)
		for i := 0; i < p; i++ {
			P[i] = readNNums(scanner, 2)
		}

		res := solve(n, k, p, A, P)

		if res {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}

		t--
	}
}

func solve(n, k, p int, A [][]int, P [][]int) bool {
	if k < p {
		return false
	}
	xs := make(PRS, k)
	for i := 0; i < k; i++ {
		xs[i] = &PR{A[i][0], A[i][1]}
	}

	ys := make(PRS, p)
	for i := 0; i < p; i++ {
		ys[i] = &PR{P[i][0], P[i][1]}
	}

	sort.Sort(xs)
	sort.Sort(ys)

	var ii, jj int

	for ii < k && jj < p {
		if xs[ii].from == ys[jj].from {
			break
		}
		if xs[ii].from < ys[jj].from {
			ii++
		} else {
			jj++
		}
	}

	if ii == k || jj == p {
		// can not align
		return false
	}

	xs = process(xs, ii, k, n)

	ys = process(ys, jj, p, n)

	ii, jj = 0, 0
	for ii < k && jj < p {
		if xs[ii].end > ys[jj].end {
			return false
		}
		if xs[ii].end == ys[jj].end {
			jj++
		}
		ii++
	}
	return ii == k && jj == p
}

func process(xs PRS, ii int, k int, n int) PRS {
	if ii > 0 {
		zs := make(PRS, k)
		copy(zs, xs[ii:])
		copy(zs[k-ii:], xs[:ii])
		xs = zs
	}
	for i := 0; i < k; i++ {
		if i > 0 {
			if xs[i].from < xs[i-1].end {
				xs[i].from += n
			}
		}
		if xs[i].end < xs[i].from {
			xs[i].end += n
		}
	}
	return xs
}

type PR struct {
	from, end int
}

type PRS []*PR

func (this PRS) Len() int {
	return len(this)
}

func (this PRS) Less(i, j int) bool {
	return this[i].from < this[j].from
}

func (this PRS) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
