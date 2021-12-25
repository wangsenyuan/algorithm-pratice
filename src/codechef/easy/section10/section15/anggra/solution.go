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

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var tmp int64
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp
	return i
}

func readNInt64Nums(scanner *bufio.Scanner, n int) []int64 {
	res := make([]int64, n)
	x := -1
	scanner.Scan()
	for i := 0; i < n; i++ {
		x = readInt64(scanner.Bytes(), x+1, &res[i])
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		line := readNNums(scanner, 3)
		res := solve(line[0], line[1], line[2])
		fmt.Println(res)
	}
}

const MOD = 1000000007

func solve(n, m, k int) int {
	f := make([]int64, n+1)
	f[0] = 1
	for i := 1; i <= n; i++ {
		f[i] = (int64(i) * f[i-1]) % MOD
	}
	iv := make([]int64, n+1)
	iv[n] = inv(f[n])

	for i := n - 1; i >= 0; i-- {
		iv[i] = iv[i+1] * int64(i+1)
		iv[i] %= MOD
	}

	C := func(x, y int) int64 {
		res := f[x] * iv[y]
		res %= MOD
		res *= iv[x-y]
		res %= MOD
		return res
	}

	// C(n, k)  * pow(m - 1, n - k)
	var res int64

	for i := k; i <= n; i++ {
		res += C(n, i) * pow(m-1, n-i)
		res %= MOD
	}

	return int(res)
}

func inv(num int64) int64 {
	return pow(int(num), MOD-2)
}

func pow(a int, b int) int64 {
	A := int64(a)
	res := int64(1)
	for b > 0 {
		if b&1 == 1 {
			res *= A
			res %= MOD
		}
		A *= A
		A %= MOD
		b >>= 1
	}
	return res
}
