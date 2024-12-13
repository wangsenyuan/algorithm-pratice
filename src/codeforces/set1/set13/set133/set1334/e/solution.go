package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	res := process(reader)
	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d\n", x))
	}
	fmt.Print(buf.String())
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

func process(reader *bufio.Reader) []int {
	D := readNum(reader)
	n := readNum(reader)
	queries := make([][]int, n)
	for i := range n {
		queries[i] = readNNums(reader, 2)
	}
	return solve(D, queries)
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

var F [100]int
var I [100]int

func init() {
	F[0] = 1
	for i := 1; i < 100; i++ {
		F[i] = mul(i, F[i-1])
	}
	I[99] = pow(F[99], mod-2)
	for i := 98; i >= 0; i-- {
		I[i] = mul(i+1, I[i+1])
	}
}

func solve(D int, queries [][]int) []int {
	var primes []int

	for i := 2; i <= D/i; i++ {
		if D%i == 0 {
			primes = append(primes, i)
			for D%i == 0 {
				D /= i
			}
		}
	}

	if D > 1 {
		primes = append(primes, D)
	}

	find := func(u int, v int) int {
		var up int
		var dw int
		ans := 1
		for _, i := range primes {
			var cnt int
			for u%i == 0 {
				cnt++
				u /= i
			}
			for v%i == 0 {
				cnt--
				v /= i
			}
			if cnt > 0 {
				up += cnt
				ans = mul(ans, I[cnt])
			} else {
				dw -= cnt
				ans = mul(ans, I[-cnt])
			}
		}
		ans = mul(ans, F[up])
		ans = mul(ans, F[dw])
		return ans
	}

	ans := make([]int, len(queries))

	for i, cur := range queries {
		ans[i] = find(cur[0], cur[1])
	}

	return ans
}
