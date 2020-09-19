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
		var n, q int
		var p uint64
		s, _ := reader.ReadBytes('\n')
		pos := readInt(s, 0, &n)
		pos = readInt(s, pos+1, &q)
		readUint64(s, pos+1, &p)
		A := readNNums(reader, n)
		solver := NewSolver(n, A, int64(p))

		for q > 0 {
			q--
			a, b, c := readThreeNums(reader)
			if a == 1 {
				solver.Update(b, c)
			} else {
				res := solver.Query(b, c)
				buf.WriteString(fmt.Sprintf("%d %d %d %d\n", res[0], res[1], res[2], res[3]))
			}
		}
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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
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

type Node struct {
	num [4]int64
}

func (this *Node) clone() *Node {
	that := new(Node)
	that.num = [4]int64{this.num[0], this.num[1], this.num[2], this.num[3]}
	return that
}

const MAX_A = 100010

const MAX_N = 100005

var nodes [MAX_N * 4]*Node

var quard [MAX_A]*Node

func init() {
	s := make([][]int, MAX_A)
	for i := 0; i*i < MAX_A; i++ {
		for j := 0; j*j < MAX_A; j++ {
			if i*i+j*j >= MAX_A {
				break
			}
			for k := j; k*k < MAX_A; k++ {
				if i*i+j*j+k*k >= MAX_A {
					break
				}
				s[i*i+j*j+k*k] = []int{i, j, k, 0}
			}
		}
	}

	for i := 0; i < MAX_A; i++ {
		if len(s[i]) != 0 {
			continue
		}
		for j := 1; j*j < MAX_A; j++ {
			if i-j*j < 0 {
				break
			}
			if s[i-j*j] != nil && s[i-j*j][3] == 0 {
				s[i] = []int{s[i-j*j][0], s[i-j*j][1], s[i-j*j][2], j}
				break
			}
		}
	}
	for i := 0; i < MAX_A; i++ {
		quard[i] = new(Node)
		quard[i].num = [...]int64{int64(s[i][0]), int64(s[i][1]), int64(s[i][2]), int64(s[i][3])}
	}
}

type Solver struct {
	n     int
	mod   int64
	nodes []*Node
}

func (solver Solver) Mul(a, b int64) int64 {
	a0 := a % (1 << 20)
	a1 := a / (1 << 20)
	b0 := b % (1 << 20)
	b1 := b / (1 << 20)
	p := solver.mod

	ret := ((a1*b1)%p*(1<<20)%p*(1<<20) + (a1*b0+a0*b1)%p*(1<<20) + a0*b0) % p

	if ret < 0 {
		ret += p
	}

	return ret
}

func (solver *Solver) combine(a, b *Node) *Node {
	if a == nil {
		return b.clone()
	}
	if b == nil {
		return a.clone()
	}

	//(a1^2 + a2^2 + a3^2 + a4^2) * (b1^2 + b2^2 + b3^2 + b4^2) =
	//(a1b1 + a2b2 + a3b3 + a4b4)^2 + (a1b2 - a2b1 + a3b4 - a4b3)^2 +
	//(a1b3 - a2b4 - a3b1 + a4b2)^2 + (a1b4 + a2b3 - a3b2 - a4b1)^2

	mult := solver.Mul
	c := new(Node)
	c.num = [...]int64{0, 0, 0, 0}
	c.num[0] = (mult(a.num[0], b.num[0]) + mult(a.num[1], b.num[1]) + mult(a.num[2], b.num[2]) + mult(a.num[3], b.num[3])) % solver.mod
	c.num[1] = (2*solver.mod + mult(a.num[0], b.num[1]) - mult(a.num[1], b.num[0]) + mult(a.num[2], b.num[3]) - mult(a.num[3], b.num[2])) % solver.mod
	c.num[2] = (2*solver.mod + mult(a.num[0], b.num[2]) - mult(a.num[1], b.num[3]) - mult(a.num[2], b.num[0]) + mult(a.num[3], b.num[1])) % solver.mod
	c.num[3] = (2*solver.mod + mult(a.num[0], b.num[3]) + mult(a.num[1], b.num[2]) - mult(a.num[2], b.num[1]) - mult(a.num[3], b.num[0])) % solver.mod

	return c
}

func (solver *Solver) BreakApart(x int) *Node {
	tmp := new(Node)
	tmp.num = [...]int64{
		quard[x].num[0] % solver.mod,
		quard[x].num[1] % solver.mod,
		quard[x].num[2] % solver.mod,
		quard[x].num[3] % solver.mod,
	}

	return tmp
}

func NewSolver(n int, A []int, p int64) Solver {
	//nodes := make([]*Node, 4*n)

	solver := new(Solver)
	solver.nodes = nodes[:4*n]
	solver.mod = p
	solver.n = n

	var build func(nd int, l, r int)

	build = func(nd int, l, r int) {
		if l == r {
			nodes[nd] = solver.BreakApart(A[l])
			return
		}
		mid := (l + r) / 2
		build(nd*2, l, mid)
		build(nd*2+1, mid+1, r)
		nodes[nd] = solver.combine(nodes[nd*2], nodes[nd*2+1])
	}

	build(1, 0, n-1)
	return *solver
}

func (solver *Solver) Update(p int, v int) {
	p--
	nodes := solver.nodes
	var loop func(nd int, l, r int)
	loop = func(nd int, l, r int) {
		if l == r {
			nodes[nd] = solver.BreakApart(v)
			return
		}
		mid := (l + r) / 2
		if p > mid {
			loop(nd*2+1, mid+1, r)
		} else {
			loop(nd*2, l, mid)
		}
		nodes[nd] = solver.combine(nodes[nd*2], nodes[nd*2+1])
	}

	loop(1, 0, solver.n-1)
}

func (solver Solver) Query(l, r int) [4]int64 {
	l--
	r--
	var loop func(nd int, ll, rr int) *Node

	loop = func(nd int, ll, rr int) *Node {
		if rr < l || r < ll {
			return nil
		}

		if l <= ll && rr <= r {
			return solver.nodes[nd]
		}
		mid := (ll + rr) / 2
		a := loop(nd*2, ll, mid)
		b := loop(nd*2+1, mid+1, rr)
		return solver.combine(a, b)
	}

	res := loop(1, 0, solver.n-1)
	return res.num
}
