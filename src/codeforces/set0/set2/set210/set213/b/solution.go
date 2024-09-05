package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n := readNum(reader)
	a := readNNums(reader, 10)
	res := solve(n, a)
	fmt.Println(res)
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

const mod = 1000000007

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

func inverse(n int) int {
	return pow(n, mod-2)
}

const N = 110

var F [N]int
var I [N]int

func init() {
	F[0] = 1
	for i := 1; i < N; i++ {
		F[i] = mul(i, F[i-1])
	}
	I[N-1] = inverse(F[N-1])
	for i := N - 2; i >= 0; i-- {
		I[i] = mul(i+1, I[i+1])
	}
}

func nCr(n int, r int) int {
	if r < 0 || n < r {
		return 0
	}
	return mul(F[n], mul(I[r], I[n-r]))
}

func solve(n int, a []int) int {
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, 10)
		for j := 0; j < 10; j++ {
			dp[i][j] = -1
		}
	}

	var dfs func(n int, x int) (ans int)

	dfs = func(n int, x int) (ans int) {
		if x == 9 {
			if n >= a[x] {
				return 1
			}
			return 0
		}
		if dp[n][x] >= 0 {
			return dp[n][x]
		}
		defer func() {
			dp[n][x] = ans
		}()

		n2 := n
		if x == 0 {
			n2--
		}

		for i := a[x]; i <= n; i++ {
			tmp := dfs(n-i, x+1)
			tmp = mul(tmp, nCr(n2, i))
			ans = add(ans, tmp)
		}
		return ans
	}

	var res int
	for i := 1; i <= n; i++ {
		res = add(res, dfs(i, 0))
	}

	return res
}

func bruteForce(n int, a []int) []int {
	arr := make([]int, n)

	var res []int

	cnt := make([]int, 10)

	for {
		if arr[n-1] < 9 {
			arr[n-1]++
		} else {
			arr[n-1] = 0
			carry := 1
			for i := n - 2; i >= 0 && carry > 0; i-- {
				arr[i] += carry
				carry = arr[i] / 10
				arr[i] %= 10
			}
			if carry == 1 {
				break
			}
		}

		clear(cnt)
		var num int
		for i := 0; i < n; i++ {
			num = num*10 + arr[i]
			if num > 0 {
				cnt[arr[i]]++
			}
		}
		if num == 0 {
			cnt[0]++
		}
		ok := true
		for i := 0; i < 10; i++ {
			if cnt[i] < a[i] {
				ok = false
				break
			}
		}
		if ok {
			res = append(res, num)
		}
	}
	return res
}
