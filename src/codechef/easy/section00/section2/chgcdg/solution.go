package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	tmp := 0
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp
	return i
}

func readNum(scanner *bufio.Scanner) (a int) {
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &a)
	return
}

func readTwoNums(scanner *bufio.Scanner) (a int, b int) {
	scanner.Scan()
	x := readInt(scanner.Bytes(), 0, &a)
	readInt(scanner.Bytes(), x+1, &b)
	return
}

func readNNums(scanner *bufio.Scanner, n int) []int {
	res := make([]int, n)
	x := -1
	scanner.Scan()
	for i := 0; i < n; i++ {
		x = readInt(scanner.Bytes(), x+1, &res[i])
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		n := readNum(scanner)
		A := readNNums(scanner, n)
		res := solve(n, A)
		fmt.Println(res)
	}
}

const MAX_A = 1e6

var factors [][]int
var primes []int

func init() {
	set := make([]bool, 100)
	primes = make([]int, 0, 100)
	for i := 2; i < 100; i++ {
		if !set[i] {
			primes = append(primes, i)
			for j := i * i; j < 100; j += i {
				set[j] = true
			}
		}
	}
	factors = make([][]int, MAX_A+1)
	factors[0] = make([]int, len(primes))
	factors[1] = make([]int, len(primes))
	for x := 2; x <= MAX_A; x++ {
		factors[x] = make([]int, len(primes))
		y := x
		for i := 0; i < len(primes) && y >= primes[i]; i++ {
			var cnt int
			for y%primes[i] == 0 {
				cnt++
				y /= primes[i]
			}
			factors[x][i] = cnt
		}
	}
}

func solve(n int, A []int) int {
	g := A[0]
	for i := 1; i < n; i++ {
		g = gcd(g, A[i])
	}
	for i := 0; i < n; i++ {
		A[i] /= g
	}

	check := func(h, j int) bool {
		a, b := 0, 0
		for i := 0; i < n; i++ {
			k := factors[A[i]][j]
			if k == 0 {
				a++
			} else if k > h {
				b += (k - h) >> 1
			} else {
				b -= h - k
			}
		}
		b -= a * h
		return b >= 0
	}

	find := func(j int) int {
		// has no prime factor p
		var left, right int
		for i := 0; i < n; i++ {
			right = max(right, factors[A[i]][j]+1)
		}
		for left < right {
			mid := left + (right-left)>>1
			if !check(mid, j) {
				right = mid
			} else {
				left = mid + 1
			}
		}
		return pow(primes[j], left-1)
	}

	ans := 1
	for i := 0; i < len(primes); i++ {
		// p := primes[i]
		ans *= find(i)
	}
	return ans * g
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func pow(a, n int) int {
	x := 1
	y := a
	for n > 0 {
		if n&1 == 1 {
			x *= y
		}
		y = y * y
		n >>= 1
	}
	return x
}
