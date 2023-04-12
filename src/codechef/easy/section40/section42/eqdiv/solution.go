package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	k := readNum(reader)

	tc := readNum(reader)

	nums := make([]int, tc)

	for i := 0; i < tc; i++ {
		nums[i] = readNum(reader)
	}

	diff, res := solve(k, nums)

	var buf bytes.Buffer

	for i := 0; i < tc; i++ {
		buf.WriteString(fmt.Sprintf("%d\n", diff[i]))
		buf.WriteString(res[i])
		buf.WriteByte('\n')
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

const S = 20 * 1000 * 1000

var dp []*BitSet
var answer []int
var belong [][]int
var cost []int

func pow(x int, k int) int {
	res := 1
	for i := 0; i < k; i++ {
		res *= x
	}
	return res
}

func calculateDp(n int, k int) {

	answer = make([]int, n)
	cost = make([]int, n)
	dp = make([]*BitSet, n)
	belong = make([][]int, n)

	for i := 0; i < n; i++ {
		dp[i] = NewBitSet(2 * S)
		belong[i] = make([]int, 0, 1)
	}
	dp[0].Set(S)
	for i := 1; i < n; i++ {
		// i  ** k
		cost[i] = pow(i, k)
		a := dp[i-1].Copy()
		a.RightShift(cost[i])
		b := dp[i-1].Copy()
		b.LeftShift(cost[i])
		a.Union(b)
		// dp[i] = dp[i-1] << cost[i] | dp[i-1] >> cost[i]
		dp[i] = a
	}

	for i := 1; i < n; i++ {
		for j := S; ; j++ {
			if dp[i].IsSet(j) {
				answer[i] = j - S
				for l := i; l >= 1; l-- {
					if j-cost[l] >= 0 && dp[l-1].IsSet(j-cost[l]) {
						belong[i] = append(belong[i], l)
						j -= cost[l]
					} else {
						j += cost[l]
					}
				}
				reverse(belong[i])
				break
			}
		}
	}
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func solve(k int, nums []int) ([]int, []string) {
	// when k = 1
	// sum = (n + 1) * n / 2
	// n = 1e6
	// dp 也不行
	calculateDp(1<<(k+2), k)
	bal := make([]int, len(nums))
	res := make([]string, len(nums))
	for i := 0; i < len(nums); i++ {
		bal[i], res[i] = solveN(nums[i], k)
	}
	return bal, res
}

func solveN(n int, k int) (int, string) {
	buf := make([]byte, n)
	for i := 0; i < n; i++ {
		buf[i] = '0'
	}

	for n >= 1<<(k+2) {
		for i := 0; i < 1<<(k+1); i++ {
			if popCount(i)&1 == 1 {
				buf[n-1-i] = '1'
			}
		}
		n -= 1 << (k + 1)
	}
	for _, i := range belong[n] {
		buf[i-1] = '1'
	}
	return answer[n], string(buf)
}

func popCount(n int) int {
	var res int
	for n > 0 {
		res++
		n -= n & -n
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
