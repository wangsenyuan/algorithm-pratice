package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	S := make([]string, n)
	for i := 0; i < n; i++ {
		S[i] = readString(reader)
	}
	res := solve(S)
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

func readLenNums(r *bufio.Reader) []int {
	s, _ := r.ReadBytes('\n')
	var n int
	pos := readInt(s, 0, &n) + 1
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		pos = readInt(s, pos, &arr[i]) + 1
	}
	return arr
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

func modSub(a, b int) int {
	return modAdd(a, MOD-b)
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

func solve(S []string) int64 {
	n := len(S)
	// n <= 23
	// f([t1....tm]) = number of strings, being sub-seq of at least one of ti
	// g([t1...tm]) = number of strings, being sub-seq of all of ti
	// f = state over m inclusion-exclusion such things
	// how to get g?
	// len(S[i]) <= 10000, S[i] is sorted
	P := make([][]int, n)
	for i := 0; i < n; i++ {
		P[i] = convert(S[i])
	}
	// g = prod[0..26] pow(2, min(cnt[i])) - 1
	T := 1 << n
	H := make([]int, T)
	cnt := make([]int, 26)
	for i := 0; i < 26; i++ {
		cnt[i] = INF
	}
	for state := 1; state < T; state++ {

		for i := 0; i < n; i++ {
			if (state>>i)&1 == 1 {
				for j := 0; j < 26; j++ {
					cnt[j] = min(cnt[j], P[i][j])
				}
			}
		}
		H[state] = 1
		for i := 0; i < 26; i++ {
			H[state] = modMul(H[state], cnt[i]+1)
			cnt[i] = INF
		}
	}

	flipAll(H, n)

	for i := n - 1; i >= 0; i-- {
		for j := (1 << n) - 1; j >= 0; j-- {
			if (j>>i)&1 == 1 {
				H[j] = modSub(H[j], H[j^(1<<i)])
			}
		}
	}

	flipAll(H, n)
	var sum int

	for i := 0; i < T; i++ {
		sum = modAdd(sum, H[i])
	}

	for i := 0; i < n; i++ {
		for j := 0; j < 1<<n; j++ {
			if (j>>i)&1 == 0 {
				H[j^(1<<i)] = modAdd(H[j^(1<<i)], H[j])
			}
		}
	}

	res := make([]int, T)

	for i := 0; i < T; i++ {
		res[i] = modSub(sum, H[(T-1)^i])
	}

	var ans int64

	for i := 0; i < T; i++ {
		var c, s int
		for j := 0; j < n; j++ {
			if i&(1<<j) > 0 {
				c++
				s += j + 1
			}
		}
		ans ^= int64(res[i]) * int64(c) * int64(s)
	}

	return ans
}

func flipAll(dp []int, n int) {
	mask := (1 << n) - 1
	for i := 0; i < (1 << (n - 1)); i++ {
		j := i ^ mask
		dp[i], dp[j] = dp[j], dp[i]
	}
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

const INF = 1 << 30

func convert(s string) []int {
	res := make([]int, 26)

	for i, j := 0, 0; i < 26 && j < len(s); i++ {
		for j < len(s) && int(s[j]-'a') == i {
			j++
			res[i]++
		}
	}
	return res
}
