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
	cmds := make([]string, m)

	for i := 0; i < m; i++ {
		cmds[i] = readString(reader)
	}
	res := solve(n, cmds)
	var buf bytes.Buffer

	for _, cur := range res {
		buf.WriteString(fmt.Sprintf("%d\n", cur))
	}
	fmt.Print(buf.String())
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			return s[:i]
		}
	}
	return s
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

func solve(n int, cmds []string) []int {
	m := len(cmds)
	A := make([]int, m)
	B := make([]int, m)
	D := make([]int, m)
	X := make([]int, 2*m)
	var sz int

	dev := make([]int, 0, m)

	for i := 0; i < m; i++ {
		buf := []byte(cmds[i])
		if buf[0] == 'P' {
			pos := readInt(buf, 2, &A[i])
			readInt(buf, pos+1, &B[i])
			dev = append(dev, i)
			X[sz] = A[i]
			sz++
			X[sz] = B[i]
			sz++
		} else if buf[0] == 'M' {
			// m
			pos := readInt(buf, 2, &A[i])
			readInt(buf, pos+1, &B[i])
			A[i]--
			ind := dev[A[i]]
			D[ind] += B[i]
			X[sz] = A[ind] + D[ind]
			sz++
			X[sz] = B[ind] + D[ind]
			sz++
		} else {
			// query
			readInt(buf, 2, &A[i])
			X[sz] = A[i]
			sz++
		}
	}

	for i := 0; i < m; i++ {
		D[i] = 0
	}

	X = X[:sz]
	sort.Ints(X)
	sz = 0
	for i := 1; i <= len(X); i++ {
		if i == len(X) || X[i] > X[i-1] {
			X[sz] = X[i-1]
			sz++
		}
	}

	findIndex := func(x int) int {
		l, r := 0, sz
		for l < r {
			m := (l + r) / 2
			if X[m] >= x {
				r = m
			} else {
				l = m + 1
			}
		}
		return r
	}

	val := make([]int, sz+1)

	set := func(p int, v int) {
		p++
		for p <= sz {
			val[p] += v
			p += p & -p
		}
	}

	update := func(l, r, v int) {
		set(l, v)
		set(r, -v)
	}

	get := func(p int) int {
		p++
		var res int
		for p > 0 {
			res += val[p]
			p -= p & -p
		}
		return res
	}

	var res []int

	for i := 0; i < m; i++ {
		if cmds[i][0] == 'P' {
			update(findIndex(A[i]), findIndex(B[i])+1, 1)
		} else if cmds[i][0] == 'M' {
			ind := dev[A[i]]
			update(findIndex(A[ind]+D[ind]), findIndex(B[ind]+D[ind])+1, -1)
			D[ind] += B[i]
			update(findIndex(A[ind]+D[ind]), findIndex(B[ind]+D[ind])+1, 1)
		} else {
			res = append(res, get(findIndex(A[i])))
		}
	}

	return res
}

func solve4(n int, cmds []string) []int {
	m := len(cmds)
	A := make([]int, m)
	B := make([]int, m)
	D := make([]int, m)
	X := make([]int, 2*m)
	var sz int

	dev := make([]int, 0, m)

	for i := 0; i < m; i++ {
		buf := []byte(cmds[i])
		if buf[0] == 'P' {
			pos := readInt(buf, 2, &A[i])
			readInt(buf, pos+1, &B[i])
			dev = append(dev, i)
			X[sz] = A[i]
			sz++
			X[sz] = B[i]
			sz++
		} else if buf[0] == 'M' {
			// m
			pos := readInt(buf, 2, &A[i])
			readInt(buf, pos+1, &B[i])
			A[i]--
			ind := dev[A[i]]
			D[ind] += B[i]
			X[sz] = A[ind] + D[ind]
			sz++
			X[sz] = B[ind] + D[ind]
			sz++
		} else {
			// query
			readInt(buf, 2, &A[i])
			X[sz] = A[i]
			sz++
		}
	}

	for i := 0; i < m; i++ {
		D[i] = 0
	}

	X = X[:sz]
	sort.Ints(X)
	sz = 0
	for i := 1; i <= len(X); i++ {
		if i == len(X) || X[i] > X[i-1] {
			X[sz] = X[i-1]
			sz++
		}
	}

	findIndex := func(x int) int {
		l, r := 0, sz
		for l < r {
			m := (l + r) / 2
			if X[m] >= x {
				r = m
			} else {
				l = m + 1
			}
		}
		return r
	}

	val := make([]int, 2*sz)

	update := func(l, r, v int) {
		l += sz
		r += sz

		for l < r {
			if l&1 == 1 {
				val[l] += v
				l++
			}
			if r&1 == 1 {
				r--
				val[r] += v
			}
			l >>= 1
			r >>= 1
		}
	}

	query := func(p int) int {
		p += sz
		var res int
		for p > 0 {
			res += val[p]
			p >>= 1
		}
		return res
	}

	var res []int

	for i := 0; i < m; i++ {
		if cmds[i][0] == 'P' {
			update(findIndex(A[i]), findIndex(B[i])+1, 1)
		} else if cmds[i][0] == 'M' {
			ind := dev[A[i]]
			update(findIndex(A[ind]+D[ind]), findIndex(B[ind]+D[ind])+1, -1)
			D[ind] += B[i]
			update(findIndex(A[ind]+D[ind]), findIndex(B[ind]+D[ind])+1, 1)
		} else {
			res = append(res, query(findIndex(A[i])))
		}
	}

	return res
}

