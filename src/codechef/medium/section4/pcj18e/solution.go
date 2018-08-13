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

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	sign := int64(1)
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := int64(0)
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for tc > 0 {
		tc--

		N := readNum(scanner)
		A := readNNums(scanner, N)

		fmt.Println(solve(N, A))
	}
}

var ps PS

func init() {
	ps = make(PS, 100000)
}

func solve(N int, A []int) int {
	// ps := make(PS, N)
	for i := 0; i < N; i++ {
		ps[i] = Pair{A[i], i}
	}

	sort.Sort(ps[0:N])
	j := ps[0].second
	for i := 1; i < N; i++ {
		if ps[i].second < j {
			return N - i
		}
		j = max(j, ps[i].second)
	}
	return 0
}

func solve1(N int, A []int) int {
	ps := make(PS, N)
	for i := 0; i < N; i++ {
		ps[i] = Pair{A[i], i}
	}

	sort.Sort(ps)
	st := NewSegTree(N)
	for i := 0; i < N; i++ {
		p := ps[i]
		j := p.second
		x := st.GetMax(j+1, N-1)
		if x > 0 {
			return N - i
		}
		st.Update(j, p.first)
	}
	return 0
}

type Pair struct {
	first, second int
}

type PS []Pair

func (ps PS) Len() int {
	return len(ps)
}

func (ps PS) Less(i, j int) bool {
	return ps[i].first < ps[j].first
}

func (ps PS) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

type SegTree struct {
	arr  []int
	size int
}

func NewSegTree(n int) SegTree {
	arr := make([]int, 4*n)
	return SegTree{arr, n}
}

func (st *SegTree) Update(pos int, val int) {
	arr := st.arr
	var loop func(i int, start, end int)
	loop = func(i int, start, end int) {
		if start == end {
			arr[i] = val
			return
		}
		mid := (start + end) >> 1
		if pos <= mid {
			loop(2*i+1, start, mid)
		} else {
			loop(2*i+2, mid+1, end)
		}
		arr[i] = max(arr[2*i+1], arr[2*i+2])
	}
	loop(0, 0, st.size-1)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func (st SegTree) GetMax(left int, right int) int {
	arr := st.arr

	var loop func(i int, start, end int) int

	loop = func(i int, start, end int) int {
		if right < start || end < left {
			return 0
		}
		if left <= start && end <= right {
			return arr[i]
		}
		mid := (start + end) >> 1
		a := loop(2*i+1, start, mid)
		b := loop(2*i+2, mid+1, end)
		return max(a, b)
	}
	return loop(0, 0, st.size-1)
}
