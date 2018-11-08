package main

import (
	"bufio"
	"fmt"
	"math"
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
		m, n := readTwoNums(scanner)
		res := solve(m, n)
		for _, num := range res {
			fmt.Println(num)
		}
		fmt.Println()
	}
}

const MAX_N = 100000

var primes []int

func init() {
	set := make([]bool, MAX_N)
	primes = make([]int, MAX_N)
	var i int
	for num := 2; num < MAX_N; num++ {
		if !set[num] {
			primes[i] = num
			i++
			x := int64(num)
			for y := x * x; y < MAX_N; y += x {
				set[int(y)] = true
			}
		}
	}
	primes = primes[:i]
}

func solve(m, n int) []int {
	res := make([]int, n-m+1)

	var i int

	for num := m; num <= n; num++ {
		if checkPrime(num) {
			res[i] = num
			i++
		}
	}

	return res[:i]
}

func checkPrime(num int) bool {
	if num == 1 {
		return false
	}
	sqrt := int(math.Sqrt(float64(num)))

	for i := 0; i < len(primes) && primes[i] <= sqrt; i++ {
		if num%primes[i] == 0 {
			return false
		}
	}
	return true
}
