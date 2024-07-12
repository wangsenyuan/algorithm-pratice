package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, m := readTwoNums(reader)
	segs := make([][]int, n)
	for i := 0; i < n; i++ {
		segs[i] = readNNums(reader, 3)
	}
	res := solve(n, m, segs)
	fmt.Println(res)
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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

func solve(n int, m int, segments [][]int) int {
	sort.Slice(segments, func(i, j int) bool {
		return segments[i][2] < segments[j][2]
	})

	tr := BuildTree(m - 1)

	var best = 1 << 30

	for i, j := 0, 0; j < n; j++ {
		l, r, w := segments[j][0], segments[j][1], segments[j][2]
		l--
		r -= 2
		tr.Add(l, r, w)
		for i <= j && tr.nodes[0].val > 0 {
			li, ri, wi := segments[i][0], segments[i][1], segments[i][2]
			li--
			ri -= 2
			tr.Add(li, ri, -wi)
			if tr.nodes[0].val == 0 {
				// unconnect
				tr.Add(li, ri, wi)
				break
			}
			i++
		}
		if tr.nodes[0].val > 0 {
			best = min(best, w-segments[i][2])
		}
	}

	return best
}

type node struct {
	val  int
	lazy int
}

type Tree struct {
	nodes []*node
	n     int
}

func BuildTree(n int) *Tree {
	arr := make([]*node, 4*n)
	for i := 0; i < 4*n; i++ {
		arr[i] = new(node)
		arr[i].val = 0
		arr[i].lazy = 0
	}

	return &Tree{arr, n}
}

func (tr *Tree) pushValue(i int, v int) {
	// val 是最小值，全部刚更新的时候，最小值也更新了
	tr.nodes[i].val += v
	tr.nodes[i].lazy += v
}

func (tr *Tree) push(i int, l int, r int) {
	if l < r && tr.nodes[i].lazy != 0 {
		tr.pushValue(i*2+1, tr.nodes[i].lazy)
		tr.pushValue(i*2+2, tr.nodes[i].lazy)
		tr.nodes[i].lazy = 0
	}
}

func (tr *Tree) pull(i int) {
	tr.nodes[i].val = min(tr.nodes[i*2+1].val, tr.nodes[i*2+2].val)
}

func (tr *Tree) Add(L int, R int, v int) {
	var loop func(i int, l int, r int)
	loop = func(i int, l int, r int) {
		// out of bounds
		if R < l || r < L {
			return
		}
		tr.push(i, l, r)
		if L <= l && r <= R {
			tr.pushValue(i, v)
			return
		}
		mid := (l + r) / 2
		loop(2*i+1, l, mid)
		loop(2*i+2, mid+1, r)
		tr.pull(i)
	}

	loop(0, 0, tr.n-1)
}