func solve3(n int, cmds []string) []int {
	m := len(cmds)
	A := make([]int, m)
	B := make([]int, m)
	D := make([]int, m)
	X := make([]int, 2*m)
	var sz int

	dev := make([]int, 0, m)

	for i := 0; i < m; i++ {
		buf := []byte(cmds[i])
		if buf[0] == 'P' {
			pos := readInt(buf, 2, &A[i])
			readInt(buf, pos+1, &B[i])
			dev = append(dev, i)
			X[sz] = A[i]
			sz++
			X[sz] = B[i]
			sz++
		} else if buf[0] == 'M' {
			// m
			pos := readInt(buf, 2, &A[i])
			readInt(buf, pos+1, &B[i])
			A[i]--
			ind := dev[A[i]]
			X[sz] = A[ind] + B[i]
			sz++
			X[sz] = B[ind] + B[i]
			sz++
		} else {
			// query
			readInt(buf, 2, &A[i])
			X[sz] = A[i]
			sz++
		}
	}

	for i := 0; i < m; i++ {
		D[i] = 0
	}

	X = X[:sz]
	sort.Ints(X)
	sz = 0
	for i := 1; i <= len(X); i++ {
		if i == len(X) || X[i] > X[i-1] {
			X[sz] = X[i-1]
			sz++
		}
	}

	findIndex := func(x int) int {
		l, r := 0, sz
		for l < r {
			m := (l + r) / 2
			if X[m] >= x {
				r = m
			} else {
				l = m + 1
			}
		}
		return r
	}

	var h int

	for 1<<uint(h) < sz {
		h++
	}

	left := make([]int, h*m)
	right := make([]int, h*m)
	cnt := make([]int, h*m)

	for i := 0; i < h*m; i++ {
		left[i] = -1
		right[i] = -1
	}
	var pos int
	newNode := func() int {
		pos++
		return pos
	}

	var update func(i int, L int, R int, l int, r int, v int)

	update = func(i int, L int, R int, l int, r int, v int) {
		if l > R || r < L {
			return
		}
		if l <= L && R <= r {
			cnt[i] += v
			return
		}
		M := (L + R) / 2

		if l <= M {
			if left[i] < 0 {
				left[i] = newNode()
			}
			update(left[i], L, M, l, r, v)
		}
		if r > M {
			if right[i] < 0 {
				right[i] = newNode()
			}
			update(right[i], M+1, R, l, r, v)
		}
	}

	var query func(i int, L int, R int, p int) int

	query = func(i int, L int, R int, p int) int {
		if L == R {
			return cnt[i]
		}
		ans := cnt[i]
		M := (L + R) / 2
		if p <= M && left[i] > 0 {
			ans += query(left[i], L, M, p)
		} else if p > M && right[i] > 0 {
			ans += query(right[i], M+1, R, p)
		}
		return ans
	}

	var res []int

	for i := 0; i < m; i++ {
		if cmds[i][0] == 'P' {
			l, r := findIndex(A[i]), findIndex(B[i])
			update(0, 0, sz-1, l, r, 1)
		} else if cmds[i][0] == 'M' {
			ind := dev[A[i]]
			update(0, 0, sz-1, findIndex(A[ind]+D[ind]), findIndex(B[ind]+D[ind]), -1)
			D[ind] += B[i]
			update(0, 0, sz-1, findIndex(A[ind]+D[ind]), findIndex(B[ind]+D[ind]), 1)
		} else {
			res = append(res, query(0, 0, sz-1, findIndex(A[i])))
		}
	}

	return res
}

