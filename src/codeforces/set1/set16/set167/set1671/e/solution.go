package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	s := readString(reader)
	res := solve(n, s)
	fmt.Println(res)
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

func solve(n int, s string) int {
	// N := (1 << n) - 1
	// 如果x的左子树和右子树完全一样，则交换后还是一样的
	// 但左子树和右子树，即使不一样的情况下，也可能在交换后，产生重复的结果
	// 对于x节点，它的children的size = size(x), x的位置时确定的
	// n = 2, N = 3
	N := (1 << n) - 1

	var dfs func(u int) (int64, string)

	dfs = func(u int) (int64, string) {
		l := u << 1
		r := l + 1

		if r > N {
			// a leaf
			return 1, s[u-1 : u]
		}

		a, ak := dfs(l)
		b, bk := dfs(r)

		c := modMul(a, b, MOD)

		if ak != bk {
			c = modAdd(c, c, MOD)
		}
		ck := min_string(ak+s[u-1:u]+bk, bk+s[u-1:u]+ak)
		return c, ck
	}
	ans, _ := dfs(1)

	return int(ans)
}
func min_string(a, b string) string {
	if a < b {
		return a
	}
	return b
}

func solve1(n int, s string) int {
	// N := (1 << n) - 1
	// 如果x的左子树和右子树完全一样，则交换后还是一样的
	// 但左子树和右子树，即使不一样的情况下，也可能在交换后，产生重复的结果
	// 对于x节点，它的children的size = size(x), x的位置时确定的
	// n = 2, N = 3
	N := (1 << n) - 1

	hash := make([]Key, N+1)

	bs := Key{[...]int64{B[0], B[1], B[2]}}

	var dfs func(u int) int64

	dfs = func(u int) int64 {
		l := u << 1
		r := l + 1

		hash[u] = NewKey(int(s[u-1]-'A') + 1)

		if r > N {
			// a leaf
			return 1
		}

		a := dfs(l)
		b := dfs(r)
		c := modMul(a, b, MOD)

		if hash[l] != hash[r] {
			c = modAdd(c, c, MOD)
		}
		hash[u] = hash[u].Add(bs.Pow(hash[l])).Add(bs.Pow(hash[r]))
		return c
	}

	return int(dfs(1))
}

const MOD = 998244353

var B = []int64{100007, 1000009, 10013}
var M = []int64{100000000007, 100000000009, 1000000000013}

type Key struct {
	arr [3]int64
}

func NewKey(v int) Key {
	var arr [3]int64
	for i := 0; i < 3; i++ {
		arr[i] = int64(v)
	}
	return Key{arr}
}

func (this Key) Add(that Key) Key {
	var arr [3]int64
	for i := 0; i < 3; i++ {
		arr[i] = modAdd(this.arr[i], that.arr[i], M[i])
	}
	return Key{arr}
}

func modAdd(a, b int64, mod int64) int64 {
	a += b
	if a >= mod {
		a -= mod
	}
	if a < 0 {
		a += mod
	}
	return a
}

func modMul(a, b, mod int64) int64 {
	return a * b % mod
}

func pow(a, b int64, mod int64) int64 {
	var r int64 = 1
	for b > 0 {
		if b&1 == 1 {
			r = modMul(r, a, mod)
		}
		a = modMul(a, a, mod)
		b >>= 1
	}
	return r
}

func (this Key) Pow(that Key) Key {
	var arr [3]int64
	for i := 0; i < 3; i++ {
		arr[i] = pow(this.arr[i], that.arr[i], M[i])
	}
	return Key{arr}
}
