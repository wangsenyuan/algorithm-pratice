package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"os"
	"sort"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)
	var buf bytes.Buffer
	for tc > 0 {
		tc--
		n, m := readTwoNums(reader)
		A := readNNums(reader, n)
		Q := make([][]int, m)
		for i := 0; i < m; i++ {
			Q[i] = readNNums(reader, 2)
		}
		res := solve(A, Q)
		for _, num := range res {
			buf.WriteString(fmt.Sprintf("%d\n", num))
		}
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

const MAX = 100010

var BLK_SZ = 230

var factors [][]int

var cnt []int
var arr []int

func init() {
	factors = make([][]int, MAX)
	for i := 1; i < MAX; i++ {
		factors[i] = make([]int, 0, 2)
		for j := 1; j*j <= i; j++ {
			if i%j == 0 {
				factors[i] = append(factors[i], j)
				k := i / j
				if k != j {
					factors[i] = append(factors[i], k)
				}
			}
		}
	}
	cnt = make([]int, MAX)
	arr = make([]int, 2*MAX)
}

func set(p int, v int) {
	p += MAX
	arr[p] = v
	for p > 1 {
		arr[p>>1] = max(arr[p], arr[p^1])
		p >>= 1
	}
}

func add(x int) {
	for _, f := range factors[x] {
		cnt[f]++
		if cnt[f] == 2 {
			set(f, f)
		}
	}
}

func sub(x int) {
	for _, f := range factors[x] {
		cnt[f]--
		if cnt[f] == 1 {
			set(f, 0)
		}
	}
}

func solve(A []int, Q [][]int) []int {
	n := len(A)
	BLK_SZ = int(math.Sqrt(float64(n))) + 1
	m := len(Q)
	qs := make([]Query, m)
	for i := 0; i < m; i++ {
		l, r := Q[i][0], Q[i][1]
		l--
		qs[i] = Query{l, r, i}
	}

	sort.Sort(QS(qs))

	ans := make([]int, m)

	L, R := 0, 0
	for _, q := range qs {
		for R < q.r {
			add(A[R])
			R++
		}
		for R > q.r {
			R--
			sub(A[R])
		}
		for L < q.l {
			sub(A[L])
			L++
		}
		for L > q.l {
			L--
			add(A[L])
		}
		ans[q.ind] = arr[1]
	}

	for i := 0; i < MAX; i++ {
		cnt[i] = 0
	}
	for i := 0; i < 2*MAX; i++ {
		arr[i] = 0
	}

	return ans
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

type Query struct {
	l, r int
	ind  int
}

func (this Query) Less(that Query) bool {
	if this.r/BLK_SZ != that.r/BLK_SZ {
		return this.r/BLK_SZ < that.r/BLK_SZ
	}
	return this.l < that.l
}

type QS []Query

func (qs QS) Len() int {
	return len(qs)
}

func (qs QS) Less(i, j int) bool {
	return qs[i].Less(qs[j])
}

func (qs QS) Swap(i, j int) {
	qs[i], qs[j] = qs[j], qs[i]
}
