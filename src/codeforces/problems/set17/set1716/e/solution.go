package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var buf bytes.Buffer

	n := readNum(reader)
	A := readNNums(reader, 1<<n)
	q := readNum(reader)
	Q := make([]int, q)
	for i := 0; i < q; i++ {
		Q[i] = readNum(reader)
	}

	res := solve(n, A, Q)

	for _, ans := range res {
		buf.WriteString(fmt.Sprintf("%d\n", ans))
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

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
}

func solve(n int, A []int, Q []int) []int64 {
	tree := make([][]Node, 2<<n)

	var build func(v int, l int, r int)
	build = func(v int, l int, r int) {
		tree[v] = make([]Node, r-l)
		if l == r-1 {
			tree[v][0] = NewNode(A[l])
		} else {
			mid := (l + r) / 2
			build(2*v+1, l, mid)
			build(2*v+2, mid, r)
			for i := 0; i < mid-l; i++ {
				tree[v][i] = combine(tree[v*2+1][i], tree[v*2+2][i])
				tree[v][i+(mid-l)] = combine(tree[v*2+2][i], tree[v*2+1][i])
			}
		}
	}

	build(0, 0, len(A))

	ans := make([]int64, len(Q))
	var cur int

	for i, x := range Q {
		cur ^= (1 << x)
		ans[i] = tree[0][cur].ans
	}

	return ans
}

type Node struct {
	sum  int64
	pref int64
	suf  int64
	ans  int64
}

func NewNode(x int) Node {
	sum := int64(x)
	pref := max(sum, 0)
	suf := pref
	ans := pref
	return Node{sum, pref, suf, ans}
}

func combine(l, r Node) Node {
	sum := l.sum + r.sum
	pref := max(l.pref, l.sum+r.pref)
	suf := max(l.suf+r.sum, r.suf)
	ans := max(max(l.ans, r.ans), l.suf+r.pref)

	return Node{sum, pref, suf, ans}
}

func max(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}
