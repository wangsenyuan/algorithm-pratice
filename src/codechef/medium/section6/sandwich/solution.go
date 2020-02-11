package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(scanner *bufio.Scanner) (a int) {
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &a)
	return
}

func readTwoNums(scanner *bufio.Scanner) (a int, b int) {
	res := readNNums(scanner, 2)
	a, b = res[0], res[1]
	return
}

func readNNums(scanner *bufio.Scanner, n int) []int {
	res := make([]int, n)
	x := 0
	scanner.Scan()
	for i := 0; i < n; i++ {
		for x < len(scanner.Bytes()) && scanner.Bytes()[x] == ' ' {
			x++
		}
		x = readInt(scanner.Bytes(), x, &res[i])
	}
	return res
}

func fillNNums(scanner *bufio.Scanner, n int, res []int) {
	x := 0
	scanner.Scan()
	for i := 0; i < n; i++ {
		for x < len(scanner.Bytes()) && scanner.Bytes()[x] == ' ' {
			x++
		}
		x = readInt(scanner.Bytes(), x, &res[i])
	}
}

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
}

func main() {
	var tc int
	fmt.Scanf("%d", &tc)

	for tc > 0 {
		tc--
		var n, k, m uint64
		fmt.Scanf("%d %d %d", &n, &k, &m)
		fmt.Println(solve(n, k, m))
	}
}

const MAX_M = 1000007

var minPrimeFactors []int

var fact []uint64

func init() {
	minPrimeFactors = make([]int, MAX_M)

	set := make([]bool, MAX_M)

	for x := 2; x < MAX_M; x++ {
		if !set[x] {
			minPrimeFactors[x] = x
			X := int64(x)
			Y := X * X
			for Y < MAX_M {
				y := int(Y)
				if minPrimeFactors[y] == 0 {
					minPrimeFactors[y] = x
				}
				set[y] = true
				Y += X
			}
		}
	}

	fact = make([]uint64, MAX_M)
}

func expo(a, b uint64, mod uint64) uint64 {
	r := uint64(1)

	for b != 0 {
		if b&1 == 1 {
			r *= a
			r %= mod
		}
		a *= a
		a %= mod
		b >>= 1
	}
	return r
}

func preFact(mod uint64, p uint64) {
	fact[0] = 1
	for i := 1; i < int(mod); i++ {
		if uint64(i)%p != 0 {
			fact[i] = (fact[i-1] * uint64(i)) % mod
		} else {
			fact[i] = fact[i-1]
		}
	}
}

func fmod(n, mod, p uint64) uint64 {
	if n == 0 {
		return 1
	}
	ans := expo(fact[mod-1], n/mod, mod)
	ans *= fact[n%mod]
	ans %= mod
	ans *= fmod(n/p, mod, p)
	ans %= mod
	return ans
}

func gcdExtend(a, b int64, x, y *int64) int64 {
	if a == 0 {
		*x = 0
		*y = 1
		return b
	}

	var x1, y1 int64

	g := gcdExtend(b%a, a, &x1, &y1)

	*x = y1 - (b/a)*x1
	*y = x1

	return g
}

func modExtend(a, m int64) int64 {
	var x, y int64
	g := gcdExtend(a, m, &x, &y)
	fmt.Fprintf(os.Stderr, "%d\n", g)
	return (x%m + m) % m
}

func solve(N, K, M uint64) (int, uint64) {
	n := (N + K - 1) / K

	if K > N {
		return int(n), 0
	}

	l := n * K

	if l == N {
		return int(n), 1
	}
	// answer is nCr(l - N + n - 1, n - 1), distribute (l - N) shortings into (n - 1) segements. it is a stars and bars problem.
	// https://en.wikipedia.org/wiki/Stars_and_bars_(combinatorics)

	a := l - N + n - 1
	b := n - 1

	// res = nCr(a, b) % m, using chinese reminder theorem
	// let m = p0 ^^ a0 * p1 ^^ a1 * p2 ^^ a2
	// then calculate r0 = nCr(a, b, p0 ^^ a0), ... r2 = nCr(a, b, p2 ^^ a2)
	// x = CRT(num, rem)
	pfac := primeFactorization(M)

	var divs []uint64
	var rem []uint64

	for i := 0; i < len(pfac); i++ {
		first, second := pfac[i][0], pfac[i][1]
		mod := expo(first, second, M+1)
		divs = append(divs, mod)

		term1, term2, term3 := a, b, a-b
		var x uint64

		for term1 != 0 {
			x += term1 / first
			x -= term2 / first
			x -= term3 / first
			term1 /= first
			term2 /= first
			term3 /= first
		}
		if x >= second {
			rem = append(rem, 0)
			continue
		}
		preFact(mod, first)

		t1 := fmod(a, mod, first)
		t2 := fmod(b, mod, first)
		t3 := fmod(a-b, mod, first)

		it2 := uint64(modExtend(int64(t2), int64(mod)))
		it3 := uint64(modExtend(int64(t3), int64(mod)))
		f := (((t1 * it2) % mod) * it3) % mod
		f *= expo(first, x, mod)
		f %= mod
		rem = append(rem, f)
	}

	return int(n), CRT(divs, rem)
}

func CRT(divs, rem []uint64) uint64 {
	prod := uint64(1)
	for i := 0; i < len(divs); i++ {
		prod *= divs[i]
	}

	var sum uint64

	for i := 0; i < len(divs); i++ {
		pp := prod / divs[i]
		tmp := pp * rem[i]
		tmp *= uint64(modExtend(int64(pp), int64(divs[i])))
		sum += tmp
	}
	return sum % prod
}

func primeFactorization(m uint64) [][]uint64 {

	var factors []int

	num := int(m)

	for num > 1 {
		p := minPrimeFactors[num]
		factors = append(factors, p)
		num /= p
	}

	sort.Ints(factors)

	var res [][]uint64
	var j int

	for i := 1; i <= len(factors); i++ {
		if i == len(factors) || factors[i] > factors[i-1] {
			res = append(res, []uint64{uint64(factors[i-1]), uint64(i - j)})
			j = i
		}
	}

	return res
}
