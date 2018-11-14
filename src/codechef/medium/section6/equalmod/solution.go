package main

import (
	"bufio"
	"fmt"
	"math"
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

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		n := readNum(scanner)
		A := readNNums(scanner, n)
		B := readNNums(scanner, n)
		fmt.Println(solve(n, A, B))
	}
}

func solve(n int, A []int, B []int) int64 {
	ps := make(Pairs, n)
	mx := math.MaxInt32
	for i := 0; i < n; i++ {
		p := Pair{A[i] % B[i], B[i]}
		mx = min(mx, B[i]-1)
		ps[i] = p
	}

	sort.Sort(ps)

	lft := make([]int64, n+1)
	for i := 0; i < n; i++ {
		lft[i+1] = lft[i] + int64(ps[i].first)
	}

	rt := make([]int64, n+1)
	for i := n - 1; i >= 0; i-- {
		rt[i] = rt[i+1] + int64((ps[i].second-ps[i].first)%ps[i].second)
	}

	ans := rt[0]

	for i := 0; i < n; i++ {
		if ps[i].first <= mx {
			ans = min2(ans, int64(ps[i].first)*int64(n)-lft[i+1]+rt[i+1])
		}
	}
	return ans
}

type Pair struct {
	first, second int
}

type Pairs []Pair

func (ps Pairs) Len() int {
	return len(ps)
}

func (ps Pairs) Less(i, j int) bool {
	return ps[i].first < ps[j].first
}

func (ps Pairs) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func min2(a, b int64) int64 {
	if a <= b {
		return a
	}
	return b
}
