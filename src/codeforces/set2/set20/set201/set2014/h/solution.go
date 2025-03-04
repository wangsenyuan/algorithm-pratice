package main

import (
	"bufio"
	"bytes"
	"math/rand"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer
	tc := readNum(reader)
	for range tc {
		res := process(reader)
		for _, x := range res {
			if x {
				buf.WriteString("YES\n")
			} else {
				buf.WriteString("NO\n")
			}
		}
	}
	buf.WriteTo(os.Stdout)
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
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

const X = 1000010
const N = 200010

var h1 *Hash
var h2 *Hash

func init() {
	h1 = NewHash(X, N)
	h2 = NewHash(X, N)
}

func process(reader *bufio.Reader) []bool {
	n, m := readTwoNums(reader)
	a := readNNums(reader, n)
	qs := make([][]int, m)
	for i := range m {
		qs[i] = readNNums(reader, 2)
	}
	return solve(a, qs)
}

func solve(a []int, qs [][]int) []bool {
	n := len(a)

	for i, v := range a {
		h1.Add(i, v)
		h2.Add(i, v)
	}

	ans := make([]bool, len(qs))
	for i, cur := range qs {
		l, r := cur[0], cur[1]
		v1 := h1.Get(l-1, r-1)
		v2 := h2.Get(l-1, r-1)

		ans[i] = v1 == 0 && v2 == 0
	}

	h1.Reset(n)
	h2.Reset(n)

	return ans
}

type Hash struct {
	val []int
	xor []int
}

func NewHash(x int, n int) *Hash {
	val := make([]int, x+1)
	for i := range x + 1 {
		val[i] = rand.Int()
	}
	xor := make([]int, n+1)
	return &Hash{val, xor}
}

func (hash *Hash) Add(i int, x int) {
	hash.xor[i+1] = hash.xor[i] ^ hash.val[x]
}

func (hash *Hash) Get(l int, r int) int {
	return hash.xor[r+1] ^ hash.xor[l]
}

func (hash *Hash) Reset(n int) {
	clear(hash.xor[:n+1])
}
