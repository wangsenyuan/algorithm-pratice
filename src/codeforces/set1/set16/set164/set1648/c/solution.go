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

	tc := 1

	for tc > 0 {
		tc--
		n, m := readTwoNums(reader)
		A := readNNums(reader, n)
		B := readNNums(reader, m)
		res := solve(A, B)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
}

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
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

const MOD = 998244353

func modMul(a, b int) int {
	return int(int64(a) * int64(b) % MOD)
}

func modAdd(a, b int) int {
	a += b
	if a >= MOD {
		a -= MOD
	}
	return a
}

func pow(a, b int) int {
	r := 1
	for b > 0 {
		if b&1 == 1 {
			r = modMul(r, a)
		}
		a = modMul(a, a)
		b >>= 1
	}
	return r
}

func inverse(a int) int {
	return pow(a, MOD-2)
}

const N = 200010

var F [N]int
var I [N]int

func init() {
	F[0] = 1
	for i := 1; i < N; i++ {
		F[i] = modMul(i, F[i-1])
	}
	I[N-1] = inverse(F[N-1])

	for i := N - 2; i >= 0; i-- {
		I[i] = modMul(i+1, I[i+1])
	}
}

func nCr(n int, r int) int {
	if n < r || r < 0 {
		return 0
	}

	return modMul(F[n], modMul(I[r], I[n-r]))
}

func solve(S, T []int) int {
	sz, id := getNumPos(S, T)
	n := len(S)

	sc := make(map[int]int)
	for i := 0; i < n; i++ {
		S[i] = id[S[i]]
		sc[S[i]]++
	}

	sum := NewBIT(sz, 0, modAdd)
	cur := n
	// tot = S的不同的排列数
	tot := 1
	for k, v := range sc {
		sum.Update(k, v)
		tot = modMul(tot, nCr(cur, v))
		cur -= v
	}

	var ans int

	for i := 0; i <= n; i++ {
		if i == len(T) {
			// no way to get a smaller one
			break
		}
		if i == n {
			// i < len(T) and S == T for now
			ans = modAdd(ans, 1)
			break
		}
		x := id[T[i]]
		// 需要知道有多少个数字小于x
		y := sum.Get(0, x-1)
		c := modMul(tot, inverse(n-i))
		c = modMul(c, y)
		// 这里少了一部分是剩余部分的全排列
		// 假设x = 3, 选择了2
		// c = sum(tot / nCr(n, cnt[y]) * nCr(n-1, cnt[y] - 1) ) y < x
		//   = sum(tot / n * cnt[y])
		//   = tot / n * sum(cnt[y])

		ans = modAdd(ans, c)
		// tot = tot / nCr(n-i, sc[x]) * nCr(n-i-1, sc[x] - 1)
		tot = modMul(tot, inverse(n-i))
		tot = modMul(tot, sc[x])
		sc[x]--
		if sc[x] < 0 {
			// no x to make S[i] = T[i]
			break
		}
		sum.Update(x, -1)
	}

	return ans
}

type BIT struct {
	arr []int
	op  func(int, int) int
}

func NewBIT(n int, iv int, op func(int, int) int) *BIT {
	arr := make([]int, n+1)
	for i := 0; i <= n; i++ {
		arr[i] = iv
	}
	return &BIT{arr, op}
}

func (b *BIT) Update(p int, v int) {
	p++
	for ; p < len(b.arr); p += p & -p {
		b.arr[p] = b.op(b.arr[p], v)
	}
}

func (b *BIT) Get(iv int, p int) int {
	res := iv
	p++
	for ; p > 0; p -= p & -p {
		res = b.op(res, b.arr[p])
	}
	return res
}

func getNumPos(S []int, T []int) (int, []int) {
	nums := make(map[int]int)
	var X int
	for i := 0; i < len(S); i++ {
		nums[S[i]]++
		if S[i] > X {
			X = S[i]
		}
	}

	for i := 0; i < len(T); i++ {
		nums[T[i]]++
		if T[i] > X {
			X = T[i]
		}
	}

	arr := make([]int, 0, len(nums))

	for k := range nums {
		arr = append(arr, k)
	}

	sort.Ints(arr)

	id := make([]int, X+1)

	for i, x := range arr {
		id[x] = i
	}
	return len(arr), id
}
