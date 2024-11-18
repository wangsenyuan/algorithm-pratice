package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, k := readTwoNums(reader)
	h := readNNums(reader, n)
	res := solve(n, k, h)
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

func inverse(x int) int {
	return pow(x, mod-2)
}

func solve(n int, k int, h []int) int {
	F := make([]int, n+1)
	F[0] = 1
	for i := 1; i <= n; i++ {
		F[i] = mul(i, F[i-1])
	}

	I := make([]int, n+1)
	I[n] = inverse(F[n])
	for i := n - 1; i >= 0; i-- {
		I[i] = mul(i+1, I[i+1])
	}

	nCr := func(n int, r int) int {
		if r < 0 || n < r {
			return 0
		}
		return mul(F[n], mul(I[r], I[n-r]))
	}

	var t int

	for i := 0; i < n; i++ {
		if h[i] != h[(i+1)%n] {
			t++
		}
	}

	res := pow(k, n-t)

	i2 := inverse(2)

	var sum int
	for i := 0; i < t; i++ {
		// 选择i个位数，不改变分数（每个有k-2个选择）
		// 共t个位置，所有有 nCr(t, i)种选择
		tmp := mul(pow(k-2, i), nCr(t, i))
		if (t-i)%2 == 1 {
			tmp = mul(tmp, pow(2, t-i-1))
		} else {
			tmp2 := sub(pow(2, t-i), nCr(t-i, (t-i)/2))
			tmp2 = mul(tmp2, i2)
			tmp = mul(tmp, tmp2)
		}
		sum = add(sum, tmp)
	}

	return mul(res, sum)
}
