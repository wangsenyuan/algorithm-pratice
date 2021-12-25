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
		res := solve(scanner.Text())
		fmt.Println(res)
	}

}

const MOD = 1e9 + 7

func solve(N string) int {
	L := len(N)
	p1 := pow(10, L-1)
	p2 := (p1 * 10) % MOD

	var num int

	for i := 0; i < L; i++ {
		num = num*10 + int(N[i]-'0')
		num %= MOD
	}

	var ans = int64(num)

	for i := 0; i < L-1; i++ {
		a0 := int(N[i] - '0')
		x := num - a0*p1
		num = x*10 + a0
		ans = (ans*int64(p2) + int64(num)) % MOD
	}

	return int(ans)
}

func pow(a int, n int) int {
	x := int64(a)
	y := int64(1)
	for n > 0 {
		if n&1 == 1 {
			y = (y * x) % MOD
		}
		x = (x * x) % MOD
		n >>= 1
	}
	return int(y)
}
