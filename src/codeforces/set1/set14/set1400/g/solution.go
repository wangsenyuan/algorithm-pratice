package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, m := readTwoNums(reader)
	L, R := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		L[i], R[i] = readTwoNums(reader)
	}
	E := make([][]int, m)
	for i := 0; i < m; i++ {
		E[i] = readNNums(reader, 2)
	}
	fmt.Println(solve(n, m, L, R, E))
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

const N = 300043
const MOD = 998244353

var fact [N]int64
var ifact [N]int64

func init() {
	fact[0] = 1
	for i := 1; i < N; i++ {
		fact[i] = int64(i) * fact[i-1] % MOD
	}
	ifact[N-1] = inverse(fact[N-1])

	for i := N - 2; i >= 0; i-- {
		ifact[i] = int64(i+1) * ifact[i+1] % MOD
	}
}

func inverse(a int64) int64 {
	return pow(a, MOD-2)
}

func pow(a int64, b int64) int64 {
	var r int64 = 1
	for b > 0 {
		if b&1 == 1 {
			r = r * a % MOD
		}
		a = a * a % MOD
		b >>= 1
	}
	return r
}

func nCr(n, r int) int64 {
	if n < 0 || r < 0 || n < r {
		return 0
	}
	res := fact[n] * ifact[r] % MOD
	res = res * ifact[n-r] % MOD
	return res
}

func solve(n int, m int, L []int, R []int, E [][]int) int {
	cnt := make([]int, n+2)
	for i := 0; i < n; i++ {
		cnt[L[i]]++
		cnt[R[i]+1]--
	}

	for i := 0; i < n; i++ {
		cnt[i+1] += cnt[i]
	}

	P := make([][]int64, 2*m+1)
	for i := 0; i < len(P); i++ {
		P[i] = make([]int64, n+1)
	}

	for i := 1; i <= n; i++ {
		for j := 0; j <= 2*m; j++ {
			P[j][i] = P[j][i-1] + nCr(cnt[i]-j, i-j)
			if P[j][i] >= MOD {
				P[j][i] -= MOD
			}
		}
	}

	var ans int64

	for mask := 0; mask < 1<<m; mask++ {
		var sign int64 = 1
		used := make(map[int]bool)
		for i := 0; i < m; i++ {
			if mask&(1<<i) > 0 {
				u, v := E[i][0], E[i][1]
				u--
				v--
				used[u] = true
				used[v] = true
				sign = sign * (MOD - 1) % MOD
			}
		}
		l, r := 1, n
		for k := range used {
			l = max(l, L[k])
			r = min(r, R[k])
		}
		if r < l {
			continue
		}
		ans = add(ans, mul(sign, add(P[len(used)][r], -P[len(used)][l-1])))
	}

	return int(ans)
}

func add(a, b int64) int64 {
	return ((a+b)%MOD + MOD) % MOD
}

func mul(a, b int64) int64 {
	return (a * b) % MOD
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
