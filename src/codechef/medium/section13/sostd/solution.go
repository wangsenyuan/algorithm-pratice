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

	tc := readNum(reader)
	var buf bytes.Buffer
	for tc > 0 {
		tc--
		n := readNum(reader)
		s, _ := reader.ReadBytes('\n')
		var pos int
		for i := 0; i < n; i++ {
			pos = readInt(s, pos, &arr[i].first) + 1
		}
		s, _ = reader.ReadBytes('\n')
		pos = 0
		for i := 0; i < n; i++ {
			pos = readInt(s, pos, &arr[i].second) + 1
		}
		res := solve(n, arr[:n])
		buf.WriteString(fmt.Sprintf("%d\n", res))
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
	tot     int64
	subsets int64
	lazy    int64
}

const MAX_V = 1000004
const MOD = 1000000007
const MX = 1 << 18

var tree [MAX_V * 4]*Node
var arr [MX]Pair

func init() {
	for i := 0; i < len(tree); i++ {
		tree[i] = new(Node)
		tree[i].lazy = 1
	}
}

func mul(x int, coef int64) {
	tree[x].tot = (tree[x].tot * coef) % MOD
	tree[x].subsets = (tree[x].subsets * coef) % MOD
	tree[x].lazy = (tree[x].lazy * coef) % MOD
}

func push(x int) {
	if tree[x].lazy == 1 {
		return
	}
	mul(2*x, tree[x].lazy)
	mul(2*x+1, tree[x].lazy)
	tree[x].lazy = 1
}

func update(x int, l, r int, st, en int, used *[]int) {
	*used = append(*used, x)
	if st > r || en < l {
		return
	}

	if st <= l && r <= en {
		mul(x, 2)
		return
	}
	push(x)
	mid := (l + r) / 2
	update(2*x, l, mid, st, en, used)
	update(2*x+1, mid+1, r, st, en, used)
	tree[x].tot = (tree[2*x].tot + tree[2*x+1].tot) % MOD
	tree[x].subsets = (tree[2*x].subsets + tree[2*x+1].subsets) % MOD
}

func merge(a, b *Node) *Node {
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}

	c := new(Node)
	c.tot = (a.tot + b.tot) % MOD
	c.subsets = (a.subsets + b.subsets) % MOD
	return c
}

func query(x int, l, r int, st, en int) *Node {
	if r < st || en < l {
		return nil
	}
	if st <= l && r <= en {
		return tree[x]
	}

	push(x)
	mid := (l + r) / 2
	a := query(2*x, l, mid, st, en)
	b := query(2*x+1, mid+1, r, st, en)

	return merge(a, b)
}

func add(x int, l, r int, pos int, toAdd int64, used *[]int) {
	if pos < l || r < pos {
		return
	}
	*used = append(*used, x)
	if l == r {
		tree[x].subsets += toAdd
		tree[x].subsets %= MOD
		tree[x].tot = (int64(l) * tree[x].subsets) % MOD
		return
	}

	push(x)
	mid := (l + r) / 2
	if pos <= mid {
		add(2*x, l, mid, pos, toAdd, used)
	} else {
		add(2*x+1, mid+1, r, pos, toAdd, used)
	}
	tree[x].tot = (tree[2*x].tot + tree[2*x+1].tot) % MOD
	tree[x].subsets = (tree[2*x].subsets + tree[2*x+1].subsets) % MOD
}

func solve(n int, arr []Pair) int {
	sort.Sort(Pairs(arr))
	var used []int
	var res int64
	for i := 0; i < n; i++ {
		cur := arr[i]
		v := cur.second
		st, en := v+1, MAX_V
		if st <= en {
			res += (int64(cur.first) * query(1, 1, MAX_V, st, en).tot) % MOD
			res %= MOD
			update(1, 1, MAX_V, st, en, &used)
		}
		st, en = 1, v
		var toAdd int64 = 1
		if st <= en {
			toAdd += query(1, 1, MAX_V, st, en).subsets
			toAdd %= MOD
		}
		res += ((int64(cur.first) * int64(cur.second)) % MOD * toAdd) % MOD
		res %= MOD
		add(1, 1, MAX_V, v, toAdd, &used)
	}

	for _, x := range used {
		tree[x].tot = 0
		tree[x].subsets = 0
		tree[x].lazy = 1
	}

	return int(res)
}

type Pair struct {
	first, second int
}

type Pairs []Pair

func (ps Pairs) Len() int {
	return len(ps)
}

func (ps Pairs) Less(i, j int) bool {
	return ps[i].first < ps[j].first || ps[i].first == ps[j].first && ps[i].second < ps[j].second
}

func (ps Pairs) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}
