package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	price := readString(reader)
	fmt.Println(solve(price))
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

const MOD = 1e9 + 7

func modMul(a, b int) int {
	return int(int64(a) * int64(b) % MOD)
}

func modDiv(a, b int) int {
	return modMul(a, inverse(b))
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

func solve(price string) int {
	// 假设删除的是[i..j]
	// 那么这时候的数字应该是，就是 suf[j...]
	// + (suf[0] - suf[i]) / pow(10, i) * pow(10, j)
	// 比如 123456 删除345后，126
	// = 6 + (123456 - 3456) / pow(10, 4) * pow(10, 1)
	// f(i, j) = (g(n) - g(i)) / pow(10, i) * pow(10, j) + g(j)
	// = (g(n) - g(i)) * pow(10, j) / pow(10, i) + g(j)
	// x(i) = (g(n) - g(i)) / pow(10, i)
	// f(i, j) = x(i) * pow(10, j) + g(j)
	// 所有前面的j考虑进去的话
	// = x(i) sum(pow(10, 0) + pow(10, 1) ...) + sum(g)
	price = reverse(price)

	n := len(price)
	pw := make([]int, n+1)
	pw[0] = 1
	for i := 1; i <= n; i++ {
		pw[i] = modMul(10, pw[i-1])
	}

	num := make([]int, n+1)
	for i := 1; i <= n; i++ {
		x := int(price[i-1] - '0')
		num[i] = modAdd(modMul(x, pw[i-1]), num[i-1])
	}
	var res int
	s10 := 1
	var sum int
	for i := 1; i <= n; i++ {
		x := modDiv(modSub(num[n], num[i]), pw[i])
		tmp := modMul(x, s10)
		tmp = modAdd(tmp, sum)
		res = modAdd(res, tmp)
		s10 = modAdd(s10, pw[i])
		sum = modAdd(sum, num[i])
	}

	return res
}

func reverse(s string) string {
	buf := []byte(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		buf[i], buf[j] = buf[j], buf[i]
	}
	return string(buf)
}