func solve2(n int, cmds []string) []int {
	m := len(cmds)

	var h int

	for 1<<uint(h) < n {
		h++
	}

	left := make([]int, h*m)
	right := make([]int, h*m)
	cnt := make([]int, h*m)

	for i := 0; i < h*m; i++ {
		left[i] = -1
		right[i] = -1
	}
	var sz int
	newNode := func() int {
		sz++
		return sz
	}

	var update func(i int, L int, R int, l int, r int, v int)

	update = func(i int, L int, R int, l int, r int, v int) {
		if l > R || r < L {
			return
		}
		if l <= L && R <= r {
			cnt[i] += v
			return
		}
		M := (L + R) / 2

		if l <= M {
			if left[i] < 0 {
				left[i] = newNode()
			}
			update(left[i], L, M, l, r, v)
		}
		if r > M {
			if right[i] < 0 {
				right[i] = newNode()
			}
			update(right[i], M+1, R, l, r, v)
		}
	}

	var query func(i int, L int, R int, p int) int

	query = func(i int, L int, R int, p int) int {
		if L == R {
			return cnt[i]
		}
		ans := cnt[i]
		M := (L + R) / 2
		if p <= M && left[i] > 0 {
			ans += query(left[i], L, M, p)
		} else if p > M && right[i] > 0 {
			ans += query(right[i], M+1, R, p)
		}
		return ans
	}

	var res []int

	var devices []Device

	for _, cur := range cmds {
		buf := []byte(cur)
		if buf[0] == 'P' {
			var l, r int
			pos := readInt(buf, 2, &l)
			readInt(buf, pos+1, &r)
			update(0, 0, n, l, r, 1)
			devices = append(devices, Device{l, r})
		} else if buf[0] == 'M' {
			// m
			var i, d int
			pos := readInt(buf, 2, &i)
			readInt(buf, pos+1, &d)
			dev := devices[i-1]

			update(0, 0, n, dev.l, dev.r, -1)

			dev = Device{dev.l + d, dev.r + d}

			update(0, 0, n, dev.l, dev.r, 1)

			devices[i-1] = dev
		} else {
			// query
			var x int
			readInt(buf, 2, &x)
			y := query(0, 0, n, x)
			res = append(res, y)
		}
	}

	return res
}

func solve1(n int, cmds []string) []int {
	// range update, point query
	// n is to large, can't use BIT

	root := NewNode(0, n)

	var res []int

	var devices []Device

	for _, cur := range cmds {
		buf := []byte(cur)
		if buf[0] == 'P' {
			var l, r int
			pos := readInt(buf, 2, &l)
			readInt(buf, pos+1, &r)
			Update(root, n+1, l, r, 1)
			devices = append(devices, Device{l, r})
		} else if buf[0] == 'M' {
			// m
			var i, d int
			pos := readInt(buf, 2, &i)
			readInt(buf, pos+1, &d)
			dev := devices[i-1]

			Update(root, n+1, dev.l, dev.r, -1)

			dev = Device{max(0, dev.l+d), min(dev.r+d, n)}

			Update(root, n+1, dev.l, dev.r, 1)

			devices[i-1] = dev
		} else {
			// query
			var x int
			readInt(buf, 2, &x)
			y := Query(root, n+1, x)
			res = append(res, y)
		}
	}

	return res
}

type Node struct {
	left  *Node
	right *Node
	lazy  int
	L     int
	R     int
	cnt   int
}

func NewNode(L, R int) *Node {
	node := new(Node)
	node.L = L
	node.R = R
	node.lazy = 0
	node.cnt = 0
	return node
}

func (node *Node) pushDown() {
	if node.L == node.R {
		// no children
		return
	}
	if node.left == nil {
		m := (node.L + node.R) / 2
		node.left = NewNode(node.L, m)
		node.right = NewNode(m+1, node.R)
	}
	node.left.lazy += node.lazy
	node.left.cnt += node.lazy
	node.right.lazy += node.lazy
	node.right.cnt += node.lazy

	node.lazy = 0
}

func Update(root *Node, n int, L int, R int, v int) {

	var loop func(node *Node, l int, r int)

	loop = func(node *Node, l int, r int) {
		if R < l || r < L {
			// out of range
			return
		}

		if L <= l && r <= R {
			// total in the range
			node.lazy += v
			node.cnt += v
			return
		}

		node.pushDown()
		m := (l + r) / 2
		if L <= m {
			loop(node.left, l, m)
		}
		if m < R {
			loop(node.right, m+1, r)
		}
	}

	loop(root, 0, n-1)
}

func Query(root *Node, n int, p int) int {

	var loop func(node *Node, l int, r int) int

	loop = func(node *Node, l int, r int) int {
		if l == r {
			return node.cnt
		}
		node.pushDown()
		m := (l + r) / 2
		if p <= m {
			return loop(node.left, l, m)
		}
		return loop(node.right, m+1, r)
	}

	return loop(root, 0, n-1)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

type Device struct {
	l int
	r int
}
