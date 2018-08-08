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

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	tc := readNum(scanner)

	for tc > 0 {
		tc--
		N, K := readTwoNums(scanner)
		fmt.Println(N, K)
	}
}

const MOD = 1000000007

func solve(N, K int) int {
	res := int64(K)
	res = (res * int64(K-1)) % MOD
	num := int64(K-2)*int64(K-2) + int64(K-1)
	num %= MOD
	res = (res * pow(num, N-1)) % MOD
	return int(res)
}

func pow(x int64, n int) int64 {
	a := x
	b := int64(1)
	for n > 0 {
		if n&1 == 1 {
			b = b * a
			b %= MOD
		}
		a = (a * a) % MOD
		n >>= 1
	}
	return b
}
