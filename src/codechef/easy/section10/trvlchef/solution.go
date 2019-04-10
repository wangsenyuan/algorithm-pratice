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
		n, D := readTwoNums(scanner)
		C := readNNums(scanner, n)
		res := solve(n, D, C)
		if res {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func solve(n int, D int, C []int) bool {
	ps := make([]Pair, n)
	for i := 0; i < n; i++ {
		ps[i] = Pair{i, C[i]}
	}

	sort.Sort(Pairs(ps))

	pos := 0
	for i := 0; i < n; i++ {
		if ps[i].first == 0 {
			pos = i
			break
		}
	}

	v := true
	vis := make([]bool, n)
	vis[pos] = true

	for pos >= 1 {
		nxt := max(0, pos-2)
		v = v && abs(C[pos]-C[nxt]) <= D
		pos = nxt
		vis[pos] = true
	}

	for i := 0; i < n; i++ {
		if !vis[i] {
			v = v && abs(C[i]-C[pos]) <= D
			pos = i
		}
	}

	return v
}

type Pair struct {
	first  int
	second int
}

type Pairs []Pair

func (ps Pairs) Len() int {
	return len(ps)
}

func (ps Pairs) Less(i, j int) bool {
	return ps[i].second < ps[j].second
}

func (ps Pairs) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
