package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(scanner *bufio.Scanner) (a int) {
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &a)
	return
}

func readTwoNums(scanner *bufio.Scanner) (a int, b int) {
	res := readNNums(scanner, 2)
	a, b = res[0], res[1]
	return
}

func readThreeNums(scanner *bufio.Scanner) (a int, b int, c int) {
	res := readNNums(scanner, 3)
	a, b, c = res[0], res[1], res[2]
	return
}

func readNNums(scanner *bufio.Scanner, n int) []int {
	res := make([]int, n)
	x := 0
	scanner.Scan()
	for i := 0; i < n; i++ {
		for x < len(scanner.Bytes()) && scanner.Bytes()[x] == ' ' {
			x++
		}
		x = readInt(scanner.Bytes(), x, &res[i])
	}
	return res
}

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
}

func main() {
	reader := bufio.NewScanner(os.Stdin)

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n, k := readTwoNums(reader)

		segs := make([][]int, n)
		for i := 0; i < n; i++ {
			segs[i] = readNNums(reader, 2)
		}

		fmt.Println(solve(n, k, segs))
	}
}

const MOD = 1000000007
const MAX_A = 1000000001
const MAX_N = 400007

var F []int64
var IF []int64

func init() {
	F = make([]int64, MAX_N)
	F[0] = 1
	for i := 1; i < MAX_N; i++ {
		F[i] = int64(i) * F[i-1]
		F[i] %= MOD
	}
	IF = make([]int64, MAX_N)
	IF[MAX_N-1] = pow(F[MAX_N-1], MOD-2)

	for i := MAX_N - 2; i >= 0; i-- {
		IF[i] = int64(i+1) * IF[i+1]
		IF[i] %= MOD
		if (F[i]*IF[i])%MOD != 1 {
			panic("wrong")
		}
	}
}

func nCr(n int, r int) int64 {
	if n < r {
		return 0
	}
	res := F[n]
	res *= IF[r]
	res %= MOD
	res *= IF[n-r]
	res %= MOD
	return res
}

func pow(a, b int64) int64 {
	var res int64 = 1
	for b > 0 {
		if b&1 == 1 {
			res *= a
			res %= MOD
		}
		a = (a * a) % MOD
		b >>= 1
	}
	return res
}

func solve(n int, k int, seg [][]int) int {
	ps := make([]Pair, n)
	for i := 0; i < n; i++ {
		ps[i] = Pair{seg[i][0], seg[i][1]}
	}

	sort.Sort(Pairs(ps))

	var ans int64

	root := NewNode(0, MAX_A)

	for i := 0; i < n; i++ {
		cur := ps[i]

		for root.GetSize() > 0 {
			r := root.GetMin()
			if r.start >= cur.first {
				break
			}
			root.Set(r.start, -r.count)
		}

		ans += nCr(root.GetSize(), k-1)

		if ans >= MOD {
			ans -= MOD
		}
		root.Set(cur.second, 1)
	}

	ans = (nCr(n, k) - ans + MOD) % MOD

	return int(ans)
}

type Pair struct {
	first, second int
}

type Pairs []Pair

func (ps Pairs) Len() int {
	return len(ps)
}

func (ps Pairs) Less(i, j int) bool {
	return ps[i].first < ps[j].first || (ps[i].first == ps[j].first && ps[i].second > ps[j].second)
}

func (ps Pairs) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

type Node struct {
	start, end  int
	left, right *Node
	count       int
}

func NewNode(start, end int) *Node {
	node := new(Node)
	node.start = start
	node.end = end
	node.count = 0
	return node
}

func (node *Node) push() {
	if node.start < node.end && node.left == nil {
		mid := (node.start + node.end) / 2
		node.left = NewNode(node.start, mid)
		node.right = NewNode(mid+1, node.end)
	}
}

func (node *Node) Set(pos int, v int) {
	node.count += v

	if node.start == node.end {
		return
	}
	node.push()

	mid := (node.start + node.end) / 2

	if pos <= mid {
		node.left.Set(pos, v)
	} else {
		node.right.Set(pos, v)
	}
}

func (node *Node) GetMin() *Node {
	if node.start == node.end {
		return node
	}

	if node.left.count > 0 {
		return node.left.GetMin()
	}

	return node.right.GetMin()
}

func (node *Node) GetSize() int {
	return node.count
}
