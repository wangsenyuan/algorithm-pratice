package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math/rand"
	"os"
)

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

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m := readTwoNums(reader)

	s, _ := reader.ReadString('\n')

	solver := NewSolver(n, s)

	var buf bytes.Buffer

	for m > 0 {
		m--
		t, l, r := readThreeNums(reader)
		res := solver.Solve(t, l, r)
		if t == 2 {
			buf.WriteString(fmt.Sprintf("%d\n", res))
		}
	}
	fmt.Print(buf.String())
}

const MOD = 1000000007
const MAX_N = 100333

var F [MAX_N]int64
var I [MAX_N]int64

func init() {
	F[0] = 1
	for i := 1; i < MAX_N; i++ {
		F[i] = int64(i) * F[i-1] % MOD
	}
	I[MAX_N-1] = pow(F[MAX_N-1], MOD-2)
	for i := MAX_N - 2; i >= 0; i-- {
		I[i] = int64(i+1) * I[i+1] % MOD
	}
}

func pow(a, b int64) int64 {
	var res int64 = 1
	for b > 0 {
		if b&1 == 1 {
			res = (res * a) % MOD
		}
		a = (a * a) % MOD
		b >>= 1
	}
	return res
}

type Solver struct {
	n int
	t *Node
}

func NewSolver(n int, s string) *Solver {
	var t *Node
	for i := 0; i < n; i++ {
		node := new(Node)
		node.sum[int(s[i]-'a')]++
		node.value[int(s[i]-'a')]++
		node.priority = rand.Int()
		merge(&t, t, node)
	}
	return &Solver{n, t}
}

func (solver *Solver) Solve(tp int, l int, r int) int {
	l--
	r--
	var t1, t2, t3 *Node
	split(solver.t, &t1, &t2, l)
	split(t2, &t2, &t3, r-l+1)
	if tp == 1 {
		t2.reversed ^= 1
		merge(&(solver.t), t1, t2)
		merge(&(solver.t), solver.t, t3)
		return 0
	}
	var odd int
	for j := 0; j < H; j++ {
		cnt[j] = t2.Sum(j)
		odd += cnt[j] & 1
	}
	merge(&(solver.t), t1, t2)
	merge(&(solver.t), solver.t, t3)
	if odd > 1 {
		return 0
	}
	for j := 0; j < H; j++ {
		if cnt[j]&1 == 1 {
			cnt[j]--
		}
	}

	var w int
	for j := 0; j < H; j++ {
		w += cnt[j] / 2
	}

	ans := F[w]
	for j := 0; j < H; j++ {
		ans = (ans * I[cnt[j]/2]) % MOD
	}
	return int(ans)
}

const H = 10

var cnt [H]int

type Node struct {
	priority int
	cnt      int
	value    [H]int
	sum      [H]int

	left, right *Node
	reversed    int
}

func (node *Node) Count() int {
	if node == nil {
		return 0
	}
	return node.cnt
}

func (node *Node) Sum(i int) int {
	if node == nil {
		return 0
	}
	return node.sum[i]
}

func push(node *Node) {
	if node == nil {
		return
	}
	if node.reversed > 0 {
		if node.left != nil {
			node.left.reversed ^= 1
		}
		if node.right != nil {
			node.right.reversed ^= 1
		}
		node.left, node.right = node.right, node.left
		node.reversed ^= 1
	}
	node.cnt = node.left.Count() + node.right.Count() + 1
	for i := 0; i < H; i++ {
		node.sum[i] = node.value[i] + node.left.Sum(i) + node.right.Sum(i)
	}
}

func split(t *Node, l **Node, r **Node, key int) {
	push(t)
	if t == nil {
		*l = nil
		*r = nil
		return
	}
	if key <= t.left.Count() {
		split(t.left, l, &(t.left), key)
		*r = t
	} else {
		split(t.right, &(t.right), r, key-t.left.Count()-1)
		*l = t
	}
	push(*l)
	push(*r)
}

func merge(t **Node, l *Node, r *Node) {
	push(l)
	push(r)
	if l == nil || r == nil {
		if l != nil {
			*t = l
		} else {
			*t = r
		}
	} else if l.priority > r.priority {
		merge(&(l.right), l.right, r)
		*t = l
	} else {
		merge(&(r.left), l, r.left)
		*t = r
	}
	push(*t)
}
