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
		A := readNNums(reader, n)
		res := solve(A)
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

const ZERO = 2000000
const MAX = ZERO + 2002001

func solve(A []int) int {
	has := make([]int, 2001)
	sort.Ints(A)

	sum := getSum(A)

	bs := NewBitSet(MAX)

	bs.Set(ZERO)

	for _, a := range A {
		sum -= a
		add := a - sum
		for i := 0; i <= a; i++ {
			if ZERO+i-add < MAX && bs.IsSet(ZERO+i-add) {
				has[i] = 1
			}
		}
		tmp := bs.Copy()
		tmp.RightShift(a)
		bs.LeftShift(a)
		bs.Union(tmp)
	}

	ans := getSum(has)

	return ans
}

func getSum(arr []int) int {
	var res int
	for _, num := range arr {
		res += num
	}
	return res
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
