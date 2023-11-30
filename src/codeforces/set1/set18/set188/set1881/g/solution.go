package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n, k := readTwoNums(reader)
		x := readString(reader)[:n]
		Q := make([][]int, k)
		for i := 0; i < k; i++ {
			s, _ := reader.ReadBytes('\n')
			var tp int
			pos := readInt(s, 0, &tp) + 1
			if tp == 1 {
				Q[i] = make([]int, 4)
			} else {
				Q[i] = make([]int, 3)
			}
			Q[i][0] = tp

			for j := 1; j < len(Q[i]); j++ {
				pos = readInt(s, pos, &Q[i][j]) + 1
			}
		}
		res := solve(x, Q)
		for _, b := range res {
			if b {
				buf.WriteString("YES\n")
			} else {
				buf.WriteString("NO\n")
			}
		}
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

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')

	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
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

func solve(s string, Q [][]int) []bool {
	// 1, l, r, x 整个区间 +x
	// 2 l..r 就是判断在l...r中间，是否有2个或者3个连续的字符组成回文
	// range update + range query
	n := len(s)
	nodes := make([]*Node, 4*n)

	pullUp := func(i int, l int, r int) {
		j := 2*i + 1
		k := j + 1

		// 最多两个
		nodes[i].first = nodes[j].first
		nodes[i].last = nodes[k].last

		if len(nodes[i].first) == 1 {
			nodes[i].first += nodes[k].first[:1]
		}
		if len(nodes[i].last) == 1 {
			ln := len(nodes[j].last)
			nodes[i].last = nodes[j].last[ln-1:] + nodes[k].first
		}

		if nodes[2*i+1].pal || nodes[2*i+2].pal {
			nodes[i].pal = true
		} else {
			x := nodes[j].last + nodes[k].first
			nodes[i].pal = checkPalindrome(x)
		}

	}

	var build func(i int, l int, r int)

	build = func(i int, l int, r int) {
		nodes[i] = new(Node)
		if l == r {
			nodes[i].pal = false
			nodes[i].first = s[l : l+1]
			nodes[i].last = s[l : l+1]
			return
		}
		mid := (l + r) / 2
		build(2*i+1, l, mid)
		build(2*i+2, mid+1, r)
		pullUp(i, l, r)
	}

	changeByte := func(d byte, x int) string {
		c := int(d - 'a')
		c = (c + x) % 26
		return fmt.Sprintf("%c", c+'a')
	}

	change := func(num string, x int) string {
		if len(num) == 1 {
			return changeByte(num[0], x)
		}
		return changeByte(num[0], x) + changeByte(num[1], x)
	}

	build(0, 0, n-1)

	pushLazy := func(i int, x int) {
		nodes[i].first = change(nodes[i].first, x)
		nodes[i].last = change(nodes[i].last, x)
		nodes[i].lazy += x
	}

	push := func(i int) {
		if nodes[i].lazy != 0 {
			pushLazy(i*2+1, nodes[i].lazy)
			pushLazy(i*2+2, nodes[i].lazy)
			nodes[i].lazy = 0
		}
	}

	var update func(i int, l int, r int, L int, R int, v int)

	update = func(i int, l int, r int, L int, R int, v int) {
		if R < L || R < l || r < L {
			return
		}
		if L <= l && r <= R {
			pushLazy(i, v)
			return
		}
		push(i)
		mid := (l + r) / 2
		update(2*i+1, l, mid, L, min(R, mid), v)
		update(2*i+2, mid+1, r, max(mid+1, L), R, v)
		pullUp(i, l, r)
	}

	var query func(i int, l int, r int, L int, R int) bool

	query = func(i int, l int, r int, L int, R int) bool {
		if R < L || R < l || r < L {
			return false
		}
		if L == l && r == R {
			return nodes[i].pal
		}
		push(i)
		mid := (l + r) / 2
		a := query(2*i+1, l, mid, L, min(R, mid))
		if a {
			return true
		}
		b := query(2*i+2, mid+1, r, max(mid+1, L), R)

		if b {
			return true
		}
		j := 2*i + 1
		k := j + 1
		if L <= mid && mid+1 <= R {
			lj := len(nodes[j].last)
			if nodes[j].last[lj-1] == nodes[k].first[0] {
				return true
			}
		}

		if L <= mid-1 && mid+1 <= R {
			if nodes[j].last[0] == nodes[k].first[0] {
				return true
			}
		}

		if L <= mid && mid+2 <= R {
			lj := len(nodes[j].last)
			if nodes[j].last[lj-1] == nodes[k].first[1] {
				return true
			}
		}
		return false
	}

	var res []bool

	for _, cur := range Q {
		if cur[0] == 1 {
			// update
			l, r, x := cur[1], cur[2], cur[3]
			l--
			r--
			update(0, 0, n-1, l, r, x%26)
		} else {
			l, r := cur[1], cur[2]
			l--
			r--
			if l == r {
				res = append(res, true)
				continue
			}
			res = append(res, !query(0, 0, n-1, l, r))
		}
	}

	return res
}

func checkPalindrome(s string) bool {
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			return true
		}
		if i-2 >= 0 && s[i] == s[i-2] {
			return true
		}
	}
	return false
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
func max(a, b int) int {
	return a + b - min(a, b)
}

type Node struct {
	lazy  int
	first string
	last  string
	pal   bool
}
