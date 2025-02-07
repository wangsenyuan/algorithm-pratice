package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d\n", x))
	}
	buf.WriteTo(os.Stdout)
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

func process(reader *bufio.Reader) []int {
	n, m, q := readThreeNums(reader)
	a := readNNums(reader, n)
	qs := make([]int, q)
	for i := range q {
		qs[i] = readNum(reader)
	}
	return solve(m, a, qs)
}

type pair struct {
	first  int
	second int
}

const inf = 1e12

func solve(m int, a []int, q []int) []int {
	freq := make([]int, m)
	for _, i := range a {
		freq[i-1]++
	}

	arr := make([]pair, m)
	for i := range m {
		arr[i] = pair{freq[i], i}
	}

	slices.SortFunc(arr, func(x, y pair) int {
		if x.first != y.first {
			return x.first - y.first
		}
		return x.second - y.second
	})
	n := len(a)

	qs := make([]pair, len(q))
	for i, x := range q {
		qs[i] = pair{x, i}
	}

	slices.SortFunc(qs, func(x, y pair) int {
		return x.first - y.first
	})

	set := NewSegTree(m)
	// 大体的思路是对的，但是细节有点麻烦
	ans := make([]int, len(qs))
	var sum int
	var j int
	arr = append(arr, pair{0, 0})
	for i := 0; i < m; {
		x := arr[i].first
		for i < m && arr[i].first == x {
			sum += x
			set.Set(arr[i].second)
			i++
		}
		// 计算它下次轮到cur.second的次数
		cnt := i*arr[i].first - sum
		// 小于等于 cnt 的，肯定是在i个城市举行
		for j < len(qs) && (i == m || qs[j].first-n <= cnt) {
			tmp := qs[j].first - n - (i*x - sum)
			// 前面的需要按照下标重新排序
			// tmp > 0是一定成立的
			// 且 tmp <= i + 1 肯定是成立的
			tmp %= i
			if tmp == 0 {
				tmp = i
			}
			u := set.Get(tmp)
			ans[qs[j].second] = u + 1
			j++
		}
	}

	return ans
}

type SegTree struct {
	arr []int
	sz  int
}

func NewSegTree(n int) *SegTree {
	arr := make([]int, 4*n)

	return &SegTree{arr, n}
}

func (t *SegTree) Set(p int) {
	var loop func(i int, l int, r int)
	loop = func(i int, l int, r int) {
		if l == r {
			t.arr[i]++
			return
		}
		mid := (l + r) / 2
		if p <= mid {
			loop(2*i+1, l, mid)
		} else {
			loop(2*i+2, mid+1, r)
		}
		t.arr[i] = t.arr[2*i+1] + t.arr[2*i+2]
	}
	loop(0, 0, t.sz-1)
}

func (t *SegTree) Get(k int) int {
	var loop func(i int, l int, r int, k int) int
	loop = func(i int, l int, r int, k int) int {
		if l == r {
			return l
		}
		mid := (l + r) / 2
		if t.arr[i*2+1] >= k {
			return loop(2*i+1, l, mid, k)
		}
		return loop(2*i+2, mid+1, r, k-t.arr[2*i+1])
	}
	return loop(0, 0, t.sz-1, k)
}
