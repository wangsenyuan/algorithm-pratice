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

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	n, m := readTwoNums(scanner)

	X := make([][]int, n)

	for i := 0; i < n; i++ {
		X[i] = readNNums(scanner, 2)
	}

	Y := make([][]int, m)

	for i := 0; i < m; i++ {
		Y[i] = readNNums(scanner, 2)
	}

	fmt.Println(solve(n, m, X, Y))
}

func solve(n int, m int, X [][]int, Y [][]int) int {
	xs := make([]Pair, n)

	for i := 0; i < n; i++ {
		xs[i] = Pair{X[i][0], X[i][1]}
	}

	sort.Sort(Pairs(xs))

	ys := make([]Pair, m)

	for i := 0; i < m; i++ {
		ys[i] = Pair{Y[i][0], Y[i][1]}
	}

	sort.Sort(Pairs(ys))

	var res int

	var i, j int

	for i < m && j < n {
		a := xs[i].Value()
		b := ys[j].Value()
		if a == b {
			res++
			i++
			j++
		} else if a < b {
			i++
		} else {
			j++
		}
	}

	return res
}

type Pair struct {
	first  int
	second int
}

func (p Pair) Value() int64 {
	return int64(p.first) * int64(p.second)
}

type Pairs []Pair

func (ps Pairs) Len() int {
	return len(ps)
}

func (ps Pairs) Less(i, j int) bool {
	a := ps[i].Value()
	b := ps[j].Value()

	if a < b {
		return true
	}

	if a == b {
		return ps[i].first < ps[j].second
	}

	return false
}

func (ps Pairs) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}
