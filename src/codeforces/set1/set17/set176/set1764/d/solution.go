package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, p := readTwoNums(reader)

	res := solve(n, p)

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

func solve(n int, p int) int {
	add := func(a, b int) int {
		a += b
		if a >= p {
			a -= p
		}
		return a
	}

	sub := func(a, b int) int {
		return add(a, p-b)
	}

	mul := func(a, b int) int {
		return a * b % p
	}

	pow := func(a, b int) int {
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

	F := make([]int, n+1)
	I := make([]int, n+1)
	F[0] = 1
	for i := 1; i <= n; i++ {
		F[i] = mul(i, F[i-1])
	}

	I[n] = pow(F[n], p-2)

	for i := n - 1; i >= 0; i-- {
		I[i] = mul(i+1, I[i+1])
	}

	nCr := func(n int, r int) int {
		if n < r || r < 0 {
			return 0
		}
		return mul(F[n], mul(I[r], I[n-r]))
	}

	forbid := make([]int, n+1)
	for i := n / 2; i < n-1; i++ {
		forbid[i] = mul(nCr(n-n/2-1, n-i-1), mul(F[i], n))
	}
	forbid[n-1] = F[n]
	var ans int

	for i := n / 2; i < n; i++ {
		ans = add(ans, sub(forbid[i], mul(forbid[i-1], n-i+1)))
	}

	return ans
}

func solve1(n int, p int) int {
	add := func(a, b int) int {
		a += b
		if a >= p {
			a -= p
		}
		return a
	}

	mul := func(a, b int) int {
		return a * b % p
	}

	pow := func(a, b int) int {
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

	F := make([]int, n+1)
	I := make([]int, n+1)
	F[0] = 1
	for i := 1; i <= n; i++ {
		F[i] = mul(i, F[i-1])
	}

	I[n] = pow(F[n], p-2)

	for i := n - 1; i >= 0; i-- {
		I[i] = mul(i+1, I[i+1])
	}

	nCr := func(n int, r int) int {
		if n < r || r < 0 {
			return 0
		}
		return mul(F[n], mul(I[r], I[n-r]))
	}

	var res int

	t := n / 2

	for i := t; i <= n-2; i++ {
		var cur int
		for j := 0; j <= n-2-i; j++ {
			cur = add(cur, mul(nCr(n-i-2, j), F[i+j-1]))
		}
		cur = mul(cur, 2*t-i)
		res = add(res, cur)
	}

	res = mul(res, n)

	if n%2 == 0 {
		res = add(res, mul(n, F[n-2]))
	}

	return res
}
