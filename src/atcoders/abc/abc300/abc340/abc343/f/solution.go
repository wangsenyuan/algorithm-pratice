package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
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
	n, m := readTwoNums(reader)
	a := readNNums(reader, n)
	queries := make([][]int, m)
	for i := 0; i < m; i++ {
		queries[i] = readNNums(reader, 3)
	}
	return solve(a, queries)
}

func solve(a []int, queries [][]int) []int {
	tr := NewTree(a)

	var ans []int

	for _, cur := range queries {
		if cur[0] == 1 {
			p, x := cur[1], cur[2]
			p--
			tr.Update(p, x)
		} else {
			l, r := cur[1], cur[2]
			l--
			r--
			tmp := tr.Get(l, r)
			ans = append(ans, tmp[1].second)
		}
	}
	return ans
}

type pair struct {
	first  int
	second int
}

type Tree struct {
	val [][2]pair
	sz  int
}

func NewTree(arr []int) *Tree {
	n := len(arr)

	val := make([][2]pair, 4*n)

	var build func(i int, l int, r int)
	build = func(i int, l int, r int) {
		if l == r {
			val[i] = [2]pair{{arr[l], 1}, {-1, 0}}
			return
		}
		mid := (l + r) / 2
		build(2*i+1, l, mid)
		build(2*i+2, mid+1, r)
		val[i] = merge(val[2*i+1], val[2*i+2])
	}

	build(0, 0, n-1)
	tr := new(Tree)
	tr.sz = n
	tr.val = val
	return tr
}

func merge(a [2]pair, b [2]pair) [2]pair {
	var c [2]pair
	c[0] = a[0]
	c[1] = a[1]
	for i := 0; i < 2; i++ {
		if b[i].first > c[0].first {
			c[1] = c[0]
			c[0] = b[i]
		} else if b[i].first == c[0].first {
			c[0].second += b[i].second
		} else if b[i].first > c[1].first {
			c[1] = b[i]
		} else if b[i].first == c[1].first {
			c[1].second += b[i].second
		}
	}
	return c
}

func (tr *Tree) Update(p int, v int) {
	var loop func(i int, l int, r int)
	loop = func(i int, l int, r int) {
		if l == r {
			tr.val[i][0].first = v
			return
		}
		mid := (l + r) / 2
		if p <= mid {
			loop(2*i+1, l, mid)
		} else {
			loop(2*i+2, mid+1, r)
		}
		tr.val[i] = merge(tr.val[2*i+1], tr.val[2*i+2])
	}
	loop(0, 0, tr.sz-1)
}

func (tr *Tree) Get(L int, R int) [2]pair {
	var loop func(i int, l int, r int) [2]pair
	loop = func(i int, l int, r int) [2]pair {
		if R < l || r < L {
			return [2]pair{{-1, 0}, {-1, 0}}
		}
		if L <= l && r <= R {
			return tr.val[i]
		}
		mid := (l + r) / 2
		a := loop(2*i+1, l, mid)
		b := loop(2*i+2, mid+1, r)
		return merge(a, b)
	}
	return loop(0, 0, tr.sz-1)
}
