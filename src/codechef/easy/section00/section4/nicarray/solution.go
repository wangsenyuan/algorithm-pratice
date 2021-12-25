package main

import (
	"bufio"
	"fmt"
	"os"
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
	scanner := bufio.NewScanner(os.Stdin)
	tc := readNum(scanner)

	for tc > 0 {
		tc--
		n, s := readTwoNums(scanner)
		fillNNums(scanner, n, nums[:n])
		fmt.Println(solve(n, s, nums[:n]))
	}
}

const MOD = 1e9 + 7

var f [51]int64
var fi [51]int64

var nums [51]int

func init() {
	f[0] = 1
	fi[0] = 1

	for i := 1; i < 51; i++ {
		f[i] = (int64(i) * f[i-1]) % MOD
		fi[i] = pow(f[i], MOD-2)
	}
}

func pow(a int64, n int64) int64 {
	res := int64(1)
	for n > 0 {
		if n&1 == 1 {
			res = (res * a) % MOD
		}
		a = (a * a) % MOD
		n >>= 1
	}
	return res
}

func solve(n int, S int, A []int) int {
	var m int
	for i := 0; i < n; i++ {
		if A[i] == -1 {
			continue
		}
		S -= A[i]
		A[m] = A[i]
		m++
	}

	if S < 0 {
		return 0
	}

	var count func(pos int, s int, v int) int64
	count = func(pos int, s int, v int) int64 {
		if pos == n {
			if s > 0 {
				return 0
			}
			var ans int64
			for i := 0; i < n; i++ {
				for j := i + 1; j < n; j++ {
					ans += int64(gcd(A[i], A[j]))
				}
			}
			cnt := 1
			for i := m + 1; i <= n; i++ {
				if i < n && A[i] == A[i-1] {
					cnt++
					continue
				}
				ans = (ans * fi[cnt]) % MOD
				cnt = 1
			}
			return (ans * f[n-m]) % MOD
		}

		var ans int64
		for i := v; i <= s-(n-pos-1); i++ {
			A[pos] = i
			ans += count(pos+1, s-i, i)
			if ans >= MOD {
				ans -= MOD
			}
		}
		return ans
	}

	return int(count(m, S, 1))
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
