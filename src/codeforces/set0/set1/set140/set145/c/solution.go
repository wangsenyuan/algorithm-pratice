package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
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

func process(reader *bufio.Reader) int {
	n, k := readTwoNums(reader)
	a := readNNums(reader, n)
	return solve(k, a)
}

const mod = 1000000007

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func mul(num ...int) int {
	r := 1
	for _, x := range num {
		r *= x
		r %= mod
	}
	return r
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

func inv(num int) int {
	return pow(num, mod-2)
}

func solve(k int, a []int) int {

	n := len(a)
	F := make([]int, n+1)
	F[0] = 1
	for i := 1; i <= n; i++ {
		F[i] = mul(i, F[i-1])
	}
	I := make([]int, n+1)
	I[n] = inv(F[n])

	for i := n - 1; i >= 0; i-- {
		I[i] = mul(i+1, I[i+1])
	}

	nCr := func(n int, r int) int {
		if r < 0 || n < r {
			return 0
		}
		return mul(F[n], I[n-r], I[r])
	}

	var cnt int
	var arr []int

	for _, num := range a {
		if checkLucky(num) {
			arr = append(arr, num)
		} else {
			cnt++
		}
	}

	sort.Ints(arr)
	// dp[i] = dp[i] + dp[i-1] * (freq[i])
	// 问题是这样子的复杂性是 n * k 的
	// 等等， 就是 lucky数的个数，不可能很多, <= 1e9,的最大的是7*1e8
	// 一共8个位置，每个位置2个可能性， 1<<8， 也就是256个数
	// 也就是说，不是 n * k, 而是 256 * k
	dp := make([]int, k+1)
	dp[0] = 1
	var w int
	for i := 0; i < len(arr); {
		j := i
		for i < len(arr) && arr[i] == arr[j] {
			i++
		}
		d := i - j
		w++
		for x := min(w, k); x > 0; x-- {
			dp[x] = add(dp[x], mul(d, dp[x-1]))
		}
	}

	var res int

	for u := 0; u <= k; u++ {
		res = add(res, mul(dp[u], nCr(cnt, k-u)))
	}

	return res
}

func checkLucky(num int) bool {
	for num > 0 {
		r := num % 10
		if r != 4 && r != 7 {
			return false
		}
		num /= 10
	}
	return true
}
