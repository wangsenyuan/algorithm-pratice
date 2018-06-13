package main

import (
	"sort"
	"bufio"
	"os"
	"fmt"
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
		fmt.Println(solve(n, A))
		t--
	}
}

func solve(n int, A []int) int64 {
	D := make([]int, n-1)
	for i := 0; i < n-1; i++ {
		D[i] = A[i] - A[i+1]
	}
	DN := len(D)
	sfs := make(SFS, DN)
	for i := 0; i < DN; i++ {
		sfs[i] = SF{i, make([]int, 2)}
	}
	m := sort.Search(30, func(i int) bool {
		return 1<<uint(i) > 2*DN
	})
	p := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		p[i] = make([]int, DN)
	}
	for i := 0; i < DN; i++ {
		p[0][i] = D[i]
	}

	var step = 1
	for cnt := 1; cnt < 2*DN; step, cnt = step+1, cnt*2 {
		for i := 0; i < DN; i++ {
			sfs[i].index = i
			sfs[i].rank[0] = p[step-1][i]
			sfs[i].rank[1] = -1000
			if i+cnt < len(D) {
				sfs[i].rank[1] = p[step-1][i+cnt]
			}
		}
		sort.Sort(sfs)

		for i := 0; i < DN; i++ {
			p[step][sfs[i].index] = i
			if i > 0 && sfs[i].rank[0] == sfs[i-1].rank[0] && sfs[i].rank[1] == sfs[i-1].rank[1] {
				p[step][sfs[i].index] = p[step][sfs[i-1].index]
			}
		}
	}

	longestCommonPrefix := func(x, y int) int {
		var res int
		for k := step - 1; k >= 0 && x < DN && y < DN; k-- {
			if p[k][x] == p[k][y] {
				res += 1 << uint(k)
				x += 1 << uint(k)
				y += 1 << uint(k)
			}
		}
		return res
	}

	var lcpSum int64
	for i := 1; i < DN; i++ {
		lcpSum += int64(longestCommonPrefix(sfs[i].index, sfs[i-1].index))
	}

	nn := int64(DN)
	return nn*(nn+1)/2 - lcpSum
}

type SF struct {
	index int
	rank  []int
}

type SFS []SF

func (this SFS) Len() int {
	return len(this)
}

func (this SFS) Less(i, j int) bool {
	a, b := this[i], this[j]
	return a.rank[0] < b.rank[0] || (a.rank[0] == b.rank[0] && a.rank[1] < b.rank[1])
}

func (this SFS) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
