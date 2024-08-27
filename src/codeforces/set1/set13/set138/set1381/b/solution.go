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
		n := readNum(reader)
		p := readNNums(reader, 2*n)
		res := solve(p)
		if res {
			buf.WriteString("YES\n")
		} else {
			buf.WriteString("NO\n")
		}
	}

	fmt.Print(buf.String())
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

func solve(p []int) bool {
	n2 := len(p)
	n := n2 / 2

	pos := make([]int, n2)
	for i, x := range p {
		pos[x-1] = i
	}

	if n2-pos[n2-1] > n {
		return false
	}

	dp := NewBitSet(n + 1)
	dp.Set(n2 - pos[n2-1])

	for x, r := n2-2, pos[n2-1]; x >= 0; x-- {
		i := pos[x]
		if i > r {
			continue
		}

		diff := r - i
		if diff > n {
			return false
		}
		fp := dp.Copy()

		dp.LeftShift(diff)
		dp.Set(diff)
		dp.Union(fp)

		r = i
	}

	return dp.IsSet(n)
}

type BitSet struct {
	sz  int
	arr []uint64
}

func NewBitSet(n int) *BitSet {
	sz := (n + 63) / 64
	arr := make([]uint64, sz)
	return &BitSet{sz, arr}
}

func (bs *BitSet) Set(p int) {
	i, j := p/64, p%64
	i = bs.sz - 1 - i
	bs.arr[i] |= 1 << uint64(j)
}

func (bs *BitSet) IsSet(p int) bool {
	i, j := p/64, p%64
	i = bs.sz - 1 - i
	return (bs.arr[i]>>uint64(j))&1 == 1
}

func (bs *BitSet) Count() int {
	var res int
	for i := 0; i < bs.sz; i++ {
		res += countDigit(bs.arr[i])
	}
	return res
}

func countDigit(num uint64) int {
	var res int
	if (num>>63)&1 == 1 {
		res++
	}
	tmp := int64(num & ((1 << 63) - 1))

	for tmp > 0 {
		res++
		tmp -= tmp & -tmp
	}
	return res
}

func (bs *BitSet) LeftShift(p int) {
	i, j := p/64, p%64
	if j == 0 {
		for u := 0; u+i < bs.sz; u++ {
			bs.arr[u] = bs.arr[u+i]
		}
	} else {
		for u := 0; u+i < bs.sz; u++ {
			v := u + i
			bs.arr[u] = bs.arr[v] << uint64(j)
			if v+1 < bs.sz {
				bs.arr[u] |= bs.arr[v+1] >> uint64(64-j)
			}
		}
	}
	for u := bs.sz - i; u < bs.sz; u++ {
		bs.arr[u] = 0
	}
}

func (bs *BitSet) RightShift(p int) {
	i, j := p/64, p%64
	if j == 0 {
		for u := bs.sz - 1; u-i >= 0; u-- {
			bs.arr[u] = bs.arr[u-i]
		}
	} else {
		for u := bs.sz - 1; u-i >= 0; u-- {
			v := u - i
			bs.arr[u] = bs.arr[v] >> uint64(j)
			if v-1 >= 0 {
				bs.arr[u] |= bs.arr[v-1] << uint64(64-j)
			}
		}
	}
	for u := i - 1; u >= 0; u-- {
		bs.arr[u] = 0
	}
}

func (bs *BitSet) Copy() *BitSet {
	res := NewBitSet(len(bs.arr) * 64)
	copy(res.arr, bs.arr)
	return res
}

func (this *BitSet) Union(that *BitSet) {
	for i := 0; i < len(this.arr); i++ {
		this.arr[i] |= that.arr[i]
	}
}

func (this *BitSet) Clear() {
	clear(this.arr)
}
