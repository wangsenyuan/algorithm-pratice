package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	num := readString(reader)
	k := readNum(reader)

	res := solve(num, k)

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

const mod = 1e9 + 7

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func sub(a, b int) int {
	return add(a, mod-b)
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

const N = 1001

var F [N]int
var I [N]int
var W [N]int
var MK int

func init() {
	F[0] = 1
	for i := 1; i < N; i++ {
		F[i] = mul(i, F[i-1])
	}

	I[N-1] = pow(F[N-1], mod-2)
	for i := N - 2; i >= 0; i-- {
		I[i] = mul(i+1, I[i+1])
	}
	W[0] = -100
	W[1] = 0
	for i := 2; i < N; i++ {
		var cnt int
		for tmp := i; tmp > 0; tmp -= tmp & -tmp {
			cnt++
		}
		W[i] = W[cnt] + 1
		MK = max(MK, W[i])
	}
}

func nCr(n, r int) int {
	if n < r || r < 0 {
		return 0
	}
	return mul(F[n], mul(I[r], I[n-r]))
}

func solve(num string, k int) int {
	if k > MK+2 {
		return 0
	}
	if k == 0 {
		// only 1
		return 1
	}

	var res int
	n := len(num)

	count := func(i int, cnt int) int {
		// 第i位放置0
		// 假设有面放置x个1，w[y+cnt] = k 才可以
		// 最多只能放 n - i - 1 个1
		var tmp int
		for x := 0; x < n-i; x++ {
			if W[x+cnt] == k-1 {
				tmp = add(tmp, nCr(n-i-1, x))
			}
		}
		return tmp
	}

	var cnt int
	for i := 0; i < n; i++ {
		if num[i] == '1' {
			// 当前位置放置0
			res = add(res, count(i, cnt))
			cnt++
		}
	}

	if W[cnt] == k-1 {
		res = add(res, 1)
	}
	if k == 1 {
		// 那么1肯定会包含在答案里，但是它的操作次数是0
		res = sub(res, 1)
	}

	return res
}

func bruteForce(num string, k int) int {
	var x int
	for i := 0; i < len(num); i++ {
		x *= 2
		if num[i] == '1' {
			x++
		}
	}

	var res int
	for i := 1; i <= x; i++ {
		if W[i] == k {
			res++
		}
	}
	return res
}
