package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)

	res := solve(n)

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

const mod = 1_000_000_000 + 7

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func mul(a, b int) int {
	a %= mod
	b %= mod
	return a * b % mod
}

func solve(n int) int {
	// a + b + c = n
	// sum of lcm(c, gcd(a, b))
	// 如果 gcd(a, b) = a ( b % a = 0)
	// lcm(c, a) 在a确定的时候，c也是确定的
	//   gcd(a, b) = d, 比如 a = 4, b = 6, d = 2
	//  如果c确定了, 那么 a + b = n - c
	//   然后d 显然是 n - c 的因子
	// d 确定的时候， a和b吧是否也是确定的？
	// (x + y) * d = (n - c) 且 gcd(x, y) = 1
	// 1) x + y = (n - c) / d
	// 2) gcd(x, y) = 1
	// 需要直到有多少对 x, y 满足 1 和 2；
	// 这个是那个欧拉数？ 不考虑1，有多少个数x < y, 且 gcd(x, y) = 1
	// 假设 set(y) = x < y, 且 gcd(x, y) = 1 的数
	//  Euler's totient function phi
	phi := make([]int, n+1)
	phi[1] = 1
	for i := 2; i <= n; i++ {
		phi[i] = i - 1
	}
	for i := 2; i <= n; i++ {
		for j := 2 * i; j <= n; j += i {
			phi[j] -= phi[i]
		}
	}

	lpf := make([]int, n+1)
	var primes []int
	for i := 2; i <= n; i++ {
		if lpf[i] == 0 {
			lpf[i] = i
			primes = append(primes, i)
		}
		for j := 0; j < len(primes); j++ {
			if i*primes[j] > n {
				break
			}
			lpf[i*primes[j]] = primes[j]
			if i%primes[j] == 0 {
				break
			}
		}
	}

	var ans int

	for c := 1; c <= n-2; c++ {
		x := n - c
		// gcd of ab
		// 1 & x
		tmp := mul(lcm(c, 1), phi[x])

		for d := 2; d*d <= x; d++ {
			if x%d == 0 {
				tmp = add(tmp, mul(lcm(c, d), phi[x/d]))
				if d*d != x {
					p := x / d
					tmp = add(tmp, mul(lcm(c, p), phi[x/p]))
				}
			}
		}

		ans = add(ans, tmp)
	}

	return ans
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}
