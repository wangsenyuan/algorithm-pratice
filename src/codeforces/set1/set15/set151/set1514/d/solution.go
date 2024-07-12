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
	n, m := readTwoNums(reader)
	a := readNNums(reader, n)
	queries := make([][]int, m)
	for i := 0; i < m; i++ {
		queries[i] = readNNums(reader, 2)
	}
	res := solve(a, queries)
	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d\n", x))
	}

	fmt.Print(buf.String())
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


func solve(a []int, queries [][]int) []int {
	n := len(a)
	pos := make([][]int, n+1)
	for i, num := range a {
		pos[num] = append(pos[num], i)
	}

	tr := Build(a)
	ans := make([]int, len(queries))

	for i, cur := range queries {
		l, r := cur[0] - 1, cur[1] - 1
		// v is the value most frequent
		v, _ := tr.Get(l, r)
		x := sort.SearchInts(pos[v], l)
		y := sort.SearchInts(pos[v], r + 1)

		c := y - x

		if c <= (r - l + 2) / 2 {
			ans[i] = 1
		} else {
			ans[i] = 2 *c - (r - l + 1)
		}
	}

	return ans
}


type node struct {
	val int
	cnt int
}

type tree struct {
	nodes []*node
	sz int
}

func Build(arr []int) *tree {
	n := len(arr)
	nodes := make([]*node, 4 * n)

	var loop func(i int, l int, r int) 
	loop = func(i int, l int, r int) {
		nodes[i] = new(node)
		if l ==r  {
			nodes[i].val = arr[l]
			nodes[i].cnt = 1
			return
		}
		mid := (l + r) / 2
		loop(i * 2 + 1, l, mid)
		loop(i * 2 + 2, mid + 1, r)
		if nodes[2 * i +1].val == nodes[2 * i + 2].val {
			nodes[i].val = nodes[ 2 * i +1].val
			nodes[i].cnt = nodes[2 * i + 1].cnt + nodes[2 * i + 2].cnt
		} else if nodes[2 * i + 1].cnt >= nodes[2 * i + 2].cnt {
			nodes[i].val  = nodes[ 2 * i +1].val
			nodes[i].cnt = nodes[2 * i + 1].cnt - nodes[i * 2 + 2].cnt
		} else {
			nodes[i].val = nodes[2 * i + 2].val
			nodes[i].cnt = nodes[2 * i + 2].cnt - nodes[i * 2 + 1].cnt
		}
	}

	loop(0, 0, n - 1)
	return &tree{nodes, n}
}

func (tr *tree) Get(L int, R int) (int, int) {
	var loop func(i int, l int, r int) (val int, cnt int) 

	loop = func(i int, l int, r int) (val int, cnt int) {
		if L <= l && r <= R {
			return tr.nodes[i].val, tr.nodes[i].cnt
		}
		mid := (l + r) / 2
		if R <= mid {
			return loop(2 * i +1, l, mid)
		}
		if mid < L {
			return loop(2 * i + 2, mid + 1, r)
		}
		av, ac := loop(2 * i + 1, l, mid)
		bv, bc := loop(2 * i + 2, mid + 1, r)

		if av == bv {
			return av, ac + bc
		}
		if ac >= bc {
			return av, ac - bc
		}
		return bv, bc - ac
	}

	return loop(0, 0, tr.sz - 1)
}