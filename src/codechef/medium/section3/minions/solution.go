package main

import (
	"sort"
	"bufio"
	"os"
	"fmt"
	"container/heap"
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

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var tmp int64
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp
	return i
}

func readNInt64s(scanner *bufio.Scanner, n int) []int64 {
	res := make([]int64, n)
	pos := 0
	scanner.Scan()
	for i := 0; i < n; i++ {
		for pos < len(scanner.Bytes()) && scanner.Bytes()[pos] == ' ' {
			pos++
		}
		pos = readInt64(scanner.Bytes(), pos, &res[i])
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for tc > 0 {
		n := readNum(scanner)
		A := make([]int64, n)
		B := make([]int64, n)
		for i := 0; i < n; i++ {
			tmp := readNInt64s(scanner, 2)
			A[i] = tmp[0]
			B[i] = tmp[1]
		}
		fmt.Println(solve(n, A, B))
		tc--
	}
}

func solve(n int, A []int64, B []int64) int {
	ps := make(Pairs, n)

	for i := 0; i < n; i++ {
		ps[i] = Pair{A[i], B[i]}
	}

	sort.Sort(sort.Reverse(ps))

	check := func(k int) bool {
		if k == 0 {
			return true
		}
		if k > n {
			return false
		}
		que := make(Q, 0, n)

		var sum int64
		for i := 0; i < k; i++ {
			heap.Push(&que, ps[i].second)
			sum += ps[i].second
		}

		if sum <= int64(k)*ps[k-1].first {
			return true
		}

		for i := k; i < n; i++ {
			//pop the largest
			heap.Push(&que, ps[i].second)
			sum += ps[i].second
			head := heap.Pop(&que).(int64)
			sum -= head
			if sum <= int64(k)*ps[i].first {
				return true
			}
		}
		return false
	}

	left, right, ans := 0, n+1, 0
	for left < right {
		mid := (left + right) >> 1
		if !check(mid) {
			right = mid
		} else {
			ans = mid
			left = mid + 1
		}
	}
	return ans
}

type Q []int64

func (q Q) Len() int {
	return len(q)
}

func (q Q) Less(i, j int) bool {
	// descreasing order
	return q[i] > q[j]
}

func (q Q) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q *Q) Push(item interface{}) {
	val := item.(int64)
	*q = append(*q, val)
}

func (q *Q) Pop() interface{} {
	old := *q
	n := len(old)
	res := old[n-1]
	*q = old[:n-1]
	return res
}

func solve1(n int, A []int64, B []int64) int {
	ps := make(Pairs, n)

	for i := 0; i < n; i++ {
		ps[i] = Pair{A[i], B[i]}
	}

	sort.Sort(ps)

	idx := make(Pairs2, n)
	for i := 0; i < n; i++ {
		idx[i] = Pair2{ps[i].second, i}
	}
	sort.Sort(idx)

	sz := 1
	for sz < n {
		sz <<= 1
	}

	t := make([]int64, 2*sz)
	c := make([]int, 2*sz)

	update := func(r int, val int64) {
		r += sz
		t[r] += val
		c[r] += 1
		for r > 1 {
			t[r>>1] = t[r] + t[r^1]
			c[r>>1] = c[r] + c[r^1]
			r >>= 1
		}
	}

	var query func(root int, val int64, cur int64, cnt int) int

	query = func(root int, val int64, cur int64, cnt int) int {
		if root >= sz {
			// a leaf
			if (t[root] + cur) <= int64(c[root]+cnt)*val {
				return c[root]
			}
			return 0
		} else if (t[root] + cur) <= int64(c[root]+cnt)*val {
			return c[root] + query(root|1, val, cur+t[root], cnt+c[root])
		}
		return query(root<<1, val, cur, cnt)
	}

	var ans int
	for i := n - 1; i >= 0; i-- {
		if ps[i].first >= ps[i].second {
			ans = max(ans, 1)
		}

		j := sort.Search(n, func(j int) bool {
			return idx[j].first > ps[i].second || (idx[j].first == ps[i].second && idx[j].second >= i)
		})
		// j has to be valid
		update(j, ps[i].second)

		if t[1] <= int64(c[1])*ps[i].first {
			ans = max(ans, c[1])
			continue
		}

		ans = max(ans, query(1, ps[i].first, 0, 0))
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
	first, second int64
}

type Pairs []Pair

func (this Pairs) Len() int {
	return len(this)
}

func (this Pairs) Less(i, j int) bool {
	return this[i].first < this[j].first
}

func (this Pairs) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

type Pair2 struct {
	first  int64
	second int
}

type Pairs2 []Pair2

func (this Pairs2) Len() int {
	return len(this)
}

func (this Pairs2) Less(i, j int) bool {
	return this[i].first < this[j].first || (this[i].first == this[j].first && this[i].second < this[j].second)
}

func (this Pairs2) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
