package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
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

func process(reader *bufio.Reader) int {
	n, k := readTwoNums(reader)
	a := readNNums(reader, n)
	return solve(n, k, a)
}

func solve(n int, k int, a []int) int {
	b := sortAndUnique(a)

	tr := NewSegTree(n + 1)

	freq := make([]int, len(b))
	freq2 := make([]int, n+1)
	var score int

	add := func(num int) {
		j := sort.SearchInts(b, num)
		// x 是num出现的次数
		x := freq[j]
		if freq2[x] == 1 {
			tr.Update(x, 0)
		}
		freq2[x]--
		freq[j]++
		x++
		freq2[x]++
		if freq2[x] == 1 {
			tr.Update(x, x)
		}

		score = tr.Get(1, n+1)
	}

	rem := func(num int) {
		j := sort.SearchInts(b, num)
		x := freq[j]
		if freq2[x] == 1 {
			tr.Update(x, 0)
		}
		freq2[x]--
		freq[j]--
		x--
		freq2[x]++
		if freq2[x] == 1 {
			tr.Update(x, x)
		}
		score = tr.Get(1, n+1)
	}

	var m int
	for l, r := 0, 0; l < n; l++ {
		for score < k && r < n {
			add(a[r])
			r++
		}
		if score == k {
			m += (n - r + 1)
		}
		rem(a[l])
	}
	return m
}

func sortAndUnique(arr []int) []int {
	res := make([]int, len(arr))
	copy(res, arr)
	sort.Ints(res)
	var n int
	for i := 1; i <= len(res); i++ {
		if i == len(res) || res[i-1] < res[i] {
			res[n] = res[i-1]
			n++
		}
	}
	return res[:n]
}

type SegTree struct {
	arr []int
	sz  int
}

func NewSegTree(n int) *SegTree {
	arr := make([]int, 2*n)
	return &SegTree{arr, n}
}

func (t *SegTree) Update(p int, v int) {
	p += t.sz
	t.arr[p] = v
	for p > 1 {
		t.arr[p>>1] = max(t.arr[p], t.arr[p^1])
		p >>= 1
	}
}

func (t *SegTree) Get(l int, r int) int {
	var res int
	l += t.sz
	r += t.sz

	for l < r {
		if l&1 == 1 {
			res = max(res, t.arr[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = max(res, t.arr[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}
