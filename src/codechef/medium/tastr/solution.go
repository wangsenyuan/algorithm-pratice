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

	scanner.Scan()
	a := scanner.Bytes()
	scanner.Scan()
	b := scanner.Bytes()
	fmt.Println(solve(a, b))
}

func solve(a, b []byte) int64 {
	ua := uniqueSubStrCount(a)
	ub := uniqueSubStrCount(b)
	s := make([]byte, len(a)+len(b)+1)
	copy(s, a)
	s[len(a)] = '$'
	copy(s[len(a)+1:], b)
	uab := uniqueSubStrCount(s) - int64(len(a)+1)*int64(len(b)+1)
	return 2*uab - ua - ub
}

func uniqueSubStrCount(s []byte) int64 {
	n := len(s)

	x, p, step := buildSuffixArray(n, s)

	compute := func(i int) int {
		var res int
		a, b := x[i-1], x[i]
		for k := step - 1; k >= 0 && a < n && b < n; k-- {
			if p[k][a] == p[k][b] {
				a += 1 << uint(k)
				b += 1 << uint(k)
				res += 1 << uint(k)
			}
		}

		return res
	}

	lca := make([]int, n)

	for i := 1; i < n; i++ {
		lca[i] = compute(i)
	}

	ans := int64(n - x[0])

	for i := 1; i < n; i++ {
		ans += int64(n - x[i] - lca[i])
	}

	return ans
}

func buildSuffixArray(n int, S []byte) ([]int, [][]int, int) {
	ss := make(suffixes, n)
	for i := 0; i < n; i++ {
		ss[i] = suffix{i, make([]int, 2)}
	}
	m := sort.Search(30, func(i int) bool {
		return 1<<uint(i) > 2*n
	})

	p := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		p[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		p[0][i] = int(S[i] - 'a')
	}
	stp := 1
	for cnt := 1; cnt < 2*n; stp, cnt = stp+1, cnt<<1 {
		for i := 0; i < n; i++ {
			ss[i].rank[0] = p[stp-1][i]
			ss[i].rank[1] = -1
			if i+cnt < n {
				ss[i].rank[1] = p[stp-1][i+cnt]
			}
			ss[i].index = i
		}
		sort.Sort(ss)
		for i := 0; i < n; i++ {
			p[stp][ss[i].index] = i
			if i > 0 && ss[i].rank[0] == ss[i-1].rank[0] && ss[i].rank[1] == ss[i-1].rank[1] {
				p[stp][ss[i].index] = p[stp][ss[i-1].index]
			}
		}
	}

	x := make([]int, n)

	for i := 0; i < n; i++ {
		x[i] = ss[i].index
	}

	return x, p, stp
}

type suffix struct {
	index int
	rank  []int
}

type suffixes []suffix

func (s suffixes) Len() int {
	return len(s)
}

func (s suffixes) Less(i, j int) bool {
	a, b := s[i], s[j]
	return a.rank[0] < b.rank[0] || (a.rank[0] == b.rank[0] && a.rank[1] < b.rank[1])
}

func (s suffixes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
