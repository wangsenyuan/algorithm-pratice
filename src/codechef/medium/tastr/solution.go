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

	x := buildSuffixArray(n, s)

	lca := make([]int, n)

	for i := 1; i < n; i++ {
		a, b := x[i], x[i-1]
		var cnt int
		for a < n && b < n && s[a] == s[b] {
			a++
			b++
			cnt++
		}
		lca[i] = cnt
	}

	ans := int64(n - x[0])

	for i := 1; i < n; i++ {
		ans += int64(n - x[i] - lca[i])
	}

	return ans
}

func buildSuffixArray(n int, S []byte) []int {
	ss := make(suffixes, n)

	for i := 0; i < n; i++ {
		rank := make([]int, 2)
		rank[0] = int(S[i] - 'a')
		if i < n-1 {
			rank[1] = int(S[i+1] - 'a')
		} else {
			rank[1] = -1
		}
		ss[i] = suffix{i, rank}
	}

	sort.Sort(ss)

	x := make([]int, n)

	for k := 4; k < 2*n; k *= 2 {
		var rank int
		prevRank := ss[0].rank[0]
		x[ss[0].index] = rank

		for i := 1; i < n; i++ {
			if ss[i].rank[0] == prevRank && ss[i].rank[1] == ss[i-1].rank[1] {
				prevRank = ss[i].rank[0]
				ss[i].rank[0] = rank
			} else {
				prevRank = ss[i].rank[0]
				ss[i].rank[0] = rank + 1
				rank++
			}
			x[ss[i].index] = i
		}
		for i := 0; i < n; i++ {
			next := ss[i].index + k/2
			if next < n {
				ss[i].rank[1] = ss[x[next]].rank[0]
			} else {
				ss[i].rank[1] = -1
			}
		}
		sort.Sort(ss)
	}

	for i := 0; i < n; i++ {
		x[i] = ss[i].index
	}

	return x
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
