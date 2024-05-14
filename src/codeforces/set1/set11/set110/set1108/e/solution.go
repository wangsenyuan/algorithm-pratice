package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, m := readTwoNums(reader)
	a := readNNums(reader, n)
	segs := make([][]int, m)
	for i := 0; i < m; i++ {
		segs[i] = readNNums(reader, 2)
	}

	best, res := solve(a, segs)
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n%d\n", best, len(res)))
	for i := 0; i < len(res); i++ {
		buf.WriteString(fmt.Sprintf("%d ", res[i]))
	}
	buf.WriteByte('\n')
	fmt.Print(buf.String())
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
}

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
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

func solve(a []int, segs [][]int) (best int, use []int) {

	tr := Build(a)

	n := len(a)

	mn := tr.Get(0, n-1)

	pivot := 0

	for i, num := range a {
		if best < num-mn {
			pivot = i
			best = num - mn
		}
	}
	beginAt := make([][]int, n)
	endAt := make([][]int, n)
	for i, cur := range segs {
		l, r := cur[0]-1, cur[1]-1
		beginAt[l] = append(beginAt[l], i)
		endAt[r] = append(endAt[r], i)
	}
	// i前面的最小值
	pref := make([]int, n)
	for i := 0; i < n; i++ {
		if i > 0 {
			pref[i] = tr.Get(0, i-1)
		} else {
			pref[i] = oo
		}

		for _, j := range endAt[i] {
			l, r := segs[j][0]-1, segs[j][1]-1
			tr.Update(l, r, -1)
		}
	}

	tr = Build(a)

	use_empty := true

	for i := n - 1; i >= 0; i-- {
		tmp := pref[i]
		if i+1 < n {
			tmp = min(tmp, tr.Get(i+1, n-1))
		}
		if a[i]-tmp > best {
			best = a[i] - tmp
			use_empty = false
			pivot = i
		}

		for _, j := range beginAt[i] {
			l, r := segs[j][0]-1, segs[j][1]-1
			tr.Update(l, r, -1)
		}
	}

	if use_empty {
		return
	}

	// not use_all
	for i, cur := range segs {
		l, r := cur[0]-1, cur[1]-1
		if r < pivot || l > pivot {
			use = append(use, i+1)
		}
	}

	return
}

type Node struct {
	val  []int
	lazy []int
	sz   int
}

func Build(arr []int) *Node {
	n := len(arr)
	val := make([]int, n*4)
	lazy := make([]int, n*4)

	var build func(i int, l int, r int)

	build = func(i int, l int, r int) {
		if l == r {
			val[i] = arr[l]
			return
		}
		mid := (l + r) / 2
		build(i*2, l, mid)
		build(i*2+1, mid+1, r)
		val[i] = min(val[i*2], val[i*2+1])
	}

	build(1, 0, n-1)

	return &Node{val, lazy, n}
}

func (node *Node) pushValue(i int, v int) {
	node.val[i] += v
	node.lazy[i] += v
}

func (node *Node) push(i int, l int, r int) {
	if node.lazy[i] != 0 && l < r {
		node.pushValue(i*2, node.lazy[i])
		node.pushValue(i*2+1, node.lazy[i])
		node.lazy[i] = 0
	}
}

func (root *Node) Update(L int, R int, v int) {
	var loop func(i int, l int, r int)
	loop = func(i int, l int, r int) {
		if r < L || R < l {
			return
		}
		root.push(i, l, r)
		if L <= l && r <= R {
			root.pushValue(i, v)
			return
		}
		mid := (l + r) / 2
		loop(i*2, l, mid)
		loop(i*2+1, mid+1, r)
		root.val[i] = min(root.val[i*2], root.val[i*2+1])
	}

	loop(1, 0, root.sz-1)
}

const oo = 1 << 50

func (root *Node) Get(L int, R int) int {
	var loop func(i int, l int, r int) int

	loop = func(i int, l int, r int) int {
		if r < L || R < l {
			return oo
		}
		if L <= l && r <= R {
			return root.val[i]
		}
		root.push(i, l, r)
		mid := (l + r) / 2
		a := loop(i*2, l, mid)
		b := loop(i*2+1, mid+1, r)
		return min(a, b)
	}

	return loop(1, 0, root.sz-1)
}
