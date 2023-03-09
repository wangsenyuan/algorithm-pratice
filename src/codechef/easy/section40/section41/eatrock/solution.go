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
	var buf bytes.Buffer
	tc := readNum(reader)
	for tc > 0 {
		tc--
		n := readNum(reader)
		X := readNNums(reader, n)
		W := readNNums(reader, n)
		res := solve(X, W)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	fmt.Println(buf.String())
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

const MOD = 1000000007

func add(a, b int) int {
	a += b
	if a >= MOD {
		a -= MOD
	}
	return a
}

func sub(a, b int) int {
	return add(a, MOD-b)
}

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

func inverse(a int) int {
	return pow(a, MOD-2)
}

func solve(X []int, W []int) int {
	// W[a] < W[b] < W[c] ...
	// 需要计算期望值
	// 按照W排序，只能从低w向高w移动
	// 当然也可以直接从w开始移动
	// 假设到达w前的期望值值是sw
	// 到w增加的是 k*xi - sum(xj) xj是小于xi的为止，一共有k个
	// +  sum(xjj) -（I - k) * ki 大于xi的的，一共有 I - k个
	// dp[i] = 到达i时的期望值
	// fp[i] = 从i出发的期望值
	//  k*xi - sum(xj) + ... = 集合S只有两个元素时，才能这样计算
	// 假设有三个元素 a, b, c (wa < wb < wc)
	// 如果 xa < xb < xc那么就是 xb - xa + xc - xb => xc - xa
	//     xb < xa < xc => xa - xb + xc - xb => xc + xa - 2 * xb
	//     xc < xa < xb => xb - xa + xb - xc => 2 * xb - xc - xa
	n := len(X)
	rs := make([]Pair, n)
	xs := make(map[int]int)
	for i := 0; i < n; i++ {
		rs[i] = Pair{W[i], X[i]}
		// x[i] < x[i+1]
		xs[X[i]] = i
	}

	sort.Slice(rs, func(i, j int) bool {
		return rs[i].first < rs[j].first
	})

	dp := make([]int, n)
	// dp[i] = dp[j] + pow(2, j) * dist(j, i) =

	//for i := range rs {
	//	// 找到比x小的一半
	//	// 2 * 34 - 2 * 20 +  1 * 20 - 1 * 1 = 2 * 14 + 1 * 19
	//
	//	for j := 0; j < i; j++ {
	//		if rs[j].second < rs[i].second {
	//			dp[i] = add(dp[i], add(dp[j], mul(pow(2, j), rs[i].second-rs[j].second)))
	//		} else {
	//			dp[i] = add(dp[i], add(dp[j], mul(pow(2, j), rs[j].second-rs[i].second)))
	//		}
	//	}
	//}

	// how to improve it
	// dist(i, j) = x(i) - x(j) when xi > xj
	// pw(2, j) * x(j) 没有问题
	// pw(2, j) + pw(2, k)
	pw_sum := NewBit(n)
	dis_sum := NewBit(n)
	var sum int

	for i, r := range rs {
		x := r.second
		j := xs[x]
		a := pw_sum.Get(n - 1)
		b := pw_sum.Get(j)
		c := sub(a, b)
		d := dis_sum.Get(n - 1)
		e := dis_sum.Get(j)
		f := sub(d, e)
		dp[i] = add(sum, sub(mul(b, x), e))
		dp[i] = add(dp[i], sub(f, mul(c, x)))
		sum = add(sum, dp[i])
		pw_sum.Add(j, pow(2, i))
		dis_sum.Add(j, mul(x, pow(2, i)))
	}

	var res int

	for i := 0; i < n; i++ {
		res = add(res, dp[i])
	}

	ans := mul(res, inverse(pow(2, n)))
	return ans
}

func solve1(X []int, W []int) int {
	n := len(X)
	rs := make([]Pair, n)
	pos := make(map[int]int)
	for i := 0; i < n; i++ {
		rs[i] = Pair{W[i], X[i]}
		// x[i] < x[i+1]
		pos[X[i]] = i
	}

	sort.Slice(rs, func(i, j int) bool {
		return rs[i].first < rs[j].first
	})

	freq := NewBit(n)
	sum := NewBit(n)
	var res int
	for i, r := range rs {
		x := r.second
		j := pos[x]
		a := freq.Get(n - 1)
		b := freq.Get(j)
		c := sub(a, b)
		d := sum.Get(n - 1)
		e := sum.Get(j)
		f := sub(d, e)
		cur := add(sub(mul(b, x), e), sub(f, mul(c, x)))
		cur = mul(cur, pow(2, n-i-1))
		res = add(res, cur)
		ways := pow(2, i)
		freq.Add(j, ways)
		sum.Add(j, mul(x, ways))
	}

	ans := mul(res, inverse(pow(2, n)))
	return ans
}

type Bit struct {
	arr []int
}

func NewBit(n int) *Bit {
	arr := make([]int, n+1)
	return &Bit{arr}
}

func (b *Bit) Add(p int, v int) {
	p++
	for p < len(b.arr) {
		b.arr[p] = add(b.arr[p], v)
		p += p & -p
	}
}

func (b *Bit) Get(p int) int {
	p++
	var res int
	for p > 0 {
		res = add(res, b.arr[p])
		p -= p & -p
	}
	return res
}

type Pair struct {
	first  int
	second int
}
