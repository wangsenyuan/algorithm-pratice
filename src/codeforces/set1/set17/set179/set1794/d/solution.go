package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	a := readNNums(reader, n*2)
	res := solve(a)
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

const mod = 998244353

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func mul(a, b int) int {
	return a * b % mod
}

const MAX = 1_000_000 + 10
const N = 2_000 + 100

var lpf [MAX]int
var F [N]int
var IF [N]int

func init() {
	var primes []int
	for i := 2; i < MAX; i++ {
		if lpf[i] == 0 {
			primes = append(primes, i)
			lpf[i] = i
		}
		for j := 0; j < len(primes); j++ {
			if i*primes[j] >= MAX {
				break
			}
			lpf[i*primes[j]] = primes[j]
			if i%primes[j] == 0 {
				break
			}
		}
	}

	F[0] = 1
	for i := 1; i < N; i++ {
		F[i] = mul(i, F[i-1])
	}

	IF[N-1] = pow(F[N-1], mod-2)

	for i := int(N - 2); i >= 0; i-- {
		IF[i] = mul(i+1, IF[i+1])
	}
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

func nCr(n int, r int) int {
	if n < r {
		return 0
	}
	return mul(F[n], mul(IF[r], IF[n-r]))
}

type Pair struct {
	first  int
	second int
}

func solve(a []int) int {
	sort.Ints(a)

	var arr []Pair
	n := len(a)
	var prime_cnt int
	for i := 0; i < n; {
		j := i
		for i < n && a[i] == a[j] {
			i++
		}
		arr = append(arr, Pair{i - j, a[j]})
		if lpf[a[j]] == a[j] {
			prime_cnt++
		}
	}

	if prime_cnt < n/2 {
		return 0
	}

	n /= 2

	dp := make([][]int, 2)
	for i := 0; i < 2; i++ {
		dp[i] = make([]int, n+1)
	}
	dp[0][0] = 1

	var sum int
	for i, cur := range arr {
		clear(dp[(i+1)&1])
		num := cur.second
		cnt := cur.first
		if lpf[num] == num {
			// num是质数， 既可以做指数，也可以做底数
			for j := n; j >= 0; j-- {
				var tmp int
				// 将num放在底数上
				if j > 0 && sum >= j-1 && n-(sum-(j-1)) >= cnt-1 {
					// tmp = mul(dp[i&1][j-1], (n - (j - 1)))
					// 剩余指数的位置 = n - (sum - (j - 1))
					tmp = mul(dp[i&1][j-1], nCr(n-(sum-(j-1)), cnt-1))
				}
				// 将num全部放在指数位置上
				if sum >= j && n-(sum-j) >= cnt {
					tmp = add(tmp, mul(dp[i&1][j], nCr(n-(sum-j), cnt)))
				}
				dp[(i+1)&1][j] = tmp
			}
		} else {
			// 只能放在指数位置上
			for j := min(n, sum); j >= max(0, cnt+sum-n); j-- {
				dp[(i+1)&1][j] = mul(dp[i&1][j], nCr(n-(sum-j), cnt))
			}
		}

		sum += cnt
	}

	return dp[len(arr)&1][n]
}
