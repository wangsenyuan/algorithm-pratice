package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, s := readTwoNums(reader)
	a := readNNums(reader, n)
	res := solve(a, s)

	fmt.Println(res)
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadBytes('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return string(s[:i])
		}
	}
	return string(s)
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
	res := 1
	for b > 0 {
		if b&1 == 1 {
			res = mul(res, a)
		}
		a = mul(a, a)
		b >>= 1
	}
	return res
}

var I [30]int

func init() {
	for i := 0; i < 30; i++ {
		I[i] = pow(i, mod-2)
	}
}

func comb(n int, r int) int {
	if n < r {
		return 0
	}
	if n-r < r {
		r = n - r
	}
	n %= mod
	res := 1
	for i := 0; i < r; i++ {
		res = mul(res, n-i)
		res = mul(res, I[i+1])
	}

	return res
}

func solve(a []int, s int) int {
	n := len(a)

	var ans int

	for state := 0; state < 1<<n; state++ {
		s2 := s
		for i := 0; i < n; i++ {
			if (state>>i)&1 == 1 {
				// 预先在boxi种放入 a[i]+1个（保证不符合条件)
				s2 -= a[i] + 1
			}
		}
		// 剩下的，随便放入n个盒子
		tmp := comb(s2+n-1, n-1)
		if bits.OnesCount(uint(state))%2 == 1 {
			ans = sub(ans, tmp)
		} else {
			ans = add(ans, tmp)
		}
	}

	return ans
}

func solve1(a []int, s int) int {

	n := len(a)

	var gen func(id int, pw int, coff int)

	mem := make(map[int]int)

	gen = func(id int, pw int, coff int) {
		if id == n {
			val := mem[pw]
			val = add(val, coff)
			mem[pw] = val
			return
		}
		gen(id+1, pw, coff)
		ncoff := coff * -1
		if ncoff < 0 {
			ncoff += mod
		}
		gen(id+1, pw+a[id]+1, ncoff)
	}

	gen(0, 0, 1)

	var ans int

	for p, v := range mem {
		rem := s - p
		if rem >= 0 {
			ans = add(ans, mul(comb(n+rem-1, rem), v))
		}
	}

	return ans
}
