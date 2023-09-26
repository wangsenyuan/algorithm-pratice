package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	ss := make([][]int, n)
	for i := 0; i < n; i++ {
		ss[i] = readNNums(reader, 2)
	}
	res := solve(ss)
	fmt.Println(res)
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

const mod = 998244353

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func mul(a, b int) int {
	return int(int64(a) * int64(b) % mod)
}

const X = 300000

func solve(ss [][]int) int {
	// a or b => 在a中或者在b中
	// a and   b => 在a中且在b中
	// a xor b => 在a中，但不在b中，或者在b中但不在a中
	// 完全没想法～
	// r <= 3e5
	// 考虑贡献，应该是对的，但是怎么计算贡献，是个大问题
	n := len(ss)

	t := make([]Mat, 4*n)

	pullUp := func(v int) {
		t[v] = t[v*2+1].Mul(t[v*2+2])
	}
	var build func(v int, l int, r int)
	build = func(v int, l int, r int) {
		if l == r-1 {
			t[v] = ZERO
			return
		}
		mid := (l + r) / 2
		build(2*v+1, l, mid)
		build(2*v+2, mid, r)
		pullUp(v)
	}

	var update func(v int, l int, r int, pos int, val int)
	update = func(v int, l int, r int, pos int, val int) {
		if l == r-1 {
			if val == 0 {
				t[v] = ZERO
			} else {
				t[v] = ONE
			}
			return
		}
		mid := (l + r) / 2
		if pos < mid {
			update(v*2+1, l, mid, pos, val)
		} else {
			update(v*2+2, mid, r, pos, val)
		}
		pullUp(v)
	}
	V := make([][]Pair, X+10)

	addEvent := func(p int, x int, y int) {
		if len(V[p]) == 0 {
			V[p] = make([]Pair, 0, 1)
		}
		V[p] = append(V[p], Pair{x, y})
	}

	for i, cur := range ss {
		l, r := cur[0], cur[1]
		addEvent(l, 1, i)
		addEvent(r+1, 0, i)
	}
	// 这里是n-1，因为op的size = n-1
	build(0, 0, n-1)
	var ans int
	var cur int
	for i := 0; i <= X; i++ {
		for _, x := range V[i] {
			if x.second == 0 {
				cur = x.first
			} else {
				update(0, 0, n-1, x.second-1, x.first)
			}
		}
		ans = add(ans, t[0][cur][1])
	}
	return ans
}

type Pair struct {
	first  int
	second int
}

type Mat [2][2]int

func (this Mat) Mul(that Mat) Mat {
	var res Mat
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {
				res[i][k] = add(res[i][k], mul(this[i][j], that[j][k]))
			}
		}
	}
	return res
}

var ZERO = Mat([2][2]int{{3, 0}, {1, 2}})
var ONE = Mat([2][2]int{{1, 2}, {1, 2}})
