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
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
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
	n := readNum(scanner)
	edges := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		edges[i] = readNNums(scanner, 3)
	}
	fmt.Println(solve(n, edges))
}

func solve(n int, edges [][]int) float64 {
	set := make([]int, n+1)
	cnt := make([]int, n+1)

	for i := 1; i <= n; i++ {
		cnt[i] = 1
	}

	var find func(x int) int

	find = func(x int) int {
		px := set[x]
		if px == 0 {
			set[x] = x
		} else if px != x {
			set[x] = find(px)
		}
		return set[x]
	}

	sort.Sort(Edges(edges))

	var sum int

	for _, edge := range edges {
		sum += edge[2]
	}

	var ans int
	for _, edge := range edges {
		a, b, c := edge[0], edge[1], edge[2]
		pa, pb := find(a), find(b)
		//pa can't equal pb
		ca, cb := cnt[pa], cnt[pb]
		ans += ca * cb * c
		if ca < cb {
			set[pa] = pb
			cnt[pb] += cnt[pa]
		} else {
			set[pb] = pa
			cnt[pa] += cnt[pb]
		}
	}
	total := n * (n - 1) / 2

	return float64(sum) - float64(ans)/float64(total)
}

type Edges [][]int

func (this Edges) Len() int {
	return len(this)
}

func (this Edges) Less(i, j int) bool {
	return this[i][2] < this[j][2]
}

func (this Edges) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
