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
	s := readString(reader)
	n := readNum(reader)
	queries := make([]string, n)
	for i := range n {
		queries[i] = readString(reader)
	}
	return solve(s, queries)
}

func solve(s string, queries []string) []int {

	n := len(s)
	pos := make([]*SegTree, 26)

	for i := 0; i < 26; i++ {
		pos[i] = NewSegTree(n)
	}

	buf := []byte(s)

	for i := range n {
		x := int(buf[i] - 'a')
		pos[x].Update(i, i)
	}

	get := func(l int, r int) int {
		l--
		var cnt int
		for i := 0; i < 26; i++ {
			j := pos[i].Get(0, r)
			if j >= l {
				cnt++
			}
		}
		return cnt
	}

	change := func(p int, x byte) {
		p--
		if buf[p] == x {
			return
		}
		i := int(buf[p] - 'a')
		pos[i].Update(p, -1)
		buf[p] = x
		i = int(buf[p] - 'a')
		pos[i].Update(p, p)
	}

	var ans []int

	for _, cur := range queries {
		if cur[0] == '2' {
			var l, r int
			pos := readInt([]byte(cur), 2, &l)
			readInt([]byte(cur), pos+1, &r)
			ans = append(ans, get(l, r))
		} else {
			var l int
			pos := readInt([]byte(cur), 2, &l)
			x := cur[pos+1]
			change(l, x)
		}
	}

	return ans
}

type SegTree struct {
	arr []int
	sz  int
}

func NewSegTree(n int) *SegTree {
	arr := make([]int, 2*n)
	for i := 0; i < len(arr); i++ {
		arr[i] = -1
	}
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
	res := -1
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
