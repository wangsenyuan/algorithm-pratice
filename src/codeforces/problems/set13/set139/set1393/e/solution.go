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
		s, _ := reader.ReadBytes('\n')
		var j int
		for j < len(s) {
			if s[j] < 'a' || s[j] > 'z' {
				break
			}
			j++
		}
		S[i] = string(s[:j])
	}
	fmt.Println(solve(n, S))
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

var MOD = [...]int64{1000000007, 1000000009}

const MAX_C = 1000010

var P [2]int64
var R [2]int64

var p [2][MAX_C]int64

var nxt [MAX_C]int

var A [MAX_C]int

var dp [2][MAX_C]int64

func init() {
	P[0] = 29
	P[1] = 31
	R[0] = pow(P[0], MOD[0]-2, MOD[0])
	R[1] = pow(P[1], MOD[1]-2, MOD[1])
	for j := 0; j < 2; j++ {
		p[j][0] = 1
		for i := 1; i < MAX_C; i++ {
			p[j][i] = (P[j] * p[j][i-1]) % MOD[j]
		}
	}
}

func pow(a, b, mod int64) int64 {
	var res int64 = 1
	for b > 0 {
		if b&1 == 1 {
			res = (res * a) % mod
		}
		a = (a * a) % mod
		b >>= 1
	}
	return res
}

func solve(n int, S []string) int {
	H := make([][][]int64, n)
	sorted := make([][]int, n)
	for i := 0; i < n; i++ {
		H[i] = make([][]int64, 2)
		for j := 0; j < 2; j++ {
			H[i][j] = make([]int64, 0, len(S[i])+1)
			H[i][j] = append(H[i][j], 0)
			for pos := 0; pos < len(S[i]); pos++ {
				H[i][j] = append(H[i][j], (H[i][j][pos]+p[j][pos]*int64(S[i][pos]-'a'+1))%MOD[j])
			}
		}

		nxt[len(S[i])-1] = len(S[i]) - 1

		for pos := len(S[i]) - 2; pos >= 0; pos-- {
			if S[i][pos] != S[i][pos+1] {
				nxt[pos] = pos + 1
			} else {
				nxt[pos] = nxt[pos+1]
			}
		}
		var l, r = 0, len(S[i]) - 1
		for pos := 0; pos < len(S[i]); pos++ {
			if S[i][nxt[pos]] <= S[i][pos] {
				A[l] = pos
				l++
			} else {
				A[r] = pos
				r--
			}
		}
		sorted[i] = make([]int, 0, len(S[i])+1)
		for j := 0; j < len(S[i]); j++ {
			sorted[i] = append(sorted[i], A[j])
			if A[j] == len(S[i])-1 {
				sorted[i] = append(sorted[i], len(S[i]))
			}
		}
	}

	for i := 0; i <= len(S[0]); i++ {
		dp[0][i] = 1
	}

	var getHash func(t int, i int, x int, l int) int64

	getHash = func(t int, i int, x int, l int) int64 {
		if l <= x {
			return H[i][t][l]
		}
		res1 := H[i][t][x]
		res2 := (H[i][t][l+1] - H[i][t][x+1] + MOD[t]) % MOD[t]
		res2 *= R[t]
		res2 %= MOD[t]

		return (res1 + res2) % MOD[t]
	}

	getHashPair := func(i, x, l int) Pair {
		first := getHash(0, i, x, l)
		second := getHash(1, i, x, l)
		return Pair{first, second}
	}

	var getAt func(i, x, l int) byte

	getAt = func(i, x, l int) byte {
		if l < x {
			return S[i][l]
		}
		if l+1 < len(S[i]) {
			return S[i][l+1]
		}
		return ' '
	}

	check := func(i int, x int, j int, y int) bool {
		len1 := len(S[i])
		len2 := len(S[j])
		if x != len1 {
			len1--
		}
		if y != len2 {
			len2--
		}
		var left, right = 0, min(len1, len2) + 1
		for right-left > 1 {
			mid := (left + right) / 2
			if getHashPair(i, x, mid) == getHashPair(j, y, mid) {
				left = mid
			} else {
				right = mid
			}
		}

		return getAt(i, x, left) >= getAt(j, y, left)
	}

	for i := 1; i < n; i++ {
		oks := i % 2
		var ptr int
		var sum int64
		for k := 0; k < len(sorted[i]); k++ {
			for ptr < len(sorted[i-1]) && check(i, sorted[i][k], i-1, sorted[i-1][ptr]) {
				sum += dp[oks^1][ptr]
				if sum >= MOD[0] {
					sum -= MOD[0]
				}
				ptr++
			}
			dp[oks][k] = sum
		}
	}

	var ans int64

	for i := 0; i <= len(S[n-1]); i++ {
		ans += dp[(n-1)%2][i]
		if ans >= MOD[0] {
			ans -= MOD[0]
		}
	}
	return int(ans)
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

type Pair struct {
	first, second int64
}
