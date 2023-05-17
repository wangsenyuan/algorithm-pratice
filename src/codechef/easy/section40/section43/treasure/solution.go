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
		n, m := readTwoNums(reader)
		A := make([][]int, n)
		for i := 0; i < n; i++ {
			A[i] = readNNums(reader, m)
		}
		res := solve(A)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
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

const MOD = 1e9 + 7

func mul(a, b int) int {
	return int(int64(a) * int64(b) % MOD)
}

func pow(a, b int) int {
	r := 1
	for b > 0 {
		if b&1 == 1 {
			r = mul(r, a)
		}
		a = mul(a, a)
		b >>= 1
	}
	return r
}

func solve(A [][]int) int {
	n := len(A)
	m := len(A[0])
	var ctr int
	pos := make([][]int, n)
	for i := 0; i < n; i++ {
		pos[i] = make([]int, m)
		for j := 0; j < m; j++ {
			ctr++
			pos[i][j] = ctr
		}
	}

	eqs := make([]BitSet, m*n)

	var id int

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if A[i][j] != -1 {
				eqs[id] = NewBitSet(m*n + 1)
				if A[i][j] == 1 {
					eqs[id].Set(0)
				}
				if i > 0 {
					eqs[id].Set(pos[i-1][j])
				}
				if i+1 < n {
					eqs[id].Set(pos[i+1][j])
				}
				if j > 0 {
					eqs[id].Set(pos[i][j-1])
				}
				if j+1 < m {
					eqs[id].Set(pos[i][j+1])
				}
				id++
			}
		}
	}

	var r int
	for i := 1; i <= n*m; i++ {
		first := -1
		for j := r; j < id; j++ {
			if eqs[j].IsSet(i) {
				first = j
				break
			}
		}
		if first < 0 {
			continue
		}
		if first != r {
			eqs[r] = eqs[r].Xor(eqs[first])
		}
		for j := r + 1; j < id; j++ {
			if eqs[j].IsSet(i) {
				eqs[j] = eqs[j].Xor(eqs[r])
			}
		}
		r++
	}
	for i := r; i < id; i++ {
		if eqs[i].IsSet(0) {
			return 0
		}
	}
	cnt := n*m - r
	return pow(2, cnt)
}

const BIT_SIZE = 64

type BitSet struct {
	arr []uint64
	sz  int
	n   int
}

func NewBitSet(n int) BitSet {
	sz := (n + BIT_SIZE - 1) / BIT_SIZE
	arr := make([]uint64, sz)
	return BitSet{arr, sz, n}
}

func (bs *BitSet) Set(i int) {
	arr := bs.arr
	if i >= bs.n {
		return
	}
	arr[i/BIT_SIZE] |= 1 << uint64(i%BIT_SIZE)
}

func (bs *BitSet) IsSet(i int) bool {
	arr := bs.arr
	if i >= bs.n {
		return false
	}
	return arr[i/BIT_SIZE]&(1<<uint64(i%BIT_SIZE)) > 0
}

func (bs *BitSet) Clear(i int) {
	if !bs.IsSet(i) {
		return
	}
	arr := bs.arr
	arr[i/BIT_SIZE] ^= 1 << uint64(i%BIT_SIZE)
}

func (bs BitSet) RightShift(x int) BitSet {
	arr := make([]uint64, len(bs.arr))

	a, b := x/BIT_SIZE, x%BIT_SIZE
	B := uint64(1<<uint64(b) - 1)
	var i int
	for a < bs.sz {
		cur := bs.arr[a]
		if b > 0 {
			cur >>= uint64(b)
			if a+1 < bs.sz {
				cur |= (bs.arr[a+1] & B) << (BIT_SIZE - uint64(b))
			}
		}
		arr[i] = cur
		i++
		a++
	}
	return BitSet{arr, bs.sz, bs.n}
}

func (this BitSet) And(that BitSet) BitSet {
	and := func(x, y uint64) uint64 {
		return x & y
	}
	return merge(this, that, and)
}

func (this BitSet) Or(that BitSet) BitSet {
	or := func(x, y uint64) uint64 {
		return x | y
	}
	return merge(this, that, or)
}

func (this BitSet) Xor(that BitSet) BitSet {
	xor := func(x, y uint64) uint64 {
		return x ^ y
	}
	return merge(this, that, xor)
}

func merge(this, that BitSet, fn func(uint64, uint64) uint64) BitSet {
	arr := make([]uint64, this.sz)

	for i := 0; i < this.sz; i++ {
		arr[i] = fn(this.arr[i], that.arr[i])
	}

	return BitSet{arr, this.sz, this.n}
}

func (this BitSet) Count() int {
	var ans int

	for i := 0; i < this.sz; i++ {
		cur := this.arr[i]
		for cur > 0 {
			ans++
			cur -= cur & -cur
		}
	}
	return ans
}
