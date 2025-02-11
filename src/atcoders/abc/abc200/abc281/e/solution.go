package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d ", x))
	}
	buf.WriteByte('\n')
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
	n, m, k := readThreeNums(reader)
	a := readNNums(reader, n)
	return solve(a, m, k)
}

func solve(a []int, m int, k int) []int {
	// 用avl tree也可以， 但是使用压缩后的segment tree也可以
	// 可能性能还好一些
	unique := compactAndSort(a)
	n := len(a)
	pos := make([]int, n)
	for i := 0; i < n; i++ {
		pos[i] = sort.SearchInts(unique, a[i])
	}
	tr := NewTree(len(unique))
	for i := 0; i < m; i++ {
		tr.Update(pos[i], a[i])
	}
	var res []int

	for i := m; i <= n; i++ {
		res = append(res, tr.GetFirstKSum(k))
		if i < n {
			tr.Update(pos[i], a[i])
		}
		tr.Update(pos[i-m], -a[i-m])
	}

	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func compactAndSort(arr []int) []int {
	res := make([]int, len(arr))
	copy(res, arr)
	sort.Ints(res)
	var p int
	for i := 1; i <= len(res); i++ {
		if i == len(res) || res[i] > res[i-1] {
			res[p] = res[i-1]
			p++
		}
	}
	return res[:p]
}

type Tree struct {
	cnt []int
	val []int
	sz  int
}

func NewTree(n int) *Tree {
	cnt := make([]int, 4*n)
	val := make([]int, 4*n)
	return &Tree{cnt, val, n}
}

func (tr *Tree) Update(p int, v int) {
	var loop func(i int, l int, r int)
	loop = func(i int, l int, r int) {
		if l == r {
			tr.val[i] += v
			if v > 0 {
				tr.cnt[i]++
			} else {
				tr.cnt[i]--
			}
			return
		}
		mid := (l + r) / 2
		if p <= mid {
			loop(2*i+1, l, mid)
		} else {
			loop(2*i+2, mid+1, r)
		}
		tr.cnt[i] = tr.cnt[2*i+1] + tr.cnt[2*i+2]
		tr.val[i] = tr.val[2*i+1] + tr.val[2*i+2]
	}
	loop(0, 0, tr.sz-1)
}

func (tr *Tree) GetFirstKSum(k int) int {
	var loop func(i int, l int, r int, k int) int
	loop = func(i int, l int, r int, k int) int {
		if l == r {
			if tr.cnt[i] == 0 {
				return 0
			}
			// 肯定能整除
			avg := tr.val[i] / tr.cnt[i]
			x := min(tr.cnt[i], k)
			return x * avg
		}
		mid := (l + r) / 2
		if tr.cnt[2*i+1] >= k {
			return loop(i*2+1, l, mid, k)
		}
		res := tr.val[2*i+1]
		k -= tr.cnt[2*i+1]
		return res + loop(2*i+2, mid+1, r, k)
	}
	return loop(0, 0, tr.sz-1, k)
}
