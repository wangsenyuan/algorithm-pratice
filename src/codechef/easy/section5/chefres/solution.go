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

	tc := readNum(scanner)
	for tc > 0 {
		tc--
		n, m := readTwoNums(scanner)
		L := make([]int, n)
		R := make([]int, n)
		for i := 0; i < n; i++ {
			L[i], R[i] = readTwoNums(scanner)
		}
		P := make([]int, m)
		for i := 0; i < m; i++ {
			P[i] = readNum(scanner)
		}
		res := solve(n, L, R, m, P)
		for _, num := range res {
			fmt.Println(num)
		}
	}
}

func solve(n int, L []int, R []int, m int, P []int) []int {
	ps := make(PS, n)
	for i := 0; i < n; i++ {
		ps[i] = Pair{L[i], R[i]}
	}
	sort.Sort(ps)

	res := make([]int, m)
	for i := 0; i < m; i++ {
		j := sort.Search(n, func(j int) bool {
			return ps[j].second > P[i]
		})

		if j == n {
			res[i] = -1
		} else if ps[j].first <= P[i] {
			res[i] = 0
		} else {
			res[i] = ps[j].first - P[i]
		}
	}

	return res
}

type Pair struct {
	first  int
	second int
}

type PS []Pair

func (ps PS) Len() int {
	return len(ps)
}

func (ps PS) Less(i, j int) bool {
	return ps[i].second < ps[j].second
}

func (ps PS) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}
