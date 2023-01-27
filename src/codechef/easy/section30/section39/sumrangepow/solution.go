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
	tc := readNum(reader)

	for tc > 0 {
		tc--
		n, k := readTwoNums(reader)
		A := readNNums(reader, n)
		res := solve(k, A)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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
		if s[i] == '\n' || s[i] == '\r' {
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

const MOD = 998244353
const N = 100005
const K = 6

var C [100][100]int
var pows [15][N]int

func init() {
	C[0][0] = 1
	for i := 1; i < 100; i++ {
		C[i][0] = 1
		for j := 1; j <= i; j++ {
			C[i][j] = modAdd(C[i-1][j-1], C[i-1][j])
		}
	}

	for i := 1; i < N; i++ {
		for j := 0; j < 15; j++ {
			if j == 0 {
				pows[j][i] = 1
			} else {
				pows[j][i] = modMul(i, pows[j-1][i])
			}
		}
	}
}

type Node struct {
	val   []int
	left  *Node
	right *Node
	lazy  int
	lo    int
	hi    int
}

func NewNode(lo, hi int) *Node {
	node := new(Node)
	node.lo = lo
	node.hi = hi
	node.val = make([]int, K)
	node.val[0] = hi - lo
	return node
}

func (node *Node) push() {
	if node.left == nil {
		mid := (node.lo + node.hi) / 2
		node.left = NewNode(node.lo, mid)
		node.right = NewNode(mid, node.hi)
	}
	if node.lazy != 0 {
		node.left.Add(node.lo, node.hi, node.lazy)
		node.right.Add(node.lo, node.hi, node.lazy)
		node.lazy = 0
	}
}

func merge(val []int, a []int, b []int) {
	for i := 0; i < K; i++ {
		val[i] = modAdd(a[i], b[i])
	}
}

func (node *Node) Add(L int, R int, x int) {
	if R <= node.lo || node.hi <= L {
		return
	}
	if L <= node.lo && node.hi <= R {
		node.lazy += x
		for k := K - 1; k > 0; k-- {
			var cur int
			for j := k; j >= 0; j-- {
				cur = modAdd(cur, modMul(node.val[j], modMul(C[k][j], pows[k-j][x])))
			}
			node.val[k] = cur
		}
	} else {
		node.push()
		node.left.Add(L, R, x)
		node.right.Add(L, R, x)
		merge(node.val, node.left.val, node.right.val)
	}
}

func (node *Node) Query(L int, R int) []int {
	if R <= node.lo || node.hi <= L {
		return make([]int, K)
	}
	if L <= node.lo && node.hi <= R {
		return node.val
	}
	node.push()
	a := node.left.Query(L, R)
	b := node.right.Query(L, R)
	var c []int
	merge(c, a, b)
	return c
}

func modAdd(a, b int) int {
	a += b
	if a >= MOD {
		a -= MOD
	}
	return a
}

func modMul(a, b int) int {
	r := int64(a) * int64(b)
	return int(r % MOD)
}

func solve(k int, A []int) int {
	n := len(A)
	var ans int
	prev := make([]int, n+1)
	for i := 0; i <= n; i++ {
		prev[i] = -1
	}
	seg := NewNode(0, n)
	for r := 0; r < n; r++ {
		seg.Add(prev[A[r]]+1, r+1, 1)
		tmp := seg.Query(0, n)
		ans = modAdd(ans, tmp[k])
		prev[A[r]] = r
	}

	return ans
}
