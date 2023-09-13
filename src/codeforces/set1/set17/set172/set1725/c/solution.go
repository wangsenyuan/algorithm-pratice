package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, m := readTwoNums(reader)
	d := readNNums(reader, n)
	res := solve(m, d)
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

func sub(a, b int) int {
	return add(a, mod-b)
}
func mul(a, b int) int {
	return int(int64(a) * int64(b) % mod)
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

const M = 10_00_000

var F [M]int
var I [M]int

func init() {
	F[0] = 1
	for i := 1; i < M; i++ {
		F[i] = mul(i, F[i-1])
	}
	I[M-1] = pow(F[M-1], mod-2)
	for i := M - 2; i >= 0; i-- {
		I[i] = mul(i+1, I[i+1])
	}
}

func nCr(n int, r int) int {
	if r < 0 || n < r {
		return 0
	}
	return mul(F[n], mul(I[r], I[n-r]))
}

func solve(m int, D []int) int {
	// 在直径上的两点 + 任何一点，有相同的颜色
	// 总数是 m ** n
	// - 不符合条件的部分
	// 但是怎么排除那些重复的部分呢？
	// 比如两条直径，在考虑a的时候，b有贡献，考虑b的时候a也有贡献
	var sum int64

	n := len(D)

	for i := 0; i < n; i++ {
		sum += int64(D[i])
	}

	if sum%2 == 1 || n == 1 {
		// 不可能出现直角三角形
		return pow(m, n)
	}
	sum /= 2

	var mid int
	var dist int64

	cnt := make([]int, 2)
	for l, r := 0, mid; l < n; l++ {
		for dist+int64(D[r%n]) <= sum {
			dist += int64(D[r%n])
			r++
		}
		if dist == sum {
			cnt[0]++
		}
		dist -= int64(D[l])

	}
	cnt[0] /= 2
	cnt[1] = n - 2*cnt[0]

	var res int

	for i := 0; i <= min(cnt[0], m); i++ {
		// you i 对直径有相同的颜色（每一对的颜色不一样）
		tmp := nCr(cnt[0], i)
		// 每对的颜色必须不同，从m种选出i种
		tmp = mul(tmp, nCr(m, i))
		// 顺序也有关系
		tmp = mul(tmp, F[i])
		// 剩下的直径对，有 m - i种颜色可以选择
		tmp = mul(tmp, pow(mul(m-i, m-i-1), cnt[0]-i))
		tmp = mul(tmp, pow(m-i, cnt[1]))

		res = add(res, tmp)
	}

	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
