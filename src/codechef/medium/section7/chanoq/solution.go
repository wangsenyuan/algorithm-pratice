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
		segs := make([][]int, n)
		for i := 0; i < n; i++ {
			segs[i] = readNNums(scanner, 2)
		}
		q := readNum(scanner)
		quries := make([][]int, q)
		for i := 0; i < q; i++ {
			scanner.Scan()
			var m int
			pos := readInt(scanner.Bytes(), 0, &m)
			arr := make([]int, m)
			for j := 0; j < m; j++ {
				pos = readInt(scanner.Bytes(), pos+1, &arr[j])
			}
			quries[i] = arr
		}

		res := solve(n, segs, q, quries)

		for i := 0; i < q; i++ {
			fmt.Println(res[i])
		}
	}
}

var tree *SegTree

const SIZE = 5000000

func init() {
	tree = new(SegTree)
	tree.root = make([]int, SIZE)
	tree.left = make([]int, SIZE)
	tree.right = make([]int, SIZE)
	tree.count = make([]int, SIZE)
}

func solve(n int, itvs [][]int, q int, queries [][]int) []int {
	ps := make(Pairs, n)
	for i := 0; i < n; i++ {
		ps[i] = Pair{itvs[i][0], itvs[i][1]}
	}
	sort.Sort(ps)
	tree.index = 0

	tn := 1
	for tn <= n {
		tn <<= 1
	}

	for i, j := 1, 0; i <= n; i++ {
		tree.root[i] = tree.root[i-1]
		for j < n && ps[j].first <= i {
			tree.root[i] = tree.Add(tree.root[i], ps[j].second, tn)
			j++
		}
	}

	find := func(l1, r1, l2, r2 int) int {
		a := tree.Sum(tree.root[r1], r2, tn) - tree.Sum(tree.root[r1], l2, tn)
		b := tree.Sum(tree.root[l1], r2, tn) - tree.Sum(tree.root[l1], l2, tn)
		return a - b
	}

	caseOne := func(query []int) int {
		sort.Ints(query)
		var res int
		for i := 0; i < len(query); i++ {
			for j := i; j < len(query); j += 2 {
				a := 0
				if i > 0 {
					a = query[i-1]
				}
				b := query[i]
				c := query[j] - 1
				d := n
				if j < len(query)-1 {
					d = query[j+1] - 1
				}
				res += find(a, b, c, d)
			}
		}
		return res
	}

	sum := make([]int, n+1)

	caseTwo := func(query []int) int {
		for i := 1; i <= n; i++ {
			sum[i] = 0
		}
		for i := 0; i < len(query); i++ {
			sum[query[i]]++
		}
		for i := 1; i <= n; i++ {
			sum[i] += sum[i-1]
		}
		var res int
		for i := 0; i < len(ps); i++ {
			res += 1 & (sum[ps[i].second] ^ (sum[ps[i].first-1]))
		}
		return res
	}

	B := int(math.Sqrt(float64(n)/math.Log2(float64(n))) * 0.8)
	ans := make([]int, q)
	for i := 0; i < q; i++ {
		query := queries[i]
		if len(query) <= B {
			ans[i] = caseOne(query)
		} else {
			ans[i] = caseTwo(query)
		}
	}
	return ans
}

type SegTree struct {
	root  []int
	left  []int
	right []int
	count []int
	index int
}

func (tree *SegTree) Add(x int, p int, s int) int {
	var loop func(x int, p int, s int) int
	loop = func(x int, p int, s int) int {
		tree.index++
		u := tree.index
		tree.count[u] = tree.count[x] + 1
		tree.left[u] = tree.left[x]
		tree.right[u] = tree.right[x]
		if s > 1 {
			if p < (s >> 1) {
				tree.left[u] = loop(tree.left[u], p, s>>1)
			} else {
				tree.right[u] = loop(tree.right[u], p-(s>>1), s>>1)
			}
		}
		return u
	}
	return loop(x, p, s)
}

func (tree *SegTree) Sum(x int, p int, s int) int {
	var loop func(x int, p int, s int) int

	loop = func(x int, p int, s int) int {
		if s == 1 {
			return tree.count[x]
		}
		if p < (s >> 1) {
			return loop(tree.left[x], p, s>>1)
		}
		return tree.count[tree.left[x]] + loop(tree.right[x], p-(s>>1), s>>1)
	}
	return loop(x, p, s)
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
	return ps[i].first < ps[j].first
}

func (ps Pairs) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}
