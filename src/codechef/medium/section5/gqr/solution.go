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
func main() {
	scanner := bufio.NewScanner(os.Stdin)

	n, q := readTwoNums(scanner)

	A := readNNums(scanner, n)

	// solver := NewSolver(n, A)

	// for q > 0 {
	// 	l, r := readTwoNums(scanner)
	// 	fmt.Println(solver.Query(l, r))
	// }
	queries := make([][]int, q)
	for i := 0; i < q; i++ {
		queries[i] = readNNums(scanner, 2)
	}

	res := solve(n, A, queries)

	for i := 0; i < q; i++ {
		fmt.Println(res[i])
	}
}

func solve(n int, A []int, queries [][]int) []int {
	ps := make([]Pair, len(queries))
	for i := 0; i < len(queries); i++ {
		ps[i] = Pair{queries[i][0], queries[i][1], i}
	}

	pre := make([]int, MAX_D)

	sort.Sort(PairSlice(ps))

	arr := make([]int, 2*n+2)

	update := func(pos int, v int) {
		pos += n

		if arr[pos] >= v {
			return
		}

		arr[pos] = v

		for pos > 0 {
			arr[pos>>1] = max(arr[pos], arr[pos^1])
			pos >>= 1
		}
	}

	get := func(left, right int) int {
		left += n
		right += n
		res := 0

		for left < right {
			if left&1 == 1 {
				res = max(res, arr[left])
				left++
			}
			if right&1 == 1 {
				right--
				res = max(res, arr[right])
			}
			left >>= 1
			right >>= 1
		}
		return res
	}

	ans := make([]int, len(ps))

	var j int

	for i := 1; i <= n; i++ {
		for d := 1; d*d <= A[i-1]; d++ {
			if A[i-1]%d != 0 {
				continue
			}

			if pre[d] > 0 {
				update(pre[d]-1, d)
			}
			pre[d] = i
			d2 := A[i-1] / d
			if d2 == d {
				continue
			}
			if pre[d2] > 0 {
				update(pre[d2]-1, d2)
			}
			pre[d2] = i
		}

		for j < len(ps) && ps[j].second == i {
			ans[ps[j].index] = get(ps[j].first-1, ps[j].second)
			j++
		}
	}
	return ans
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

type Pair struct {
	first  int
	second int
	index  int
}

type PairSlice []Pair

func (this PairSlice) Len() int {
	return len(this)
}

func (this PairSlice) Less(i, j int) bool {
	return this[i].second < this[j].second
}

func (this PairSlice) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

type Solver struct {
	t [][]int
}

const MAX_D = 100000007

func NewSolver(n int, A []int) Solver {
	t := make([][]int, n)
	for i := 0; i < n; i++ {
		t[i] = make([]int, n)
	}

	pre := make([]int, MAX_D)

	for i := 0; i < MAX_D; i++ {
		pre[i] = -1
	}

	for i := 0; i < n; i++ {
		a := A[i]
		for d := 1; d*d <= a; d++ {
			if a%d != 0 {
				continue
			}
			if pre[d] >= 0 {
				relax(&t[pre[d]][i], d)
			}
			pre[d] = i
			if d*d < a {
				if pre[a/d] >= 0 {
					relax(&t[pre[a/d]][i], a/d)
				}
				pre[a/d] = i
			}
		}
	}

	for l := 2; l <= n; l++ {
		for i := 0; i+l-1 < n; i++ {
			j := i + l - 1
			relax(&t[i][j], t[i][j-1])
			relax(&t[i][j], t[i+1][j])
		}
	}

	return Solver{t}
}

func (solver Solver) Query(l, r int) int {
	l--
	r--
	return solver.t[l][r]
}

func relax(a *int, b int) {
	if *a < b {
		*a = b
	}
}
