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
		scanner.Scan()
		line := scanner.Bytes()
		var N uint64
		pos := readUint64(line, 0, &N)
		S := string(line[pos+1:])
		fmt.Println(solve(N, S))
	}
}

const MOD = 1000000007

func solve(N uint64, S string) int {
	M := len(S)

	if N <= uint64(M) {
		return 0
	}

	n := int64(N - uint64(M))

	x := pow(26, n-1)
	X := (x * 25) % MOD
	Y := (x * 26) % MOD
	R := int64(M)*X + Y
	R %= MOD

	return int(R)
}

func pow(a int64, b int64) int64 {
	res := int64(1)
	for b > 0 {
		if b&1 == 1 {
			res *= a
			res %= MOD
		}
		a *= a
		a %= MOD
		b >>= 1
	}

	return res
}
