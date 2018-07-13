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

	tc := readNum(scanner)

	for tc > 0 {
		n := readNum(scanner)
		seq := make([]int, n-1)
		for i := 0; i < n-1; i++ {
			seq[i] = readNum(scanner)
		}
		res := solve(n, seq)
		for _, ans := range res {
			fmt.Printf("%d\n", ans)
		}
		tc--
	}
}

func solve(n int, seq []int) []int {
	P := make([][]int, n)

	for i := 0; i < n; i++ {
		P[i] = make([]int, 20)
	}

	level := make([]int, n)
	// connect v to u
	connect := func(u, v int) {
		P[v][0] = u
		level[v] = level[u] + 1
		for i := 1; i < 20; i++ {
			P[v][i] = P[P[v][i-1]][i-1]
		}
	}

	res := make([]int, n-1)

	lca := func(u, v int) int {
		if level[u] < level[v] {
			u, v = v, u
		}

		for i := 19; i >= 0; i-- {
			if level[u]-(1<<uint(i)) >= level[v] {
				u = P[u][i]
			}
		}

		if u == v {
			return u
		}
		for i := 19; i >= 0; i-- {
			if P[u][i] != P[v][i] {
				u = P[u][i]
				v = P[v][i]
			}
		}
		return P[u][0]
	}

	u, v, diameter := 0, 0, 0

	pathLength := func(a, b int) int {
		p := lca(a, b)
		return level[a] - level[p] + level[b] - level[p]
	}

	compute := func(w int) int {
		uw := pathLength(u, w)
		vw := pathLength(v, w)

		if uw >= diameter && uw >= vw {
			v = w
			diameter = uw
		} else if vw >= diameter && vw >= uw {
			u = w
			diameter = vw
		}
		return diameter
	}

	for i := 1; i < n; i++ {
		// from 1 not 0
		connect(seq[i-1]-1, i)
		res[i-1] = compute(i)
	}

	return res
}
