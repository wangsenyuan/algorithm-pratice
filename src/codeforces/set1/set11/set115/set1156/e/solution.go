package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n := readNum(reader)
	a := readNNums(reader, n)
	res := solve(a)
	fmt.Println(res)
}

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(reader *bufio.Reader) (a int) {
	bs, _ := reader.ReadBytes('\n')
	readInt(bs, 0, &a)
	return
}

func readTwoNums(reader *bufio.Reader) (a int, b int) {
	res := readNNums(reader, 2)
	a, b = res[0], res[1]
	return
}

func readThreeNums(reader *bufio.Reader) (a int, b int, c int) {
	res := readNNums(reader, 3)
	a, b, c = res[0], res[1], res[2]
	return
}

func readNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, n)
	x := 0
	bs, _ := reader.ReadBytes('\n')
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func solve(a []int) int {
	n := len(a)
	pos := make([]int, n+1)

	tr := NewSegTree(n)

	for i := range n {
		pos[a[i]] = i
		tr.Update(i, a[i])
	}
	var ans int

	merge := func(l int, mid int, r int) {
		v := a[mid]
		if mid-l <= r-mid {
			// iterate left
			for i := l; i < mid; i++ {
				x := a[i]
				y := v - x
				if pos[y] > mid && pos[y] <= r {
					ans++
				}
			}
		} else {
			// iterate right
			for i := mid + 1; i <= r; i++ {
				x := a[i]
				y := v - x
				if pos[y] >= l && pos[y] < mid {
					ans++
				}
			}
		}
	}

	var dfs func(l int, r int)

	dfs = func(l int, r int) {
		if l == r {
			return
		}
		x := tr.Get(l, r+1)
		// x.first = 这个区间最大的值
		i := x.second
		if l < i {
			dfs(l, i-1)
		}
		if i < r {
			dfs(i+1, r)
		}
		merge(l, i, r)
	}

	dfs(0, n-1)

	return ans
}

type pair struct {
	first  int
	second int
}

func max_pair(a, b pair) pair {
	if a.first > b.first || a.first == b.first && a.second > b.second {
		return a
	}
	return b
}

type SegTree []pair

const inf = 1 << 60

func NewSegTree(n int) SegTree {
	arr := make([]pair, 2*n)
	for i := n; i < len(arr); i++ {
		arr[i] = pair{-inf, i - n}
	}

	for i := n - 1; i > 0; i-- {
		arr[i] = max_pair(arr[i*2], arr[i*2+1])
	}
	return arr
}

func (tr SegTree) Update(p int, v int) {
	n := len(tr) / 2
	p += n
	tr[p] = pair{v, p - n}

	for p > 1 {
		tr[p>>1] = max_pair(tr[p], tr[p^1])
		p >>= 1
	}
}

func (tr SegTree) Get(l int, r int) pair {
	n := len(tr) / 2
	l += n
	r += n
	res := pair{-inf, -1}
	for l < r {
		if l&1 == 1 {
			res = max_pair(res, tr[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = max_pair(res, tr[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}
